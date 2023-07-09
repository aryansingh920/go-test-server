package routes

import (
	"github.com/gofiber/fiber/v2"
)

func HelloWorld(c *fiber.Ctx) error {
	// Implement logic to return "Hello, World!"
	return c.JSON(fiber.Map{"message": "Hello, World!"})
}
