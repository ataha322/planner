package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"planner.xyi/src/database"
	"planner.xyi/src/models"
)

func Users(c *fiber.Ctx) error {
	var users []models.User

	database.DB.Find(&users)

	return c.JSON(users)
}

//Finds a user by ID
func GetUser(c *fiber.Ctx) error {
	var user models.User
	id, _ := strconv.Atoi(c.Params("id"))

	user.Id = uint(id)
	database.DB.Find(&user)
	return c.JSON(user)

}
