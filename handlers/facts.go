package handlers

import (
	"strconv"

	"github.com/Ninani/go-orm/database"
	"github.com/Ninani/go-orm/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Home(c *fiber.Ctx) error {
	return c.SendString("Hello!")
}

func CreateFact(c *fiber.Ctx) error {
	fact := new(models.Fact)

	if err := c.BodyParser(fact); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	database.DB.Db.Create(&fact)
	return c.Status(200).JSON(fact)
}

func UpdateFact(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 0, 64)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	fact := models.Fact{Model: gorm.Model{ID: uint(id)}}
	database.DB.Db.First(&fact)

	newFact := new(models.Fact)
	if err := c.BodyParser(newFact); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	if newFact.Question != "" {
		fact.Question = newFact.Question
	}
	if newFact.Answer != "" {
		fact.Answer = newFact.Answer
	}

	database.DB.Db.Save(&fact)
	return c.Status(200).JSON(fact)
}

func ListFacts(c *fiber.Ctx) error {
	facts := []models.Fact{}
	database.DB.Db.Find(&facts)

	return c.Status(200).JSON(facts)
}
