package app

import (
	"go.uber.org/zap"
	"public-rpc/internal/adapters/storage"
	adminapi "public-rpc/internal/app/admin-api"
	publicapi "public-rpc/internal/app/public-api"
	"public-rpc/internal/app/worker"
	"public-rpc/internal/config"
)

type Application struct {
	Config  config.Config
	Logger  *zap.Logger
	Storage *storage.Storage
}

func (app *Application) RunAdminAPI() error {
	component, err := adminapi.NewAdminAPIComponent(app.Config.AdminAPIConfig, app.Logger, app.Storage)
	if err != nil {
		return err
	}
	return component.Run()
}

func (app *Application) RunPublicAPI() error {
	component, err := publicapi.NewPublicAPIComponent(app.Config.PublicAPIConfig, app.Logger, app.Storage)
	if err != nil {
		return err
	}
	return component.Run()
}

func (app *Application) RunWorker() error {
	component, err := worker.NewWorkerComponent(app.Config.WorkerConfig, app.Logger, app.Storage)
	if err != nil {
		return err
	}
	return component.Run()
}
