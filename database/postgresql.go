package database

import (
	"fmt"
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
	host := "10.96.196.199"
	port := "5432"
	user := "todouser"         // Ensure this matches your env variables
	password := "todopassword"  // Use DB_PASSWORD instead of DB_PASS
	dbname := "tododb"

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
