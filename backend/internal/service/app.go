package service

import (
	"context"
	"golang.org/x/sync/errgroup"
	"public-rpc/internal/app"
	"public-rpc/internal/config"
	"public-rpc/internal/logger"
)

func RunApplication(cfg config.Config) error {
	logger_, loggerCleanup := logger.New(cfg.LogLevel)
	defer loggerCleanup()

	logger_.Info("initializing application")

	app_ := app.Application{Config: cfg, Logger: logger_}
	componentsRunning := false

	ctx := context.Background()
	group, ctx := errgroup.WithContext(ctx)

	if cfg.AdminAPIConfig.Enable {
		componentsRunning = true
		group.Go(func() error {
			logger_.Info("running admin api")
			return app_.RunAdminAPI()
		})
	}

	if cfg.PublicAPIConfig.Enable {
		componentsRunning = true
		group.Go(func() error {
			logger_.Info("running public api")
			return app_.RunPublicAPI()
		})
	}

	if cfg.WorkerConfig.Enable {
		componentsRunning = true
		group.Go(func() error {
			logger_.Info("running worker")
			return app_.RunWorker()
		})
	}

	if !componentsRunning {
		logger_.Warn("no components are running, exiting application")
	}

	return group.Wait()
}
