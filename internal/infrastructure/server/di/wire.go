//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	handler "github.com/josepdcs/go-proposal-hexagonal-arch/internal/api/handler"
	"github.com/josepdcs/go-proposal-hexagonal-arch/internal/application/usecase"
	"github.com/josepdcs/go-proposal-hexagonal-arch/internal/infrastructure/repository"
	"github.com/josepdcs/go-proposal-hexagonal-arch/internal/infrastructure/server/config"
	"github.com/josepdcs/go-proposal-hexagonal-arch/internal/infrastructure/server/http"
)

func InitializeAPI(cfg config.Config) (*http.Server, error) {
	wire.Build(
		//db.ConnectDatabase,
		repository.NewUserInMemory,
		usecase.NewUserFinderAll,
		usecase.NewUserFinderByID,
		usecase.NewUserCreator,
		usecase.NewUserModifier,
		usecase.NewUserDeleter,
		handler.NewUserAPI,
		http.NewServer,
	)

	return &http.Server{}, nil
}
