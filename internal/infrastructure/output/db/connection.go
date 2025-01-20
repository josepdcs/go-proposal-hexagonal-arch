package db

import (
	"fmt"

	"github.com/josepdcs/go-proposal-hexagonal-arch/internal/infrastructure/input/server/config"
	"github.com/josepdcs/go-proposal-hexagonal-arch/internal/infrastructure/output/user"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase(cfg config.DB) (*gorm.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s user=%s port=%s password=%s sslmode=disable", cfg.Host, cfg.User, cfg.Port, cfg.Password)
	db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{
		SkipDefaultTransaction: true,
	})

	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&user.DBEntity{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
