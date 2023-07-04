package main

import (
	"fmt"
	"log"

	Apps "Bookstore/Config"
	Routes "Bookstore/Routes"

	fiber "github.com/gofiber/fiber/v2"
)

func main() {
	fmt.Println("test module implemented")
	config := Apps.FiberConfig()
	app := fiber.New(config)
	Routes.Bookroutes(app)
	log.Fatal(app.Listen(":3000"))

}
