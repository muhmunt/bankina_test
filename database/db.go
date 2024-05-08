package db

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitMysql(db string) *gorm.DB {
	database, err := gorm.Open(mysql.Open(db), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("Database Connected.")
	return database
}
