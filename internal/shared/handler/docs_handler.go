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
	app.Get("/swagger/*", swagger.New(swagger.Config{
		URL:         "/docs/openapi.yaml",
		DeepLinking: true,
	}))
}