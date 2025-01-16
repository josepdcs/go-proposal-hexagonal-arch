package repository

import (
	"context"

	"github.com/josepdcs/go-proposal-hexagonal-arch/internal/domain/entity"
	"github.com/josepdcs/go-proposal-hexagonal-arch/internal/domain/repository"
	"gorm.io/gorm"
)

// UserDBEntity represents a user entity in the database
type UserDBEntity struct {
	ID      uint   `json:"id" gorm:"unique;not null"`
	Name    string `json:"name"`
	Surname string `json:"surname"`

	gorm.Model
}

// TableName overrides the table name used by UserDBEntity to `users`
func (UserDBEntity) TableName() string {
	return "users"
}

type UserDB struct {
	DB *gorm.DB
}

// NewUserDB creates a new instance of repository.UserDB
func NewUserDB(DB *gorm.DB) repository.User {
	return &UserDB{DB}
}

// FindAll returns all users
func (r *UserDB) FindAll(ctx context.Context) ([]entity.User, error) {
	var userEntities []UserDBEntity
	err := r.DB.Find(&userEntities).Error

	users := make([]entity.User, 0, len(userEntities))
	for _, e := range userEntities {
		users = append(users, e.toEntityUser())
	}

	return users, err
}

// FindByID returns a user by ID
func (r *UserDB) FindByID(ctx context.Context, id uint) (entity.User, error) {
	var userEntity UserDBEntity
	err := r.DB.First(&userEntity, id).Error

	return userEntity.toEntityUser(), err
}

// Create creates a user
func (r *UserDB) Create(ctx context.Context, user entity.User) (entity.User, error) {
	return r.save(ctx, user)
}

// Modify modifies a user
func (r *UserDB) Modify(ctx context.Context, user entity.User) (entity.User, error) {
	return r.save(ctx, user)
}

// save saves a user
func (r *UserDB) save(ctx context.Context, user entity.User) (entity.User, error) {
	userEntity := UserDBEntity{}
	userEntity = userEntity.fromEntityUser(user)
	err := r.DB.Save(&userEntity).Error
	if err != nil {
		return entity.User{}, err
	}

	return userEntity.toEntityUser(), nil
}

// Delete deletes a user
func (r *UserDB) Delete(ctx context.Context, user entity.User) error {
	userEntity := UserDBEntity{}
	userEntity = userEntity.fromEntityUser(user)

	return r.DB.Delete(&userEntity).Error
}
