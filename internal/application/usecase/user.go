package usecase

import (
	"context"

	"github.com/thnkrn/go-gin-clean-arch/internal/domain/entity"
	"github.com/thnkrn/go-gin-clean-arch/internal/domain/repository"
	"github.com/thnkrn/go-gin-clean-arch/internal/domain/usecase"
)

type userFindAll struct {
	userRepository repository.UserRepository
}

func NewUserFindAll(userRepository repository.UserRepository) usecase.UserFindAll {
	return &userFindAll{
		userRepository: userRepository,
	}
}

func (u *userFindAll) FindAll(ctx context.Context) ([]entity.User, error) {
	users, err := u.userRepository.FindAll(ctx)
	return users, err
}

type userFindByID struct {
	userRepository repository.UserRepository
}

func NewUserFindByID(userRepository repository.UserRepository) usecase.UserFindByID {
	return &userFindByID{
		userRepository: userRepository,
	}
}

func (u *userFindByID) FindByID(ctx context.Context, id uint) (entity.User, error) {
	user, err := u.userRepository.FindByID(ctx, id)
	return user, err
}

type userCreate struct {
	userRepository repository.UserRepository
}

func NewUserCreate(userRepository repository.UserRepository) usecase.UserCreate {
	return &userCreate{
		userRepository: userRepository,
	}
}

func (u *userCreate) Create(ctx context.Context, user entity.User) (entity.User, error) {
	user, err := u.userRepository.Create(ctx, user)
	return user, err
}

type userModify struct {
	userRepository repository.UserRepository
}

func NewUserModify(userRepository repository.UserRepository) usecase.UserModify {
	return &userModify{
		userRepository: userRepository,
	}
}

func (u *userModify) Modify(ctx context.Context, user entity.User) (entity.User, error) {
	user, err := u.userRepository.Modify(ctx, user)
	return user, err
}

type userDelete struct {
	userRepository repository.UserRepository
}

func NewUserDelete(userRepository repository.UserRepository) usecase.UserDelete {
	return &userDelete{
		userRepository: userRepository,
	}
}

func (u *userDelete) Delete(ctx context.Context, user entity.User) error {
	err := u.userRepository.Delete(ctx, user)
	return err
}
