//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/josepdcs/go-proposal-hexagonal-arch/internal/api/handler"
	"github.com/josepdcs/go-proposal-hexagonal-arch/internal/application/user"
	"github.com/josepdcs/go-proposal-hexagonal-arch/internal/infrastructure/input/server/config"
	"github.com/josepdcs/go-proposal-hexagonal-arch/internal/infrastructure/input/server/http"
)

func InitializeAPI(cfg config.DB) (*http.Server, error) {
	wire.Build(
		ResolveUserRepository,
		user.NewDefaultFinderAllUseCase,
		user.NewDefaultFinderByIDUseCase,
		user.NewDefaultCreatorUseCase,
		user.NewDefaultModifierUseCase,
		user.NewDefaultDeleterUseCase,
		handler.NewUserAPI,
		http.NewServer,
	)

	return &http.Server{}, nil
}
