package routes

import (
	"go-server/controllers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Get("/helo", controllers.Hello)
	app.Get("/categories", controllers.GetCategory)
	app.Get("/categories/:id", controllers.GetCategoryById)

	app.Post("/products", controllers.AddProduct)
	app.Get("/products", controllers.GetProduct)

	app.Post("/auth/register", controllers.Register)
	app.Post("/auth/login", controllers.Login)
	app.Get("/auth/logout", controllers.Logout)
	app.Get("/auth/user", controllers.GetUser)

	app.Get("/cart", controllers.GetCart)
	app.Post("/cart", controllers.AddCart)

}
