package components

import (
	"go.uber.org/zap"
	"public-rpc/internal/config"
)

type PublicAPIComponent struct {
	Cfg    config.PublicAPIConfig
	Logger *zap.Logger
}

func (c *PublicAPIComponent) Run() error {
	return nil
}
