// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/thnkrn/go-gin-clean-arch/internal/api/handler"
	"github.com/thnkrn/go-gin-clean-arch/internal/api/server/config"
	"github.com/thnkrn/go-gin-clean-arch/internal/api/server/http"
	"github.com/thnkrn/go-gin-clean-arch/internal/application/usecase"
	"github.com/thnkrn/go-gin-clean-arch/internal/infraestructure/db"
	"github.com/thnkrn/go-gin-clean-arch/internal/infraestructure/repository"
)

// Injectors from wire.go:

func InitializeAPI(cfg config.Config) (*http.Server, error) {
	gormDB, err := db.ConnectDatabase(cfg)
	if err != nil {
		return nil, err
	}
	user := repository.NewUser(gormDB)
	userFindAll := usecase.NewUserFindAll(user)
	userFindByID := usecase.NewUserFindByID(user)
	userCreate := usecase.NewUserCreate(user)
	userModify := usecase.NewUserModify(user)
	userDelete := usecase.NewUserDelete(user)
	userAPI := handler.NewUserAPI(userFindAll, userFindByID, userCreate, userModify, userDelete)
	server := http.NewServer(userAPI)
	return server, nil
}
