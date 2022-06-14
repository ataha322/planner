package routes

import (
	"github.com/gofiber/fiber/v2"
	"planner.xyi/src/controllers"
	"planner.xyi/src/middlewares"
)

func Setup(app *fiber.App) {
	api := app.Group("api")
	admin := api.Group("admin")

	admin.Post("register", controllers.Register)
	admin.Post("login", controllers.Login)

	adminAuthenticated := admin.Use(middlewares.IsAuthenticated)
	adminAuthenticated.Get("user", controllers.User)
	adminAuthenticated.Post("logout", controllers.Logout)
	adminAuthenticated.Put("user/info", controllers.UpdateInfo)
	adminAuthenticated.Put("user/password", controllers.UpdatePassword)

	// ? I supouse we wouldn't have tasks that are not attached of a user
	adminAuthenticated.Get("user/task", controllers.Tasks)

}
