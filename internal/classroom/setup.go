package classroom

import (
	"github.com/gofiber/fiber/v2"
	"github.com/leninner/hear-backend/internal/auth/application"
	classroomApp "github.com/leninner/hear-backend/internal/classroom/application"
	classroomRepo "github.com/leninner/hear-backend/internal/classroom/infrastructure"
	"github.com/leninner/hear-backend/internal/infrastructure/db"
	sharedMiddleware "github.com/leninner/hear-backend/internal/shared/middleware"
)

func Setup(api fiber.Router, db *db.Queries, authHandler *application.Handler) *classroomApp.Handler {
	authMiddleware := sharedMiddleware.NewAuthMiddleware(authHandler.GetUseCase().ValidateAccessToken)
	
	// Admin and teacher routes
	classrooms := api.Group("/classrooms", authMiddleware.Authenticate(), authMiddleware.RequireTeacherOrAdmin())

	classroomRepository := classroomRepo.NewPostgresRepository(db)
	classroomUseCase := classroomApp.NewUseCase(classroomRepository)
	classroomHandler := classroomApp.NewHandler(classroomUseCase)
	classroomApp.SetupRoutes(classrooms, classroomHandler)

	return classroomHandler
} 