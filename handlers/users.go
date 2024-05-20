package handlers

import (
	"github.com/Ninani/go-orm/database"
	"github.com/Ninani/go-orm/models"
	"github.com/gofiber/fiber/v2"
)

func CreateUser(c *fiber.Ctx) error {
	user := new(models.User)

	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	database.DB.Db.Create(&user)
	return c.Status(200).JSON(user)
}

func ListUsers(c *fiber.Ctx) error {
	users := []models.User{}
	database.DB.Db.Preload("Quizzes").Find(&users)

	return c.Status(200).JSON(users)
}
