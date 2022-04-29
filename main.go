package main

import (
	"github.com/gofiber/fiber/v2"
	"planner.xyi/src/database"
)

func main() {

	database.Connect()

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World desu~~~~~!!!!!!!")
	})

	app.Listen(":8000")
}
