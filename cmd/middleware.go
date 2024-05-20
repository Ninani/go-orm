package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func setupMiddleware(app *fiber.App) {
	app.Use(logger.New(logger.Config{
		Format:     "${pid} ${status} - ${method} ${path}\n",
		TimeFormat: "02-Jan-2006",
		TimeZone:   "UTC",
	}))

	app.Use(recover.New())

	app.Use(limiter.New())

	// app.Use(cache.New(cache.Config{
	// 	Next: func(c *fiber.Ctx) bool {
	// 		return c.Query("noCache") == "true"
	// 	},
	// 	Expiration:   5 * time.Minute,
	// 	CacheControl: true,
	// }))
}
