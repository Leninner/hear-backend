package users

import (
	"github.com/gofiber/fiber/v2"
	"github.com/leninner/hear-backend/internal/infrastructure/db"
	userApp "github.com/leninner/hear-backend/internal/users/application"
	userRepo "github.com/leninner/hear-backend/internal/users/infrastructure"
)

func Setup(api fiber.Router, db *db.Queries) *userApp.Handler {
	users := api.Group("/users")

	userRepository := userRepo.NewPostgresRepository(db)
	userUseCase := userApp.NewUseCase(userRepository)
	userHandler := userApp.NewHandler(userUseCase)
	userApp.SetupRoutes(users, userHandler)

	return userHandler
}
