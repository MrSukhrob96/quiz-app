package db

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	DB_HOST = "192.168.10.63"
	DB_USER = "agent"
	DB_NAME = "agent_bank"
	DB_PASS = "muhammad"
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
