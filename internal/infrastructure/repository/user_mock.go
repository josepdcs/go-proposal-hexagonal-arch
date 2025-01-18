package repository

import (
	"context"

	"github.com/josepdcs/go-proposal-hexagonal-arch/internal/domain/entity"
	"github.com/josepdcs/go-proposal-hexagonal-arch/internal/domain/repository"
	"github.com/stretchr/testify/mock"
)

// MockUser is a mock implementation of repository.User by using testify mock.Mock
type MockUser struct {
	mock.Mock
}

func NewMockUser() *MockUser {
	return &MockUser{}
}

func (m *MockUser) FindAll(ctx context.Context) ([]entity.User, error) {
	args := m.Called(ctx)
	return args.Get(0).([]entity.User), args.Error(1)
}

func (m *MockUser) FindByID(ctx context.Context, id uint) (entity.User, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(entity.User), args.Error(1)
}

func (m *MockUser) Create(ctx context.Context, user entity.User) (entity.User, error) {
	return m.save(ctx, user)
}

func (m *MockUser) Modify(ctx context.Context, user entity.User) (entity.User, error) {
	return m.save(ctx, user)
}

func (m *MockUser) save(ctx context.Context, user entity.User) (entity.User, error) {
	args := m.Called(ctx, user)
	return args.Get(0).(entity.User), args.Error(1)
}

func (m *MockUser) Delete(ctx context.Context, user entity.User) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

// FakeUser is a simple fake implementation of repository.User
type FakeUser struct {
	entities []entity.User
	err      error
}

func NewFakeUser(entities []entity.User, err error) repository.User {
	return &FakeUser{
		entities: entities,
		err:      err,
	}
}

func (m *FakeUser) FindAll(ctx context.Context) ([]entity.User, error) {
	return m.entities, m.err
}

func (m *FakeUser) FindByID(ctx context.Context, id uint) (entity.User, error) {
	var u entity.User
	for _, e := range m.entities {
		if e.ID == id {
			u = e
			break
		}
	}
	return u, m.err
}

func (m *FakeUser) Create(ctx context.Context, user entity.User) (entity.User, error) {
	var u entity.User
	for _, e := range m.entities {
		if e.Name == user.Name && e.Surname == user.Surname {
			u = e
			break
		}
	}
	return u, m.err
}

func (m *FakeUser) Modify(ctx context.Context, user entity.User) (entity.User, error) {
	var u entity.User
	for _, e := range m.entities {
		if e.ID == user.ID {
			u = e
			break
		}
	}
	return u, m.err
}

func (m *FakeUser) Delete(ctx context.Context, user entity.User) error {
	return m.err
}
