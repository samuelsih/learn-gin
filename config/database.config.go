package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

//SetupDatabaseConnection create new connection to the database
func SetupDatabaseConnection() *gorm.DB {
	errENV := godotenv.Load()
	if errENV != nil {
		panic("Failed to load / recognize .env file")
	}

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=local", dbUser, dbPassword, dbHost, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to MYSQL")
	}

	return db
}

//CloseDatabaseConnection close the database when needed
func CloseDatabaseConnection(database *gorm.DB) {
	sqlDB, err := database.DB()

	if err != nil {
		panic("Failed to close the database")
	}

	sqlDB.Close()
}