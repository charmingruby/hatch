package main

import (
	"HATCH_APP/config"
	"HATCH_APP/internal/common/http/rest"
	"HATCH_APP/internal/note"

	"HATCH_APP/pkg/db/postgres"
	"HATCH_APP/pkg/o11y/logging"
	"HATCH_APP/pkg/validator"
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/jmoiron/sqlx"
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

	logging.Init()

	_ = godotenv.Load()

	logging.Log.Info("config: loading...")

	cfg, err := config.Load()
	if err != nil {
		logging.Log.Error("config: error loading config", "error", err)

		return err
	}

	logging.Log.Info("config: loaded")

	logging.Log.Info("postgres: connecting...")

	db, err := postgres.Connect(ctx, cfg.PostgresURL)
	if err != nil {
		logging.Log.Error("postgres: connection error", "error", err)

		return err
	}

	logging.Log.Info("postgres: connected")

	val := validator.New()

	srv, r := rest.NewServer(cfg, val, db)

	logging.Log.Info("note: creating module...")

	if err := note.Register(r, db); err != nil {
		logging.Log.Error("note: module error", "error", err)

		return err
	}

	logging.Log.Info("note: module created")

	shutdownErrCh := make(chan error, 1)

	go shutdown(ctx, shutdownErrCh, srv, db)

	logging.Log.Info("server: running...", "port", cfg.RestServerPort)

	if err := srv.Start(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		logging.Log.Info("server: server start error", "error", err)

		return err
	}

	logging.Log.Info("shutdown: signal received, starting graceful shutdown...")

	err = <-shutdownErrCh
	if err != nil {
		logging.Log.Error("shutdown: shutdown error", "error", err)

		return err
	}

	logging.Log.Info("shutdown: gracefully shutdown")

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
