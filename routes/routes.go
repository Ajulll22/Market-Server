package routes

import (
	"go-server/controllers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Get("/helo", controllers.Hello)
	app.Get("/users", controllers.GetCategory)
	app.Get("/users/:id", controllers.GetCategoryById)

	app.Get("/products", controllers.AddProduct)
}
