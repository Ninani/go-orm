package handlers

import (
	"math/rand"
	"strconv"

	"github.com/Ninani/go-orm/database"
	"github.com/Ninani/go-orm/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CreateQuiz(c *fiber.Ctx) error {
	userId, err := strconv.ParseUint(c.Params("user_id"), 0, 64)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	// quizSize can be parametrized in the future
	quizSize := 3

	facts := []models.Fact{}
	database.DB.Db.Find(&facts)

	quizFacts := []models.Fact{}
	for i := 0; i < quizSize; i++ {
		quizFacts = append(quizFacts, facts[rand.Intn(len(facts))])
	}

	quiz := &models.Quiz{
		UserRefer: uint(userId),
		Facts:     quizFacts,
		Result:    0,
	}

	database.DB.Db.Create(&quiz)
	return c.Status(200).JSON(quiz)
}

func ShowQuiz(c *fiber.Ctx) error {
	quizId, err := strconv.ParseUint(c.Params("quiz_id"), 0, 64)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	quiz := models.Quiz{Model: gorm.Model{ID: uint(quizId)}}
	database.DB.Db.Preload("Facts").First(&quiz)

	return c.Status(200).JSON(quiz)
}
