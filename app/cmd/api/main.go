package main

import (
	"time"

	"HATCH_APP/config"
	"HATCH_APP/internal/health"
	"HATCH_APP/internal/note"
	"HATCH_APP/internal/shared/http/rest"
	"HATCH_APP/pkg/database/postgres"
	"HATCH_APP/pkg/logger"

	"github.com/joho/godotenv"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
)

func main() {
	log := logger.New()

	if err := godotenv.Load(); err != nil {
		log.Warn("env: missing", "error", err)
	}

	fx.New(
		fx.Supply(log),
		config.Module,
		postgres.Module,
		rest.Module,
		health.Module,
		note.Module,
		fx.WithLogger(func() fxevent.Logger {
			return log
		}),
		fx.StartTimeout(30*time.Second),
		fx.StopTimeout(15*time.Second),
	).Run()
}
