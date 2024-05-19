package main

import (
	"github.com/Ninani/go-orm/handlers"
	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	app.Get("/home", handlers.Home)
	app.Post("/fact", handlers.CreateFact)
	app.Put("/fact/:id", handlers.UpdateFact)
	app.Get("/facts", handlers.ListFacts)
}
