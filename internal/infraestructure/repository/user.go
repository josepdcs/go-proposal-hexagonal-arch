package repository

import (
	"context"

	"github.com/thnkrn/go-gin-clean-arch/internal/domain/entity"
	"github.com/thnkrn/go-gin-clean-arch/internal/domain/repository"
	"gorm.io/gorm"
)

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) repository.UserRepository {
	return &userRepository{DB}
}

func (r *userRepository) FindAll(ctx context.Context) ([]entity.User, error) {
	var users []entity.User
	err := r.DB.Find(&users).Error

	return users, err
}

func (r *userRepository) FindByID(ctx context.Context, id uint) (entity.User, error) {
	var user entity.User
	err := r.DB.First(&user, id).Error

	return user, err
}

func (r *userRepository) Create(ctx context.Context, user entity.User) (entity.User, error) {
	err := r.DB.Save(&user).Error

	return user, err
}

func (r *userRepository) Modify(ctx context.Context, user entity.User) (entity.User, error) {
	err := r.DB.Save(&user).Error

	return user, err
}

func (r *userRepository) Delete(ctx context.Context, user entity.User) error {
	err := r.DB.Delete(&user).Error

	return err
}
