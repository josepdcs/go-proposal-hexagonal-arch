//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	handler "github.com/thnkrn/go-gin-clean-arch/internal/api/handler"
	"github.com/thnkrn/go-gin-clean-arch/internal/api/server/config"
	"github.com/thnkrn/go-gin-clean-arch/internal/api/server/http"
	"github.com/thnkrn/go-gin-clean-arch/internal/application/usecase"
	"github.com/thnkrn/go-gin-clean-arch/internal/infraestructure/db"
	"github.com/thnkrn/go-gin-clean-arch/internal/infraestructure/repository"
)

func InitializeAPI(cfg config.Config) (*http.ServerHTTP, error) {
	wire.Build(
		db.ConnectDatabase,
		repository.NewUserRepository,
		usecase.NewUserFindAllUseCase,
		usecase.NewUserUseCase,
		handler.NewUserHandler,
		http.NewServerHTTP,
	)

	return &http.ServerHTTP{}, nil
}
