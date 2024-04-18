package apiRoutes

import (
	"github.com/gofiber/fiber/v2"
	"tour-api/internal/handlers/apiHandlers"
)

func SetupApiRoutes(router fiber.Router) {

	api := router.Group("/api")

	api.Post("/", apiHandlers.CreateNotes)
	api.Get("/", apiHandlers.GetNotes)
	api.Get("/:noteId", apiHandlers.GetNote)
	api.Put("/:id", apiHandlers.UpdateNote)
	api.Delete("/:id", apiHandlers.DeleteNote)
}
