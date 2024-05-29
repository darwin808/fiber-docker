package main

import (
	"github.com/darwin808/pg-fiber/handlers"
	"github.com/gofiber/fiber/v3"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", handlers.Home)
	app.Get("/fact", handlers.ListFacts)
	app.Post("/fact", handlers.CreateFact)
}
