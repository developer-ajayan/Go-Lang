package Middleware

import "github.com/gofiber/fiber/v2"

func TokenVerifyMiddleware(c *fiber.Ctx) error {
	apikey := c.Get("apikey") // get from header or query parmas
	if apikey == "" || apikey != "12345" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Missing or invalid API key",
		})
	}

	return c.Next()
}
