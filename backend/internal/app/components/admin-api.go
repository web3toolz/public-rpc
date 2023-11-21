package components

import (
	"go.uber.org/zap"
	"public-rpc/internal/config"
)

type AdminAPIComponent struct {
	Cfg    config.AdminAPIConfig
	Logger *zap.Logger
}

func (c *AdminAPIComponent) Run() error {
	return nil
}
