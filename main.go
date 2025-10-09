package main

import (
	"time"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		response := map[string]interface{}{
			"message":   "My name is Olexander Shestopalov",
			"timestamp": time.Now().Unix(),
		}
		return c.JSON(response)
	})

	app.Listen(":8080")
}

