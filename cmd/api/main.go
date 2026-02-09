package main

import (
	"HATCH_APP/config"
	"HATCH_APP/internal/note"
	"HATCH_APP/pkg/database"
	"HATCH_APP/pkg/http/rest"
	"HATCH_APP/pkg/o11y"
	"HATCH_APP/pkg/validator"
	"context"
	"errors"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/jmoiron/sqlx"
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

	log := o11y.Init()

	log.Info("config: loading...")

	cfg, err := config.Load()
	if err != nil {
		log.Error("config: error loading config", "error", err)

		return err
	}

	log.Info("config: loaded")

	log.Info("postgres: connecting...")

	db, err := database.PostgresConnect(ctx, cfg.PostgresURL)
	if err != nil {
		log.Error("postgres: connection error", "error", err)

		return err
	}

	log.Info("postgres: connected")

	val := validator.New()

	srv, r := rest.NewServer(cfg, val, db)

	log.Info("note: creating module...")

	if err := note.Register(r, db); err != nil {
		log.Error("note: module error", "error", err)

		return err
	}

	log.Info("note: module created")

	shutdownErrCh := make(chan error, 1)

	go shutdown(ctx, shutdownErrCh, srv, db)

	log.Info("server: running...", "port", cfg.RestServerPort)

	if err := srv.Start(); err != nil {
		log.Info("server: server start error", "error", err)

		return err
	}

	log.Info("shutdown: signal received, starting graceful shutdown...")

	err = <-shutdownErrCh
	if err != nil {
		log.Error("shutdown: shutdown error", "error", err)

		return err
	}

	log.Info("shutdown: gracefully shutdown")

	return nil
}

func shutdown(
	ctx context.Context,
	errShutdown chan error,
	srv *rest.Server,
	db *sqlx.DB,
) {
	<-ctx.Done()

	ctxTimeout, cancel := context.WithTimeout(context.Background(), SHUTDOWN_TIMEOUT)
	defer cancel()

	err := srv.Close(ctxTimeout)
	switch {
	case err == nil:
	case errors.Is(err, context.DeadlineExceeded):
		errShutdown <- errors.New("deadline exceeded, forcing shutdown")
		return
	default:
		errShutdown <- errors.New("forcing shutdown")
		return
	}

	if err := db.Close(); err != nil {
		errShutdown <- err
		return
	}

	errShutdown <- nil
}
