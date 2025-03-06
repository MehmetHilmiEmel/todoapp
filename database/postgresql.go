package database

import (
	"fmt"
	"os"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var once sync.Once
var connection *gorm.DB

func Connection() *gorm.DB {
	once.Do(func() {
		connection = initialize()
	})

	return connection
}

func initialize() *gorm.DB {
	// Environment variables
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")          // Ensure this matches your env variables
	password := os.Getenv("DB_PASSWORD")  // Use DB_PASSWORD instead of DB_PASS
	dbname := os.Getenv("DB_NAME")

	// Check if any required environment variable is missing
	if host == "" || port == "" || user == "" || password == "" || dbname == "" {
		panic("Database environment variables are not set properly!")
	}

	// Connection string for PostgreSQL
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Europe/Istanbul",
		host, port, user, password, dbname,
	)

	// Initialize connection
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to database: %v", err))
	}

	return db
}
