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
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable timezone=Europe/Istanbul",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("POSTGRES_USER"),  // Kubernetes'teki POSTGRES_USER
		os.Getenv("POSTGRES_PASSWORD"),  // Kubernetes'teki POSTGRES_PASSWORD
		os.Getenv("POSTGRES_DB"),  // Kubernetes'teki POSTGRES_DB
	)

	connection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return nil
	}

	return connection
}
