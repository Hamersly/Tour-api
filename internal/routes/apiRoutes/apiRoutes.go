package apiRoutes

import (
	"github.com/gofiber/fiber/v2"
	"tour-api/internal/handlers/apiHandlers"
)

func SetupApiRoutes(router fiber.Router) {

	api := router.Group("/api")

	api.Post("/", apiHandlers.CreateRoute)
	api.Get("/", apiHandlers.GetRoutes)
	api.Get("/:noteId", apiHandlers.GetRoute)
	api.Put("/:id", apiHandlers.UpdateRoute)
	api.Delete("/:id", apiHandlers.DeleteRoute)
}
