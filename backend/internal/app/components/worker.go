package components

import (
	"go.uber.org/zap"
	"public-rpc/internal/config"
)

type WorkerComponent struct {
	Cfg    config.WorkerConfig
	Logger *zap.Logger
}

func (c *WorkerComponent) Run() error {
	return nil
}
