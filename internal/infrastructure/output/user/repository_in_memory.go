package user

import (
	"context"
	"sync"

	"github.com/josepdcs/go-proposal-hexagonal-arch/internal/domain/model/entity"
	"github.com/josepdcs/go-proposal-hexagonal-arch/internal/domain/model/errors"
	"github.com/josepdcs/go-proposal-hexagonal-arch/internal/domain/port/output/user"
)

// InMemoryEntity represents a user entity in the in-memory database
type InMemoryEntity struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
}

// InMemoryRepository represents a user repository in the in-memory database
type InMemoryRepository struct {
	DB sync.Map
}

// NewInMemoryRepository creates a new instance of InMemoryRepository
func NewInMemoryRepository() user.Repository {
	u := &InMemoryRepository{}

	// Add some initial DBRepository
	u.DB.Store(uint(1), InMemoryEntity{ID: 1, Name: "John", Surname: "Doe"})
	u.DB.Store(uint(2), InMemoryEntity{ID: 2, Name: "Jane", Surname: "Doe"})
	u.DB.Store(uint(3), InMemoryEntity{ID: 3, Name: "Alice", Surname: "Smith"})

	return u
}

// FindAll returns all users
func (r *InMemoryRepository) FindAll(ctx context.Context) ([]entity.User, error) {
	var userEntities []InMemoryEntity
	// get all users from the in-memory database
	r.DB.Range(func(key, value interface{}) bool {
		userEntities = append(userEntities, value.(InMemoryEntity))
		return true
	})

	users := make([]entity.User, 0, len(userEntities))
	for _, e := range userEntities {
		users = append(users, e.toEntityUser())
	}

	return users, nil
}

// FindByID returns a user by ID
func (r *InMemoryRepository) FindByID(ctx context.Context, id uint) (entity.User, error) {
	var userEntity InMemoryEntity
	// get the user by ID from the in-memory database
	value, ok := r.DB.Load(id)
	if !ok {
		return entity.User{}, errors.ErrUserNotFound
	}

	userEntity = value.(InMemoryEntity)
	return userEntity.toEntityUser(), nil
}

// Create creates a user
func (r *InMemoryRepository) Create(ctx context.Context, user entity.User) (entity.User, error) {
	// get last ID from memory
	lastID := uint(0)
	r.DB.Range(func(key, value interface{}) bool {
		ID := key.(uint)
		if ID > lastID {
			lastID = ID
		}
		return true
	})
	user.ID = lastID + 1
	return r.save(ctx, user)
}

// Modify modifies a user
func (r *InMemoryRepository) Modify(ctx context.Context, user entity.User) (entity.User, error) {
	_, err := r.FindByID(ctx, user.ID)
	if err != nil {
		return entity.User{}, err
	}
	return r.save(ctx, user)
}

// save saves a user
func (r *InMemoryRepository) save(ctx context.Context, user entity.User) (entity.User, error) {
	userEntity := InMemoryEntity{}.fromEntityUser(user)
	r.DB.Store(user.ID, userEntity)

	return userEntity.toEntityUser(), nil
}

// Delete deletes a user
func (r *InMemoryRepository) Delete(ctx context.Context, user entity.User) error {
	r.DB.Delete(user.ID)

	return nil
}
