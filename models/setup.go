package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	database, err := gorm.Open(mysql.Open("root:@tcp(127.0.0.1:3306)/go_restapi_login"))

	if err != nil {
		panic(err)
	}

	database.AutoMigrate(&User{})

	DB = database
}
