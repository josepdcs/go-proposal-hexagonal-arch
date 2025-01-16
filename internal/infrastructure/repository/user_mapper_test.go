package repository

import (
	"testing"

	"github.com/josepdcs/go-proposal-hexagonal-arch/internal/domain/entity"
	"github.com/stretchr/testify/assert"
)

func TestUserInMemoryEntity_toEntityUser(t *testing.T) {
	userEntity := UserInMemoryEntity{
		ID:      1,
		Name:    "John",
		Surname: "Doe",
	}
	user := userEntity.toEntityUser()
	assert.Equal(t, userEntity.ID, user.ID)
	assert.Equal(t, userEntity.Name, user.Name)
	assert.Equal(t, userEntity.Surname, user.Surname)
}

func TestUserInMemoryEntity_fromEntityUser(t *testing.T) {
	user := entity.User{
		ID:      1,
		Name:    "John",
		Surname: "Doe",
	}
	userEntity := UserInMemoryEntity{}
	userEntity = userEntity.fromEntityUser(user)
	assert.Equal(t, user.ID, userEntity.ID)
	assert.Equal(t, user.Name, userEntity.Name)
	assert.Equal(t, user.Surname, userEntity.Surname)
}

func TestUserDBEntity_toEntityUser(t *testing.T) {
	userDBEntity := UserDBEntity{
		ID:      1,
		Name:    "John",
		Surname: "Doe",
	}
	user := userDBEntity.toEntityUser()
	assert.Equal(t, userDBEntity.ID, user.ID)
	assert.Equal(t, userDBEntity.Name, user.Name)
	assert.Equal(t, userDBEntity.Surname, user.Surname)
}

func TestUserDBEntity_fromEntityUser(t *testing.T) {
	user := entity.User{
		ID:      1,
		Name:    "John",
		Surname: "Doe",
	}
	userDBEntity := UserDBEntity{}
	userDBEntity = userDBEntity.fromEntityUser(user)
	assert.Equal(t, user.ID, userDBEntity.ID)
	assert.Equal(t, user.Name, userDBEntity.Name)
	assert.Equal(t, user.Surname, userDBEntity.Surname)
}
