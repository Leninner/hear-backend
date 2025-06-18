package application

import "github.com/gofiber/fiber/v2"

func SetupRoutes(router fiber.Router, handler *Handler) {
	router.Post("/", handler.CreateUniversity)
	router.Get("/", handler.GetAllUniversities)
	router.Get("/:id", handler.GetUniversity)
	router.Put("/:id", handler.UpdateUniversity)
	router.Delete("/:id", handler.DeleteUniversity)
}

func SetupPublicRoutes(router fiber.Router, handler *Handler) {
	router.Get("/", handler.GetAllUniversities)
	router.Get("/:id", handler.GetUniversity)
}

func SetupAdminRoutes(router fiber.Router, handler *Handler) {
	router.Post("/", handler.CreateUniversity)
	router.Put("/:id", handler.UpdateUniversity)
	router.Delete("/:id", handler.DeleteUniversity)
} 