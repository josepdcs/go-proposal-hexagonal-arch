package repository

import (
	"context"

	"github.com/thnkrn/go-gin-clean-arch/internal/domain/entity"
	"github.com/thnkrn/go-gin-clean-arch/internal/domain/repository"
	"gorm.io/gorm"
)

// UserDBEntity represents a user entity in the database
type UserDBEntity struct {
	ID      uint   `json:"id" gorm:"unique;not null"`
	Name    string `json:"name"`
	Surname string `json:"surname"`

	gorm.Model
}

// toEntityUser converts a UserDBEntity to an entity.User
func toEntityUser(u UserDBEntity) entity.User {
	return entity.User{
		ID:      u.ID,
		Name:    u.Name,
		Surname: u.Surname,
	}
}

// toUserDBEntity converts an entity.User to a UserDBEntity
func toUserDBEntity(u entity.User) UserDBEntity {
	return UserDBEntity{
		ID:      u.ID,
		Name:    u.Name,
		Surname: u.Surname,
	}
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
		users = append(users, toEntityUser(e))
	}

	return users, err
}

// FindByID returns a user by ID
func (r *UserDB) FindByID(ctx context.Context, id uint) (entity.User, error) {
	var userEntity UserDBEntity
	err := r.DB.First(&userEntity, id).Error

	return toEntityUser(userEntity), err
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
	userEntity := toUserDBEntity(user)
	err := r.DB.Save(&userEntity).Error

	return toEntityUser(userEntity), err
}

// Delete deletes a user
func (r *UserDB) Delete(ctx context.Context, user entity.User) error {
	userEntity := toUserDBEntity(user)
	err := r.DB.Delete(&userEntity).Error

	return err
}
