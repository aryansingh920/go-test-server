package routes

import (
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", HelloWorld)
	app.Get("/files", getAllFiles)
	app.Post("/files", createFile)
}

func getAllFiles(c *fiber.Ctx) error {
	// Implement logic to fetch and return all files
	return c.SendString("Get all files")
}

func createFile(c *fiber.Ctx) error {
	// Implement logic to create a new file
	return c.SendString("Create a file")
}
