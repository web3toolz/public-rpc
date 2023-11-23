package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"time"
)

type PublicAPIConfig struct {
	Enable  bool          `envconfig:"PUBLIC_API_ENABLE" default:"true"`
	Host    string        `envconfig:"PUBLIC_API_HOST" default:"0.0.0.0"`
	Port    string        `envconfig:"PUBLIC_API_PORT" default:"8000"`
	Timeout time.Duration `envconfig:"PUBLIC_API_TIMEOUT" default:"60s"`
	Path    string        `envconfig:"PUBLIC_API_PATH" default:"/"`
}

type AdminAPIConfig struct {
	Enable bool   `envconfig:"ADMIN_API_ENABLE" default:"true"`
	Host   string `envconfig:"ADMIN_API_HOST" default:"0.0.0.0"`
	Port   string `envconfig:"ADMIN_API_PORT" default:"8001"`
	Path   string `envconfig:"ADMIN_API_PATH" default:"/"`
}

type WorkerConfig struct {
	Enable   bool          `envconfig:"WORKER_ENABLE" default:"true"`
	Interval time.Duration `envconfig:"WORKER_INTERVAL" default:"10s"`
}

type MongoDBStorageConfig struct {
	Uri        string `envconfig:"STORAGE_MONGODB_URI" default:"mongodb://localhost:27017"`
	Database   string `envconfig:"STORAGE_MONGODB_DATABASE" default:"main"`
	Collection string `envconfig:"STORAGE_MONGODB_COLLECTION" default:"main"`
}

type StorageConfig struct {
	Type string `envconfig:"STORAGE_TYPE" default:"mongodb"`

	MongoDB MongoDBStorageConfig `envconfig:""`
}

type Config struct {
	LogLevel        string          `envconfig:"LOG_LEVEL" default:"info"`
	PublicAPIConfig PublicAPIConfig `envconfig:""`
	AdminAPIConfig  AdminAPIConfig  `envconfig:""`
	WorkerConfig    WorkerConfig    `envconfig:""`
	StorageConfig   StorageConfig   `envconfig:""`
}

func LoadConfigFromEnv(prefix string) (*Config, error) {
	var cfg Config

	_ = godotenv.Load(".env")

	err := envconfig.Process(prefix, &cfg)

	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
