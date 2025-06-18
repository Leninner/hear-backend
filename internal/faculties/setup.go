package faculties

import (
	"github.com/gofiber/fiber/v2"
	"github.com/leninner/hear-backend/internal/auth/application"
	facultiesApp "github.com/leninner/hear-backend/internal/faculties/application"
	facultiesRepo "github.com/leninner/hear-backend/internal/faculties/infrastructure"
	"github.com/leninner/hear-backend/internal/infrastructure/db"
	sharedMiddleware "github.com/leninner/hear-backend/internal/shared/middleware"
)

func Setup(api fiber.Router, db *db.Queries, authHandler *application.Handler) *facultiesApp.Handler {
	authMiddleware := sharedMiddleware.NewAuthMiddleware(authHandler.GetUseCase().ValidateAccessToken)
	
	// Public routes (GET) - accessible to all authenticated users
	publicRoutes := api.Group("/faculties", authMiddleware.Authenticate())
	
	// Admin routes (POST, PUT, DELETE) - accessible only to admins
	adminRoutes := api.Group("/admin/faculties", authMiddleware.Authenticate(), authMiddleware.RequireAdmin())

	facultiesRepository := facultiesRepo.NewPostgresRepository(db)
	facultiesUseCase := facultiesApp.NewUseCase(facultiesRepository)
	facultiesHandler := facultiesApp.NewHandler(facultiesUseCase)
	
	// Setup routes for both groups
	facultiesApp.SetupPublicRoutes(publicRoutes, facultiesHandler)
	facultiesApp.SetupAdminRoutes(adminRoutes, facultiesHandler)

	return facultiesHandler
} 