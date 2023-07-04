package routes

import (
	controllers "Bookstore/Controllers"

	fiber "github.com/gofiber/fiber/v2"
)

func Bookroutes(app *fiber.App) {
	// Create routes group.
	route := app.Group("/")
	route.Get("/", controllers.Home)
	route.Get("/test", controllers.Test)
}
