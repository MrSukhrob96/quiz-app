package db

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	DB_HOST = "localhost"
	DB_USER = "root"
	DB_NAME = "agent_bank"
	DB_PASS = "root"
	DB_PORT = "3306"
)

func ConnectDatabase() *gorm.DB {

	dsn := DB_USER + ":" + DB_PASS + "@tcp" + "(" + DB_HOST + ":" + DB_PORT + ")/" + DB_NAME + "?" + "parseTime=true&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return nil
	}

	return db
}
