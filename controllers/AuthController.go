package controllers

import (
	"go-server/controllers/requests"
	"go-server/databases"
	"go-server/models"
	"go-server/utils"

	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

const SecretKey = "secret"

func Register(c *fiber.Ctx) error {
	request := new(requests.RegisRequest)
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
	var user models.User
	databases.DB.Where("email_user = ?", request.Email_user).First(&user)
	if user.Id_user != 0 {
		c.Status(fiber.StatusUnprocessableEntity)
		return c.JSON(fiber.Map{
			"message": "Email Telah Terpakai",
		})
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(request.Password_user), 14)
	user.Email_user = request.Email_user
	user.Nama_user = request.Nama_user
	user.Password_user = string(password)
	databases.DB.Create(&user)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Berhasil Menambahkan",
	})
}

func Login(c *fiber.Ctx) error {
	request := new(requests.LoginRequest)
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
	var user models.User
	databases.DB.Where("email_user = ?", request.Email_user).First(&user)
	if user.Id_user == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "User Tidak Ditemukan",
		})
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password_user), []byte(request.Password_user)); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Password Tidak Sesuai",
		})
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(user.Id_user)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})
	token, err := claims.SignedString([]byte(SecretKey))
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":   "Berhasil Login",
		"nama_user": user.Nama_user,
	})
}

func Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success",
	})
}

func GetUser(c *fiber.Ctx) error {
	user, err := utils.Auth(c)
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}
	if user.Id_user == 0 {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	return c.Status(fiber.StatusOK).JSON(&user)
}

func EditProfil(c *fiber.Ctx) error {
	user, err := utils.Auth(c)
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}
	if user.Id_user == 0 {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}
	request := new(requests.EditProfilRequest)
	if err := c.BodyParser(request); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	errors := utils.Validate(*request)
	if errors != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"error":   &errors,
			"message": "Data Tidak Sesuai",
		})
	}

	user.Alamat_user = request.Alamat_user
	databases.DB.Save(&user)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Berhasil Update",
	})
}
