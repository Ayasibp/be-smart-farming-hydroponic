package logger

import (
	"log/slog"
	"os"
)

var logger *slog.Logger

func Init(logFilePath string) error {
	logFile, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return err
	}

	logger = slog.New(slog.NewTextHandler(logFile, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))

	slog.SetDefault(logger)

	return nil
}

func Info(msg string, process string, details interface{}) {
	logger.Info(msg,
		slog.String("process", process),
		slog.Any("details", details))
}

func Debug(msg string, process string, details interface{}) {
	logger.Debug(msg,
		slog.String("process", process),
		slog.Any("details", details))
}

func Warn(msg string, process string, details interface{}) {
	logger.Warn(msg,
		slog.String("process", process),
		slog.Any("details", details))
}

func Error(msg string, process string, details interface{}) {
	logger.Error(msg,
		slog.String("process", process),
		slog.Any("details", details))
}
