package server

import (
	"fmt"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
	"github.com/ruhulfbr/go-echo-ddd-boilerplate/config"
	"github.com/ruhulfbr/go-echo-ddd-boilerplate/internal/infrastructure/logger/slogx"
)

func Run() error {
	cfg, err := loadEnv()
	if err != nil {
		return err
	}

	loggerCleanup, err := initLogger(cfg)
	if err != nil {
		return err
	}
	defer loggerCleanup()

	gormDB, err := InitDatabase(cfg)
	if err != nil {
		return err
	}
	defer CloseDatabase(gormDB)

	app, err := InitServer(cfg, gormDB)
	if err != nil {
		return err
	}

	return StartServer(app, cfg)
}

func loadEnv() (*config.Config, error) {
	_ = godotenv.Load()

	var cfg config.Config
	if err := env.Parse(&cfg); err != nil {
		return nil, fmt.Errorf("parse env: %w", err)
	}

	return &cfg, nil
}

func initLogger(cfg *config.Config) (func(), error) {
	if err := slogx.Init(cfg.Logger); err != nil {
		return nil, fmt.Errorf("init logger: %w", err)
	}
	return func() {}, nil
}
