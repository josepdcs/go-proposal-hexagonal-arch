package repository

import (
	"cmp"
	"context"
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thnkrn/go-gin-clean-arch/internal/domain/entity"
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
