package main

import (
	"github.com/Ninani/go-orm/database"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.ConnectDb()
	app := fiber.New()

	setupMiddleware(app)
	setupRoutes(app)

	app.Listen(":3000")
}
