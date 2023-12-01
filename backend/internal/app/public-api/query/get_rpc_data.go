package query

import (
	"context"
	"go.uber.org/zap"
	"public-rpc/internal/adapters/storage"
	"public-rpc/models"
)

type (
	GetRPCDataQuery struct {
		Network string
	}
	GetRPCDataHandler struct {
		Logger  *zap.Logger
		storage *storage.Storage
	}
)

func NewGetRPCDataHandler(logger *zap.Logger, storage *storage.Storage) GetRPCDataHandler {
	return GetRPCDataHandler{
		Logger:  logger,
		storage: storage,
	}
}

func (h *GetRPCDataHandler) Handle(ctx context.Context, query GetRPCDataQuery) ([]models.RPC, error) {
	storage_ := *h.storage

	if query.Network == "" {
		return storage_.ListRPCByNetwork(query.Network)
	}

	return storage_.ListRPC()
}
