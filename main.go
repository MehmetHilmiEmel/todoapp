package main

import (
	_ "todo-app/database"
	"todo-app/server"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	server.Run()
}
