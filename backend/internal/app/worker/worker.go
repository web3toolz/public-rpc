package worker

import (
	"go.uber.org/zap"
	"public-rpc/internal/adapters/storage"
	"public-rpc/internal/config"
)

type WorkerComponent struct {
	Cfg     config.WorkerConfig
	Logger  *zap.Logger
	Storage *storage.Storage
}

func (c *WorkerComponent) Run() error {
	return nil
}
