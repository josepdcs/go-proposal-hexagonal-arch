package user

import (
	"cmp"
	"context"
	"slices"
	"testing"

	"github.com/josepdcs/go-proposal-hexagonal-arch/internal/domain/model/entity"
	"github.com/josepdcs/go-proposal-hexagonal-arch/internal/infrastructure/output/user"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestDefaultFinderAllUseCase_Find(t *testing.T) {
	tests := []struct {
		name  string
		given func() *user.MockRepository
		when  func(repository *user.MockRepository) ([]entity.User, error)
		then  func([]entity.User, error)
	}{
		{
			name: "should find all users",
			given: func() *user.MockRepository {
				repository := user.NewMockRepository()
				users := []entity.User{
					{ID: 1, Name: "John", Surname: "Doe"},
					{ID: 2, Name: "Jane", Surname: "Doe"},
					{ID: 3, Name: "Alice", Surname: "Smith"},
				}
				repository.On("FindAll", context.Background()).Return(users, nil)
				return repository
			},
			when: func(repository *user.MockRepository) ([]entity.User, error) {
				return NewDefaultFinderAllUseCase(repository).Find(context.Background())
			},
			then: func(users []entity.User, err error) {
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
			},
		},
		{
			name: "should not find users",
			given: func() *user.MockRepository {
				repository := user.NewMockRepository()
				var users []entity.User
				repository.On("FindAll", context.Background()).Return(users, errors.New("not found"))
				return repository
			},
			when: func(repository *user.MockRepository) ([]entity.User, error) {
				return NewDefaultFinderAllUseCase(repository).Find(context.Background())
			},
			then: func(users []entity.User, err error) {
				assert.Error(t, err)
				assert.Len(t, users, 0)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Given
			mockUser := tt.given()

			// When
			users, err := tt.when(mockUser)

			// Then
			tt.then(users, err)
		})
	}
}

func TestDefaultFinderByIDUseCase_Find(t *testing.T) {
	tests := []struct {
		name  string
		given func() *user.MockRepository
		when  func(repository *user.MockRepository) (entity.User, error)
		then  func(entity.User, error)
	}{
		{
			name: "should find repository by ID",
			given: func() *user.MockRepository {
				repository := user.NewMockRepository()
				u := entity.User{ID: 1, Name: "John", Surname: "Doe"}
				repository.On("FindByID", context.Background(), u.ID).Return(u, nil)
				return repository
			},
			when: func(repository *user.MockRepository) (entity.User, error) {
				return NewDefaultFinderByIDUseCase(repository).Find(context.Background(), 1)
			},
			then: func(user entity.User, err error) {
				assert.NoError(t, err)
				assert.NotNil(t, user)
				assert.Equal(t, "John", user.Name)
				assert.Equal(t, "Doe", user.Surname)
			},
		},
		{
			name: "should not find repository by ID",
			given: func() *user.MockRepository {
				repository := user.NewMockRepository()
				repository.On("FindByID", context.Background(), mock.Anything).Return(entity.User{}, errors.New("not found"))
				return repository
			},
			when: func(repository *user.MockRepository) (entity.User, error) {
				return NewDefaultFinderByIDUseCase(repository).Find(context.Background(), 1)
			},
			then: func(user entity.User, err error) {
				assert.Error(t, err)
				assert.Equal(t, entity.User{}, user)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Given
			mockUser := tt.given()

			// When
			userEntity, err := tt.when(mockUser)

			// Then
			tt.then(userEntity, err)
		})
	}
}

func TestDefaultCreatorUseCase_Create(t *testing.T) {
	tests := []struct {
		name  string
		given func() *user.MockRepository
		when  func(repository *user.MockRepository) (entity.User, error)
		then  func(entity.User, error)
	}{
		{
			name: "should create repository",
			given: func() *user.MockRepository {
				repository := user.NewMockRepository()
				u := entity.User{Name: "John", Surname: "Doe"}
				created := entity.User{ID: 1, Name: "John", Surname: "Doe"}
				repository.On("save", context.Background(), u).Return(created, nil)
				return repository
			},
			when: func(repository *user.MockRepository) (entity.User, error) {
				return NewDefaultCreatorUseCase(repository).Create(context.Background(), entity.User{Name: "John", Surname: "Doe"})
			},
			then: func(user entity.User, err error) {
				assert.NoError(t, err)
				assert.NotNil(t, user)
				assert.Equal(t, uint(1), user.ID)
				assert.Equal(t, "John", user.Name)
				assert.Equal(t, "Doe", user.Surname)
			},
		},
		{
			name: "should not create repository",
			given: func() *user.MockRepository {
				repository := user.NewMockRepository()
				repository.On("save", context.Background(), mock.Anything).Return(entity.User{}, errors.New("not created"))
				return repository
			},
			when: func(repository *user.MockRepository) (entity.User, error) {
				return NewDefaultCreatorUseCase(repository).Create(context.Background(), entity.User{})
			},
			then: func(user entity.User, err error) {
				assert.Error(t, err)
				assert.Equal(t, entity.User{}, user)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Given
			mockUser := tt.given()

			// When
			userEntity, err := tt.when(mockUser)

			// Then
			tt.then(userEntity, err)
		})
	}
}

func TestDefaultModifierUseCase_Modify(t *testing.T) {
	tests := []struct {
		name  string
		given func() *user.MockRepository
		when  func(repository *user.MockRepository) (entity.User, error)
		then  func(entity.User, error)
	}{
		{
			name: "should modify repository",
			given: func() *user.MockRepository {
				repository := user.NewMockRepository()
				u := entity.User{ID: 1, Name: "John", Surname: "Doe"}
				repository.On("save", context.Background(), u).Return(u, nil)
				return repository
			},
			when: func(repository *user.MockRepository) (entity.User, error) {
				return NewDefaultModifierUseCase(repository).Modify(context.Background(), entity.User{ID: 1, Name: "John", Surname: "Doe"})
			},
			then: func(user entity.User, err error) {
				assert.NoError(t, err)
				assert.NotNil(t, user)
				assert.Equal(t, uint(1), user.ID)
				assert.Equal(t, "John", user.Name)
				assert.Equal(t, "Doe", user.Surname)
			},
		},
		{
			name: "should not modify repository",
			given: func() *user.MockRepository {
				repository := user.NewMockRepository()
				repository.On("save", context.Background(), mock.Anything).Return(entity.User{}, errors.New("not modified"))
				return repository
			},
			when: func(repository *user.MockRepository) (entity.User, error) {
				return NewDefaultModifierUseCase(repository).Modify(context.Background(), entity.User{})
			},
			then: func(user entity.User, err error) {
				assert.Error(t, err)
				assert.Equal(t, entity.User{}, user)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Given
			repository := tt.given()

			// When
			userEntity, err := tt.when(repository)

			// Then
			tt.then(userEntity, err)
		})
	}
}

func TestDefaultDeleterUseCase_Delete(t *testing.T) {
	tests := []struct {
		name  string
		given func() *user.MockRepository
		when  func(repository *user.MockRepository) error
		then  func(error)
	}{
		{
			name: "should delete repository",
			given: func() *user.MockRepository {
				repository := user.NewMockRepository()
				u := entity.User{ID: 1, Name: "John", Surname: "Doe"}
				repository.On("Delete", context.Background(), u).Return(nil)
				return repository
			},
			when: func(repository *user.MockRepository) error {
				return NewDefaultDeleterUseCase(repository).Delete(context.Background(), entity.User{ID: 1, Name: "John", Surname: "Doe"})
			},
			then: func(err error) {
				assert.NoError(t, err)
			},
		},
		{
			name: "should not delete repository",
			given: func() *user.MockRepository {
				repository := user.NewMockRepository()
				repository.On("Delete", context.Background(), mock.Anything).Return(errors.New("not deleted"))
				return repository
			},
			when: func(repository *user.MockRepository) error {
				return NewDefaultDeleterUseCase(repository).Delete(context.Background(), entity.User{})
			},
			then: func(err error) {
				assert.Error(t, err)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Given
			repository := tt.given()

			// When
			err := tt.when(repository)

			// Then
			tt.then(err)
		})
	}
}
