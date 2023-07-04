package controllers

import (
	fiber "github.com/gofiber/fiber/v2"
)

func Home(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}
func Test(c *fiber.Ctx) error {
	return c.SendString("Hello, World! Test url working")
}

func GetBooks(c *fiber.Ctx) error {
	return c.SendString("Hello, World! get books")
}
