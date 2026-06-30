package main

import (
	"context"
	"os"

	"github.com/linc593823915/atlas/internal/config"
	"github.com/linc593823915/atlas/internal/logger"
)

// main initializes the shared logger and runs the cookbook PDF command.
func main() {
	if err := setupLogger(); err != nil {
		logger.Error(context.Background(), "setup logger failed", "error", err)
		os.Exit(1)
	}
	if err := run(context.Background(), os.Args[1:]); err != nil {
		logger.Error(context.Background(), "cookbook pdf generation failed", "error", err)
		os.Exit(1)
	}
}

// setupLogger loads project config and configures the internal logger.
func setupLogger() error {
	manager := config.NewManager()
	if err := manager.Load(); err != nil {
		return err
	}
	return logger.WithOptions(manager.Logger().Level, manager.Logger().Format)
}
