package bootstrap

import (
	"fmt"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
	"github.com/ruhulfbr/go-echo-ddd-boilerplate/config"
)

func LoadEnv() (*config.Config, error) {
	_ = godotenv.Load()

	var cfg config.Config
	if err := env.Parse(&cfg); err != nil {
		return nil, fmt.Errorf("parse env: %w", err)
	}
	return &cfg, nil
}
