package logger

import (
	"context"
	"log/slog"
)

func Info(ctx context.Context, msg string, fields ...any) {
	emit(ctx, slog.LevelInfo, msg, fields...)
}

func Error(ctx context.Context, msg string, fields ...any) {
	emit(ctx, slog.LevelError, msg, fields...)
}

func Debug(ctx context.Context, msg string, fields ...any) {
	emit(ctx, slog.LevelDebug, msg, fields...)
}

func Warn(ctx context.Context, msg string, fields ...any) {
	emit(ctx, slog.LevelWarn, msg, fields...)
}
