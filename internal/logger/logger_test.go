package logger

import (
	"context"
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
