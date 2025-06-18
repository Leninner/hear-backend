package application

import "github.com/gofiber/fiber/v2"

func SetupRoutes(api fiber.Router, handler *Handler) {
	api.Post("/login", handler.Login)
	api.Post("/register", handler.Register)
	api.Post("/refresh", handler.RefreshToken)
	api.Post("/logout", handler.Logout)
} 