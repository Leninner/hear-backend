package application

import "github.com/gofiber/fiber/v2"

func SetupRoutes(router fiber.Router, handler *Handler) {
	router.Post("/", handler.CreateClassroom)
	router.Get("/", handler.GetAllClassrooms)
	router.Get("/faculty/:facultyId", handler.GetClassroomsByFaculty)
	router.Get("/location", handler.GetClassroomsByLocation)
	router.Get("/:id", handler.GetClassroom)
	router.Put("/:id", handler.UpdateClassroom)
	router.Delete("/:id", handler.DeleteClassroom)
} 