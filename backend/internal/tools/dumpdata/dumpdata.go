package dumpdata

import (
	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"public-rpc/internal/adapters/storage"
	"public-rpc/internal/config"
	"public-rpc/internal/logger"
	"public-rpc/models"
)

func saveToYaml(rpcs *[]models.RPC, filepath string) error {
	yamlData, err := yaml.Marshal(rpcs)

	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filepath, yamlData, 0644)

	if err != nil {
		return err
	}

	return nil
}

func DumpData(cfg config.Config, filepath string) error {
	logger_, loggerCleanup := logger.New(cfg.LogLevel)
	defer loggerCleanup()

	logger_.Info("initializing storage")

	storagePointer, err := storage.InitializeStorage(cfg.StorageConfig)

	if err != nil {
		logger_.Error("failed to initialize storage", zap.Error(err))
		return err
	}

	storage_ := *storagePointer

	rpcs, err := storage_.ListRPC()

	if err != nil {
		logger_.Error("failed to list RPCs", zap.Error(err))
		return err
	}

	err = saveToYaml(&rpcs, filepath)

	if err != nil {
		logger_.Error("failed to save to yaml", zap.Error(err))
		return err
	}

	logger_.Info("dumped data to yaml", zap.String("filepath", filepath))

	return nil
}
