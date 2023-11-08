package main

import (
	"log"

	Apps "Bookstore/Config"
	Routes "Bookstore/Routes"

	fiber "github.com/gofiber/fiber/v2"
)

func main() {

	config := Apps.FiberConfig()
	app := fiber.New(config)

	// routes
	Routes.Bookroutes(app)

	log.Fatal(app.Listen(":4000"))

}
