package router

import (
	"github.com/gofiber/fiber/v2"
	"tour-api/internal/routes/apiRoutes"
)

func SetupRoutes(app *fiber.App) {
	mainGroup := app.Group("/")
	apiRoutes.SetupApiRoutes(mainGroup)
}
