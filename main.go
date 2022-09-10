package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"go-server/config"
	"go-server/databases"
	"go-server/routes"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed to load config ", err)
	}

	app := fiber.New()
	app.Use(cors.New())

	databases.Conect(c.DBUrl)

	routes.Setup(app)

	app.Listen(c.Port)
}
