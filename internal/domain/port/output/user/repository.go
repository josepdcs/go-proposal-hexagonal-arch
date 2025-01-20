package user

import (
	"context"

	"github.com/josepdcs/go-proposal-hexagonal-arch/internal/domain/model/entity"
)

type Repository interface {
	FindAll(ctx context.Context) ([]entity.User, error)
	FindByID(ctx context.Context, id uint) (entity.User, error)
	Create(ctx context.Context, user entity.User) (entity.User, error)
	Modify(ctx context.Context, user entity.User) (entity.User, error)
	Delete(ctx context.Context, user entity.User) error
}
