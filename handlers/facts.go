package handlers

import (
	"github.com/bytedance/sonic"
	"github.com/darwin808/pg-fiber/database"
	"github.com/darwin808/pg-fiber/models"
	"github.com/gofiber/fiber/v3"
)

func Home(c fiber.Ctx) error {
	return c.SendString("Hello, darwin222222222222222 👋!")
}

func ListFacts(c fiber.Ctx) error {
	facts := []models.Fact{}
	database.DB.Db.Find(&facts)

	return c.Status(200).JSON(facts)
}

func CreateFact(c fiber.Ctx) error {
	fact := new(models.Fact)
	// Read the request body
	body := c.Body()

	// Unmarshal the request body using sonic
	if err := sonic.Unmarshal(body, fact); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	database.DB.Db.Create(&fact)

	return c.Status(200).JSON(fact)
}
