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
	app.Get("/products/:id", controllers.GetProductById)
	app.Put("/products/:id", controllers.EditProduct)
	app.Delete("/products/:id", controllers.DeleteProduct)

	app.Post("/auth/register", controllers.Register)
	app.Post("/auth/login", controllers.Login)
	app.Get("/auth/logout", controllers.Logout)
	app.Get("/auth/user", controllers.GetUser)
	app.Put("/auth/user", controllers.EditProfil)

	app.Get("/cart", controllers.GetCart)
	app.Post("/cart", controllers.AddCart)
	app.Delete("/cart/:id", controllers.DeleteCart)
	app.Put("/cart/:id", controllers.EditCart)
	app.Post("/cart/checked", controllers.GetCartChecked)

	app.Post("/transaction", controllers.AddTransaction)
	app.Get("/transaction", controllers.GetTransaction)
	app.Get("/transaction/:id", controllers.GetTransactionById)

	app.Get("/test", controllers.Test)

}
