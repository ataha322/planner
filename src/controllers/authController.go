package controllers

import "github.com/gofiber/fiber/v2"

func Register(c *fiber.Ctx) error {

	return c.JSON(fiber.Map{
		"message": "hello",
	})

}
