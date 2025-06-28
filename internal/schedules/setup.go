package schedules

import (
	"github.com/gofiber/fiber/v2"
	"github.com/leninner/hear-backend/internal/auth/application"
	coursesRepo "github.com/leninner/hear-backend/internal/courses/infrastructure"
	"github.com/leninner/hear-backend/internal/infrastructure/db"
	schedulesApp "github.com/leninner/hear-backend/internal/schedules/application"
	schedulesRepo "github.com/leninner/hear-backend/internal/schedules/infrastructure"
	sharedMiddleware "github.com/leninner/hear-backend/internal/shared/middleware"
)

func Setup(api fiber.Router, db *db.Queries, authHandler *application.Handler) *schedulesApp.Handler {
	authMiddleware := sharedMiddleware.NewAuthMiddleware(authHandler.GetUseCase().ValidateAccessToken)
	
	// Admin and teacher routes
	schedules := api.Group("/", authMiddleware.Authenticate(), authMiddleware.RequireAnyRole("admin", "teacher", "student"))

	schedulesRepository := schedulesRepo.NewPostgresRepository(db)
	coursesRepository := coursesRepo.NewPostgresRepository(db)
	schedulesUseCase := schedulesApp.NewUseCase(schedulesRepository, coursesRepository)
	schedulesHandler := schedulesApp.NewHandler(schedulesUseCase)
	schedulesApp.SetupRoutes(schedules, schedulesHandler)

	return schedulesHandler
} 