//go:build integration_test
// +build integration_test

package ittests

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"path/filepath"
	"testing"
	"time"

	json "github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
	"github.com/josepdcs/go-proposal-hexagonal-arch/internal/api/handler"
	"github.com/josepdcs/go-proposal-hexagonal-arch/internal/domain/entity"
	"github.com/josepdcs/go-proposal-hexagonal-arch/internal/infrastructure/app"
	"github.com/josepdcs/go-proposal-hexagonal-arch/internal/infrastructure/server/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	tc "github.com/testcontainers/testcontainers-go/modules/compose"
	"github.com/testcontainers/testcontainers-go/wait"
)

const (
	LoginEndpoint = "http://localhost:8080/login"
	UsersEndpoint = "http://localhost:8080/api/users"

	BearerToken = "Bearer "
)

type TokenResponse struct {
	Token string `json:"token"`
}

type UserAPITestITSuite struct {
	suite.Suite
	token string
}

func TestUserTestITSuite(t *testing.T) {
	suite.Run(t, new(UserAPITestITSuite))
}

func (st *UserAPITestITSuite) SetupSuite() {
	err := st.startUpDockerCompose()
	assert.NoError(st.T(), err)

	fullPath, _ := filepath.Abs(filepath.Join("testdata", "config.yml"))
	st.T().Setenv(config.ConfigPathEnv, fullPath)
	st.T().Setenv(config.ConfigOverridePathEnv, "")

	// Start the application in the background to avoid blocking the test suite
	go func() {
		err = app.Start()
		require.NoError(st.T(), err)
	}()

	// Wait for the application to be ready
	time.Sleep(1 * time.Second)

	// Login to get the token
	st.doLogin()
}

func (st *UserAPITestITSuite) doLogin() {
	req, err := http.NewRequest(http.MethodPost, LoginEndpoint, nil)
	assert.NoError(st.T(), err)

	client := &http.Client{}

	resp, err := client.Do(req)
	assert.NoError(st.T(), err)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		assert.NoError(st.T(), err)
	}(resp.Body)

	assert.Equal(st.T(), http.StatusOK, resp.StatusCode)

	body, err := io.ReadAll(resp.Body)
	assert.NoError(st.T(), err)

	var tokenResponse TokenResponse
	err = json.Unmarshal(body, &tokenResponse)
	assert.NoError(st.T(), err)
	assert.NotEmpty(st.T(), tokenResponse.Token)

	st.token = tokenResponse.Token
}

func (st *UserAPITestITSuite) TearDownSuite() {
	fmt.Println("Stopping UserAPITestITSuite...")

	err := app.ShutdownWithTimeout(5 * time.Second)
	assert.NoError(st.T(), err)
}

func (st *UserAPITestITSuite) startUpDockerCompose() error {
	fullPath, err := filepath.Abs(filepath.Join("testdata", "docker-compose.yml"))
	assert.NoError(st.T(), err)

	compose, err := tc.NewDockerCompose(fullPath)
	assert.NoError(st.T(), err, "NewDockerComposeAPI()")

	st.T().Cleanup(func() {
		assert.NoError(st.T(), compose.Down(context.Background(), tc.RemoveOrphans(true), tc.RemoveImagesLocal), "compose.Down()")
	})

	composeContext, cancel := context.WithCancel(context.Background())
	st.T().Cleanup(cancel)

	compose.WaitForService("db", wait.ForExposedPort())
	err = compose.Up(composeContext, tc.Wait(true))
	require.NoError(st.T(), err, "compose.Up()")

	return nil
}

func (st *UserAPITestITSuite) TestApiUsersFindAll() {
	req, err := http.NewRequest(http.MethodGet, UsersEndpoint, nil)
	assert.NoError(st.T(), err)
	req.Header.Set(fiber.HeaderAuthorization, BearerToken+st.token)

	client := &http.Client{}

	resp, err := client.Do(req)
	assert.NoError(st.T(), err)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		assert.NoError(st.T(), err)
	}(resp.Body)

	assert.Equal(st.T(), http.StatusOK, resp.StatusCode)

	body, err := io.ReadAll(resp.Body)
	assert.NoError(st.T(), err)

	var userResponses []handler.UserDTO
	err = json.Unmarshal(body, &userResponses)
	assert.NoError(st.T(), err)
	assert.Len(st.T(), userResponses, 3)
	assert.Equal(st.T(), "John", userResponses[0].Name)
	assert.Equal(st.T(), "Doe", userResponses[0].Surname)
	assert.Equal(st.T(), "Jane", userResponses[1].Name)
	assert.Equal(st.T(), "Doe", userResponses[1].Surname)
	assert.Equal(st.T(), "Alice", userResponses[2].Name)
	assert.Equal(st.T(), "Smith", userResponses[2].Surname)
}

