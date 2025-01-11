package repository

import (
	"context"

	"github.com/thnkrn/go-gin-clean-arch/internal/domain/entity"
)

type User interface {
	FindAll(ctx context.Context) ([]entity.User, error)
	FindByID(ctx context.Context, id uint) (entity.User, error)
	Create(ctx context.Context, user entity.User) (entity.User, error)
	Modify(ctx context.Context, user entity.User) (entity.User, error)
	Delete(ctx context.Context, user entity.User) error
}
