package main

import (
	"time"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()


	// When someone accesses http://<host>:1024/, this function runs.
	app.Get("/", func(c *fiber.Ctx) error {
		response := map[string]interface{}{
			"message":   "My name is Olexander Shestopalov",
			"timestamp": time.Now().Unix(),
		}
		return c.JSON(response)// Return the response as a JSON object to the client.
	})

	app.Listen(":80")
}

