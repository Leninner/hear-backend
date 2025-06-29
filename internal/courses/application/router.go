package application

import "github.com/gofiber/fiber/v2"

func SetupRoutes(router fiber.Router, handler *Handler) {
	// Course routes
	router.Post("/", handler.CreateCourse)
	router.Get("/", handler.GetAllCourses)
	router.Get("/:id", handler.GetCourse)
	router.Put("/:id", handler.UpdateCourse)
	router.Delete("/:id", handler.DeleteCourse)
	
	// Section routes
	router.Post("/sections", handler.CreateSection)
	router.Get("/sections/:id", handler.GetSection)
	router.Get("/:courseId/sections", handler.GetSectionsByCourse)
	router.Get("/teacher/:teacherId/sections", handler.GetSectionsByTeacher)
	router.Get("/student/:studentId/sections", handler.GetSectionsWithSchedulesByStudent)
	router.Put("/sections/:id", handler.UpdateSection)
	router.Delete("/sections/:id", handler.DeleteSection)
	router.Post("/sections/:id/enroll", handler.EnrollInSection)
	router.Get("/sections/:id/enrollments", handler.GetEnrollmentsWithDetailsBySection)
} 