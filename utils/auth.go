package utils

import (
	"go-server/databases"
	"go-server/models"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

const SecretKey = "secret"

func Auth(c *fiber.Ctx) (models.User, error) {
	cookie := c.Cookies("jwt")
	var user models.User
	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})
	if err != nil {
		return user, err
	}
	claims := token.Claims.(*jwt.StandardClaims)

	error := databases.DB.First(&user, claims.Issuer).Error
	if user.Nama_user == "" {
		return user, error
	}
	return user, nil
}
