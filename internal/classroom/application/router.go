package application

import "github.com/gofiber/fiber/v2"

func SetupRoutes(router fiber.Router, handler *Handler) {
	router.Post("/", handler.CreateClassroom)
	router.Get("/", handler.GetAllClassrooms)
	router.Get("/:id", handler.GetClassroom)
	router.Get("/location/:location", handler.GetClassroomsByLocation)
	router.Put("/:id", handler.UpdateClassroom)
	router.Delete("/:id", handler.DeleteClassroom)
} 