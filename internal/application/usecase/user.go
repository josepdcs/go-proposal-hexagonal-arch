package usecase

import (
	"context"

	"github.com/thnkrn/go-gin-clean-arch/internal/domain/entity"
	"github.com/thnkrn/go-gin-clean-arch/internal/domain/repository"
	"github.com/thnkrn/go-gin-clean-arch/internal/domain/usecase"
)

// UserFinderAll use case
type UserFinderAll struct {
	user repository.User
}

// NewUserFinderAll creates a new usecase.UserFinderAll instance
func NewUserFinderAll(user repository.User) usecase.UserFinderAll {
	return &UserFinderAll{
		user: user,
	}
}

// Find returns all users or an error if something goes wrong
func (u *UserFinderAll) Find(ctx context.Context) ([]entity.User, error) {
	users, err := u.user.FindAll(ctx)
	return users, err
}

// UserFinderByID use case
type UserFinderByID struct {
	user repository.User
}

// NewUserFinderByID creates a new usecase.UserFinderByID instance
func NewUserFinderByID(user repository.User) usecase.UserFinderByID {
	return &UserFinderByID{
		user: user,
	}
}

// Find returns a user by ID or an error if something goes wrong
func (u *UserFinderByID) Find(ctx context.Context, id uint) (entity.User, error) {
	user, err := u.user.FindByID(ctx, id)
	return user, err
}

// UserCreator defines the use case for creating a user
type UserCreator struct {
	user repository.User
}

// NewUserCreator creates a new usecase.UserCreator instance
func NewUserCreator(user repository.User) usecase.UserCreator {
	return &UserCreator{
		user: user,
	}
}

// Create creates a user and returns the created user or an error if something goes wrong
func (u *UserCreator) Create(ctx context.Context, user entity.User) (entity.User, error) {
	user, err := u.user.Create(ctx, user)
	return user, err
}

// UserModifier defines the use case for modifying a user
type UserModifier struct {
	user repository.User
}

// NewUserModifier creates a new usecase.UserModifier instance
func NewUserModifier(user repository.User) usecase.UserModifier {
	return &UserModifier{
		user: user,
	}
}

// Modify modifies a user and returns the modified user or an error if something goes wrong
func (u *UserModifier) Modify(ctx context.Context, user entity.User) (entity.User, error) {
	return u.user.Modify(ctx, user)
}

// UserDeleter defines the use case for deleting a user
type UserDeleter struct {
	user repository.User
}

// NewUserDeleter creates a new usecase.UserDeleter instance
func NewUserDeleter(user repository.User) usecase.UserDeleter {
	return &UserDeleter{
		user: user,
	}
}

// Delete deletes a user and returns an error if something goes wrong
func (u *UserDeleter) Delete(ctx context.Context, user entity.User) error {
	err := u.user.Delete(ctx, user)
	return err
}
