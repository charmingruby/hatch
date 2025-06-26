package main

import (
	"github/charmingruby/pack/config"
	"github/charmingruby/pack/internal/device"
	"github/charmingruby/pack/pkg/database/postgres"
	"github/charmingruby/pack/pkg/http/rest"
	"github/charmingruby/pack/pkg/telemetry/logger"
	"os"
	"os/signal"

	"github.com/joho/godotenv"
)

func main() {
	log := logger.New()

	if err := godotenv.Load(); err != nil {
		log.Warn("failed to find .env file", "error", err)
	}

	cfg, err := config.New()
	if err != nil {
		log.Error("failed to loading environment variables", "error", err)
		failAndExit()
	}

	db, err := postgres.New(cfg.PostgresURL)
	if err != nil {
		log.Error("failed to connect to Postgres", "error", err)
		failAndExit()
	}

	srv, r := rest.New(cfg.RestServerPort)

	if err := device.New(r, db.Conn); err != nil {
		log.Error("failed to start device module", "error", err)
		failAndExit()

	}

	go func() {
		log.Info("rest server is running...")

		if err := srv.Start(); err != nil {
			log.Error("failed starting rest server", "error", err)
			failAndExit()
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	// proper shutdown
}

func failAndExit() {
	// proper shutdown

	os.Exit(1)
}
