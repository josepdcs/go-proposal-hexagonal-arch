package usecase

import (
	"context"

	"github.com/josepdcs/go-proposal-hexagonal-arch/internal/domain/entity"
)

// UserFinderAll defines the use case for finding all users
type UserFinderAll interface {
	// Find returns all users or an error if something goes wrong
	Find(ctx context.Context) ([]entity.User, error)
}

// UserFinderByID defines the use case for finding a user by ID
type UserFinderByID interface {
	// Find returns a user by ID or an error if something goes wrong
	Find(ctx context.Context, id uint) (entity.User, error)
}

// UserCreator defines the use case for creating a user
type UserCreator interface {
	// Create creates a user and returns the created user or an error if something goes wrong
	Create(ctx context.Context, user entity.User) (entity.User, error)
}

// UserModifier defines the use case for modifying a user
type UserModifier interface {
	// Modify modifies a user and returns the modified user or an error if something goes wrong
	Modify(ctx context.Context, user entity.User) (entity.User, error)
}

// UserDeleter defines the use case for deleting a user
type UserDeleter interface {
	// Delete deletes a user and returns an error if something goes wrong
	Delete(ctx context.Context, user entity.User) error
}
