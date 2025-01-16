package di

import (
	"github.com/josepdcs/go-proposal-hexagonal-arch/internal/domain/repository"
	"github.com/josepdcs/go-proposal-hexagonal-arch/internal/infrastructure/db"
	infrarepo "github.com/josepdcs/go-proposal-hexagonal-arch/internal/infrastructure/repository"
	"github.com/josepdcs/go-proposal-hexagonal-arch/internal/infrastructure/server/config"
)

func ResolveUserRepository(cfg config.Config) (repository.User, error) {
	if cfg.DBType != config.InMemoryDB {
		DB, err := db.ConnectDatabase(cfg)
		if err != nil {
			return nil, err
		}
		return infrarepo.NewUserDB(DB), nil
	}
	return infrarepo.NewUserInMemory(), nil
}
