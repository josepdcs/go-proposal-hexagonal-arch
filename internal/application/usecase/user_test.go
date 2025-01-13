package usecase

import (
	"cmp"
	"context"
	"slices"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/thnkrn/go-gin-clean-arch/internal/domain/entity"
	"github.com/thnkrn/go-gin-clean-arch/internal/infrastructure/repository"
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
				m.On("FindAll", mock.Anything).Return(users, nil)
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
				m.On("FindAll", mock.Anything).Return(users, errors.New("not found"))
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
