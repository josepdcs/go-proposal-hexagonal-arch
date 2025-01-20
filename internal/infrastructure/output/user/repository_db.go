package user

import (
	"context"

	"github.com/josepdcs/go-proposal-hexagonal-arch/internal/domain/model/entity"
	"github.com/josepdcs/go-proposal-hexagonal-arch/internal/domain/port/output/user"
	"gorm.io/gorm"
)

// DBEntity represents a user entity in the database
type DBEntity struct {
	ID      uint   `json:"id" gorm:"unique;not null"`
	Name    string `json:"name"`
	Surname string `json:"surname"`

	gorm.Model
}

// TableName overrides the table name used by DBEntity to `users`
func (DBEntity) TableName() string {
	return "users"
}

type DBRepository struct {
	DB *gorm.DB
}

// NewDBRepository creates a new instance of user.Repository
func NewDBRepository(DB *gorm.DB) user.Repository {
	return &DBRepository{DB: DB}
}

// FindAll returns all users
func (r *DBRepository) FindAll(ctx context.Context) ([]entity.User, error) {
	var userEntities []DBEntity
	err := r.DB.Find(&userEntities).Error

	users := make([]entity.User, 0, len(userEntities))
	for _, e := range userEntities {
		users = append(users, e.toEntityUser())
	}

	return users, err
}

// FindByID returns a user by ID
func (r *DBRepository) FindByID(ctx context.Context, id uint) (entity.User, error) {
	var userEntity DBEntity
	err := r.DB.First(&userEntity, id).Error

	return userEntity.toEntityUser(), err
}

// Create creates a user
func (r *DBRepository) Create(ctx context.Context, user entity.User) (entity.User, error) {
	return r.save(ctx, user)
}

// Modify modifies a user
func (r *DBRepository) Modify(ctx context.Context, user entity.User) (entity.User, error) {
	return r.save(ctx, user)
}

// save saves a user
func (r *DBRepository) save(ctx context.Context, user entity.User) (entity.User, error) {
	userEntity := DBEntity{}
	userEntity = userEntity.fromEntityUser(user)
	err := r.DB.Save(&userEntity).Error
	if err != nil {
		return entity.User{}, err
	}

	return userEntity.toEntityUser(), nil
}

// Delete deletes a user
func (r *DBRepository) Delete(ctx context.Context, user entity.User) error {
	userEntity := DBEntity{}
	userEntity = userEntity.fromEntityUser(user)

	return r.DB.Delete(&userEntity).Error
}
