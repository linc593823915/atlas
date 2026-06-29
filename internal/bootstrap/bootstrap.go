package bootstrap

import (
	"context"

	"github.com/linc593823915/atlas/internal/config"
	"github.com/linc593823915/atlas/internal/logger"
	"github.com/linc593823915/atlas/internal/runtime"
)

type Bootstrap struct {
	config  *config.Config
	runtime *runtime.Runtime
	ctx     context.Context
}

func New(ctx context.Context) (*Bootstrap, error) {
	er := logger.WithOptions(logger.LevelInfo, logger.FormatText)
	if er != nil {
		return nil, er
	}
	configInstance, err := config.LoadConfig()
	if err != nil {
		return nil, err
	}
	runtimeInstance := runtime.CreateRuntime()
	return &Bootstrap{
		ctx:     ctx,
		config:  configInstance,
		runtime: runtimeInstance,
	}, nil
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
