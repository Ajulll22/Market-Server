package controllers

import (
	"go-server/controllers/requests"
	"go-server/databases"
	"go-server/models"
	"go-server/utils"

	"github.com/gofiber/fiber/v2"
)

func AddCart(c *fiber.Ctx) error {
	user, err := utils.Auth(c)
	if err != nil {
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
	user, err := utils.Auth(c)
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}
	var carts []models.Cart

	error := databases.DB.Joins("JOIN products on products.id_product = carts.id_product").
		Joins("JOIN categories on categories.id_category = products.id_category").
		Preload("Product.Category").Where("id_user = ?", user.Id_user).
		Find(&carts).Error
	if error != nil {
		return fiber.NewError(fiber.StatusNotFound, error.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&carts)
}

func DeleteCart(c *fiber.Ctx) error {
	_, err := utils.Auth(c)
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	id, error := c.ParamsInt("id")
	if error == nil {
		var cart models.Cart
		err := databases.DB.First(&cart, id).Error
		if err != nil {
			return fiber.NewError(fiber.StatusNotFound, err.Error())
		}

		databases.DB.Delete(&cart)
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Deleted",
		})
	}
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"message": "Error Parameter",
	})
}

func EditCart(c *fiber.Ctx) error {
	_, err := utils.Auth(c)
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	id, err := c.ParamsInt("id")
	if err == nil {
		request := new(requests.EditCartRequest)

		if err := c.BodyParser(request); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": err.Error(),
			})

		}

		errors := utils.Validate(*request)
		if errors != nil {
			return c.Status(fiber.StatusBadRequest).JSON(errors)

		}

		var cart models.Cart
		err := databases.DB.Preload("Product").First(&cart, id).Error
		if err != nil {
			return fiber.NewError(fiber.StatusNotFound, err.Error())
		}

		jumlah, _ := request.Jumlah.Int64()
		cart.Jumlah = int(jumlah)
		cart.Total_harga = cart.Product.Harga_product * int(jumlah)

		databases.DB.Save(&cart)

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message":     "Updated",
			"total_harga": &cart.Total_harga,
		})
	}
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"message": "Error Parameter",
	})
}

func GetCartChecked(c *fiber.Ctx) error {
	user, err := utils.Auth(c)
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	request := new(requests.CartCheckedRequest)
	if err := c.BodyParser(request); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})

	}
	var carts []models.Cart
	if len(request.Id_check) == 0 {
		response := []string{}
		return c.Status(fiber.StatusOK).JSON(response)
	}
	databases.DB.Select("id_cart", "total_harga").Where("id_user = ?", user.Id_user).Find(&carts, request.Id_check)
	return c.Status(fiber.StatusOK).JSON(&carts)
}

func Test(c *fiber.Ctx) error {
	user, err := utils.Auth(c)
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}
	return c.JSON(&user)
}
