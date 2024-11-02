package database

import (
	"fmt"
	"golang-fiber-cicd/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "golang-fiber-cicd:piDKpyxKEXya3msh@tcp(172.17.0.2:3306)/golang-fiber-cicd?charset=utf8mb4&parseTime=True&loc=Local"
	connection, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}

	fmt.Println("Database connected!")
	connection.AutoMigrate(&models.User{})
	DB = connection
}
