package application

import "github.com/gofiber/fiber/v2"

func SetupRoutes(api fiber.Router, handler *Handler) {
	users := api.Group("/users")

	users.Post("/", handler.Create)
	users.Get("/:id", handler.GetByID)
	users.Get("/", handler.GetAll)
} 