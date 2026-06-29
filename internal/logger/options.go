package logger

import (
	"fmt"
	"io"
	"log/slog"
	"os"
)

const (
	LevelDebug = "debug"
	LevelInfo  = "info"
	LevelWarn  = "warn"
	LevelError = "error"

	FormatText = "text"
	FormatJson = "json"
)

type options struct {
	level  string
	format string
}

var defaultOptions = &options{
	level:  LevelInfo,
	format: FormatText,
}

func WithOptions(level, format string) error {
	logLevel, err := parseLevel(level)
	if err != nil {
		return err
	}
	if err := setLogLevel(logLevel); err != nil {
		return err
	}
	handler, err := newHandler(format, os.Stdout, logLevel)
	if err != nil {
		return err
	}
	slog.SetDefault(slog.New(handler))
	defaultOptions = &options{
		level:  level,
		format: format,
	}
	return nil
}

func parseLevel(level string) (slog.Level, error) {
	switch level {
	case LevelDebug:
		return slog.LevelDebug, nil
	case LevelInfo:
		return slog.LevelInfo, nil
	case LevelWarn:
		return slog.LevelWarn, nil
	case LevelError:
		return slog.LevelError, nil
	default:
		return 0, fmt.Errorf("invalid level: %s", level)
	}
}

func setLogLevel(level slog.Level) error {
	slog.SetLogLoggerLevel(level)
	return nil
}

func newHandler(format string, writer io.Writer, level slog.Level) (slog.Handler, error) {
	switch format {
	case FormatJson:
		return slog.NewJSONHandler(writer, handlerOptions(level)), nil
	case FormatText:
		return slog.NewTextHandler(writer, handlerOptions(level)), nil
	default:
		return nil, fmt.Errorf("invalid format: %s", format)
	}
}

func handlerOptions(level slog.Level) *slog.HandlerOptions {
	return &slog.HandlerOptions{
		AddSource: true,
		Level:     level,
	}
}
