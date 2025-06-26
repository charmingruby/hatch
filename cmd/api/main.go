package main

import (
	"context"
	"github/charmingruby/pack/config"
	"github/charmingruby/pack/internal/device"
	"github/charmingruby/pack/pkg/database/postgres"
	"github/charmingruby/pack/pkg/http/rest"
	"github/charmingruby/pack/pkg/telemetry/logger"
	"os"
	"os/signal"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	log := logger.New()

	if err := godotenv.Load(); err != nil {
		log.Warn("failed to find .env file", "error", err)
	}

	log.Info("loading environment variables...")

	cfg, err := config.New()
	if err != nil {
		log.Error("failed to loading environment variables", "error", err)
		failAndExit(log, nil, nil)
	}

	log.Info("environment variables loaded")

	log.Info("connecting to Postgres...")

	db, err := postgres.New(cfg.PostgresURL)
	if err != nil {
		log.Error("failed to connect to Postgres", "error", err)
		failAndExit(log, nil, nil)
	}

	log.Info("connected to Postgres")

	srv, r := rest.New(cfg.RestServerPort)

	log.Info("starting device module...")

	if err := device.New(r, db.Conn); err != nil {
		log.Error("failed to start device module", "error", err)
		failAndExit(log, db, nil)
	}

	log.Info("device module started")

	go func() {
		log.Info("rest server is running...")

		if err := srv.Start(); err != nil {
			log.Error("failed starting rest server", "error", err)
			failAndExit(log, db, srv)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	log.Info("received an interrupt signal")

	log.Info("starting graceful shutdown...")

	gracefulShutdown(log, db, srv)

	log.Info("gracefully shutdown")
}

func failAndExit(log *logger.Logger, db *postgres.Client, srv *rest.Server) {
	gracefulShutdown(log, db, srv)
	os.Exit(1)
}

func gracefulShutdown(log *logger.Logger, db *postgres.Client, srv *rest.Server) {
	parentCtx := context.Background()

	var hasError bool

	if db != nil {
		ctx, cancel := context.WithTimeout(parentCtx, 10*time.Second)
		defer cancel()

		if err := db.Close(ctx); err != nil {
			log.Error("error closing Postgres connection", "error", err)
			hasError = true
		}
	}

	if srv != nil {
		ctx, cancel := context.WithTimeout(parentCtx, 15*time.Second)
		defer cancel()

		if err := srv.Stop(ctx); err != nil {
			log.Error("error closing rest server", "error", err)
			hasError = true
		}
	}

	if hasError {
		os.Exit(1)
	}

	os.Exit(0)
}
