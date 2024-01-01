package database

import (
	"fmt"

	"github.com/rohitsmart/studio/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	dsn := "rohit:Rohit@3066@tcp(127.0.0.1:3306)/gostudio?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database")
	}

	DB.AutoMigrate(&model.User{})
	fmt.Println("Database connected")
}
