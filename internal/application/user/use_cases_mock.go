package user

import (
	"context"

	"github.com/josepdcs/go-proposal-hexagonal-arch/internal/domain/model/entity"
	"github.com/stretchr/testify/mock"
)

type MockFinderAllUseCase struct {
	mock.Mock
}

func NewMockFinderAllUseCase() *MockFinderAllUseCase {
	return &MockFinderAllUseCase{}
}

func (m *MockFinderAllUseCase) Find(ctx context.Context) ([]entity.User, error) {
	args := m.Called(ctx)
	return args.Get(0).([]entity.User), args.Error(1)
}

type MockUserFinderByIDUseCase struct {
	mock.Mock
}

func NewMockUserFinderByIDUseCase() *MockUserFinderByIDUseCase {
	return &MockUserFinderByIDUseCase{}
}

func (m *MockUserFinderByIDUseCase) Find(ctx context.Context, id uint) (entity.User, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(entity.User), args.Error(1)
}

type MockUserCreatorUseCase struct {
	mock.Mock
}

func NewMockUserCreatorUseCase() *MockUserCreatorUseCase {
	return &MockUserCreatorUseCase{}
}

func (m *MockUserCreatorUseCase) Create(ctx context.Context, user entity.User) (entity.User, error) {
	args := m.Called(ctx, user)
	return args.Get(0).(entity.User), args.Error(1)
}

type MockUserModifierUseCase struct {
	mock.Mock
}

func NewMockUserModifierUseCase() *MockUserModifierUseCase {
	return &MockUserModifierUseCase{}
}

func (m *MockUserModifierUseCase) Modify(ctx context.Context, user entity.User) (entity.User, error) {
	args := m.Called(ctx, user)
	return args.Get(0).(entity.User), args.Error(1)
}

type MockUserDeleterUseCase struct {
	mock.Mock
}

func NewMockUserDeleterUseCase() *MockUserDeleterUseCase {
	return &MockUserDeleterUseCase{}
}

func (m *MockUserDeleterUseCase) Delete(ctx context.Context, user entity.User) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}
