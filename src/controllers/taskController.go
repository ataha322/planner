package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"planner.xyi/src/database"
	"planner.xyi/src/models"
	"sort"
	"strconv"
	"strings"
	"time"
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

	go database.ClearCache("tasks_frontend", "tasks_backend")

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

	go database.ClearCache("tasks_frontend", "tasks_backend")

	return c.JSON(task)
}

func DeleteTask(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	task := models.Task{}
	task.Id = uint(id)

	database.DB.Delete(&task)

	go database.ClearCache("tasks_frontend", "tasks_backend")

	return nil
}

func TaskFrontend(c *fiber.Ctx) error {
	var tasks []models.Task
	var ctx = context.Background()

	result, err := database.Cache.Get(ctx, "tasks_frontend").Result()

	if err != nil {
		fmt.Println(err.Error())

		database.DB.Find(&tasks)

		bytes, err := json.Marshal(tasks)

		if err != nil {
			panic(err)
		}

		if errKey := database.Cache.Set(ctx, "tasks_frontend", bytes, 24*time.Hour).Err(); err != nil {
			panic(errKey)
		}
	} else {
		json.Unmarshal([]byte(result), &tasks)
	}

	return c.JSON(tasks)
}

func TaskBackend(c *fiber.Ctx) error {
	var tasks []models.Task
	var ctx = context.Background()

	result, err := database.Cache.Get(ctx, "tasks_backend").Result()

	if err != nil {
		fmt.Println(err.Error())

		database.DB.Find(&tasks)

		bytes, err := json.Marshal(tasks)

		if err != nil {
			panic(err)
		}

		database.Cache.Set(ctx, "tasks_backend", bytes, 30*time.Minute)
	} else {
		json.Unmarshal([]byte(result), &tasks)
	}

	var searchedTasks []models.Task

	if s := c.Query("s"); s != "" {
		lower := strings.ToLower(s)
		for _, task := range tasks {
			if strings.Contains(strings.ToLower(task.TaskName), lower) || strings.Contains(strings.ToLower(task.TaskDescription), lower) {
				searchedTasks = append(searchedTasks, task)
			}
		}
	} else {
		searchedTasks = tasks
	}

	if sortParam := c.Query("sort"); sortParam != "" {
		sortLower := strings.ToLower(sortParam)
		if sortLower == "asc" {
			sort.Slice(searchedTasks, func(i, j int) bool {
				return searchedTasks[i].TaskDeadline.Before(searchedTasks[j].TaskDeadline)
			})
		} else if sortLower == "desc" {
			sort.Slice(searchedTasks, func(i, j int) bool {
				return searchedTasks[i].TaskDeadline.After(searchedTasks[j].TaskDeadline)
			})
		}
	}

	var total = len(searchedTasks)
	page, _ := strconv.Atoi(c.Query("page", "1"))
	perPage := 9
	var data []models.Task = searchedTasks

	if total <= page*perPage && total >= (page-1)*perPage {
		data = searchedTasks[(page-1)*perPage : total]
	} else if total >= page*perPage {
		data = searchedTasks[(page-1)*perPage : page*perPage]
	} else {
		data = []models.Task{}
	}

	return c.JSON(fiber.Map{
		"data":      data,
		"total":     total,
		"page":      page,
		"last_page": total/perPage + 1,
	})
}
