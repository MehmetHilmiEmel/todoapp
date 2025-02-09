package database

import "todo-app/models"

func init() {
	conn := Connection()
	if conn == nil {
		panic("Database connection failed!")
	}

	conn.AutoMigrate(&models.Task{})
}
