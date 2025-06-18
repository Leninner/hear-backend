package auth

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/leninner/hear-backend/internal/auth/application"
	"github.com/leninner/hear-backend/internal/auth/infrastructure"
	"github.com/leninner/hear-backend/internal/infrastructure/db"
	userRepo "github.com/leninner/hear-backend/internal/users/infrastructure"
)

func Setup(api fiber.Router, db *db.Queries) *application.Handler {
	auth := api.Group("/auth")

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		jwtSecret = "default-secret-key-change-in-production"
	}

	userRepository := userRepo.NewPostgresRepository(db)
	tokenRepository := infrastructure.NewPostgresTokenRepository(db)
	authUseCase := application.NewUseCase(userRepository, tokenRepository, jwtSecret)
	authHandler := application.NewHandler(authUseCase)
	application.SetupRoutes(auth, authHandler)

	return authHandler
} 