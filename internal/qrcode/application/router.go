package application

import "github.com/gofiber/fiber/v2"

func SetupRoutes(router fiber.Router, handler *Handler) {
	router.Post("/", handler.CreateQRCode)
	router.Get("/:id", handler.GetQRCode)
	router.Get("/code/:code", handler.GetQRCodeByCode)
	router.Get("/course-section/:courseSectionId", handler.GetCourseSectionQRCodes)
	router.Get("/course-section/:courseSectionId/active", handler.GetActiveQRCode)
	router.Put("/:id", handler.UpdateQRCode)
	router.Delete("/:id", handler.DeleteQRCode)
} 