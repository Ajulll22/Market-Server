package controllers

import (
	"fmt"
	"go-server/controllers/requests"
	"go-server/databases"
	"go-server/models"
	"go-server/utils"
	"math"
	"os"
	"strings"
	"time"

	"path/filepath"

	"github.com/gofiber/fiber/v2"
)

func GetProduct(c *fiber.Ctx) error {
	dataQuery := new(requests.QueryProducts)
	if err := c.QueryParser(dataQuery); err != nil {
		return err
	}
	if dataQuery.Limit == 0 {
		dataQuery.Limit = 10
	}

	var products []models.Product
	totalRows := databases.DB.Select("id_product").Where("products.id_category LIKE ?", "%"+dataQuery.Id_category+"%").Where(databases.DB.Where("nama_product LIKE ?", "%"+dataQuery.Search+"%").Or("deskripsi_product LIKE ?", "%"+dataQuery.Search+"%")).Find(&products).RowsAffected
	offset := dataQuery.Limit * dataQuery.Page

	databases.DB.Limit(dataQuery.Limit).Offset(offset).Joins("Category").Where("products.id_category LIKE ?", "%"+dataQuery.Id_category+"%").Where(databases.DB.Where("nama_product LIKE ?", "%"+dataQuery.Search+"%").Or("deskripsi_product LIKE ?", "%"+dataQuery.Search+"%")).Order("id_product desc").Find(&products)
	totalPage := math.Ceil(float64(totalRows) / float64(dataQuery.Limit))
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"result":    &products,
		"page":      dataQuery.Page,
		"limit":     dataQuery.Limit,
		"totalPage": totalPage,
		"totalRows": totalRows,
	})
}

func GetProductById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Error Parameter",
		})
	}
	var product models.Product
	databases.DB.First(&product, id)
	if product.Id_product == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Product Tidak Ada",
		})
	}

	return c.Status(fiber.StatusOK).JSON(&product)
}

func AddProduct(c *fiber.Ctx) error {
	err := utils.Admin(c)
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}
	fmt.Print(err)
	request := new(requests.AddProductRequest)
	if err := c.BodyParser(request); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})

	}
	errors := utils.Validate(*request)
	if errors != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": errors,
		})
	}
	var category databases.Category
	databases.DB.First(&category, request.Id_category)
	if category.Id_category == 0 {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Kategori Tidak Ada",
		})
	}
	if _, err := c.FormFile("file"); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	allowFile := []string{".png", ".jpg", ".jpeg", ".webp"}

	harga, _ := request.Harga_product.Int64()
	file, _ := c.FormFile("file")
	ext := filepath.Ext(file.Filename)
	changeName := "product-" + time.Now().Format("20060102150405") + ext

	if !utils.Contains(allowFile, strings.ToLower(ext)) {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": "File Tidak Sesuai",
		})
	}
	if file.Size > 1001737 {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": "Ukuran File Terlalu Besar",
		})
	}

	c.SaveFile(file, fmt.Sprintf("./public/images/%s", changeName))

	product := databases.Product{
		Nama_product:      request.Nama_product,
		Deskripsi_product: request.Deskripsi_product,
		Gambar_product:    changeName,
		Url_product:       c.BaseURL() + "/images/" + changeName,
		Id_category:       category.Id_category,
		Harga_product:     int(harga),
	}
	databases.DB.Create(&product)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Berhasil Menambahkan",
	})
}

func EditProduct(c *fiber.Ctx) error {
	err := utils.Admin(c)
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
	var product models.Product
	databases.DB.First(&product, id)
	if product.Id_product == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Product Tidak Ada",
		})
	}

	request := new(requests.EditProductRequest)
	if err := c.BodyParser(request); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	errors := utils.Validate(*request)
	if errors != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": errors,
		})
	}
	var category databases.Category
	databases.DB.First(&category, request.Id_category)
	if category.Id_category == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Kategori Tidak Ada",
		})
	}
	harga, _ := request.Harga_product.Int64()

	product.Nama_product = request.Nama_product
	product.Deskripsi_product = request.Deskripsi_product
	product.Id_category = category.Id_category
	product.Harga_product = int(harga)
	file, _ := c.FormFile("file")
	if file != nil {
		allowFile := []string{".png", ".jpg", ".jpeg", ".webp"}

		ext := filepath.Ext(file.Filename)
		changeName := "product-" + time.Now().Format("20060102150405") + ext

		if !utils.Contains(allowFile, strings.ToLower(ext)) {
			return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
				"message": "File Tidak Sesuai",
			})
		}
		if file.Size > 1001737 {
			return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
				"message": "Ukuran File Terlalu Besar",
			})
		}
		if _, err := os.Stat("./public/images/" + product.Gambar_product); err == nil {
			os.Remove("./public/images/" + product.Gambar_product)
		}

		product.Gambar_product = changeName
		product.Url_product = c.BaseURL() + "/images/" + changeName
		c.SaveFile(file, fmt.Sprintf("./public/images/%s", changeName))
	}
	databases.DB.Save(&product)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Berhasil Update",
	})
}

func DeleteProduct(c *fiber.Ctx) error {
	err := utils.Admin(c)
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
	var product models.Product
	databases.DB.First(&product, id)
	if product.Id_product == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Product Tidak Ada",
		})
	}
	if error := databases.DB.Delete(&product).Error; error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": error.Error(),
		})
	}
	if _, err := os.Stat("./public/images/" + product.Gambar_product); err == nil {
		os.Remove("./public/images/" + product.Gambar_product)
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Berhasil Menghapus",
	})
}
