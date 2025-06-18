package application

import "github.com/gofiber/fiber/v2"

func SetupRoutes(router fiber.Router, handler *Handler) {
	router.Post("/", handler.CreateQRCode)
	router.Get("/:id", handler.GetQRCode)
	router.Get("/code/:code", handler.GetQRCodeByCode)
	router.Get("/course/:courseId", handler.GetCourseQRCodes)
	router.Get("/course/:courseId/active", handler.GetActiveQRCode)
	router.Put("/:id", handler.UpdateQRCode)
	router.Delete("/:id", handler.DeleteQRCode)
} 