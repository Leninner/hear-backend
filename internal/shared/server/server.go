package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/leninner/hear-backend/internal/di"
	"github.com/leninner/hear-backend/internal/shared/middleware"
	userApp "github.com/leninner/hear-backend/internal/users/application"
)

func SetupServer(container *di.Container) {
	app := container.App

	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(cors.New())

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "ok",
		})
	})

	app.Get("/docs/openapi.yaml", func(c *fiber.Ctx) error {
		return c.SendFile("docs/openapi.yaml")
	})

	container.DocsHandler.Register(app)
	userApp.SetupRoutes(app, container.UserHandler)

	middleware.EndpointsLogger(app)
}