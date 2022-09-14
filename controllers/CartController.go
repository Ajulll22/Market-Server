package controllers

import (
	"go-server/controllers/requests"
	"go-server/databases"
	"go-server/models"
	"go-server/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func AddCart(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")
	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}
	claims := token.Claims.(*jwt.StandardClaims)
	var user models.User
	databases.DB.First(&user, claims.Issuer)
	if user.Nama_user == "" {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	request := new(requests.AddCartRequest)
	if err := c.BodyParser(request); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})

	}
	errors := utils.Validate(*request)
	if errors != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(errors)
	}

	var product models.Product
	databases.DB.First(&product, request.Id_product)
	if product.Id_product == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Product Tidak Ditemukan",
		})
	}

	var cart models.Cart
	databases.DB.Where("id_product = ? AND id_user = ?", request.Id_product, user.Id_user).First(&cart)
	if cart.Id_cart != 0 {
		cart.Jumlah = cart.Jumlah + 1
		cart.Total_harga = cart.Jumlah * product.Harga_product

		databases.DB.Save(&cart)
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Success",
		})
	}
	cart.Jumlah = 1
	cart.Total_harga = product.Harga_product
	cart.Id_product = product.Id_product
	cart.Id_user = user.Id_user
	databases.DB.Create(&cart)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Success",
	})
}

func GetCart(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")
	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}
	claims := token.Claims.(*jwt.StandardClaims)
	var user models.User
	databases.DB.First(&user, claims.Issuer)
	if user.Nama_user == "" {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	var carts []models.Cart

	error := databases.DB.Joins("JOIN products on products.id_product = carts.id_product").
		Joins("JOIN categories on categories.id_category = products.id_category").
		Preload("Product.Category").
		Find(&carts).Error
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, error.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&carts)
}
