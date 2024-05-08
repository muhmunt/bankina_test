package main

import (
	db "go-technical-test-bankina/database"
	"go-technical-test-bankina/src/web/server"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	dbMysql := db.InitMysql(os.Getenv("MYSQL_CONNECTION"))

	server.Run(dbMysql)
}
