package logging

import (
	"log/slog"
	"os"
)

func NewLogger() *slog.Logger {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: getLogLevel()}))
	logger.Info("Starting Service Wheres-My-Subtitle")
	return logger
}
func getLogLevel() slog.Level {
	levelEnv := os.Getenv("LOG_LEVEL")
	var level slog.Level

	switch levelEnv {
	case "DEBUG":
		level = slog.LevelDebug
	case "INFO":
		level = slog.LevelInfo
	case "WARN":
		level = slog.LevelWarn
	case "ERROR":
		level = slog.LevelError
	default:
		level = slog.LevelInfo // Default to INFO if not set
	}
	return level
}
