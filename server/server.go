package server

import (
	"encoding/json"
	"todo-app/app/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

var adminConfig = fiber.Config{
	Prefork:     false,
	BodyLimit:   1024 * 1024 * 1024,
	JSONEncoder: json.Marshal,
	JSONDecoder: json.Unmarshal,
	ErrorHandler: func(c *fiber.Ctx, err error) error {
		code := fiber.StatusInternalServerError
		if e, ok := err.(*fiber.Error); ok {
			code = e.Code
		}
		return c.Status(code).JSON(fiber.Map{"error": err.Error()})
	},
}

func Run() {
	app := fiber.New(adminConfig)
	app.Use(compress.New())
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*", // Allows requests from any origin
		AllowMethods: "GET,POST,PATCH,DELETE",
		AllowHeaders: "Content-Type",
	}))
	routes.Routes(app)
	app.Listen(":3000")
}
