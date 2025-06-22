package application

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/leninner/hear-backend/internal/schedules/domain"
	"github.com/leninner/hear-backend/internal/shared/response"
)

type Handler struct {
	useCase *UseCase
}

func NewHandler(useCase *UseCase) *Handler {
	return &Handler{
		useCase: useCase,
	}
}

func (h *Handler) CreateSchedule(c *fiber.Ctx) error {
	dto := new(domain.CreateScheduleDTO)
	if err := c.BodyParser(dto); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid request body format")
	}

	schedule, err := h.useCase.CreateSchedule(dto)
	if err != nil {
		switch e := err.(type) {
		case *domain.ValidationErrors:
			return response.Error(c, fiber.StatusBadRequest, e.GetErrors())
		case *domain.DomainError:
			switch e.Type {
			case domain.ErrorTypeValidation:
				return response.Error(c, fiber.StatusBadRequest, e.Message)
			default:
				return response.Error(c, fiber.StatusInternalServerError, "An unexpected error occurred")
			}
		default:
			return response.Error(c, fiber.StatusInternalServerError, "Failed to create schedule")
		}
	}

	return response.Success(c, "Schedule created successfully", schedule)
}

func (h *Handler) GetSchedule(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, domain.ErrInvalidID.Error())
	}

	schedule, err := h.useCase.GetSchedule(id)
	if err != nil {
		switch e := err.(type) {
		case *domain.DomainError:
			switch e.Type {
			case domain.ErrorTypeNotFound:
				return response.Error(c, fiber.StatusNotFound, e.Message)
			default:
				return response.Error(c, fiber.StatusInternalServerError, "An unexpected error occurred")
			}
		default:
			return response.Error(c, fiber.StatusInternalServerError, "Failed to retrieve schedule")
		}
	}

	return response.Success(c, "Schedule retrieved successfully", schedule)
}

func (h *Handler) GetSchedulesByCourse(c *fiber.Ctx) error {
	courseID, err := uuid.Parse(c.Params("courseId"))
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid course ID format")
	}

	schedules, err := h.useCase.GetSchedulesByCourse(courseID)
	if err != nil {
		switch e := err.(type) {
		case *domain.DomainError:
			switch e.Type {
			case domain.ErrorTypeNotFound:
				return response.Error(c, fiber.StatusNotFound, e.Message)
			default:
				return response.Error(c, fiber.StatusInternalServerError, "An unexpected error occurred")
			}
		default:
			return response.Error(c, fiber.StatusInternalServerError, "Failed to retrieve course schedules")
		}
	}

	return response.Success(c, "Course schedules retrieved successfully", schedules)
}

func (h *Handler) GetSchedulesBySection(c *fiber.Ctx) error {
	sectionID, err := uuid.Parse(c.Params("sectionId"))
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid section ID format")
	}

	schedules, err := h.useCase.GetSchedulesBySection(sectionID)
	if err != nil {
		switch e := err.(type) {
		case *domain.DomainError:
			switch e.Type {
			case domain.ErrorTypeNotFound:
				return response.Error(c, fiber.StatusNotFound, e.Message)
			default:
				return response.Error(c, fiber.StatusInternalServerError, "An unexpected error occurred")
			}
		default:
			return response.Error(c, fiber.StatusInternalServerError, "Failed to retrieve section schedules")
		}
	}

	return response.Success(c, "Section schedules retrieved successfully", schedules)
}

func (h *Handler) GetSchedulesByClassroom(c *fiber.Ctx) error {
	classroomID, err := uuid.Parse(c.Params("classroomId"))
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid classroom ID format")
	}

	schedules, err := h.useCase.GetSchedulesByClassroom(classroomID)
	if err != nil {
		switch e := err.(type) {
		case *domain.DomainError:
			switch e.Type {
			case domain.ErrorTypeNotFound:
				return response.Error(c, fiber.StatusNotFound, e.Message)
			default:
				return response.Error(c, fiber.StatusInternalServerError, "An unexpected error occurred")
			}
		default:
			return response.Error(c, fiber.StatusInternalServerError, "Failed to retrieve classroom schedules")
		}
	}

	return response.Success(c, "Classroom schedules retrieved successfully", schedules)
}

func (h *Handler) UpdateSchedule(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, domain.ErrInvalidID.Error())
	}

	dto := new(domain.UpdateScheduleDTO)
	if err := c.BodyParser(dto); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid request body format")
	}

	if err := h.useCase.UpdateSchedule(id, dto); err != nil {
		switch e := err.(type) {
		case *domain.ValidationErrors:
			return response.Error(c, fiber.StatusBadRequest, e.GetErrors())
		case *domain.DomainError:
			switch e.Type {
			case domain.ErrorTypeValidation:
				return response.Error(c, fiber.StatusBadRequest, e.Message)
			case domain.ErrorTypeNotFound:
				return response.Error(c, fiber.StatusNotFound, e.Message)
			default:
				return response.Error(c, fiber.StatusInternalServerError, "An unexpected error occurred")
			}
		default:
			return response.Error(c, fiber.StatusInternalServerError, "Failed to update schedule")
		}
	}

	return response.Success(c, "Schedule updated successfully", nil)
}

func (h *Handler) DeleteSchedule(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, domain.ErrInvalidID.Error())
	}

	if err := h.useCase.DeleteSchedule(id); err != nil {
		switch e := err.(type) {
		case *domain.DomainError:
			switch e.Type {
			case domain.ErrorTypeNotFound:
				return response.Error(c, fiber.StatusNotFound, e.Message)
			default:
				return response.Error(c, fiber.StatusInternalServerError, "An unexpected error occurred")
			}
		default:
			return response.Error(c, fiber.StatusInternalServerError, "Failed to delete schedule")
		}
	}

	return response.Success(c, "Schedule deleted successfully", nil)
} 