package admin_api

import (
	"go.uber.org/zap"
	"public-rpc/internal/adapters/storage"
	"public-rpc/internal/config"
)

type AdminAPIComponent struct {
	Cfg     config.AdminAPIConfig
	Logger  *zap.Logger
	Storage *storage.Storage
}

func (c *AdminAPIComponent) Run() error {
	return nil
}
