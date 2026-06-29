package logger

import (
	"bytes"
	"context"
	"log/slog"
	"strings"
	"testing"
)

func TestDebug(t *testing.T) {
	err := WithOptions(LevelDebug, FormatText)
	if err != nil {
		t.Fatal(err)
	}
	Debug(context.Background(), "test", "name", "linc")
}

func TestError(t *testing.T) {
	err := WithOptions(LevelDebug, FormatText)
	if err != nil {
		t.Fatal(err)
	}
	Error(context.Background(), "test", "name", "linc")
}

func TestInfo(t *testing.T) {
	err := WithOptions(LevelDebug, FormatText)
	if err != nil {
		t.Fatal(err)
	}
	Info(context.Background(), "test", "name", "linc")
}

func TestWarn(t *testing.T) {
	err := WithOptions(LevelDebug, FormatText)
	if err != nil {
		t.Fatal(err)
	}
	Warn(context.Background(), "test", "name", "linc")
}

func TestInfoSourceUsesCallerLocation(t *testing.T) {
	var buffer bytes.Buffer
	handler, err := newHandler(FormatText, &buffer, slog.LevelInfo)
	if err != nil {
		t.Fatal(err)
	}

	previous := slog.Default()
	slog.SetDefault(slog.New(handler))
	t.Cleanup(func() {
		slog.SetDefault(previous)
	})

	Info(context.Background(), "test", "name", "linc")

	output := buffer.String()
	if !strings.Contains(output, "logger_test.go:") {
		t.Fatalf("source should point to caller, got %q", output)
	}
}

func TestWarnLevelSuppressesInfo(t *testing.T) {
	var buffer bytes.Buffer
	handler, err := newHandler(FormatText, &buffer, slog.LevelWarn)
	if err != nil {
		t.Fatal(err)
	}

	previous := slog.Default()
	slog.SetDefault(slog.New(handler))
	t.Cleanup(func() {
		slog.SetDefault(previous)
	})

	Info(context.Background(), "info message")
	Warn(context.Background(), "warn message")

	output := buffer.String()
	if strings.Contains(output, "info message") {
		t.Fatalf("info log should be filtered at warn level, got %q", output)
	}
	if !strings.Contains(output, "warn message") {
		t.Fatalf("warn log should be emitted, got %q", output)
	}
}
