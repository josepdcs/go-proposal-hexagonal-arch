package repository

import (
	"context"

	"github.com/stretchr/testify/mock"
	"github.com/thnkrn/go-gin-clean-arch/internal/domain/entity"
	"github.com/thnkrn/go-gin-clean-arch/internal/domain/repository"
)

// UserMock is a mock implementation of repository.User by using testify mock.Mock
type UserMock struct {
	mock.Mock
}

func NewUserMock() *UserMock {
	return &UserMock{}
}

func (m *UserMock) FindAll(ctx context.Context) ([]entity.User, error) {
	args := m.Called(ctx)
	return args.Get(0).([]entity.User), args.Error(1)
}

func (m *UserMock) FindByID(ctx context.Context, id uint) (entity.User, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(entity.User), args.Error(1)
}

func (m *UserMock) Create(ctx context.Context, user entity.User) (entity.User, error) {
	return m.save(ctx, user)
}

func (m *UserMock) Modify(ctx context.Context, user entity.User) (entity.User, error) {
	return m.save(ctx, user)
}

func (m *UserMock) save(ctx context.Context, user entity.User) (entity.User, error) {
	args := m.Called(ctx, user)
	return args.Get(0).(entity.User), args.Error(1)
}

func (m *UserMock) Delete(ctx context.Context, user entity.User) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

// UserFake is a simple fake implementation of repository.User
type UserFake struct {
	entities []entity.User
	err      error
}

func NewMockUserInMemory(entities []entity.User, err error) repository.User {
	return &UserFake{
		entities: entities,
		err:      err,
	}
}

func (m *UserFake) FindAll(ctx context.Context) ([]entity.User, error) {
	return m.entities, m.err
}

func (m *UserFake) FindByID(ctx context.Context, id uint) (entity.User, error) {
	var u entity.User
	for _, e := range m.entities {
		if e.ID == id {
			u = e
			break
		}
	}
	return u, m.err
}

func (m *UserFake) Create(ctx context.Context, user entity.User) (entity.User, error) {
	var u entity.User
	for _, e := range m.entities {
		if e.Name == user.Name && e.Surname == user.Surname {
			u = e
			break
		}
	}
	return u, m.err
}

func (m *UserFake) Modify(ctx context.Context, user entity.User) (entity.User, error) {
	var u entity.User
	for _, e := range m.entities {
		if e.ID == user.ID {
			u = e
			break
		}
	}
	return u, m.err
}

func (m *UserFake) Delete(ctx context.Context, user entity.User) error {
	return m.err
}
