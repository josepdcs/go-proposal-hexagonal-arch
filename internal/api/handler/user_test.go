package handler

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	json "github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
	"github.com/josepdcs/go-proposal-hexagonal-arch/internal/application/user"
	"github.com/josepdcs/go-proposal-hexagonal-arch/internal/domain/model/entity"
	domerrors "github.com/josepdcs/go-proposal-hexagonal-arch/internal/domain/model/errors"
	"github.com/josepdcs/go-proposal-hexagonal-arch/internal/infrastructure/input/server/testutil"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

const (
	ApiUsersEndpoint = "/api/users"
)

func TestUserAPI_FindAll(t *testing.T) {
	tests := []struct {
		name  string
		given func() *fiber.App
		when  func(a *fiber.App) (*http.Response, error)
		then  func(t *testing.T, resp *http.Response, err error)
	}{
		{
			name: "should find user by ID",
			given: func() *fiber.App {
				a := testutils.App()
				c := testutils.AcquireFiberCtx(a)

				mockFinderAllUseCase := user.NewMockFinderAllUseCase()
				mockFinderAllUseCase.On("Find", c.UserContext()).Return([]entity.User{
					{ID: 1, Name: "John", Surname: "Doe"},
					{ID: 2, Name: "Jane", Surname: "Doe"},
					{ID: 3, Name: "Alice", Surname: "Smith"},
				}, nil)
				api := NewUserAPI(
					mockFinderAllUseCase,
					nil,
					nil,
					nil,
					nil)

				a.Get(ApiUsersEndpoint, api.FindAll)
				return a
			},
			when: func(a *fiber.App) (*http.Response, error) {
				req := httptest.NewRequest(http.MethodGet, ApiUsersEndpoint, nil)
				return a.Test(req, -1)
			},
			then: func(t *testing.T, resp *http.Response, err error) {
				assert.NoError(t, err)
				assert.Equal(t, http.StatusOK, resp.StatusCode)

				body, err := io.ReadAll(resp.Body)
				assert.NoError(t, err)

				var userResponses []UserDTO
				err = json.Unmarshal(body, &userResponses)
				assert.NoError(t, err)

				assert.Len(t, userResponses, 3)
				assert.Equal(t, "John", userResponses[0].Name)
				assert.Equal(t, "Doe", userResponses[0].Surname)
				assert.Equal(t, "Jane", userResponses[1].Name)
				assert.Equal(t, "Doe", userResponses[1].Surname)
				assert.Equal(t, "Alice", userResponses[2].Name)
				assert.Equal(t, "Smith", userResponses[2].Surname)

				err = resp.Body.Close()
				assert.NoError(t, err)
			},
		},
		{
			name: "should not find users",
			given: func() *fiber.App {
				a := testutils.App()
				c := testutils.AcquireFiberCtx(a)

				mockFinderAllUseCase := user.NewMockFinderAllUseCase()
				mockFinderAllUseCase.On("Find", c.UserContext()).Return([]entity.User{}, errors.New("not found"))
				api := NewUserAPI(
					mockFinderAllUseCase,
					nil,
					nil,
					nil,
					nil)

				a.Get(ApiUsersEndpoint, api.FindAll)
				return a
			},
			when: func(a *fiber.App) (*http.Response, error) {
				req := httptest.NewRequest(http.MethodGet, ApiUsersEndpoint, nil)
				return a.Test(req, -1)
			},
			then: func(t *testing.T, resp *http.Response, err error) {
				assert.NoError(t, err)
				assert.Equal(t, http.StatusNotFound, resp.StatusCode)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Given
			a := tt.given()

			// When
			resp, err := tt.when(a)

			// Then
			tt.then(t, resp, err)
		})
	}
}

