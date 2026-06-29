package logger

import "log/slog"

type Logger struct {
	*slog.Logger
}

func InitLogger() *Logger {
	return &Logger{}
}
