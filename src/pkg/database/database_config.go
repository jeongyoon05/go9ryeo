package database

import (
	"fmt"
	"log"
	"os"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DBConn *gorm.DB
)

func InitDatabase() {
	var dbErr error

	envErr := godotenv.Load(".env")

	if envErr != nil {
		log.Fatal("Error loading .env file")
	}

	var (
		username = os.Getenv("DB_USERNAME")
		password = os.Getenv("DB_PASSWORD")
		dbName = os.Getenv("DB_NAME")
		dbHost = os.Getenv("DB_HOST")
	)
	
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4", username, password, dbHost, dbName)

	DBConn, dbErr = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	
	if dbErr != nil {
		log.Fatal("faild to connect database")
	}
	fmt.Print("Connection Opend to Database")
}