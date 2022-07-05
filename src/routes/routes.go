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

	userAuthenticated := user.Use(middlewares.IsAuthenticated)
	userAuthenticated.Get("user", controllers.User)
	userAuthenticated.Post("logout", controllers.Logout)
	userAuthenticated.Put("users/info", controllers.UpdateInfo)
	userAuthenticated.Put("users/password", controllers.UpdatePassword)
}
