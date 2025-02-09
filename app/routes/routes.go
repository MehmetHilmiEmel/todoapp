package routes

import (
	"todo-app/app/controllers"

	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {
	taskGroup := app.Group("/tasks")
	{
		taskGroup.Get("/", controllers.Index)
		taskGroup.Post("/", controllers.Create)
		taskGroup.Patch("/:id", controllers.Update)
		taskGroup.Delete("/:id", controllers.Delete)
	}
}
