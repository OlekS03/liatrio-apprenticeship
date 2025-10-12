package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		response := map[string]interface{}{
			"message":   "My name is Olexander Shestopalov",
			"timestamp": time.Now().UnixMilli(),
		}

		// Marshal to minified JSON (no spaces or newlines)
		minifiedJSON, err := json.Marshal(response)
		if err != nil {
			return c.Status(500).SendString(`{"error":"failed to encode JSON"}`)
		}

		c.Set("Content-Type", "application/json")
		return c.Send(minifiedJSON)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}

	fmt.Println("Server listening on port:", port)
	app.Listen(":" + port)
}
