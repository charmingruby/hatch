// Package config is used for environment variables parsing and validation.
package config

import "github.com/caarlos0/env"

// Config holds the structure of expected environment variables.
//
// Each field should include an `env` struct tag that defines the environment variable name.
//
// To make a field required, add the ",required" option.
type Config struct {
	RestServerPort string `env:"REST_SERVER_PORT,required"`
	PostgresURL    string `env:"POSTGRES_URL,required"`
	MQTTURL        string `env:"MQTT_URL,required"`
	LogLevel       string `env:"LOG_LEVEL"`
}

// New parses environment variables into a Config struct.
//
// Returns:
//   - *Config: the parsed environment variables.
//   - error: if there is an error on validating environment variables.
func New() (*Config, error) {
	var cfg Config

	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
