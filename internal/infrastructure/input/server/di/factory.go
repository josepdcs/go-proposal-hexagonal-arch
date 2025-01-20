package di

import (
	"github.com/josepdcs/go-proposal-hexagonal-arch/internal/domain/port/output/user"
	"github.com/josepdcs/go-proposal-hexagonal-arch/internal/infrastructure/input/server/config"
	"github.com/josepdcs/go-proposal-hexagonal-arch/internal/infrastructure/output/db"
	user2 "github.com/josepdcs/go-proposal-hexagonal-arch/internal/infrastructure/output/user"
)

// ResolveUserRepository resolves the user repository based on the configuration
func ResolveUserRepository(cfg config.DB) (user.Repository, error) {
	if cfg.Type != config.InMemoryDB {
		DB, err := db.ConnectDatabase(cfg)
		if err != nil {
			return nil, err
		}
		return user2.NewDBRepository(DB), nil
	}
	return user2.NewInMemoryRepository(), nil
}