func (st *UserAPITestITSuite) TestApiUsersFindById() {
	req, err := http.NewRequest(http.MethodGet, UsersEndpoint+"/1", nil)
	assert.NoError(st.T(), err)
	req.Header.Set(fiber.HeaderAuthorization, BearerToken+st.token)

	client := &http.Client{}

	resp, err := client.Do(req)
	assert.NoError(st.T(), err)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		assert.NoError(st.T(), err)
	}(resp.Body)

	assert.Equal(st.T(), http.StatusOK, resp.StatusCode)

	body, err := io.ReadAll(resp.Body)
	assert.NoError(st.T(), err)

	var userResponses handler.UserDTO
	err = json.Unmarshal(body, &userResponses)
	assert.NoError(st.T(), err)
	assert.Equal(st.T(), uint(1), userResponses.ID)
	assert.Equal(st.T(), "John", userResponses.Name)
	assert.Equal(st.T(), "Doe", userResponses.Surname)
}

func (st *UserAPITestITSuite) TestApiUsersFindByIdNotFound() {
	req, err := http.NewRequest(http.MethodGet, UsersEndpoint+"/999", nil)
	assert.NoError(st.T(), err)
	req.Header.Set(fiber.HeaderAuthorization, BearerToken+st.token)

	client := &http.Client{}

	resp, err := client.Do(req)
	assert.NoError(st.T(), err)
	assert.Equal(st.T(), http.StatusNotFound, resp.StatusCode)
}

func (st *UserAPITestITSuite) TestApiUsersCreateModifyAndDelete() {
	// Create a new user
	user := entity.User{
		Name:    "John",
		Surname: "Doe",
	}
	body, err := json.Marshal(user)
	assert.NoError(st.T(), err)
	req, err := http.NewRequest(http.MethodPost, UsersEndpoint, bytes.NewReader(body))
	assert.NoError(st.T(), err)
	req.Header.Add(fiber.HeaderContentType, fiber.MIMEApplicationJSONCharsetUTF8)
	req.Header.Set(fiber.HeaderAuthorization, BearerToken+st.token)

	client := &http.Client{}

	resp, err := client.Do(req)
	assert.NoError(st.T(), err)
	assert.Equal(st.T(), http.StatusOK, resp.StatusCode)

	body, err = io.ReadAll(resp.Body)
	assert.NoError(st.T(), err)

	var userResponses handler.UserDTO
	err = json.Unmarshal(body, &userResponses)
	assert.NoError(st.T(), err)
	assert.Equal(st.T(), uint(4), userResponses.ID)
	assert.Equal(st.T(), "John", userResponses.Name)
	assert.Equal(st.T(), "Doe", userResponses.Surname)

	err = resp.Body.Close()
	assert.NoError(st.T(), err)

	// Modify the user
	user = entity.User{
		ID:      userResponses.ID,
		Name:    "John Modified",
		Surname: "Doe Modified",
	}
	body, err = json.Marshal(user)
	assert.NoError(st.T(), err)
	req, err = http.NewRequest(http.MethodPut, UsersEndpoint+"/4", bytes.NewReader(body))
	assert.NoError(st.T(), err)
	req.Header.Add(fiber.HeaderContentType, fiber.MIMEApplicationJSONCharsetUTF8)
	req.Header.Set(fiber.HeaderAuthorization, BearerToken+st.token)

	resp, err = client.Do(req)
	assert.NoError(st.T(), err)
	assert.Equal(st.T(), http.StatusOK, resp.StatusCode)

	body, err = io.ReadAll(resp.Body)
	assert.NoError(st.T(), err)

	err = json.Unmarshal(body, &userResponses)
	assert.NoError(st.T(), err)
	assert.Equal(st.T(), uint(4), userResponses.ID)
	assert.Equal(st.T(), "John Modified", userResponses.Name)
	assert.Equal(st.T(), "Doe Modified", userResponses.Surname)

	err = resp.Body.Close()
	assert.NoError(st.T(), err)

	// Delete the user
	req, err = http.NewRequest(http.MethodDelete, UsersEndpoint+"/4", nil)
	assert.NoError(st.T(), err)
	req.Header.Set(fiber.HeaderAuthorization, BearerToken+st.token)

	resp, err = client.Do(req)
	assert.NoError(st.T(), err)
	assert.Equal(st.T(), http.StatusNoContent, resp.StatusCode)

	// Check that the user has been deleted
	req, err = http.NewRequest(http.MethodGet, UsersEndpoint+"/4", nil)
	assert.NoError(st.T(), err)
	req.Header.Set(fiber.HeaderAuthorization, BearerToken+st.token)

	resp, err = client.Do(req)
	assert.NoError(st.T(), err)
	assert.Equal(st.T(), http.StatusNotFound, resp.StatusCode)
}
