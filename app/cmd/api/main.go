package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"time"

	"HATCH_APP/config"
	"HATCH_APP/internal/health"
	"HATCH_APP/internal/note"
	"HATCH_APP/internal/shared/http/rest"
	"HATCH_APP/pkg/database/postgres"
	"HATCH_APP/pkg/logger"

	"github.com/joho/godotenv"
)

func main() {
	logger.New()

	if err := godotenv.Load(); err != nil {
		slog.Warn("env: missing", "error", err)
	}

	slog.Info("config: loading")
	cfg, err := config.New()
	if err != nil {
		slog.Error("config: error", "error", err)
		failAndExit(nil, nil)
	}
	slog.Info("config: loaded")

	slog.Info("postgres: connecting")
	db, err := postgres.New(cfg.PostgresURL)
	if err != nil {
		slog.Error("postgres: error", "error", err)
		failAndExit(nil, nil)
	}
	slog.Info("postgres: ready")

	srv, r := rest.New(cfg.RestServerPort)

	health.New(r, db)

	if err := note.New(r, db.Conn); err != nil {
		slog.Error("note: error", "error", err)
		failAndExit(nil, db)
	}
	slog.Info("note: ready")

	go func() {
		slog.Info("server: runnig", "port", cfg.RestServerPort)
		if err := srv.Start(); err != nil {
			slog.Error("server: error", "error", err)
			failAndExit(srv, db)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	slog.Info("signal: received")
	slog.Info("shutdown: starting")

	code := gracefulShutdown(srv, db)

	slog.Info("shutdown: complete", "code", code)
	os.Exit(code)
}

func failAndExit(srv *rest.Server, db *postgres.Client) {
	gracefulShutdown(srv, db)
	os.Exit(1)
}

func gracefulShutdown(srv *rest.Server, db *postgres.Client) int {
	parentCtx := context.Background()
	var hasError bool

	if srv != nil {
		ctx, cancel := context.WithTimeout(parentCtx, 15*time.Second)
		defer cancel()

		if err := srv.Close(ctx); err != nil {
			slog.Error("server: error", "error", err)
			hasError = true
		} else {
			slog.Info("server: stopped")
		}
	}

	if db != nil {
		if err := db.Close(); err != nil {
			slog.Error("postgres: error", "error", err)
			hasError = true
		} else {
			slog.Info("postgres: stopped")
		}
	}

	if hasError {
		return 1
	}

	return 0
}
