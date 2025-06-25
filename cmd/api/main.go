package main

import (
	"github/charmingruby/gew/config"
	"github/charmingruby/gew/pkg/telemetry/logger"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	log := logger.New()

	if err := godotenv.Load(); err != nil {
		log.Warn("failed to find .env file", "error", err)
	}

	_, err := config.New()
	if err != nil {
		log.Error("failed to loading environment variables", "error", err)
		failAndExit()
	}
}

func failAndExit() {
	// proper shutdown

	os.Exit(1)
}
