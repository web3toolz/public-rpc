package cache

import "public-rpc/models"

type Cache interface {
	ListRPC() ([]models.RPC, bool)
	ListRPCWithFilters(chain string, network string) ([]models.RPC, bool)
	SetListRPC(rpc []models.RPC) error
	SetListRPCWithFilters(chain string, network string, rpc []models.RPC) error
}
