package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"planner.xyi/src/middlewares"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"planner.xyi/src/database"
	"planner.xyi/src/models"
)

// Tasks : returns an array of tasks specific to the current UserId
func Tasks(c *fiber.Ctx) error {
	var tasks []models.Task

	id, _ := middlewares.GetUserId(c) //getting user id

	database.DB.Where("user_id = ?", id).Find(&tasks)

	return c.JSON(tasks)
}

// CreateTask : creates a task and assigns a current UserId to it
func CreateTask(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	id, _ := middlewares.GetUserId(c) //getting UserId

	time, _ := time.Parse(time.RFC1123, data["task_deadline"]) //type time.Time, parsing from json
	//first argument is format, second is the deadline itself
	//RFC1123 format looks as following: "Mon, 02 Jan 2006 15:04:05 MST"

	task := models.Task{
		TaskName:        data["task_name"],
		TaskDescription: data["task_description"],
		TaskDeadline:    time,
		UserId:          id,
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

// DeleteTask : deletes a task by its id
func DeleteTask(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	task := models.Task{}
	task.Id = uint(id)

	database.DB.Delete(&task)

	go database.ClearCache("tasks_frontend", "tasks_backend")

	return nil
}

// TaskFrontend : caching tasks
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

//TaskBackend : Sort, Search, display in pages
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
