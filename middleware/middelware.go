package middleware

import (
	"github.com/gofiber/fiber/v2"
)

func ExampleMiddleware(c *fiber.Ctx) error {
	// Perform any middleware operations here
	return c.Next()
}
