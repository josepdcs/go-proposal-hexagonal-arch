package repository

import "github.com/thnkrn/go-gin-clean-arch/internal/domain/entity"

// toEntityUser converts a UserDBEntity to an entity.User
func (ub UserDBEntity) toEntityUser() entity.User {
	return entity.User{
		ID:      ub.ID,
		Name:    ub.Name,
		Surname: ub.Surname,
	}
}

// fromEntityUser converts an entity.User to a UserDBEntity
func (ub UserDBEntity) fromEntityUser(u entity.User) UserDBEntity {
	ub.ID = u.ID
	ub.Name = u.Name
	ub.Surname = u.Surname
	return ub
}

// toEntityUser converts a UserInMemoryEntity to an entity.User
func (um UserInMemoryEntity) toEntityUser() entity.User {
	return entity.User{
		ID:      um.ID,
		Name:    um.Name,
		Surname: um.Surname,
	}
}

// fromEntityUser converts an entity.User to a UserInMemoryEntity
func (um UserInMemoryEntity) fromEntityUser(u entity.User) UserInMemoryEntity {
	um.ID = u.ID
	um.Name = u.Name
	um.Surname = u.Surname
	return um
}
