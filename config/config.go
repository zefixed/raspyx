package config

import (
	"fmt"
	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type (
	Config struct {
		App  App
		Log  Log
		HTTP HTTP
	}
	App struct {
		Name    string `env:"APP_NAME,required"`
		Version string `env:"APP_VERSION,required"`
	}
	Log struct {
		Level string `env:"LOG_LEVEL,required"`
		Type  string `env:"LOG_TYPE,required"`
	}

	HTTP struct {
		Port string `env:"HTTP_PORT,required"`
	}
)

func NewConfig() (*Config, error) {
	cfg := &Config{}
	err := godotenv.Load("local.env")
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	if err := env.Parse(cfg); err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	return cfg, nil
}
