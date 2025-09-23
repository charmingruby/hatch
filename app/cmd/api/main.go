package main

import (
	"context"
	"os"
	"os/signal"
	"time"

	"PACK_APP/config"
	"PACK_APP/internal/health"
	"PACK_APP/internal/note"
	"PACK_APP/internal/shared/http/rest"
	"PACK_APP/pkg/database/postgres"
	"PACK_APP/pkg/logger"

	"github.com/joho/godotenv"
)

func main() {
	log := logger.New()

	if err := godotenv.Load(); err != nil {
		log.Warn("env: missing", "err", err)
	}

	log.Info("config: loading")
	cfg, err := config.New()
	if err != nil {
		log.Error("config: error", "err", err)
		failAndExit(log, nil, nil)
	}
	log.Info("config: loaded")

	logLevel := logger.ChangeLevel(cfg.LogLevel)
	log.Info("logger: set", "level", logLevel)

	log.Info("postgres: connecting")
	db, err := postgres.New(log, cfg.PostgresURL)
	if err != nil {
		log.Error("postgres: error", "err", err)
		failAndExit(log, nil, nil)
	}
	log.Info("postgres: ready")

	srv, r := rest.New(cfg.RestServerPort)

	health.New(r, db)

	if err := note.New(log, r, db.Conn); err != nil {
		log.Error("note: error", "err", err)
		failAndExit(log, nil, db)
	}
	log.Info("note: ready")

	go func() {
		log.Info("server: starting", "port", cfg.RestServerPort)
		if err := srv.Start(); err != nil {
			log.Error("server: error", "err", err)
			failAndExit(log, srv, db)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	log.Info("signal: received")
	log.Info("shutdown: starting")

	code := gracefulShutdown(log, srv, db)

	log.Info("shutdown: complete", "code", code)
	os.Exit(code)
}

func failAndExit(log *logger.Logger, srv *rest.Server, db *postgres.Client) {
	gracefulShutdown(log, srv, db)
	os.Exit(1)
}

func gracefulShutdown(log *logger.Logger, srv *rest.Server, db *postgres.Client) int {
	parentCtx := context.Background()
	var hasError bool

	if srv != nil {
		ctx, cancel := context.WithTimeout(parentCtx, 15*time.Second)
		defer cancel()

		if err := srv.Close(ctx); err != nil {
			log.Error("server: error", "err", err)
			hasError = true
		} else {
			log.Info("server: stopped")
		}
	}

	if db != nil {
		if err := db.Close(); err != nil {
			log.Error("postgres: error", "err", err)
			hasError = true
		} else {
			log.Info("postgres: stopped")
		}
	}

	if hasError {
		return 1
	}

	return 0
}
