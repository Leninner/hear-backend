package application

import "github.com/gofiber/fiber/v2"

func SetupRoutes(router fiber.Router, handler *Handler) {
	// General schedule routes
	router.Post("/schedules", handler.CreateSchedule)
	router.Get("/schedules/:id", handler.GetSchedule)
	router.Put("/schedules/:id", handler.UpdateSchedule)
	router.Delete("/schedules/:id", handler.DeleteSchedule)
	
	// Course-specific schedule routes
	router.Get("/courses/:courseId/schedules", handler.GetSchedulesByCourse)
	
	// Section-specific schedule routes
	router.Get("/sections/:sectionId/schedules", handler.GetSchedulesBySection)
	
	// Classroom-specific schedule routes
	router.Get("/classrooms/:classroomId/schedules", handler.GetSchedulesByClassroom)
} 