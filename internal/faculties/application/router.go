package application

import "github.com/gofiber/fiber/v2"

func SetupRoutes(router fiber.Router, handler *Handler) {
	router.Post("/", handler.CreateFaculty)
	router.Get("/", handler.GetAllFaculties)
	router.Get("/university/:universityId", handler.GetFacultiesByUniversity)
	router.Get("/:id", handler.GetFaculty)
	router.Put("/:id", handler.UpdateFaculty)
	router.Delete("/:id", handler.DeleteFaculty)
}

func SetupPublicRoutes(router fiber.Router, handler *Handler) {
	router.Get("/", handler.GetAllFaculties)
	router.Get("/university/:universityId", handler.GetFacultiesByUniversity)
	router.Get("/:id", handler.GetFaculty)
}

func SetupAdminRoutes(router fiber.Router, handler *Handler) {
	router.Post("/", handler.CreateFaculty)
	router.Put("/:id", handler.UpdateFaculty)
	router.Delete("/:id", handler.DeleteFaculty)
} 