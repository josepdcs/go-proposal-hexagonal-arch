package handler

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func Login(c *fiber.Ctx) error {
	// implement login logic here
	// user := c.Query("user")
	// pass := c.Query("pass")

	// // Throws Unauthorized error
	// if user != "john" || pass != "lark" {
	// 	return c.AbortWithStatus(http.StatusUnauthorized)
	// }

	// Create the Claims
	// claims := jwt.MapClaims{
	// 	"name":  "John Lark",
	// 	"admin": true,
	// 	"exp":   time.Now().Add(time.Hour * 72).Unix(),
	// }

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		ExpiresAt: &jwt.NumericDate{Time: time.Now().Add(time.Hour * 72)},
	})

	ss, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.NewError(fiber.StatusInternalServerError, err.Error()))
	}

	return c.JSON(fiber.Map{
		"token": ss,
	})
}
