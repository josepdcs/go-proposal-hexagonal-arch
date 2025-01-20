package usecase

import (
	"context"

	"github.com/josepdcs/go-proposal-hexagonal-arch/internal/domain/entity"
	"github.com/stretchr/testify/mock"
)

type MockUserFinderAll struct {
	mock.Mock
}

func NewMockUserFinderAll() *MockUserFinderAll {
	return &MockUserFinderAll{}
}

func (m *MockUserFinderAll) Find(ctx context.Context) ([]entity.User, error) {
	args := m.Called(ctx)
	return args.Get(0).([]entity.User), args.Error(1)
}

type MockUserFinderByID struct {
	mock.Mock
}

func NewMockUserFinderByID() *MockUserFinderByID {
	return &MockUserFinderByID{}
}

func (m *MockUserFinderByID) Find(ctx context.Context, id uint) (entity.User, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(entity.User), args.Error(1)
}

type MockUserCreator struct {
	mock.Mock
}

func NewMockUserCreator() *MockUserCreator {
	return &MockUserCreator{}
}

func (m *MockUserCreator) Create(ctx context.Context, user entity.User) (entity.User, error) {
	args := m.Called(ctx, user)
	return args.Get(0).(entity.User), args.Error(1)
}

type MockUserModifier struct {
	mock.Mock
}

func NewMockUserModifier() *MockUserModifier {
	return &MockUserModifier{}
}

func (m *MockUserModifier) Modify(ctx context.Context, user entity.User) (entity.User, error) {
	args := m.Called(ctx, user)
	return args.Get(0).(entity.User), args.Error(1)
}

type MockUserDeleter struct {
	mock.Mock
}

func NewMockUserDeleter() *MockUserDeleter {
	return &MockUserDeleter{}
}

func (m *MockUserDeleter) Delete(ctx context.Context, user entity.User) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}
