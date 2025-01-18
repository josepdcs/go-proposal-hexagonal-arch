package usecase

import (
	"context"

	"github.com/josepdcs/go-proposal-hexagonal-arch/internal/domain/entity"
	"github.com/stretchr/testify/mock"
)

type MockUserFinderAll struct {
	mock.Mock
}

func (m *MockUserFinderAll) Find(ctx context.Context) ([]entity.User, error) {
	args := m.Called(ctx)
	return args.Get(0).([]entity.User), args.Error(1)
}

type MockUserFinderByID struct {
	mock.Mock
}

func (m *MockUserFinderByID) Find(ctx context.Context, id uint) (entity.User, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(entity.User), args.Error(1)
}

type MockUserCreator struct {
	mock.Mock
}

func (m *MockUserCreator) Create(ctx context.Context, user entity.User) (entity.User, error) {
	args := m.Called(ctx, user)
	return args.Get(0).(entity.User), args.Error(1)
}

type MockUserModifier struct {
	mock.Mock
}

func (m *MockUserModifier) Modify(ctx context.Context, user entity.User) (entity.User, error) {
	args := m.Called(ctx, user)
	return args.Get(0).(entity.User), args.Error(1)
}

type MockUserDeleter struct {
	mock.Mock
}

func (m *MockUserDeleter) Delete(ctx context.Context, user entity.User) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}
