package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"time"
	"tour-api/database"
	"tour-api/router"
)

func main() {
	app := fiber.New()

	app.Use(cache.New(cache.Config{
		Next: func(c *fiber.Ctx) bool {
			return c.Query("refresh") == "true"
		},
		Expiration: 6 * time.Second,
	}))

	database.ConnectDB()

	router.SetupRoutes(app)

	app.Listen(":3000")
}
