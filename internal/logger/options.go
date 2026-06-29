package logger

import (
	"fmt"
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
	switch level {
	case LevelDebug:
		slog.SetLogLoggerLevel(slog.LevelDebug)
	case LevelInfo:
		slog.SetLogLoggerLevel(slog.LevelInfo)
	case LevelWarn:
		slog.SetLogLoggerLevel(slog.LevelWarn)
	case LevelError:
		slog.SetLogLoggerLevel(slog.LevelError)
	default:
		return fmt.Errorf("invalid level: %s", level)

	}

	switch format {
	case FormatJson:
		setLogJsonFormat()
	case FormatText:
		setLogTextFormat()
	default:
		return fmt.Errorf("invalid format: %s", format)
	}

	defaultOptions = &options{
		level:  level,
		format: format,
	}
	return nil
}

func setLogTextFormat() {
	slog.SetDefault(slog.New(slog.NewTextHandler(
		os.Stdout, &slog.HandlerOptions{
			AddSource: true,
		})))
}

func setLogJsonFormat() {
	slog.SetDefault(slog.New(slog.NewJSONHandler(
		os.Stdout, &slog.HandlerOptions{
			AddSource: true,
		})))
}
