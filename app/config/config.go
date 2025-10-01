package config

import (
	"github.com/caarlos0/env"
	"go.uber.org/fx"
)

type Config struct {
	RestServerPort string `env:"REST_SERVER_PORT,required"`
	PostgresURL    string `env:"POSTGRES_URL,required"`
}

func New() (*Config, error) {
	var cfg Config

	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

var Module = fx.Module("config",
	fx.Provide(New),
)
