package controllers

import (
	"go-server/databases"
	"go-server/models"

	"github.com/gofiber/fiber/v2"
)

func Hello(c *fiber.Ctx) error {

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"msg": "Hello",
	})
}

func GetCategory(c *fiber.Ctx) error {
	var categories []models.Category

	databases.DB.Find(&categories)

	return c.Status(fiber.StatusOK).JSON(&categories)
}

func GetCategoryById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Parameter Harus Integer",
		})
	}
	var category models.Category

	databases.DB.First(&category, id)

	if category.Id_category == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Data Tidak Ditemukan",
		})
	}

	return c.Status(fiber.StatusOK).JSON(&category)
}
