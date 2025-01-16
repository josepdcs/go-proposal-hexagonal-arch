package repository

import (
	"cmp"
	"context"
	"slices"
	"testing"

	"github.com/josepdcs/go-proposal-hexagonal-arch/internal/domain/entity"
	"github.com/stretchr/testify/assert"
)

func TestUserInMemory_FindAll(t *testing.T) {
	repo := NewUserInMemory()
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

func TestUserInMemory_FindByID(t *testing.T) {
	repo := NewUserInMemory()
	user, err := repo.FindByID(context.Background(), 1)
	assert.NoError(t, err)
	assert.Equal(t, "John", user.Name)
	assert.Equal(t, "Doe", user.Surname)
}

func TestUserInMemory_FindByID_NotFound(t *testing.T) {
	repo := NewUserInMemory()
	_, err := repo.FindByID(context.Background(), 4)
	assert.ErrorIs(t, err, ErrUserNotFound)
}

func TestUserInMemory_Create(t *testing.T) {
	repo := NewUserInMemory()
	user, err := repo.Create(context.Background(), entity.User{
		Name:    "Alice",
		Surname: "Smith",
	})
	assert.NoError(t, err)
	assert.Equal(t, "Alice", user.Name)
	assert.Equal(t, "Smith", user.Surname)
}

func TestUserInMemory_Modify(t *testing.T) {
	repo := NewUserInMemory()
	user, err := repo.Modify(context.Background(), entity.User{
		ID:      1,
		Name:    "Alice",
		Surname: "Smith",
	})
	assert.NoError(t, err)
	assert.Equal(t, "Alice", user.Name)
	assert.Equal(t, "Smith", user.Surname)
}

func TestUserInMemory_Modify_NotFound(t *testing.T) {
	repo := NewUserInMemory()
	_, err := repo.Modify(context.Background(), entity.User{
		ID:      4,
		Name:    "Alice",
		Surname: "Smith",
	})
	assert.ErrorIs(t, err, ErrUserNotFound)
}

func TestUserInMemory_Delete(t *testing.T) {
	repo := NewUserInMemory()
	err := repo.Delete(context.Background(), entity.User{
		ID: 1,
	})
	assert.NoError(t, err)
}
