package repository

import (
	"context"

	"github.com/thnkrn/go-gin-clean-arch/internal/domain/entity"
	"github.com/thnkrn/go-gin-clean-arch/internal/domain/repository"
	"gorm.io/gorm"
)

// DBUserEntity represents a user entity in the database
type DBUserEntity struct {
	ID      uint   `json:"id" gorm:"unique;not null"`
	Name    string `json:"name"`
	Surname string `json:"surname"`

	gorm.Model
}

// toEntityUser converts a DBUserEntity to an entity.User
func toEntityUser(u DBUserEntity) entity.User {
	return entity.User{
		ID:      u.ID,
		Name:    u.Name,
		Surname: u.Surname,
	}
}

// toDBUserEntity converts an entity.User to a DBUserEntity
func toDBUserEntity(u entity.User) DBUserEntity {
	return DBUserEntity{
		ID:      u.ID,
		Name:    u.Name,
		Surname: u.Surname,
	}
}

type UserRepository struct {
	DB *gorm.DB
}

// NewUserRepository creates a new instance of repository.UserRepository
func NewUserRepository(DB *gorm.DB) repository.UserRepository {
	return &UserRepository{DB}
}

// FindAll returns all users
func (r *UserRepository) FindAll(ctx context.Context) ([]entity.User, error) {
	var userEntities []DBUserEntity
	err := r.DB.Find(&userEntities).Error

	users := make([]entity.User, 0, len(userEntities))
	for _, e := range userEntities {
		users = append(users, toEntityUser(e))
	}

	return users, err
}

// FindByID returns a user by ID
func (r *UserRepository) FindByID(ctx context.Context, id uint) (entity.User, error) {
	var userEntity DBUserEntity
	err := r.DB.First(&userEntity, id).Error

	return toEntityUser(userEntity), err
}

// Create creates a user
func (r *UserRepository) Create(ctx context.Context, user entity.User) (entity.User, error) {
	return r.save(ctx, user)
}

// Modify modifies a user
func (r *UserRepository) Modify(ctx context.Context, user entity.User) (entity.User, error) {
	return r.save(ctx, user)
}

// save saves a user
func (r *UserRepository) save(ctx context.Context, user entity.User) (entity.User, error) {
	userEntity := toDBUserEntity(user)
	err := r.DB.Save(&userEntity).Error

	return toEntityUser(userEntity), err
}

// Delete deletes a user
func (r *UserRepository) Delete(ctx context.Context, user entity.User) error {
	userEntity := toDBUserEntity(user)
	err := r.DB.Delete(&userEntity).Error

	return err
}
