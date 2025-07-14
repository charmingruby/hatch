package main

import (
	"context"
	"fmt"
	"github/charmingruby/pack/config"
	"github/charmingruby/pack/internal/device"
	"github/charmingruby/pack/internal/platform"
	"github/charmingruby/pack/pkg/delivery/http/rest"
	"github/charmingruby/pack/pkg/telemetry/logger"
	"github/charmingruby/pack/pkg/validator"
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
		failAndExit(log, nil)
	}

	log.Info("environment variables loaded")

	logLevel := logger.ChangeLevel(cfg.LogLevel)

	log.Info("log level configured", "level", logLevel)

	srv, r := rest.New(cfg.RestServerPort)

	val := validator.New()

	if err := device.New(log, r, val); err != nil {
		log.Error("failed to start device module", "error", err)
		failAndExit(log, nil)
	}

	platform.New(r)

	go func() {
		log.Info("rest server is running...")

		if err := srv.Start(); err != nil {
			log.Error("failed starting rest server", "error", err)
			failAndExit(log, srv)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	log.Info("received an interrupt signal")

	log.Info("starting graceful shutdown...")

	signal := gracefulShutdown(log, srv)

	log.Info(fmt.Sprintf("gracefully shutdown, with code %d", signal))

	os.Exit(signal)
}

func failAndExit(log *logger.Logger, srv *rest.Server) {
	gracefulShutdown(log, srv)
	os.Exit(1)
}

func gracefulShutdown(log *logger.Logger, srv *rest.Server) int {
	parentCtx := context.Background()

	var hasError bool

	if srv != nil {
		ctx, cancel := context.WithTimeout(parentCtx, 15*time.Second)
		defer cancel()

		if err := srv.Stop(ctx); err != nil {
			log.Error("error closing rest server", "error", err)
			hasError = true
		}
	}

	if hasError {
		return 1
	}

	return 0
}
