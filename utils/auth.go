package utils

import (
	"go-server/databases"
	"go-server/models"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

const SecretKey = "secret"

func Auth(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")
	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})
	if err != nil {
		return err
	}
	claims := token.Claims.(*jwt.StandardClaims)
	var user models.User
	error := databases.DB.First(&user, claims.Issuer).Error
	if user.Nama_user == "" {
		return error
	}
	return nil
}
