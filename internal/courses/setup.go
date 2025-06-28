package courses

import (
	"github.com/gofiber/fiber/v2"
	"github.com/leninner/hear-backend/internal/auth/application"
	coursesApp "github.com/leninner/hear-backend/internal/courses/application"
	coursesRepo "github.com/leninner/hear-backend/internal/courses/infrastructure"
	"github.com/leninner/hear-backend/internal/infrastructure/db"
	sharedMiddleware "github.com/leninner/hear-backend/internal/shared/middleware"
)

func Setup(api fiber.Router, db *db.Queries, authHandler *application.Handler) *coursesApp.Handler {
	authMiddleware := sharedMiddleware.NewAuthMiddleware(authHandler.GetUseCase().ValidateAccessToken)
	
	// Admin and teacher routes
	courses := api.Group("/courses", authMiddleware.Authenticate(), authMiddleware.RequireAnyRole("admin", "teacher", "student"))

	coursesRepository := coursesRepo.NewPostgresRepository(db)
	coursesUseCase := coursesApp.NewUseCase(coursesRepository)
	coursesHandler := coursesApp.NewHandler(coursesUseCase)
	coursesApp.SetupRoutes(courses, coursesHandler)

	return coursesHandler
} 