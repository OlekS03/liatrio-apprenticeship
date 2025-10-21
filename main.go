package main

import (
	"time"
	"os"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New() // Create a new Fiber web application instance

	app.Get("/", func(c *fiber.Ctx) error {
		// Minified JSON string
		response := fmt.Sprintf(
			`{"message":"My name is Oleksandr Shestopalov","timestamp":%d}`,
			time.Now().UnixMilli(), // Current Unix time in milliseconds
		)

		// Set the Content-Type header to indicate JSON response
		c.Set("Content-Type", "application/json") 

		// Send the JSON string as the HTTP response body
		return c.SendString(response)
	})

	// If not set, default to port 80
	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}

	// Log a message to indicate which port the server is listening on
	fmt.Println("Server listening on port:", port)

	// Start the Fiber application and listen on the specified port
	app.Listen(":" + port)
}
