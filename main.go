package main

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/aryansingh920/go-test-server/middleware"
	"github.com/aryansingh920/go-test-server/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Use the example middleware for all routes
	app.Use(middleware.ExampleMiddleware)

	// Setup routes
	routes.SetupRoutes(app)

	// Read the port number from environment variables
	port := os.Getenv("PORT")
	if port == "" {
		port = "3001" // Default to port 3000 if PORT environment variable is not set
	}

	// Start the server
	go func() {
		if err := app.Listen(":" + port); err != nil {
			fmt.Printf("Error starting the server: %s\n", err.Error())
		}
	}()

	fmt.Printf("Server running at http://localhost:%s\n", port)

	// Wait for the server to be interrupted
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	fmt.Println("Server stopped.")
}
