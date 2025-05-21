package application

import "github.com/gofiber/fiber/v2"

func SetupRoutes(app *fiber.App, handler *Handler) {
	users := app.Group("/api/users")
	users.Post("/", handler.Create)
	users.Get("/:id", handler.GetByID)
	users.Get("/", handler.GetAll)
} 