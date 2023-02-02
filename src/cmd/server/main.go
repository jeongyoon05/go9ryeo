package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jeongyoon05/go9ryeo/src/pkg/database"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	database.InitDatabase()

	app.Listen(":3000")
}
