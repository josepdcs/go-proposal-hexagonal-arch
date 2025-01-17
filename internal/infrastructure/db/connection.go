package db

import (
	"fmt"

	"github.com/josepdcs/go-proposal-hexagonal-arch/internal/infrastructure/repository"
	"github.com/josepdcs/go-proposal-hexagonal-arch/internal/infrastructure/server/config"
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

	err = db.AutoMigrate(&repository.UserDBEntity{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
