package usecase

import (
	"context"

	"github.com/thnkrn/go-gin-clean-arch/internal/domain/entity"
	"github.com/thnkrn/go-gin-clean-arch/internal/domain/repository"
	"github.com/thnkrn/go-gin-clean-arch/internal/domain/usecase"
)

// UserFindAll use case
type UserFindAll struct {
	user repository.User
}

// NewUserFindAll creates a new usecase.UserFindAll instance
func NewUserFindAll(user repository.User) usecase.UserFindAll {
	return &UserFindAll{
		user: user,
	}
}

// FindAll returns all users
func (u *UserFindAll) FindAll(ctx context.Context) ([]entity.User, error) {
	users, err := u.user.FindAll(ctx)
	return users, err
}

// UserFindByID use case
type UserFindByID struct {
	user repository.User
}

// NewUserFindByID creates a new usecase.UserFindByID instance
func NewUserFindByID(user repository.User) usecase.UserFindByID {
	return &UserFindByID{
		user: user,
	}
}

// FindByID returns a user by ID
func (u *UserFindByID) FindByID(ctx context.Context, id uint) (entity.User, error) {
	user, err := u.user.FindByID(ctx, id)
	return user, err
}

// UserCreate use case
type UserCreate struct {
	user repository.User
}

// NewUserCreate creates a new usecase.UserCreate instance
func NewUserCreate(user repository.User) usecase.UserCreate {
	return &UserCreate{
		user: user,
	}
}

// Create creates a new user
func (u *UserCreate) Create(ctx context.Context, user entity.User) (entity.User, error) {
	user, err := u.user.Create(ctx, user)
	return user, err
}

// UserModify use case
type UserModify struct {
	user repository.User
}

// NewUserModify creates a new usecase.UserModify instance
func NewUserModify(user repository.User) usecase.UserModify {
	return &UserModify{
		user: user,
	}
}

// Modify modifies a user
func (u *UserModify) Modify(ctx context.Context, user entity.User) (entity.User, error) {
	user, err := u.user.Modify(ctx, user)
	return user, err
}

// UserDelete use case
type UserDelete struct {
	user repository.User
}

// NewUserDelete creates a new usecase.UserDelete instance
func NewUserDelete(user repository.User) usecase.UserDelete {
	return &UserDelete{
		user: user,
	}
}

// Delete deletes a user
func (u *UserDelete) Delete(ctx context.Context, user entity.User) error {
	err := u.user.Delete(ctx, user)
	return err
}
