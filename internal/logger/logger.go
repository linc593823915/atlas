package logger

import (
	"context"
	"log/slog"
)

func Info(ctx context.Context, msg string, fields ...any) {
	slog.InfoContext(ctx, msg, fields...)
}

func Error(ctx context.Context, msg string, fields ...any) {
	slog.ErrorContext(ctx, msg, fields...)
}

func Debug(ctx context.Context, msg string, fields ...any) {
	slog.DebugContext(ctx, msg, fields...)
}
func Warn(ctx context.Context, msg string, fields ...any) {
	slog.WarnContext(ctx, msg, fields...)
}
