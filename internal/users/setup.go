package users

import (
	"github.com/gofiber/fiber/v2"
	"github.com/leninner/hear-backend/internal/auth/application"
	"github.com/leninner/hear-backend/internal/infrastructure/db"
	sharedMiddleware "github.com/leninner/hear-backend/internal/shared/middleware"
	userApp "github.com/leninner/hear-backend/internal/users/application"
	userRepo "github.com/leninner/hear-backend/internal/users/infrastructure"
)

func Setup(api fiber.Router, db *db.Queries, authHandler *application.Handler) *userApp.Handler {
	authMiddleware := sharedMiddleware.NewAuthMiddleware(authHandler.GetUseCase().ValidateAccessToken)
	
	// Admin routes only
	adminRoutes := api.Group("/admin", authMiddleware.Authenticate(), authMiddleware.RequireAdmin())
	users := adminRoutes.Group("/users")

	userRepository := userRepo.NewPostgresRepository(db)
	userUseCase := userApp.NewUseCase(userRepository)
	userHandler := userApp.NewHandler(userUseCase)
	userApp.SetupRoutes(users, userHandler)

	return userHandler
}
