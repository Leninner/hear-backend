package application

import (
	"time"

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
	var input struct {
		CourseID  uuid.UUID `json:"courseId"`
		Code      string    `json:"code"`
		ExpiresAt time.Time `json:"expiresAt"`
	}

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid input",
		})
	}

	qrcode, err := h.useCase.CreateQRCode(input.CourseID, input.Code, input.ExpiresAt)
	if err != nil {
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

func (h *Handler) GetCourseQRCodes(c *fiber.Ctx) error {
	courseID, err := uuid.Parse(c.Params("courseId"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid course ID",
		})
	}

	qrcodes, err := h.useCase.GetCourseQRCodes(courseID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch course QR codes",
		})
	}

	return c.JSON(qrcodes)
}

func (h *Handler) GetActiveQRCode(c *fiber.Ctx) error {
	courseID, err := uuid.Parse(c.Params("courseId"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid course ID",
		})
	}

	qrcode, err := h.useCase.GetActiveQRCode(courseID)
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

	var qrcode domain.QRCode
	if err := c.BodyParser(&qrcode); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid input",
		})
	}

	qrcode.ID = id
	if err := h.useCase.UpdateQRCode(&qrcode); err != nil {
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