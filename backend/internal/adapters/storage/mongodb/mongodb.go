package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"public-rpc/internal/config"
	"time"
)

type MongoDBStorage struct {
	Config config.StorageConfig
	client *mongo.Client
}

func (storage *MongoDBStorage) Close() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return storage.client.Disconnect(ctx)
}

func (storage *MongoDBStorage) database() *mongo.Database {
	return storage.client.Database(storage.Config.MongoDB.Database)
}

func (storage *MongoDBStorage) collection() *mongo.Collection {
	return storage.database().Collection(storage.Config.MongoDB.Collection)
}

func Init(cfg config.StorageConfig) (*MongoDBStorage, error) {
	if cfg.MongoDB.Uri == "" {
		return nil, fmt.Errorf("mongodb uri is required")
	} else if cfg.MongoDB.Database == "" {
		return nil, fmt.Errorf("mongodb database is required")
	} else if cfg.MongoDB.Collection == "" {
		return nil, fmt.Errorf("mongodb collection is required")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.MongoDB.Uri))

	if err != nil {
		return nil, err
	}

	storage := MongoDBStorage{
		Config: cfg,
		client: client,
	}

	err = client.Ping(ctx, nil)

	if err != nil {
		return nil, fmt.Errorf("failed to ping DB: %w", err)
	}

	return &storage, nil
}
