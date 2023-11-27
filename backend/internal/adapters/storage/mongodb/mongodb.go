package mongodb

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"public-rpc/internal/config"
	"public-rpc/models"
	"time"
)

type MongoDBStorage struct {
	Config config.StorageConfig
	client *mongo.Client
}

func (storage *MongoDBStorage) Close() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return storage.client.Disconnect(ctx)
}

func (storage *MongoDBStorage) database() *mongo.Database {
	return storage.client.Database(storage.Config.MongoDB.Database)
}

func (storage *MongoDBStorage) collection() *mongo.Collection {
	return storage.database().Collection(storage.Config.MongoDB.Collection)
}

func (storage *MongoDBStorage) ListRPC() ([]models.RPC, error) {
	return storage.ListRPCByNetwork("")
}

func (storage *MongoDBStorage) ListRPCByNetwork(network string) ([]models.RPC, error) {
	var data []models.RPC
	query := bson.D{}
	ctx := context.Background()

	coll := storage.collection()

	if network != "" {
		query = bson.D{{"network", network}}
	}

	cur, err := coll.Find(ctx, query)

	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	if err = cur.All(ctx, &data); err != nil {
		return nil, err
	}

	return data, nil
}

func (storage *MongoDBStorage) GetRPCById(id string) (*models.RPC, error) {
	var data models.RPC
	query := bson.D{{"_id", id}}
	ctx := context.Background()

	coll := storage.collection()

	err := coll.FindOne(ctx, query).Decode(&data)

	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return &data, nil
}

func (storage *MongoDBStorage) GetRPCByHttpOrWs(httpOrWsUrl string) (*models.RPC, error) {
	var data models.RPC
	query := bson.D{{"$or", bson.A{
		bson.D{{"http", httpOrWsUrl}},
		bson.D{{"ws", httpOrWsUrl}},
	}}}
	ctx := context.Background()

	coll := storage.collection()

	err := coll.FindOne(ctx, query).Decode(&data)

	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return &data, nil
}

func (storage *MongoDBStorage) CreateRPC(rpc models.RPC) (*models.RPC, error) {
	ctx := context.Background()

	coll := storage.collection()

	result, err := coll.InsertOne(ctx, rpc)

	if err != nil {
		return nil, err
	}

	rpc.Id = result.InsertedID.(string)

	return &rpc, nil
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
