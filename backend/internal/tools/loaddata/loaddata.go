package loaddata

import (
	"fmt"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"public-rpc/internal/adapters/storage"
	"public-rpc/internal/config"
	"public-rpc/internal/logger"
	"public-rpc/models"
	"time"
)

func loadFileData(filepath string) ([]models.RPC, error) {
	yamlFile, err := ioutil.ReadFile(filepath)

	if err != nil {
		return nil, err
	}

	var data []models.RPC

	err = yaml.Unmarshal(yamlFile, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func LoadData(cfg config.Config, filepath string) error {

	logger_, loggerCleanup := logger.New(cfg.LogLevel)
	defer loggerCleanup()

	logger_.Info("initializing storage")

	storagePointer, err := storage.InitializeStorage(cfg.StorageConfig)

	if err != nil {
		logger_.Error("failed to initialize storage", zap.Error(err))
		return err
	}

	storage_ := *storagePointer

	loadedDataFromFile, err := loadFileData(filepath)

	if err != nil {
		logger_.Error("failed to load data", zap.Error(err))
		return err
	}

	counter := 0
	httpOrWs := ""

	for _, rpc := range loadedDataFromFile {
		if rpc.HTTP != "" {
			httpOrWs = rpc.HTTP
		} else if rpc.WS != "" {
			httpOrWs = rpc.WS
		} else {
			httpOrWs = ""
		}

		if httpOrWs == "" {
			logger_.Warn("RPC has no endpoint, skipping", zap.Any("rpc", rpc))
			continue
		}

		rpc.Id = uuid.NewString()

		rpc.AddedAt = time.Now()
		rpc.Status = models.StatusActive

		rpcInDb, err := storage_.GetRPCByHttpOrWs(httpOrWs)

		if err != nil {
			logger_.Warn("failed to load rpc from db", zap.String("url", httpOrWs), zap.Any("rpc", rpc))
			continue
		}

		if rpcInDb != nil {
			logger_.Info("rpc already exists, skipping", zap.Any("rpc", rpc))
		} else {
			_, err = storage_.CreateRPC(rpc)

			if err != nil {
				logger_.Warn("failed to create rpc", zap.Error(err), zap.Any("rpc", rpc))
			} else {
				counter += 1
			}
		}

	}

	logger_.Info(fmt.Sprintf("added %d new rpcs in storage, finishing job", counter))

	_ = storage_.Close()

	return nil
}
