package app

import (
	"go.uber.org/zap"
	"public-rpc/internal/adapters/storage"
	admin_api "public-rpc/internal/app/admin-api"
	public_api "public-rpc/internal/app/public-api"
	"public-rpc/internal/app/worker"
	"public-rpc/internal/config"
)

type Application struct {
	Config  config.Config
	Logger  *zap.Logger
	Storage *storage.Storage
}

func (app *Application) RunAdminAPI() error {
	component := admin_api.AdminAPIComponent{Cfg: app.Config.AdminAPIConfig, Logger: app.Logger, Storage: app.Storage}
	return component.Run()
}

func (app *Application) RunPublicAPI() error {
	component := public_api.PublicAPIComponent{Cfg: app.Config.PublicAPIConfig, Logger: app.Logger, Storage: app.Storage}
	return component.Run()
}

func (app *Application) RunWorker() error {
	component := worker.WorkerComponent{Cfg: app.Config.WorkerConfig, Logger: app.Logger, Storage: app.Storage}
	return component.Run()
}