func TestUserAPI_FindByID(t *testing.T) {
	tests := []struct {
		name  string
		given func() *fiber.App
		when  func(a *fiber.App) (*http.Response, error)
		then  func(t *testing.T, resp *http.Response, err error)
	}{
		{
			name: "should find user by ID",
			given: func() *fiber.App {
				a := testutils.App()
				c := testutils.AcquireFiberCtx(a)

				mockUserFinderByIDUseCase := user.NewMockUserFinderByIDUseCase()
				mockUserFinderByIDUseCase.On("Find", c.UserContext(), uint(1)).Return(entity.User{ID: 1, Name: "John", Surname: "Doe"}, nil)
				api := NewUserAPI(
					nil,
					mockUserFinderByIDUseCase,
					nil,
					nil,
					nil)

				a.Get("/api/users/:id", api.FindByID)
				return a
			},
			when: func(a *fiber.App) (*http.Response, error) {
				req := httptest.NewRequest(http.MethodGet, ApiUsersEndpoint+"/1", nil)
				return a.Test(req, -1)
			},
			then: func(t *testing.T, resp *http.Response, err error) {
				assert.NoError(t, err)
				assert.Equal(t, http.StatusOK, resp.StatusCode)

				body, err := io.ReadAll(resp.Body)
				assert.NoError(t, err)

				var userResponses UserDTO
				err = json.Unmarshal(body, &userResponses)
				assert.NoError(t, err)

				assert.Equal(t, uint(1), userResponses.ID)
				assert.Equal(t, "John", userResponses.Name)
				assert.Equal(t, "Doe", userResponses.Surname)

				err = resp.Body.Close()
				assert.NoError(t, err)
			},
		},
		{
			name: "should not find user by ID",
			given: func() *fiber.App {
				a := testutils.App()
				c := testutils.AcquireFiberCtx(a)

				mockUserFinderByIDUseCase := user.NewMockUserFinderByIDUseCase()
				mockUserFinderByIDUseCase.On("Find", c.UserContext(), uint(999)).Return(entity.User{}, errors.New("not found"))
				api := NewUserAPI(
					nil,
					mockUserFinderByIDUseCase,
					nil,
					nil,
					nil)

				a.Get("/api/users/:id", api.FindByID)
				return a
			},
			when: func(a *fiber.App) (*http.Response, error) {
				req := httptest.NewRequest(http.MethodGet, ApiUsersEndpoint+"/999", nil)
				return a.Test(req, -1)
			},
			then: func(t *testing.T, resp *http.Response, err error) {
				assert.NoError(t, err)
				assert.Equal(t, http.StatusNotFound, resp.StatusCode)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Given
			a := tt.given()

			// When
			resp, err := tt.when(a)

			// Then
			tt.then(t, resp, err)
		})
	}
}

