// Package logger provides a preconfigured structured logger.
package logger

import (
	"log/slog"
	"os"
)

// Predefined log levels for easier configuration.
const (
	DebugLevel = "debug"
	InfoLevel  = "info"
	ErrorLevel = "error"
	WarnLevel  = "warn"
)

// Logger is an alias for slog.Logger to simplify usage across the codebase.
type Logger = slog.Logger

// logger holds the current global logger instance used by the application.
var logger *Logger

// New initializes a new logger instance with debug level.
//
// This should be called at the start of the application, before loading config.
//
// Returns:
//   - *Logger: the configured logger instance.
func New() *Logger {
	log := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))

	logger = log
	slog.SetDefault(logger)

	return logger
}

// ChangeLevel updates the log level of the global logger at runtime, replacing the handler.
//
// Accepts one of the predefined level strings (case-sensitive: "debug", "info", "warn", "error".
//
// Defaults to "info" if an unrecognized value is passed.
//
// Parameters:
//   - string: log level.
//
// Returns:
//   - string: the log level configured.
func ChangeLevel(level string) string {
	var slogLevel slog.Level

	switch level {
	case DebugLevel:
		slogLevel = slog.LevelDebug
	case InfoLevel:
		slogLevel = slog.LevelInfo
	case WarnLevel:
		slogLevel = slog.LevelWarn
	case ErrorLevel:
		slogLevel = slog.LevelError
	default:
		slogLevel = slog.LevelInfo
	}

	log := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slogLevel,
	}))

	logger = log
	slog.SetDefault(logger)

	return slogLevel.String()
}
