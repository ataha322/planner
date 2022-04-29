package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() {
	_, err := gorm.Open(mysql.Open("root:root@tcp(db:3306)/plannerData"), &gorm.Config{})

	if err != nil {
		panic("Could no find the database")
	}
}
