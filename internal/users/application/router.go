package application

import "github.com/gofiber/fiber/v2"

func SetupRoutes(api fiber.Router, handler *Handler) {
	api.Post("/", handler.Create)
	api.Get("/:id", handler.GetByID)
	api.Get("/", handler.GetAll)
} 