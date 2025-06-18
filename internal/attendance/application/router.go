package application

import "github.com/gofiber/fiber/v2"

func SetupRoutes(router fiber.Router, handler *Handler) {
	router.Post("/", handler.CreateAttendance)
	router.Get("/:id", handler.GetAttendance)
	router.Get("/student/:studentId", handler.GetStudentAttendance)
	router.Get("/class-schedule/:classScheduleId", handler.GetClassScheduleAttendance)
	router.Put("/:id", handler.UpdateAttendance)
	router.Delete("/:id", handler.DeleteAttendance)
} 