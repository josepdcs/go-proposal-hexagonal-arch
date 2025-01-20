package user

import (
	"context"

	"github.com/josepdcs/go-proposal-hexagonal-arch/internal/domain/model/entity"
)

// FinderAllUseCase defines the use case for finding all users
type FinderAllUseCase interface {
	// Find returns all users or an error if something goes wrong
	Find(ctx context.Context) ([]entity.User, error)
}

// FinderByIDUseCase defines the use case for finding a user by ID
type FinderByIDUseCase interface {
	// Find returns a user by ID or an error if something goes wrong
	Find(ctx context.Context, id uint) (entity.User, error)
}

// CreatorUseCase defines the use case for creating a user
type CreatorUseCase interface {
	// Create creates a user and returns the created user or an error if something goes wrong
	Create(ctx context.Context, user entity.User) (entity.User, error)
}

// ModifierUseCase defines the use case for modifying a user
type ModifierUseCase interface {
	// Modify modifies a user and returns the modified user or an error if something goes wrong
	Modify(ctx context.Context, user entity.User) (entity.User, error)
}

// DeleterUseCase defines the use case for deleting a user
type DeleterUseCase interface {
	// Delete deletes a user and returns an error if something goes wrong
	Delete(ctx context.Context, user entity.User) error
}
