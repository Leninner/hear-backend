package application

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/leninner/hear-backend/internal/attendance/domain"
)

type Handler struct {
	useCase *UseCase
}

func NewHandler(useCase *UseCase) *Handler {
	return &Handler{
		useCase: useCase,
	}
}

func (h *Handler) CreateAttendance(c *fiber.Ctx) error {
	var input struct {
		StudentID       uuid.UUID           `json:"studentId"`
		ClassScheduleID uuid.UUID           `json:"classScheduleId"`
		Status          domain.AttendanceStatus `json:"status"`
		Date            time.Time           `json:"date"`
	}

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid input",
		})
	}

	attendance, err := h.useCase.CreateAttendance(input.StudentID, input.ClassScheduleID, input.Status, input.Date)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create attendance",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(attendance)
}

func (h *Handler) GetAttendance(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid attendance ID",
		})
	}

	attendance, err := h.useCase.GetAttendance(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Attendance not found",
		})
	}

	return c.JSON(attendance)
}

func (h *Handler) GetStudentAttendance(c *fiber.Ctx) error {
	studentID, err := uuid.Parse(c.Params("studentId"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid student ID",
		})
	}

	attendances, err := h.useCase.GetStudentAttendance(studentID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch student attendance",
		})
	}

	return c.JSON(attendances)
}

func (h *Handler) GetClassScheduleAttendance(c *fiber.Ctx) error {
	classScheduleID, err := uuid.Parse(c.Params("classScheduleId"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid class schedule ID",
		})
	}

	attendances, err := h.useCase.GetClassScheduleAttendance(classScheduleID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch class schedule attendance",
		})
	}

	return c.JSON(attendances)
}

func (h *Handler) UpdateAttendance(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid attendance ID",
		})
	}

	var attendance domain.Attendance
	if err := c.BodyParser(&attendance); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid input",
		})
	}

	attendance.ID = id
	if err := h.useCase.UpdateAttendance(&attendance); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update attendance",
		})
	}

	return c.JSON(attendance)
}

func (h *Handler) DeleteAttendance(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid attendance ID",
		})
	}

	if err := h.useCase.DeleteAttendance(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete attendance",
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
} 