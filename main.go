package main

import (
	"time"
	"os"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Define the GET endpoint for "/"
	app.Get("/", func(c *fiber.Ctx) error {
		response := map[string]interface{}{
			"message":   "My name is Olexander Shestopalov",
			"timestamp": time.Now().UnixMilli(),
		}
		return c.JSON(response)
	})

	// Get the PORT environment variable (Cloud Run sets this automatically)
	port := os.Getenv("PORT")
	if port == "" {
		port = "80" // Default for local dev
	}

	fmt.Println("Server listening on port:", port)
	app.Listen(":" + port)
}
