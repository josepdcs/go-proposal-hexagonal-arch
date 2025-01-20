package user

import (
	"cmp"
	"context"
	"slices"
	"testing"

	"github.com/josepdcs/go-proposal-hexagonal-arch/internal/domain/model/entity"
	"github.com/josepdcs/go-proposal-hexagonal-arch/internal/domain/model/errors"
	"github.com/stretchr/testify/assert"
)

func TestInMemoryRepository_FindAll(t *testing.T) {
	repo := NewInMemoryRepository()
	users, err := repo.FindAll(context.Background())
	assert.NoError(t, err)
	assert.Len(t, users, 3)

	slices.SortFunc(users, func(i, j entity.User) int {
		return cmp.Compare(i.ID, j.ID)
	})

	assert.Equal(t, "John", users[0].Name)
	assert.Equal(t, "Doe", users[0].Surname)
	assert.Equal(t, "Jane", users[1].Name)
	assert.Equal(t, "Doe", users[1].Surname)
	assert.Equal(t, "Alice", users[2].Name)
	assert.Equal(t, "Smith", users[2].Surname)
}

func TestInMemoryRepository_FindByID(t *testing.T) {
	repo := NewInMemoryRepository()
	user, err := repo.FindByID(context.Background(), 1)
	assert.NoError(t, err)
	assert.Equal(t, "John", user.Name)
	assert.Equal(t, "Doe", user.Surname)
}

func TestInMemoryRepository_FindByID_NotFound(t *testing.T) {
	repo := NewInMemoryRepository()
	_, err := repo.FindByID(context.Background(), 4)
	assert.ErrorIs(t, err, errors.ErrUserNotFound)
}

func TestInMemoryRepository_Create(t *testing.T) {
	repo := NewInMemoryRepository()
	user, err := repo.Create(context.Background(), entity.User{
		Name:    "Alice",
		Surname: "Smith",
	})
	assert.NoError(t, err)
	assert.Equal(t, "Alice", user.Name)
	assert.Equal(t, "Smith", user.Surname)
}

func TestInMemoryRepository_Modify(t *testing.T) {
	repo := NewInMemoryRepository()
	user, err := repo.Modify(context.Background(), entity.User{
		ID:      1,
		Name:    "Alice",
		Surname: "Smith",
	})
	assert.NoError(t, err)
	assert.Equal(t, "Alice", user.Name)
	assert.Equal(t, "Smith", user.Surname)
}

func TestInMemoryRepository_Modify_NotFound(t *testing.T) {
	repo := NewInMemoryRepository()
	_, err := repo.Modify(context.Background(), entity.User{
		ID:      4,
		Name:    "Alice",
		Surname: "Smith",
	})
	assert.ErrorIs(t, err, errors.ErrUserNotFound)
}

func TestInMemoryRepository_Delete(t *testing.T) {
	repo := NewInMemoryRepository()
	err := repo.Delete(context.Background(), entity.User{
		ID: 1,
	})
	assert.NoError(t, err)
}
