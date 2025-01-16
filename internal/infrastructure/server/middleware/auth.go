package middleware

import (
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/pkg/errors"
)

func Authorization(c *fiber.Ctx) error {
	s := c.Get("Authorization")

	token := strings.TrimPrefix(s, "Bearer ")

	if err := validateToken(token); err != nil {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"error": "invalid token",
		})
	}

	return c.Next()
}

func validateToken(token string) error {
	_, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		return []byte("secret"), nil
	})

	return err
}
