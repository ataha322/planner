package routes

import (
	"github.com/gofiber/fiber/v2"
	"planner.xyi/src/controllers"
	"planner.xyi/src/middlewares"
)

func Setup(app *fiber.App) {
	api := app.Group("api")
	user := api.Group("user")

	user.Post("register", controllers.Register)
	user.Post("login", controllers.Login)
	user.Get("task/frontend", controllers.TaskFrontend)
	user.Get("task/backend", controllers.TaskBackend)

	userAuthenticated := user.Use(middlewares.IsAuthenticated)
	userAuthenticated.Get("user", controllers.User)
	userAuthenticated.Post("logout", controllers.Logout)
	userAuthenticated.Put("users/info", controllers.UpdateInfo)
	userAuthenticated.Put("users/password", controllers.UpdatePassword)
	userAuthenticated.Get("users/")
	userAuthenticated.Get("tasks", controllers.Tasks)
	userAuthenticated.Post("tasks", controllers.CreateTask)
	userAuthenticated.Get("tasks/:id", controllers.GetTask)
	userAuthenticated.Put("tasks/:id", controllers.UpdateTask)
	userAuthenticated.Delete("tasks/:id", controllers.DeleteTask)
	// userAuthenticated.Get("users/:id/links", controllers.Link)
	// userAuthenticated.Post("links", controllers.CreateLink)
}
