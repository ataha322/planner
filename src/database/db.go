package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"planner.xyi/src/models"
)

var DB *gorm.DB

func Connect() {
	var err error

	dns := "root:root@tcp(db:3306)/plannerData"

	DB, err = gorm.Open(mysql.Open(dns), &gorm.Config{})

	if err != nil {
		panic("Could no find the database")
	}
}

func AutoMigrate() {
	DB.AutoMigrate(models.User{})
}
