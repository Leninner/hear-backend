package application

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/leninner/hear-backend/internal/qrcode/domain"
)

type Handler struct {
	useCase *UseCase
}

func NewHandler(useCase *UseCase) *Handler {
	return &Handler{
		useCase: useCase,
	}
}

func (h *Handler) CreateQRCode(c *fiber.Ctx) error {
	var dto domain.CreateQRCodeDTO

	if err := c.BodyParser(&dto); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid input",
		})
	}

	qrcode, err := h.useCase.CreateQRCode(&dto)
	if err != nil {
		if validationErrors, ok := err.(*domain.ValidationErrors); ok {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":   "Validation failed",
				"details": validationErrors.GetErrors(),
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create QR code",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(qrcode)
}

func (h *Handler) GetQRCode(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid QR code ID",
		})
	}

	qrcode, err := h.useCase.GetQRCode(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "QR code not found",
		})
	}

	return c.JSON(qrcode)
}

func (h *Handler) GetQRCodeByCode(c *fiber.Ctx) error {
	code := c.Params("code")
	qrcode, err := h.useCase.GetQRCodeByCode(code)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "QR code not found",
		})
	}

	return c.JSON(qrcode)
}

func (h *Handler) GetCourseSectionQRCodes(c *fiber.Ctx) error {
	courseSectionID, err := uuid.Parse(c.Params("courseSectionId"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid course section ID",
		})
	}

	qrcodes, err := h.useCase.GetCourseSectionQRCodes(courseSectionID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch course section QR codes",
		})
	}

	return c.JSON(qrcodes)
}

func (h *Handler) GetActiveQRCode(c *fiber.Ctx) error {
	courseSectionID, err := uuid.Parse(c.Params("courseSectionId"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid course section ID",
		})
	}

	qrcode, err := h.useCase.GetActiveQRCode(courseSectionID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "No active QR code found",
		})
	}

	return c.JSON(qrcode)
}

func (h *Handler) UpdateQRCode(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid QR code ID",
		})
	}

	var dto domain.UpdateQRCodeDTO
	if err := c.BodyParser(&dto); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid input",
		})
	}

	qrcode, err := h.useCase.UpdateQRCode(id, &dto)
	if err != nil {
		if validationErrors, ok := err.(*domain.ValidationErrors); ok {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":   "Validation failed",
				"details": validationErrors.GetErrors(),
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update QR code",
		})
	}

	return c.JSON(qrcode)
}

func (h *Handler) DeleteQRCode(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid QR code ID",
		})
	}

	if err := h.useCase.DeleteQRCode(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete QR code",
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
} 