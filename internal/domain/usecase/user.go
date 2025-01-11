package usecase

import (
	"context"

	"github.com/thnkrn/go-gin-clean-arch/internal/domain/entity"
)

// UserFindAll use case for finding all users
type UserFindAll interface {
	// FindAll returns all users
	FindAll(ctx context.Context) ([]entity.User, error)
}

// UserFindByID use case for finding a user by ID
type UserFindByID interface {
	// FindByID returns a user by ID
	FindByID(ctx context.Context, id uint) (entity.User, error)
}

// UserCreate use case for creating a user
type UserCreate interface {
	// Create creates a user
	Create(ctx context.Context, user entity.User) (entity.User, error)
}

// UserModify use case for modifying a user
type UserModify interface {
	// Modify modifies a user
	Modify(ctx context.Context, user entity.User) (entity.User, error)
}

// UserDelete use case for deleting a user
type UserDelete interface {
	// Delete deletes a user
	Delete(ctx context.Context, user entity.User) error
}
