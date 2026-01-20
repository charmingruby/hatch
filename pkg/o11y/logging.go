package o11y

import (
	"context"
	"log/slog"
	"os"
)

var Log *Logger

type Logger = slog.Logger

type ctxKey struct{}

func Init() *Logger {
	Log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))

	return Log
}

func WithLogger(ctx context.Context, log *Logger) context.Context {
	return context.WithValue(ctx, ctxKey{}, log)
}

func FromContext(ctx context.Context) *Logger {
	if ctx == nil {
		return Log
	}

	if log, ok := ctx.Value(ctxKey{}).(*Logger); ok {
		return log
	}

	return Log
}
