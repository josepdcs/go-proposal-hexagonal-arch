package handler

import (
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/suite"
)

type UserTestITSuite struct {
	suite.Suite
	app *fiber.App
}

func TestUserTestITSuite(t *testing.T) {
	suite.Run(t, new(UserTestITSuite))
}
