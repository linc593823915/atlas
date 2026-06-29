package bootstrap

import (
	"context"

	"github.com/linc593823915/atlas/internal/config"
	"github.com/linc593823915/atlas/internal/logger"
	"github.com/linc593823915/atlas/internal/runtime"
)

type Bootstrap struct {
	config  *config.Config
	logger  *logger.Logger
	runtime *runtime.Runtime
	ctx     context.Context
}

func New(ctx context.Context) (*Bootstrap, error) {
	configInstance, err := config.LoadConfig()
	if err != nil {
		return nil, err
	}
	loggerInstance := logger.InitLogger()
	runtimeInstance := runtime.CreateRuntime()
	return &Bootstrap{
		ctx:     ctx,
		config:  configInstance,
		logger:  loggerInstance,
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
