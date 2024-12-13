package database

import (
	"airbnb/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {

	dsn := "root:LOVE@123anish@tcp(127.0.0.1:3306)/airbnb_metrics?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the MySQL database!")
	}
	DB = db

	err = DB.AutoMigrate(&models.Room{})
	if err != nil {
		panic("Failed to migrate the database schema!")
	}
}
