package di

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
	docsApp "github.com/leninner/hear-backend/internal/shared/handler"
	users "github.com/leninner/hear-backend/internal/users"
	userApp "github.com/leninner/hear-backend/internal/users/application"
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
	
	// SHARED MODULES
	docsHandler := docsApp.NewDocsHandler()
	docsHandler.Register(app)

	api := app.Group("/api")
	
	// APP MODULES
	userHandler := users.Setup(api, db)

	return &Container{
		App:            app,
		UserHandler:    userHandler,
		DocsHandler:    docsHandler,
	}
} 