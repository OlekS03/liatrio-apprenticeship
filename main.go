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
		message := "My name is Olexander Shestopalov"
		timestamp := time.Now().UnixMilli()

		jsonResponse := fmt.Sprintf(`{"message":"%s","timestamp":%d}`, message, timestamp)
		c.Set("Content-Type", "application/json")
		return c.Send([]byte(jsonResponse))
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "80" // Default for local dev
	}

	fmt.Println("Server listening on port:", port)
	app.Listen(":" + port)
}
