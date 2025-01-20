package user

import (
	"context"

	"github.com/josepdcs/go-proposal-hexagonal-arch/internal/domain/model/entity"
	"github.com/josepdcs/go-proposal-hexagonal-arch/internal/domain/port/output/user"
	"github.com/stretchr/testify/mock"
)

// MockRepository is a mock implementation of user.Repository by using testify mock.Mock
type MockRepository struct {
	mock.Mock
}

func NewMockRepository() *MockRepository {
	return &MockRepository{}
}

func (m *MockRepository) FindAll(ctx context.Context) ([]entity.User, error) {
	args := m.Called(ctx)
	return args.Get(0).([]entity.User), args.Error(1)
}

func (m *MockRepository) FindByID(ctx context.Context, id uint) (entity.User, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(entity.User), args.Error(1)
}

func (m *MockRepository) Create(ctx context.Context, user entity.User) (entity.User, error) {
	return m.save(ctx, user)
}

func (m *MockRepository) Modify(ctx context.Context, user entity.User) (entity.User, error) {
	return m.save(ctx, user)
}

func (m *MockRepository) save(ctx context.Context, user entity.User) (entity.User, error) {
	args := m.Called(ctx, user)
	return args.Get(0).(entity.User), args.Error(1)
}

func (m *MockRepository) Delete(ctx context.Context, user entity.User) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

// FakeRepository is a simple fake implementation of user.Repository
type FakeRepository struct {
	entities []entity.User
	err      error
}

func NewFakeUser(entities []entity.User, err error) user.Repository {
	return &FakeRepository{
		entities: entities,
		err:      err,
	}
}

func (m *FakeRepository) FindAll(ctx context.Context) ([]entity.User, error) {
	return m.entities, m.err
}

func (m *FakeRepository) FindByID(ctx context.Context, id uint) (entity.User, error) {
	var u entity.User
	for _, e := range m.entities {
		if e.ID == id {
			u = e
			break
		}
	}
	return u, m.err
}

func (m *FakeRepository) Create(ctx context.Context, user entity.User) (entity.User, error) {
	var u entity.User
	for _, e := range m.entities {
		if e.Name == user.Name && e.Surname == user.Surname {
			u = e
			break
		}
	}
	return u, m.err
}

func (m *FakeRepository) Modify(ctx context.Context, user entity.User) (entity.User, error) {
	var u entity.User
	for _, e := range m.entities {
		if e.ID == user.ID {
			u = e
			break
		}
	}
	return u, m.err
}

func (m *FakeRepository) Delete(ctx context.Context, user entity.User) error {
	return m.err
}
