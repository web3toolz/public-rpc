package memory

import (
	"fmt"
	"github.com/patrickmn/go-cache"
	"public-rpc/models"
	"time"
)

const (
	ttl             = 1 * time.Minute
	cleanupInterval = 10 * time.Minute
	listRPCKey      = "list.rpc"
)

type MemoryCache struct {
	cache *cache.Cache
}

func (m MemoryCache) ListRPC() ([]models.RPC, bool) {
	data, found := m.cache.Get(listRPCKey)
	if !found {
		return nil, false
	}
	return data.([]models.RPC), found
}

func (m MemoryCache) ListRPCWithFilters(chain string, network string) ([]models.RPC, bool) {
	key := fmt.Sprintf("%s.%s.%s", listRPCKey, chain, network)
	data, found := m.cache.Get(key)
	if !found {
		return nil, false
	}
	return data.([]models.RPC), found
}

func (m MemoryCache) SetListRPC(rpc []models.RPC) error {
	m.cache.Set(listRPCKey, rpc, cache.DefaultExpiration)
	return nil
}

func (m MemoryCache) SetListRPCWithFilters(chain string, network string, rpc []models.RPC) error {
	key := fmt.Sprintf("%s.%s.%s", listRPCKey, chain, network)
	m.cache.Set(key, rpc, cache.DefaultExpiration)
	return nil
}

func Init() (*MemoryCache, error) {
	return &MemoryCache{cache: cache.New(ttl, cleanupInterval)}, nil
}
