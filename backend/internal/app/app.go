package app

import (
	"go.uber.org/zap"
	"public-rpc/internal/adapters/storage"
	"public-rpc/internal/app/components"
	"public-rpc/internal/config"
)

type Application struct {
	Config  config.Config
	Logger  *zap.Logger
	Storage *storage.Storage
}

func (app *Application) RunAdminAPI() error {
	component := components.AdminAPIComponent{Cfg: app.Config.AdminAPIConfig, Logger: app.Logger}
	return component.Run()
}

func (app *Application) RunPublicAPI() error {
	component := components.PublicAPIComponent{Cfg: app.Config.PublicAPIConfig, Logger: app.Logger}
	return component.Run()
}

func (app *Application) RunWorker() error {
	component := components.WorkerComponent{Cfg: app.Config.WorkerConfig, Logger: app.Logger}
	return component.Run()
}
