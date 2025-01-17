package di

import (
	"github.com/josepdcs/go-proposal-hexagonal-arch/internal/domain/repository"
	"github.com/josepdcs/go-proposal-hexagonal-arch/internal/infrastructure/db"
	infrarepo "github.com/josepdcs/go-proposal-hexagonal-arch/internal/infrastructure/repository"
	"github.com/josepdcs/go-proposal-hexagonal-arch/internal/infrastructure/server/config"
)

// ResolveUserRepository resolves the user repository based on the configuration
func ResolveUserRepository(cfg config.DB) (repository.User, error) {
	if cfg.Type != config.InMemoryDB {
		DB, err := db.ConnectDatabase(cfg)
		if err != nil {
			return nil, err
		}
		return infrarepo.NewUserDB(DB), nil
	}
	return infrarepo.NewUserInMemory(), nil
}
