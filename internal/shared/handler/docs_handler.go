package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

type DocsHandler struct{}

func NewDocsHandler() *DocsHandler {
	return &DocsHandler{}
}

func (h *DocsHandler) Register(app *fiber.App) {
	app.Get("/docs/openapi.yaml", func(c *fiber.Ctx) error {
		return c.SendFile("docs/openapi.yaml")
	})

	app.Get("/documentation/*", swagger.New(swagger.Config{
		URL:         "/docs/openapi.yaml",
		DeepLinking: true,
		Title:       "Hear API",
	}))
}