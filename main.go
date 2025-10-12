package main

import (
	"time"
	"os"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		// Minified JSON string (no spaces/newlines)
		response := fmt.Sprintf(
			`{"message":"My name is Olexander Shestopalov","timestamp":%d}`,
			time.Now().UnixMilli(),
		)

		c.Set("Content-Type", "application/json")
		return c.SendString(response)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}

	fmt.Println("Server listening on port:", port)
	app.Listen(":" + port)
}
