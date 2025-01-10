package usecase

import (
	"context"

	"github.com/thnkrn/go-gin-clean-arch/internal/domain/entity"
)

type UserUseCase interface {
	UserFindAllUseCase
	UserFindByIDUseCase
	UserCreateUseCase
	UserModifyUseCase
	UserDeleteUseCase
}

type UserFindAllUseCase interface {
	FindAll(ctx context.Context) ([]entity.User, error)
}

type UserFindByIDUseCase interface {
	FindByID(ctx context.Context, id uint) (entity.User, error)
}

type UserCreateUseCase interface {
	Create(ctx context.Context, user entity.User) (entity.User, error)
}

type UserModifyUseCase interface {
	Modify(ctx context.Context, user entity.User) (entity.User, error)
}

type UserDeleteUseCase interface {
	Delete(ctx context.Context, user entity.User) error
}
