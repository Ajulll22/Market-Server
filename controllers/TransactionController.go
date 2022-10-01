package controllers

import (
	"go-server/controllers/requests"
	"go-server/databases"
	"go-server/models"
	"go-server/utils"
	"time"

	"github.com/gofiber/fiber/v2"
)

func AddTransaction(c *fiber.Ctx) error {
	user, err := utils.Auth(c)
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}
	request := new(requests.AddTransactionRequest)
	if err := c.BodyParser(request); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})

	}
	if len(request.Id_cart) == 0 {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"message": "Pilih Keranjang",
		})
	}
	if len(user.Alamat_user) < 5 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Alamat tidak valid",
		})
	}

	var carts []models.Cart
	databases.DB.Where("id_user = ?", user.Id_user).
		Find(&carts, request.Id_cart)
	if len(carts) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Keranjang Tidak Ditemukan",
		})
	}
	total_trx := 0
	for _, cart := range carts {
		total_trx = total_trx + cart.Total_harga
	}
	transaction := models.Transaction{
		Total_trx:  total_trx,
		Id_user:    user.Id_user,
		Alamat_trx: user.Alamat_user,
		Created_at: time.Now(),
	}
	databases.DB.Create(&transaction)

	for _, cart := range carts {
		trx_item := models.Trx_item{
			Jumlah_item: cart.Jumlah,
			Harga_item:  cart.Total_harga,
			Id_product:  cart.Id_product,
			Id_trx:      transaction.Id_trx,
		}
		databases.DB.Create(&trx_item)
	}
	databases.DB.Delete(&carts)

	return c.Status(fiber.StatusOK).JSON(carts)
}

func GetTransaction(c *fiber.Ctx) error {
	user, err := utils.Auth(c)
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}
	var transactions []models.Transaction
	error := databases.DB.
		Preload("Trx_item.Product.Category").Order("created_at desc").Where("id_user = ?", user.Id_user).
		Find(&transactions).Error
	if error != nil {
		return fiber.NewError(fiber.StatusNotFound, error.Error())
	}
	return c.Status(fiber.StatusOK).JSON(transactions)
}

func GetTransactionById(c *fiber.Ctx) error {
	user, err := utils.Auth(c)
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Error Parameter",
		})
	}
	var transaction models.Transaction
	error := databases.DB.Joins("JOIN trx_items on trx_items.id_trx = transactions.id_trx").
		Joins("JOIN products on products.id_product = trx_items.id_product").
		Joins("JOIN categories on categories.id_category = products.id_category").
		Preload("Trx_item.Product.Category").Where("id_user = ?", user.Id_user).
		First(&transaction, id).Error
	if error != nil {
		return fiber.NewError(fiber.StatusNotFound, error.Error())
	}
	return c.Status(fiber.StatusOK).JSON(transaction)
}
