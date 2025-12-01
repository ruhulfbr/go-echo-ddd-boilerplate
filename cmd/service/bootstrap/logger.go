package bootstrap

import (
	"fmt"

	"github.com/ruhulfbr/go-echo-ddd-boilerplate/config"
	"github.com/ruhulfbr/go-echo-ddd-boilerplate/internal/infrastructure/logger/slogx"
)

func InitLogger(cfg *config.Config) (func(), error) {
	if err := slogx.Init(cfg.Logger); err != nil {
		return nil, fmt.Errorf("init logger: %w", err)
	}
	return func() {}, nil
}
