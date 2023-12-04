package storage

import (
	"fmt"
	"public-rpc/internal/adapters/storage/mongodb"
	"public-rpc/internal/config"
	"public-rpc/models"
)

type Storage interface {
	ListRPC() ([]models.RPC, error)
	ListRPCWithFilters(chain string, network string) ([]models.RPC, error)
	GetRPCById(id string) (*models.RPC, error)
	GetRPCByHttpOrWs(httpOrWsUrl string) (*models.RPC, error)
	CreateRPC(rpc models.RPC) (*models.RPC, error)
	Close() error
}

func InitializeStorage(cfg config.StorageConfig) (*Storage, error) {
	var storage Storage

	if cfg.Type == "mongodb" {
		mongodbStorage, err := mongodb.Init(cfg)
		if err != nil {
			return nil, err
		}
		storage = mongodbStorage
	} else {
		return nil, fmt.Errorf("storage type %s is not supported", cfg.Type)
	}

	return &storage, nil
}
