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
	response := fmt.Sprintf(
		`{"message":"My name is Olexander Shestopalov","timestamp":%d}`,
		time.Now().UnixMilli(),
	)

	c.Response().Header.SetContentType("application/json; charset=utf-8")
	c.Response().SetBodyRaw([]byte(response))
	return nil
})


	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}

	fmt.Println("Server listening on port:", port)
	app.Listen(":" + port)
}
