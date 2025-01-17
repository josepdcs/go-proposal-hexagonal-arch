package ittest

import (
	"context"
	"fmt"
	"net/http"
	"path/filepath"
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/josepdcs/go-proposal-hexagonal-arch/internal/infrastructure/server/config"
	"github.com/josepdcs/go-proposal-hexagonal-arch/internal/infrastructure/server/di"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	tc "github.com/testcontainers/testcontainers-go/modules/compose"
)

type UserAPITestITSuite struct {
	suite.Suite
	app *fiber.App
}

func TestUserTestITSuite(t *testing.T) {
	suite.Run(t, new(UserAPITestITSuite))
}

func (st *UserAPITestITSuite) SetupSuite() {
	err := st.startUpDockerCompose()
	assert.NoError(st.T(), err)

	fullPath, _ := filepath.Abs(filepath.Join("testdata", "config.yml"))
	st.T().Setenv(config.ConfigPathEnv, fullPath)

	cfg, err := config.Load()
	require.NoError(st.T(), err)

	server, err := di.InitializeAPI(cfg.DB)
	require.NoError(st.T(), err)

	server.Start()
}

func (st *UserAPITestITSuite) TearDownSuite() {
	fmt.Println("Stopping UserAPITestITSuite...")

	err := st.app.ShutdownWithTimeout(5 * time.Second)
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

	require.NoError(st.T(), compose.Up(composeContext, tc.Wait(true)), "compose.Up()")

	return nil
}

func (st *UserAPITestITSuite) TestApiUsersFindAll() {
	req, err := http.NewRequest(http.MethodPost, "/api/users", nil)
	assert.NoError(st.T(), err)
	req.Header.Set("Authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzcyNzQxMDh9.rFaq0VmO9kTTBiTJo-54pnE6YylHn91do5Yc-Vf0F4o")

	resp, err := st.app.Test(req)
	assert.NoError(st.T(), err)

	assert.Equal(st.T(), http.StatusOK, resp.StatusCode)
}
