package main

import (
	"HATCH_APP/config"
	"HATCH_APP/internal/note"
	"HATCH_APP/internal/shared/http/rest"
	"HATCH_APP/pkg/db/postgres"
	"HATCH_APP/pkg/telemetry"
	"HATCH_APP/pkg/validator"
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
)

const SHUTDOWN_TIMEOUT = 30 * time.Second

func main() {
	if err := run(); err != nil {
		os.Exit(1)
	}
}

func run() error {
	ctx, stop := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)
	defer stop()

	log := telemetry.NewLogger()

	_ = godotenv.Load()

	log.Info("config: loading...")
	cfg, err := config.Load()
	if err != nil {
		log.Error("config: error loading config", "error", err)

		return err
	}
	log.Info("config: loaded")

	log.Info("postgres: connecting...")
	db, err := postgres.New(ctx, cfg.PostgresURL)
	if err != nil {
		log.Error("postgres: connection error", "error", err)

		return err
	}
	log.Info("postgres: connected")

	val := validator.New()

	srv, r := rest.NewServer(log, cfg, val, db)

	log.Info("note: creating module...")
	if err := note.NewModule(log, r, db.Conn); err != nil {
		log.Error("note: module error", "error", err)

		return err
	}

	errShutdown := make(chan error, 1)

	go shutdown(ctx, errShutdown, log, srv, db)

	log.Info("server: running...", "port", cfg.RestServerPort)

	if err := srv.Start(); err != nil {
		log.Info("server: server start error", "error", err)

		return err
	}

	err = <-errShutdown
	if err != nil {
		return err
	}

	return nil
}

func shutdown(
	ctx context.Context,
	errShutdown chan error,
	log *telemetry.Logger,
	srv *rest.Server,
	db *postgres.Client,
) {
	<-ctx.Done()

	log.Info("shutdown: signal received, starting graceful shutdown...")

	ctxTimeout, cancel := context.WithTimeout(context.Background(), SHUTDOWN_TIMEOUT)
	defer cancel()

	log.Info("shutdown: stopping server...")
	if err := srv.Close(ctxTimeout); err != nil {
		log.Error("shutdown: server error", "error", err)

		errShutdown <- err
		return
	}

	log.Info("shutdown: closing database...")
	if err := db.Close(); err != nil {
		log.Error("shutdown: database error", "error", err)

		errShutdown <- err
		return
	}

	log.Info("shutdown: completed successfully")

	errShutdown <- nil
}
