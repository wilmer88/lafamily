package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	
	"gorm.io/gorm"
)

const DB_USERNAME = "root"
const DB_PASSWORD = "morter706"
const DB_NAME = "wilmerfamily"
const DB_HOST = "localhost"
const DB_PORT = "3306"

var Db *gorm.DB

func InitDb() *gorm.DB {
	Db = connectDB()
	return Db
}

func connectDB() *gorm.DB {
	var err error
	dsn := "root:morter706@tcp(127.0.0.1:3306)/wilmerfamily?parseTime=true&loc=Local"
	fmt.Println("dsn :is workig ", dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Error connecting to database : error=%v", err)
		return nil
	}

	return db
}
