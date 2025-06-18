package courses

import (
	"github.com/gofiber/fiber/v2"
	courseApp "github.com/leninner/hear-backend/internal/courses/application"
	courseRepo "github.com/leninner/hear-backend/internal/courses/infrastructure"
	"github.com/leninner/hear-backend/internal/infrastructure/db"
)

func Setup(api fiber.Router, db *db.Queries) *courseApp.Handler {
	courses := api.Group("/courses")

	courseRepository := courseRepo.NewPostgresRepository(db)
	courseUseCase := courseApp.NewUseCase(courseRepository)
	courseHandler := courseApp.NewHandler(courseUseCase)
	courseApp.SetupRoutes(courses, courseHandler)

	return courseHandler
} 