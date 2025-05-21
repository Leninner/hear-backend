package users

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
	userApp "github.com/leninner/hear-backend/internal/users/application"
	userRepo "github.com/leninner/hear-backend/internal/users/infrastructure"
)

func Setup(api fiber.Router, db *sql.DB) *userApp.Handler {
	users := api.Group("/users")

	userRepository := userRepo.NewPostgresRepository(db)
	userUseCase := userApp.NewUseCase(userRepository)
	userHandler := userApp.NewHandler(userUseCase)
	userApp.SetupRoutes(users, userHandler)

	return userHandler
}
