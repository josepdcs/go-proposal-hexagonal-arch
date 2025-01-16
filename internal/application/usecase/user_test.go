package usecase

import (
	"cmp"
	"context"
	"slices"
	"testing"

	"github.com/josepdcs/go-proposal-hexagonal-arch/internal/domain/entity"
	"github.com/josepdcs/go-proposal-hexagonal-arch/internal/infrastructure/repository"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUserFinderAll_Find(t *testing.T) {
	tests := []struct {
		name  string
		given func() *repository.UserMock
		when  func(userMock *repository.UserMock) ([]entity.User, error)
		then  func([]entity.User, error)
	}{
		{
			name: "should find all users",
			given: func() *repository.UserMock {
				m := repository.NewUserMock()
				users := []entity.User{
					{ID: 1, Name: "John", Surname: "Doe"},
					{ID: 2, Name: "Jane", Surname: "Doe"},
					{ID: 3, Name: "Alice", Surname: "Smith"},
				}
				m.On("FindAll", context.Background()).Return(users, nil)
				return m
			},
			when: func(userMock *repository.UserMock) ([]entity.User, error) {
				return NewUserFinderAll(userMock).Find(context.Background())
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
			given: func() *repository.UserMock {
				m := repository.NewUserMock()
				var users []entity.User
				m.On("FindAll", context.Background()).Return(users, errors.New("not found"))
				return m
			},
			when: func(userMock *repository.UserMock) ([]entity.User, error) {
				return NewUserFinderAll(userMock).Find(context.Background())
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
			userMock := tt.given()

			// When
			users, err := tt.when(userMock)

			// Then
			tt.then(users, err)
		})
	}
}

func TestUserFinderByID_Find(t *testing.T) {
	tests := []struct {
		name  string
		given func() *repository.UserMock
		when  func(userMock *repository.UserMock) (entity.User, error)
		then  func(entity.User, error)
	}{
		{
			name: "should find user by ID",
			given: func() *repository.UserMock {
				m := repository.NewUserMock()
				user := entity.User{ID: 1, Name: "John", Surname: "Doe"}
				m.On("FindByID", context.Background(), user.ID).Return(user, nil)
				return m
			},
			when: func(userMock *repository.UserMock) (entity.User, error) {
				return NewUserFinderByID(userMock).Find(context.Background(), 1)
			},
			then: func(user entity.User, err error) {
				assert.NoError(t, err)
				assert.NotNil(t, user)
				assert.Equal(t, "John", user.Name)
				assert.Equal(t, "Doe", user.Surname)
			},
		},
		{
			name: "should not find user by ID",
			given: func() *repository.UserMock {
				m := repository.NewUserMock()
				m.On("FindByID", context.Background(), mock.Anything).Return(entity.User{}, errors.New("not found"))
				return m
			},
			when: func(userMock *repository.UserMock) (entity.User, error) {
				return NewUserFinderByID(userMock).Find(context.Background(), 1)
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
			userMock := tt.given()

			// When
			user, err := tt.when(userMock)

			// Then
			tt.then(user, err)
		})
	}
}

func TestUserCreator_Create(t *testing.T) {
	tests := []struct {
		name  string
		given func() *repository.UserMock
		when  func(userMock *repository.UserMock) (entity.User, error)
		then  func(entity.User, error)
	}{
		{
			name: "should create user",
			given: func() *repository.UserMock {
				m := repository.NewUserMock()
				user := entity.User{Name: "John", Surname: "Doe"}
				created := entity.User{ID: 1, Name: "John", Surname: "Doe"}
				m.On("save", context.Background(), user).Return(created, nil)
				return m
			},
			when: func(userMock *repository.UserMock) (entity.User, error) {
				return NewUserCreator(userMock).Create(context.Background(), entity.User{Name: "John", Surname: "Doe"})
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
			name: "should not create user",
			given: func() *repository.UserMock {
				m := repository.NewUserMock()
				m.On("save", context.Background(), mock.Anything).Return(entity.User{}, errors.New("not created"))
				return m
			},
			when: func(userMock *repository.UserMock) (entity.User, error) {
				return NewUserCreator(userMock).Create(context.Background(), entity.User{})
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
			userMock := tt.given()

			// When
			user, err := tt.when(userMock)

			// Then
			tt.then(user, err)
		})
	}
}

func TestUserModifier_Modify(t *testing.T) {
	tests := []struct {
		name  string
		given func() *repository.UserMock
		when  func(userMock *repository.UserMock) (entity.User, error)
		then  func(entity.User, error)
	}{
		{
			name: "should modify user",
			given: func() *repository.UserMock {
				m := repository.NewUserMock()
				user := entity.User{ID: 1, Name: "John", Surname: "Doe"}
				m.On("save", context.Background(), user).Return(user, nil)
				return m
			},
			when: func(userMock *repository.UserMock) (entity.User, error) {
				return NewUserModifier(userMock).Modify(context.Background(), entity.User{ID: 1, Name: "John", Surname: "Doe"})
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
			name: "should not modify user",
			given: func() *repository.UserMock {
				m := repository.NewUserMock()
				m.On("save", context.Background(), mock.Anything).Return(entity.User{}, errors.New("not modified"))
				return m
			},
			when: func(userMock *repository.UserMock) (entity.User, error) {
				return NewUserModifier(userMock).Modify(context.Background(), entity.User{})
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
			userMock := tt.given()

			// When
			user, err := tt.when(userMock)

			// Then
			tt.then(user, err)
		})
	}
}

func TestUserDeleter_Delete(t *testing.T) {
	tests := []struct {
		name  string
		given func() *repository.UserMock
		when  func(userMock *repository.UserMock) error
		then  func(error)
	}{
		{
			name: "should delete user",
			given: func() *repository.UserMock {
				m := repository.NewUserMock()
				user := entity.User{ID: 1, Name: "John", Surname: "Doe"}
				m.On("Delete", context.Background(), user).Return(nil)
				return m
			},
			when: func(userMock *repository.UserMock) error {
				return NewUserDeleter(userMock).Delete(context.Background(), entity.User{ID: 1, Name: "John", Surname: "Doe"})
			},
			then: func(err error) {
				assert.NoError(t, err)
			},
		},
		{
			name: "should not delete user",
			given: func() *repository.UserMock {
				m := repository.NewUserMock()
				m.On("Delete", context.Background(), mock.Anything).Return(errors.New("not deleted"))
				return m
			},
			when: func(userMock *repository.UserMock) error {
				return NewUserDeleter(userMock).Delete(context.Background(), entity.User{})
			},
			then: func(err error) {
				assert.Error(t, err)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Given
			userMock := tt.given()

			// When
			err := tt.when(userMock)

			// Then
			tt.then(err)
		})
	}
}
