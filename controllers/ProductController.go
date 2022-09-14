package controllers

import (
	"fmt"
	"go-server/controllers/requests"
	"go-server/databases"
	"go-server/models"
	"go-server/utils"
	"math"
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

func AddProduct(c *fiber.Ctx) error {
	request := new(requests.AddProductRequest)

	if err := c.BodyParser(request); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})

	}
	errors := utils.Validate(*request)
	if errors != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(errors)

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
	allowFile := []string{".png", ".jpg", ".jpeg"}

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

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Berhasil Menambahkan",
	})
}
