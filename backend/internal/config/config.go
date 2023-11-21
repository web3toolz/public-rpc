package config

import (
	"github.com/kelseyhightower/envconfig"
	"time"
)

type PublicAPIConfig struct {
	Enable bool   `envconfig:"PUBLIC_API_ENABLE" default:"true"`
	Host   string `envconfig:"PUBLIC_API_HOST" default:"0.0.0.0"`
	Port   string `envconfig:"PUBLIC_API_PORT" default:"8000"`
}

type AdminAPIConfig struct {
	Enable bool   `envconfig:"ADMIN_API_ENABLE" default:"true"`
	Host   string `envconfig:"ADMIN_API_HOST" default:"0.0.0.0"`
	Port   string `envconfig:"ADMIN_API_PORT" default:"8001"`
}

type WorkerConfig struct {
	Enable   bool          `envconfig:"WORKER_ENABLE" default:"true"`
	Interval time.Duration `envconfig:"WORKER_INTERVAL" default:"10s"`
}

type Config struct {
	LogLevel        string `envconfig:"LOG_LEVEL" default:"info"`
	PublicAPIConfig PublicAPIConfig
	AdminAPIConfig  AdminAPIConfig
	WorkerConfig    WorkerConfig
}

func LoadConfigFromEnv(prefix string) (*Config, error) {
	var cfg Config

	err := envconfig.Process(prefix, &cfg)

	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
