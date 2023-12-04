package query

import (
	"context"
	"go.uber.org/zap"
	"public-rpc/internal/adapters/cache"
	"public-rpc/internal/adapters/storage"
	"public-rpc/models"
)

type (
	GetRPCDataQuery struct {
		Chain   string
		Network string
	}
	GetRPCDataHandler struct {
		Logger  *zap.Logger
		storage *storage.Storage
		cache   cache.Cache
	}
)

func NewGetRPCDataHandler(logger *zap.Logger, storage *storage.Storage, cache cache.Cache) GetRPCDataHandler {
	return GetRPCDataHandler{
		Logger:  logger,
		storage: storage,
		cache:   cache,
	}
}

func (h *GetRPCDataHandler) getListWithFilters(chain string, network string) ([]models.RPC, error) {
	var data []models.RPC
	var err error
	var foundInCache bool
	storage_ := *h.storage

	data, foundInCache = h.cache.ListRPCWithFilters(chain, network)

	if !foundInCache {
		data, err = storage_.ListRPCWithFilters(chain, network)

		if err != nil {
			return nil, err
		}

		err = h.cache.SetListRPCWithFilters(chain, network, data)
	}

	return data, nil
}

func (h *GetRPCDataHandler) getList() ([]models.RPC, error) {
	var data []models.RPC
	var err error
	var foundInCache bool
	storage_ := *h.storage

	data, foundInCache = h.cache.ListRPC()

	if !foundInCache {
		data, err = storage_.ListRPC()

		if err != nil {
			return nil, err
		}

		err = h.cache.SetListRPC(data)
	}

	return data, nil
}

func (h *GetRPCDataHandler) Handle(ctx context.Context, query GetRPCDataQuery) ([]models.RPC, error) {
	if query.Chain != "" || query.Network != "" {
		return h.getListWithFilters(query.Chain, query.Network)
	}
	return h.getList()
}