func TestUserAPI_Create(t *testing.T) {
	tests := []struct {
		name  string
		given func() *fiber.App
		when  func(a *fiber.App) (*http.Response, error)
		then  func(t *testing.T, resp *http.Response, err error)
	}{
		{
			name: "should create a new user",
			given: func() *fiber.App {
				a := testutils.App()
				c := testutils.AcquireFiberCtx(a)

				mockUserCreatorUseCase := user.NewMockUserCreatorUseCase()
				mockUserCreatorUseCase.On("Create", c.UserContext(), entity.User{Name: "John", Surname: "Doe"}).
					Return(entity.User{ID: 1, Name: "John", Surname: "Doe"}, nil)
				api := NewUserAPI(
					nil,
					nil,
					mockUserCreatorUseCase,
					nil,
					nil)

				a.Post(ApiUsersEndpoint, api.Create)
				return a
			},
			when: func(a *fiber.App) (*http.Response, error) {
				req := httptest.NewRequest(http.MethodPost, ApiUsersEndpoint, strings.NewReader(`{"name": "John", "surname": "Doe"}`))
				req.Header.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
				return a.Test(req, -1)
			},
			then: func(t *testing.T, resp *http.Response, err error) {
				assert.NoError(t, err)
				assert.Equal(t, http.StatusCreated, resp.StatusCode)

				body, err := io.ReadAll(resp.Body)
				assert.NoError(t, err)

				var userResponses UserDTO
				err = json.Unmarshal(body, &userResponses)
				assert.NoError(t, err)

				assert.Equal(t, uint(1), userResponses.ID)
				assert.Equal(t, "John", userResponses.Name)
				assert.Equal(t, "Doe", userResponses.Surname)

				err = resp.Body.Close()
				assert.NoError(t, err)
			},
		},
		{
			name: "should not create a new user",
			given: func() *fiber.App {
				a := testutils.App()
				c := testutils.AcquireFiberCtx(a)

				mockUserCreatorUseCase := user.NewMockUserCreatorUseCase()
				mockUserCreatorUseCase.On("Create", c.UserContext(), entity.User{Name: "John", Surname: "Doe"}).
					Return(entity.User{}, errors.New("error creating user"))
				api := NewUserAPI(
					nil,
					nil,
					mockUserCreatorUseCase,
					nil,
					nil)

				a.Post(ApiUsersEndpoint, api.Create)
				return a
			},
			when: func(a *fiber.App) (*http.Response, error) {
				req := httptest.NewRequest(http.MethodPost, ApiUsersEndpoint, strings.NewReader(`{"name": "John", "surname": "Doe"}`))
				req.Header.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
				return a.Test(req, -1)
			},
			then: func(t *testing.T, resp *http.Response, err error) {
				assert.NoError(t, err)
				assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
			},
		},
		{
			name: "should not create a new user with invalid data",
			given: func() *fiber.App {
				a := testutils.App()

				mockUserCreatorUseCase := user.NewMockUserCreatorUseCase()
				api := NewUserAPI(
					nil,
					nil,
					mockUserCreatorUseCase,
					nil,
					nil)

				a.Post(ApiUsersEndpoint, api.Create)
				return a
			},
			when: func(a *fiber.App) (*http.Response, error) {
				req := httptest.NewRequest(http.MethodPost, ApiUsersEndpoint, strings.NewReader(`{"other": "message"}`))
				return a.Test(req, -1)
			},
			then: func(t *testing.T, resp *http.Response, err error) {
				assert.NoError(t, err)
				assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Given
			a := tt.given()

			// When
			resp, err := tt.when(a)

			// Then
			tt.then(t, resp, err)
		})
	}
}

func TestUserAPI_Modify(t *testing.T) {
	tests := []struct {
		name  string
		given func() *fiber.App
		when  func(a *fiber.App) (*http.Response, error)
		then  func(t *testing.T, resp *http.Response, err error)
	}{
		{
			name: "should modify a user",
			given: func() *fiber.App {
				a := testutils.App()
				c := testutils.AcquireFiberCtx(a)

				mockUserModifierUseCase := user.NewMockUserModifierUseCase()
				mockUserModifierUseCase.On("Modify", c.UserContext(), entity.User{ID: 1, Name: "John Modified", Surname: "Doe Modified"}).
					Return(entity.User{ID: 1, Name: "John Modified", Surname: "Doe Modified"}, nil)
				api := NewUserAPI(
					nil,
					nil,
					nil,
					mockUserModifierUseCase,
					nil)

				a.Put(ApiUsersEndpoint+"/:id", api.Modify)
				return a
			},
			when: func(a *fiber.App) (*http.Response, error) {
				req := httptest.NewRequest(http.MethodPut, ApiUsersEndpoint+"/:id", strings.NewReader(`{"id": 1, "name": "John Modified", "surname": "Doe Modified"}`))
				req.Header.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
				return a.Test(req, -1)
			},
			then: func(t *testing.T, resp *http.Response, err error) {
				assert.NoError(t, err)
				assert.Equal(t, http.StatusOK, resp.StatusCode)

				body, err := io.ReadAll(resp.Body)
				assert.NoError(t, err)

				var userResponses UserDTO
				err = json.Unmarshal(body, &userResponses)
				assert.NoError(t, err)

				assert.Equal(t, uint(1), userResponses.ID)
				assert.Equal(t, "John Modified", userResponses.Name)
				assert.Equal(t, "Doe Modified", userResponses.Surname)

				err = resp.Body.Close()
				assert.NoError(t, err)
			},
		},
		{
			name: "should not modify a user",
			given: func() *fiber.App {
				a := testutils.App()
				c := testutils.AcquireFiberCtx(a)

				mockUserModifierUseCase := user.NewMockUserModifierUseCase()
				mockUserModifierUseCase.On("Modify", c.UserContext(), entity.User{ID: 1, Name: "John Modified", Surname: "Doe Modified"}).
					Return(entity.User{}, errors.New("error modifying user"))
				api := NewUserAPI(
					nil,
					nil,
					nil,
					mockUserModifierUseCase,
					nil)

				a.Put(ApiUsersEndpoint+"/:id", api.Modify)
				return a
			},
			when: func(a *fiber.App) (*http.Response, error) {
				req := httptest.NewRequest(http.MethodPut, ApiUsersEndpoint+"/:id", strings.NewReader(`{"id": 1, "name": "John Modified", "surname": "Doe Modified"}`))
				req.Header.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
				return a.Test(req, -1)
			},
			then: func(t *testing.T, resp *http.Response, err error) {
				assert.NoError(t, err)
				assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
			},
		},
		{
			name: "should not modify a user with invalid data",
			given: func() *fiber.App {
				a := testutils.App()

				mockUserModifierUseCase := user.NewMockUserModifierUseCase()
				api := NewUserAPI(
					nil,
					nil,
					nil,
					mockUserModifierUseCase,
					nil)

				a.Put(ApiUsersEndpoint+"/:id", api.Modify)
				return a
			},
			when: func(a *fiber.App) (*http.Response, error) {
				req := httptest.NewRequest(http.MethodPut, ApiUsersEndpoint+"/:id", strings.NewReader(`{"other": "message"}`))
				return a.Test(req, -1)
			},
			then: func(t *testing.T, resp *http.Response, err error) {
				assert.NoError(t, err)
				assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Given
			a := tt.given()

			// When
			resp, err := tt.when(a)

			// Then
			tt.then(t, resp, err)
		})
	}
}

