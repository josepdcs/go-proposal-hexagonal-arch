package usecase

import (
	"context"

	"github.com/thnkrn/go-gin-clean-arch/internal/domain/entity"
	"github.com/thnkrn/go-gin-clean-arch/internal/domain/repository"
	"github.com/thnkrn/go-gin-clean-arch/internal/domain/usecase"
)

// UserFindAll use case
type UserFindAll struct {
	userRepository repository.UserRepository
}

// NewUserFindAll creates a new usecase.UserFindAll instance
func NewUserFindAll(userRepository repository.UserRepository) usecase.UserFindAll {
	return &UserFindAll{
		userRepository: userRepository,
	}
}

// FindAll returns all users
func (u *UserFindAll) FindAll(ctx context.Context) ([]entity.User, error) {
	users, err := u.userRepository.FindAll(ctx)
	return users, err
}

// UserFindByID use case
type UserFindByID struct {
	userRepository repository.UserRepository
}

// NewUserFindByID creates a new usecase.UserFindByID instance
func NewUserFindByID(userRepository repository.UserRepository) usecase.UserFindByID {
	return &UserFindByID{
		userRepository: userRepository,
	}
}

// FindByID returns a user by ID
func (u *UserFindByID) FindByID(ctx context.Context, id uint) (entity.User, error) {
	user, err := u.userRepository.FindByID(ctx, id)
	return user, err
}

// UserCreate use case
type UserCreate struct {
	userRepository repository.UserRepository
}

// NewUserCreate creates a new usecase.UserCreate instance
func NewUserCreate(userRepository repository.UserRepository) usecase.UserCreate {
	return &UserCreate{
		userRepository: userRepository,
	}
}

// Create creates a new user
func (u *UserCreate) Create(ctx context.Context, user entity.User) (entity.User, error) {
	user, err := u.userRepository.Create(ctx, user)
	return user, err
}

// UserModify use case
type UserModify struct {
	userRepository repository.UserRepository
}

// NewUserModify creates a new usecase.UserModify instance
func NewUserModify(userRepository repository.UserRepository) usecase.UserModify {
	return &UserModify{
		userRepository: userRepository,
	}
}

// Modify modifies a user
func (u *UserModify) Modify(ctx context.Context, user entity.User) (entity.User, error) {
	user, err := u.userRepository.Modify(ctx, user)
	return user, err
}

// UserDelete use case
type UserDelete struct {
	userRepository repository.UserRepository
}

// NewUserDelete creates a new usecase.UserDelete instance
func NewUserDelete(userRepository repository.UserRepository) usecase.UserDelete {
	return &UserDelete{
		userRepository: userRepository,
	}
}

// Delete deletes a user
func (u *UserDelete) Delete(ctx context.Context, user entity.User) error {
	err := u.userRepository.Delete(ctx, user)
	return err
}
