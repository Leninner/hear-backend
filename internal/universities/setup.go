package universities

import (
	"github.com/gofiber/fiber/v2"
	"github.com/leninner/hear-backend/internal/auth/application"
	"github.com/leninner/hear-backend/internal/infrastructure/db"
	sharedMiddleware "github.com/leninner/hear-backend/internal/shared/middleware"
	universitiesApp "github.com/leninner/hear-backend/internal/universities/application"
	universitiesRepo "github.com/leninner/hear-backend/internal/universities/infrastructure"
)

func Setup(api fiber.Router, db *db.Queries, authHandler *application.Handler) *universitiesApp.Handler {
	authMiddleware := sharedMiddleware.NewAuthMiddleware(authHandler.GetUseCase().ValidateAccessToken)
	
	// Public routes (GET) - accessible to all authenticated users
	publicRoutes := api.Group("/universities", authMiddleware.Authenticate())
	
	// Admin routes (POST, PUT, DELETE) - accessible only to admins
	adminRoutes := api.Group("/admin/universities", authMiddleware.Authenticate(), authMiddleware.RequireAdmin())

	universitiesRepository := universitiesRepo.NewPostgresRepository(db)
	universitiesUseCase := universitiesApp.NewUseCase(universitiesRepository)
	universitiesHandler := universitiesApp.NewHandler(universitiesUseCase)
	
	// Setup routes for both groups
	universitiesApp.SetupPublicRoutes(publicRoutes, universitiesHandler)
	universitiesApp.SetupAdminRoutes(adminRoutes, universitiesHandler)

	return universitiesHandler
} 