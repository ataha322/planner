package controllers

import (
	"github.com/gofiber/fiber/v2"
	"planner.xyi/src/database"
	"planner.xyi/src/middlewares"
	"planner.xyi/src/models"
)

func Users(c *fiber.Ctx) error {
	var users []models.User

	database.DB.Find(&users)

	return c.JSON(users)
}

//Finds a user by ID and returns their taskList
//TODO: test this function
func GetTaskList(c *fiber.Ctx) error {
	id, _ := middlewares.GetUserId(c)
	var user models.User
	database.DB.Where("id = ?", id).First(&user)
	return c.JSON(user.Tasks)
}
