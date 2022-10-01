package utils

import (
	"github.com/gofiber/fiber/v2"
)

func Admin(c *fiber.Ctx) error {
	user, err := Auth(c)
	if err != nil {
		return err
	}
	if user.Level_user != 2 {
		return fiber.ErrUnauthorized
	}
	return nil
}
