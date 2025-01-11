package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func Login(c *gin.Context) {
	// implement login logic here
	// user := c.PostForm("user")
	// pass := c.PostForm("pass")

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
	// token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	/*token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
	})*/

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		ExpiresAt: &jwt.NumericDate{Time: time.Now().Add(time.Hour * 72)},
	})

	ss, err := token.SignedString([]byte("secret"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"token": ss,
	})
}
