package logging

import (
	"log/slog"
	"os"
)

func NewLogger(level string) *slog.Logger {
	options := &slog.HandlerOptions{
		Level: parseLevel(level),
	}

	return slog.New(slog.NewTextHandler(os.Stderr, options))
}

func parseLevel(level string) slog.Level {
	switch level {
	case "debug":
		return slog.LevelDebug
	case "info":
		return slog.LevelInfo
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}
