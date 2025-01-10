package usecase

import (
	"context"

	"github.com/thnkrn/go-gin-clean-arch/internal/domain/entity"
	"github.com/thnkrn/go-gin-clean-arch/internal/domain/repository"
	services "github.com/thnkrn/go-gin-clean-arch/internal/domain/usecase"
)

type userFindAllUseCase struct {
	userRepository repository.UserRepository
}

func NewUserFindAllUseCase(userRepository repository.UserRepository) services.UserFindAllUseCase {
	return &userFindAllUseCase{
		userRepository: userRepository,
	}
}

func (u *userFindAllUseCase) FindAll(ctx context.Context) ([]entity.User, error) {
	users, err := u.userRepository.FindAll(ctx)
	return users, err
}

type userUseCase struct {
	userRepository repository.UserRepository
}

func NewUserUseCase(userRepository repository.UserRepository) services.UserUseCase {
	return &userUseCase{
		userRepository: userRepository,
	}
}

func (c *userUseCase) FindAll(ctx context.Context) ([]entity.User, error) {
	users, err := c.userRepository.FindAll(ctx)
	return users, err
}

func (c *userUseCase) FindByID(ctx context.Context, id uint) (entity.User, error) {
	user, err := c.userRepository.FindByID(ctx, id)
	return user, err
}

func (c *userUseCase) Create(ctx context.Context, user entity.User) (entity.User, error) {
	user, err := c.userRepository.Create(ctx, user)

	return user, err
}

func (c *userUseCase) Modify(ctx context.Context, user entity.User) (entity.User, error) {
	user, err := c.userRepository.Modify(ctx, user)

	return user, err
}

func (c *userUseCase) Delete(ctx context.Context, user entity.User) error {
	err := c.userRepository.Delete(ctx, user)

	return err
}
