package di

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
	docsApp "github.com/leninner/hear-backend/internal/shared/handler"
	userApp "github.com/leninner/hear-backend/internal/users/application"
	userRepo "github.com/leninner/hear-backend/internal/users/infrastructure"
)

type Container struct {
	App            *fiber.App
	UserHandler    *userApp.Handler
	DocsHandler    *docsApp.DocsHandler
}

func NewContainer(db *sql.DB) *Container {
	app := fiber.New(fiber.Config{
		AppName: "Hear Backend",
	})

	api := app.Group("/api")

	userRepository := userRepo.NewPostgresRepository(db)
	userUseCase := userApp.NewUseCase(userRepository)
	userHandler := userApp.NewHandler(userUseCase)
	userApp.SetupRoutes(api, userHandler)

	docsHandler := docsApp.NewDocsHandler()
	docsHandler.Register(app)


	return &Container{
		App:            app,
		UserHandler:    userHandler,
		DocsHandler:    docsHandler,
	}
} 