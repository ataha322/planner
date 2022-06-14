package main

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"planner.xyi/src/database"
	"planner.xyi/src/routes"
)

func main() {

	_, err := gorm.Open(mysql.Open("root:root@tcp(db:3306)/plannerData"), &gorm.Config{})

	if err != nil {
		panic("Could no find the database")
	}

	database.Connect()
	database.AutoMigrate()
	app := fiber.New()

	routes.Setup(app)

	app.Listen(":8000")
}
