package application

import "github.com/gofiber/fiber/v2"

func SetupRoutes(router fiber.Router, handler *Handler) {
	router.Post("/", handler.CreateCourse)
	router.Get("/", handler.GetAllCourses)
	router.Get("/:id", handler.GetCourse)
	router.Put("/:id", handler.UpdateCourse)
	router.Delete("/:id", handler.DeleteCourse)
} 