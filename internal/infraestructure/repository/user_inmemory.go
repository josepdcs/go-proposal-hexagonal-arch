package repository

import (
	"context"
	"sync"

	"github.com/pkg/errors"
	"github.com/thnkrn/go-gin-clean-arch/internal/domain/entity"
	"github.com/thnkrn/go-gin-clean-arch/internal/domain/repository"
)

// ErrUserNotFound is the error returned when a user is not found
var ErrUserNotFound = errors.New("user not found")

// UserInMemoryEntity represents a user entity in the in-memory database
type UserInMemoryEntity struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
}

// UserInMemory represents a user repository in the in-memory database
type UserInMemory struct {
	data sync.Map
}

// NewUserInMemory creates a new instance of repository.UserInMemory
func NewUserInMemory() repository.User {
	u := &UserInMemory{}

	// Add some initial data
	u.data.Store(1, UserInMemoryEntity{ID: 1, Name: "John", Surname: "Doe"})
	u.data.Store(2, UserInMemoryEntity{ID: 2, Name: "Jane", Surname: "Doe"})
	u.data.Store(3, UserInMemoryEntity{ID: 3, Name: "Alice", Surname: "Smith"})

	return u
}

// FindAll returns all users
func (r *UserInMemory) FindAll(ctx context.Context) ([]entity.User, error) {
	var userEntities []UserInMemoryEntity
	// get all users from the in-memory database
	r.data.Range(func(key, value interface{}) bool {
		userEntities = append(userEntities, value.(UserInMemoryEntity))
		return true
	})

	users := make([]entity.User, 0, len(userEntities))
	for _, e := range userEntities {
		users = append(users, e.toEntityUser())
	}

	return users, nil
}

// FindByID returns a user by ID
func (r *UserInMemory) FindByID(ctx context.Context, id uint) (entity.User, error) {
	var userEntity UserInMemoryEntity
	// get the user by ID from the in-memory database
	value, ok := r.data.Load(id)
	if !ok {
		return entity.User{}, ErrUserNotFound
	}

	userEntity = value.(UserInMemoryEntity)
	return userEntity.toEntityUser(), nil
}

// Create creates a user
func (r *UserInMemory) Create(ctx context.Context, user entity.User) (entity.User, error) {
	return r.save(ctx, user)
}

// Modify modifies a user
func (r *UserInMemory) Modify(ctx context.Context, user entity.User) (entity.User, error) {
	return r.save(ctx, user)
}

// save saves a user
func (r *UserInMemory) save(ctx context.Context, user entity.User) (entity.User, error) {
	userEntity := UserInMemoryEntity{}
	r.data.Store(user.ID, userEntity.fromEntityUser(user))

	return userEntity.toEntityUser(), nil
}

// Delete deletes a user
func (r *UserInMemory) Delete(ctx context.Context, user entity.User) error {
	r.data.Delete(user.ID)

	return nil
}
