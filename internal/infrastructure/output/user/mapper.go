package user

import "github.com/josepdcs/go-proposal-hexagonal-arch/internal/domain/model/entity"

// toEntityUser converts a DBEntity to an entity.User
func (ub DBEntity) toEntityUser() entity.User {
	return entity.User{
		ID:      ub.ID,
		Name:    ub.Name,
		Surname: ub.Surname,
	}
}

// fromEntityUser converts an entity.User to a DBEntity
func (ub DBEntity) fromEntityUser(u entity.User) DBEntity {
	ub.ID = u.ID
	ub.Name = u.Name
	ub.Surname = u.Surname
	return ub
}

// toEntityUser converts an InMemoryEntity to an entity.User
func (um InMemoryEntity) toEntityUser() entity.User {
	return entity.User{
		ID:      um.ID,
		Name:    um.Name,
		Surname: um.Surname,
	}
}

// fromEntityUser converts an entity.User to a InMemoryEntity
func (um InMemoryEntity) fromEntityUser(u entity.User) InMemoryEntity {
	um.ID = u.ID
	um.Name = u.Name
	um.Surname = u.Surname
	return um
}
