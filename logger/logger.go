package logger

import (
	"log/slog"
	"os"
)

var (
	Logger  *slog.Logger
	logFile *os.File
)

func init() {
	file, err := os.OpenFile("application.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o666)
	if err != nil {
		slog.Error("Failed to open log file", "err", err)
		return
	}

	Logger = slog.New(slog.NewJSONHandler(file, &slog.HandlerOptions{Level: slog.LevelInfo}))
}

func Close() error {
	if logFile != nil {
		return logFile.Close()
	}
	return nil
}
