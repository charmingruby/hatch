package main

import (
	"time"

	"HATCH_APP/config"
	"HATCH_APP/internal/note"
	"HATCH_APP/internal/shared/http"
	"HATCH_APP/pkg/db/postgres"
	"HATCH_APP/pkg/telemetry"
	"HATCH_APP/pkg/validator"

	"github.com/joho/godotenv"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
)

func main() {
	log := telemetry.NewLogger()

	_ = godotenv.Load()

	fx.New(
		fx.Supply(log),
		config.Module,
		postgres.Module,
		validator.Module,
		http.Module,
		note.Module,
		fx.WithLogger(func() fxevent.Logger {
			return log
		}),
		fx.StartTimeout(30*time.Second),
		fx.StopTimeout(15*time.Second),
	).Run()
}
