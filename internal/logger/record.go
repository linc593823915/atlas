package logger

import (
	"context"
	"log/slog"
	"runtime"
	"time"
)

const skipLoggerFrames = 5

func emit(ctx context.Context, level slog.Level, msg string, fields ...any) {
	defaultLogger := slog.Default()
	if !defaultLogger.Enabled(ctx, level) {
		return
	}
	record := newRecord(level, msg, fields)
	_ = defaultLogger.Handler().Handle(ctx, record)
}

func newRecord(level slog.Level, msg string, fields []any) slog.Record {
	record := slog.NewRecord(time.Now(), level, msg, callerPC())
	record.Add(fields...)
	return record
}

func callerPC() uintptr {
	var pcs [1]uintptr
	runtime.Callers(skipLoggerFrames, pcs[:])
	return pcs[0]
}
