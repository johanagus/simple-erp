package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/johanagus/simple-erp/config"
)

func main() {

	app := fiber.New()
	config.InitDB()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Wellcome to Fiber App")
	})

	app.Listen(":8000")
}
