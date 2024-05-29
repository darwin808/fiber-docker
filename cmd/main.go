package main

import (
	"github.com/darwin808/pg-fiber/database"
	"github.com/gofiber/fiber/v3"
)

func main() {
	database.ConnectDb()

	app := fiber.New(fiber.Config{
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "test darwin",
		AppName:       "Test App darwin",
	})

	setupRoutes(app)

	app.Listen(":3000")
}
