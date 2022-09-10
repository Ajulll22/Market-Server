package controllers

import (
	"go-server/databases"
	"go-server/models"

	"github.com/gofiber/fiber/v2"
)

func GetProduct(c *fiber.Ctx) error {
	var products []models.Product

	databases.DB.Find(&products)

	return c.Status(fiber.StatusOK).JSON(&products)
}

func AddProduct(c *fiber.Ctx) error {
	product := models.Product{
		Nama_product:      "Cloud Core - DTS - Gaming Headset",
		Deskripsi_product: "Compatible with PC, Xbox Series X|S, Xbox One",
		Gambar_product:    "sdasd.jpg",
		Url_product:       "test",
		Id_category:       3,
	}
	error := databases.DB.Create(&product).Error
	if error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": error,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Berhasil Menambahkan",
	})
}
