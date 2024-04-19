package apiHandlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"tour-api/database"
	"tour-api/internal/model"
)

func GetRoutes(c *fiber.Ctx) error {
	db := database.DB
	var routes []model.Route

	// find all notes in the database
	db.Find(&routes)

	// If no noteHandler is present return an error
	if len(routes) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No notes present", "data": nil})
	}

	// Else return notes
	return c.JSON(fiber.Map{"status": "success", "message": "Notes Found", "data": routes})
}

func CreateRoute(c *fiber.Ctx) error {
	db := database.DB
	route := new(model.Route)

	// Store the body in the noteHandler and return error if encountered
	err := c.BodyParser(route)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}
	route.ID = uuid.New()
	// Create the Note and return error if encountered
	err = db.Create(&route).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create noteHandler", "data": err})
	}

	// Return the created noteHandler
	return c.JSON(fiber.Map{"status": "success", "message": "Created Note", "data": route})
}

func GetRoute(c *fiber.Ctx) error {
	db := database.DB
	var route model.Route

	// Read the param noteId
	id := c.Params("noteId")
	db.Find(&route, "id = ?", id)

	// If no such noteHandler present return an error
	if route.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No noteHandler present", "data": nil})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Notes Found", "data": route})
}

func UpdateRoute(c *fiber.Ctx) error {
	type updateRoute struct {
		Title    string  `json:"title"`
		SubTitle string  `json:"sub_title"`
		Text     string  `json:"Text"`
		Distance float32 `json:"Distance"`
	}
	db := database.DB
	var route model.Route

	// Read the param noteId
	id := c.Params("noteId")

	db.Find(&route, "id = ?", id)

	// If no such noteHandler present return an error
	if route.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No noteHandler present", "data": nil})
	}

	// Store the body containing the updated data and return error if encountered
	var updateNoteData updateRoute
	err := c.BodyParser(&updateNoteData)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	// Edit the noteHandler
	route.Title = updateNoteData.Title
	route.SubTitle = updateNoteData.SubTitle
	route.Text = updateNoteData.Text
	route.Distance = updateNoteData.Distance

	// Save the Changes
	db.Save(&route)

	// Return the updated noteHandler
	return c.JSON(fiber.Map{"status": "success", "message": "Notes Found", "data": route})
}

func DeleteRoute(c *fiber.Ctx) error {
	db := database.DB
	var route model.Route

	// Read the param noteId
	id := c.Params("noteId")

	db.Find(&route, "id = ?", id)

	// If no such noteHandler present return an error
	if route.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No noteHandler present", "data": nil})
	}

	// Delete the noteHandler and return error if encountered
	err := db.Delete(&route, "id = ?", id).Error

	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Failed to delete noteHandler", "data": nil})
	}

	// Return success message
	return c.JSON(fiber.Map{"status": "success", "message": "Deleted Note"})
}
