package server

import (
	"fmt"

	"github.com/ruhulfbr/go-echo-ddd-boilerplate/config"
	"github.com/ruhulfbr/go-echo-ddd-boilerplate/internal/infrastructure/database"
	"gorm.io/gorm"
)

func InitDatabase(cfg *config.Config) (*gorm.DB, error) {
	db, err := database.NewGormDB(cfg.DB)
	if err != nil {
		return nil, fmt.Errorf("new db connection: %w", err)
	}
	return db, nil
}

func CloseDatabase(db *gorm.DB) {
	conn, _ := db.DB()
	_ = conn.Close()
}
