package db

import (
	"fmt"

	"github.com/josepdcs/go-proposal-hexagonal-arch/internal/infrastructure/repository"
	"github.com/josepdcs/go-proposal-hexagonal-arch/internal/infrastructure/server/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase(cfg config.Config) (*gorm.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s user=%s dbname=%s port=%s password=%s", cfg.DB.Host, cfg.DB.User, cfg.DB.Name, cfg.DB.Port, cfg.DB.Password)
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
