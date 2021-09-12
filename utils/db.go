package utils

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const DbName = "chat"
const DbHost = "127.0.0.1"
const DbPort = "33060"
const DbUsername = "root"
const DbPassword = "root"

func InitDb() *gorm.DB {
	return connectDB()
}

func connectDB() *gorm.DB {
	var err error
	dsn := DbUsername + ":" + DbPassword + "@tcp" + "(" + DbHost + ":" + DbPort + ")/" + DbName + "?" + "parseTime=true&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Printf("Error connecting to database : error=%v\n", err)
		return nil
	}
	return db
}
