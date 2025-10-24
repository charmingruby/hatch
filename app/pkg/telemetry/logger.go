package telemetry

import (
	"log/slog"
	"os"
)

type Logger = slog.Logger

func NewLogger() *Logger {
	log := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))

	return log
}
