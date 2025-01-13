//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	handler "github.com/thnkrn/go-gin-clean-arch/internal/api/handler"
	"github.com/thnkrn/go-gin-clean-arch/internal/application/usecase"
	"github.com/thnkrn/go-gin-clean-arch/internal/infrastructure/repository"
	"github.com/thnkrn/go-gin-clean-arch/internal/infrastructure/server/config"
	"github.com/thnkrn/go-gin-clean-arch/internal/infrastructure/server/http"
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
