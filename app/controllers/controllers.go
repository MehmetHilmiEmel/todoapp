package controllers

import (
	"todo-app/database"
	"todo-app/models"

	"github.com/gofiber/fiber/v2"
)

func Index(c *fiber.Ctx) error {
	var tasks []models.Task
	if database.Connection().Model(&models.Task{}).Find(&tasks).Error != nil {
		return c.Status(500).SendString("Failed to fetch tasks")
	}

	return c.JSON(tasks)
}

func Create(c *fiber.Ctx) error {
	task := models.Task{}
	if err := c.BodyParser(&task); err != nil {
		return c.Status(500).SendString("Failed to parse JSON")
	}

	if database.Connection().Model(&models.Task{}).Create(&task).Error != nil {
		return c.Status(500).SendString("Failed to create task")
	}

	return c.JSON("Task created successfully")
}

func Update(c *fiber.Ctx) error {
	task := models.Task{}
	if err := c.BodyParser(&task); err != nil {
		return c.Status(500).SendString("Failed to parse JSON")
	}

	id := c.Params("id")
	if database.Connection().Model(&models.Task{}).Where("id = ?", id).Updates(&task).Error != nil {
		return c.Status(500).SendString("Failed to update task")
	}

	return c.JSON("Task updated successfully")
}

func Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	if database.Connection().Model(&models.Task{}).Where("id = ?", id).Delete(&models.Task{}).Error != nil {
		return c.Status(500).SendString("Failed to delete task")
	}

	return c.JSON("Task deleted successfully")
}
