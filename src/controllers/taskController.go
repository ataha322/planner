package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"planner.xyi/src/database"
	"planner.xyi/src/models"
)

func Tasks(c *fiber.Ctx) error {
	var tasks []models.Task

	database.DB.Find(&tasks)

	return c.JSON(tasks)

}

func CreateTask(c *fiber.Ctx) error {
	var task models.Task

	if err := c.BodyParser(&task); err != nil {
		return err
	}

	database.DB.Create(&task)

	var user models.User

	userId, _ := strconv.Atoi(c.Params("user_id"))

	user.Id = uint(userId)

	if err := c.BodyParser(&user); err != nil {
		return err
	}
	user.AddTask(task)

	database.DB.Model(&user).Updates(&user)

	return c.JSON(task)
}

func GetTask(c *fiber.Ctx) error {
	var task models.Task
	id, _ := strconv.Atoi(c.Params("id"))

	task.Id = uint(id)
	database.DB.Find(&task)
	return c.JSON(task)
}

func UpdateTask(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	task := models.Task{}
	task.Id = uint(id)
	if err := c.BodyParser(&task); err != nil {
		return err
	}
	database.DB.Model(&task).Updates(&task)

	return c.JSON(task)
}

func UpdateDescritpion(c *fiber.Ctx) error {
	var data map[string]string

	id, _ := strconv.Atoi(c.Params("id"))

	task := models.Task{}
	task.Id = uint(id)

	if err := c.BodyParser(&task); err != nil {
		return err
	}

	task.SetDescription(data["task_description"])

	database.DB.Model(&task).Updates(&task)

	return c.JSON(task)
}

func DeleteTask(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	task := models.Task{}
	task.Id = uint(id)

	// Searches for a correct user that has this task
	// and deleates itself drom user task list

	// ? Might be too much, but as far as i understand,
	// ? we previously deleated the data from our db,
	// ? not an instance of task
	var user models.User

	userId, _ := strconv.Atoi(c.Params("user_id"))

	user.Id = uint(userId)

	if err := c.BodyParser(&user); err != nil {
		return err
	}
	user.DeleteTask(task)

	database.DB.Model(&user).Updates(&user)
	database.DB.Delete(&task)

	return nil
}