func TestUserAPI_Delete(t *testing.T) {
	tests := []struct {
		name  string
		given func() *fiber.App
		when  func(a *fiber.App) (*http.Response, error)
		then  func(t *testing.T, resp *http.Response, err error)
	}{
		{
			name: "should delete a user",
			given: func() *fiber.App {
				a := testutils.App()
				c := testutils.AcquireFiberCtx(a)

				mockUserFinderByIDUseCase := user.NewMockUserFinderByIDUseCase()
				mockUserFinderByIDUseCase.On("Find", c.UserContext(), uint(1)).Return(entity.User{ID: 1, Name: "John", Surname: "Doe"}, nil)
				mockUserDeleterUseCase := user.NewMockUserDeleterUseCase()
				mockUserDeleterUseCase.On("Delete", c.UserContext(), entity.User{ID: 1, Name: "John", Surname: "Doe"}).Return(nil)
				api := NewUserAPI(
					nil,
					mockUserFinderByIDUseCase,
					nil,
					nil,
					mockUserDeleterUseCase)

				a.Delete(ApiUsersEndpoint+"/:id", api.Delete)
				return a
			},
			when: func(a *fiber.App) (*http.Response, error) {
				req := httptest.NewRequest(http.MethodDelete, ApiUsersEndpoint+"/1", nil)
				return a.Test(req, -1)
			},
			then: func(t *testing.T, resp *http.Response, err error) {
				assert.NoError(t, err)
				assert.Equal(t, http.StatusNoContent, resp.StatusCode)
			},
		},
		{
			name: "should not delete a user",
			given: func() *fiber.App {
				a := testutils.App()
				c := testutils.AcquireFiberCtx(a)

				mockUserFinderByIDUseCase := user.NewMockUserFinderByIDUseCase()
				mockUserFinderByIDUseCase.On("Find", c.UserContext(), uint(1)).Return(entity.User{ID: 1, Name: "John", Surname: "Doe"}, nil)
				mockUserDeleterUseCase := user.NewMockUserDeleterUseCase()
				mockUserDeleterUseCase.On("Delete", c.UserContext(), entity.User{ID: 1, Name: "John", Surname: "Doe"}).
					Return(errors.New("error deleting user"))
				api := NewUserAPI(
					nil,
					mockUserFinderByIDUseCase,
					nil,
					nil,
					mockUserDeleterUseCase)

				a.Delete(ApiUsersEndpoint+"/:id", api.Delete)
				return a
			},
			when: func(a *fiber.App) (*http.Response, error) {
				req := httptest.NewRequest(http.MethodDelete, ApiUsersEndpoint+"/1", nil)
				return a.Test(req, -1)
			},
			then: func(t *testing.T, resp *http.Response, err error) {
				assert.NoError(t, err)
				assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
			},
		},
		{
			name: "should not delete a user when error not found",
			given: func() *fiber.App {
				a := testutils.App()
				c := testutils.AcquireFiberCtx(a)

				mockUserFinderByIDUseCase := user.NewMockUserFinderByIDUseCase()
				mockUserFinderByIDUseCase.On("Find", c.UserContext(), uint(1)).Return(entity.User{}, domerrors.ErrUserNotFound)
				mockUserDeleterUseCase := user.NewMockUserDeleterUseCase()
				api := NewUserAPI(
					nil,
					mockUserFinderByIDUseCase,
					nil,
					nil,
					mockUserDeleterUseCase)

				a.Delete(ApiUsersEndpoint+"/:id", api.Delete)
				return a
			},
			when: func(a *fiber.App) (*http.Response, error) {
				req := httptest.NewRequest(http.MethodDelete, ApiUsersEndpoint+"/1", nil)
				return a.Test(req, -1)
			},
			then: func(t *testing.T, resp *http.Response, err error) {
				assert.NoError(t, err)
				assert.Equal(t, http.StatusNotFound, resp.StatusCode)
			},
		},
		{
			name: "should not delete a user when not found",
			given: func() *fiber.App {
				a := testutils.App()
				c := testutils.AcquireFiberCtx(a)

				mockUserFinderByIDUseCase := user.NewMockUserFinderByIDUseCase()
				mockUserFinderByIDUseCase.On("Find", c.UserContext(), uint(1)).Return(entity.User{}, nil)
				mockUserDeleterUseCase := user.NewMockUserDeleterUseCase()
				api := NewUserAPI(
					nil,
					mockUserFinderByIDUseCase,
					nil,
					nil,
					mockUserDeleterUseCase)

				a.Delete(ApiUsersEndpoint+"/:id", api.Delete)
				return a
			},
			when: func(a *fiber.App) (*http.Response, error) {
				req := httptest.NewRequest(http.MethodDelete, ApiUsersEndpoint+"/1", nil)
				return a.Test(req, -1)
			},
			then: func(t *testing.T, resp *http.Response, err error) {
				assert.NoError(t, err)
				assert.Equal(t, http.StatusNotFound, resp.StatusCode)
			},
		},
		{
			name: "should not delete a user when not error finding user",
			given: func() *fiber.App {
				a := testutils.App()
				c := testutils.AcquireFiberCtx(a)

				mockUserFinderByIDUseCase := user.NewMockUserFinderByIDUseCase()
				mockUserFinderByIDUseCase.On("Find", c.UserContext(), uint(1)).Return(entity.User{}, errors.New("error finding user"))
				mockUserDeleterUseCase := user.NewMockUserDeleterUseCase()
				api := NewUserAPI(
					nil,
					mockUserFinderByIDUseCase,
					nil,
					nil,
					mockUserDeleterUseCase)

				a.Delete(ApiUsersEndpoint+"/:id", api.Delete)
				return a
			},
			when: func(a *fiber.App) (*http.Response, error) {
				req := httptest.NewRequest(http.MethodDelete, ApiUsersEndpoint+"/1", nil)
				return a.Test(req, -1)
			},
			then: func(t *testing.T, resp *http.Response, err error) {
				assert.NoError(t, err)
				assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Given
			a := tt.given()

			// When
			resp, err := tt.when(a)

			// Then
			tt.then(t, resp, err)
		})
	}
}
