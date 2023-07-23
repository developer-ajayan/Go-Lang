package routes

import (
	controllers "Bookstore/Controllers"

	middlewares "Bookstore/Middleware"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Bookroutes(app *fiber.App) {
	// Create routes group.
	route := app.Group("/")
	route.Get("/", controllers.Home)

	// api routes
	apiroute := app.Group("/api")
	app.Use(
		logger.New(), // add Logger middleware
	)
	// apply middleware
	apiroute.Use(middlewares.TokenVerifyMiddleware)
	apiroute.Get("/books/list", controllers.GetBooks)
	apiroute.Post("/books/create", controllers.CreateBooks)
}
