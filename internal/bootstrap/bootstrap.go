package bootstrap

import (
	"context"

	"github.com/linc593823915/atlas/internal/config"
	"github.com/linc593823915/atlas/internal/logger"
	"github.com/linc593823915/atlas/internal/runtime"
)

type Bootstrap struct {
	config  *config.Manager
	runtime *runtime.Runtime
	ctx     context.Context
}

func New(ctx context.Context) (*Bootstrap, error) {
	configManager, err := buildConfigManager()
	if err != nil {
		return nil, err
	}
	runtimeInstance := runtime.CreateRuntime()
	return &Bootstrap{
		ctx:     ctx,
		config:  configManager,
		runtime: runtimeInstance,
	}, nil
}

func buildConfigManager() (*config.Manager, error) {
	manager := config.NewManager()
	if err := manager.Load(); err != nil {
		return nil, err
	}
	if err := logger.WithOptions(manager.Logger().Level, manager.Logger().Format); err != nil {
		return nil, err
	}
	return manager, nil
}

func (b *Bootstrap) Start() error {
	return b.runtime.Start()
}

func (b *Bootstrap) Stop() error {
	er := b.runtime.Stop()
	if er != nil {
		return er
	}
	return nil
}
