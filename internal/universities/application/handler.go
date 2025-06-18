package application

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/leninner/hear-backend/internal/shared/response"
	"github.com/leninner/hear-backend/internal/universities/domain"
)

type Handler struct {
	useCase *UseCase
}

func NewHandler(useCase *UseCase) *Handler {
	return &Handler{
		useCase: useCase,
	}
}

func (h *Handler) CreateUniversity(c *fiber.Ctx) error {
	university := new(domain.CreateUniversityDTO)
	if err := c.BodyParser(university); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid request body format")
	}

	createdUniversity, err := h.useCase.CreateUniversity(university)
	if err != nil {
		switch e := err.(type) {
		case *domain.ValidationErrors:
			return response.Error(c, fiber.StatusBadRequest, e.GetErrors())
		case *domain.DomainError:
			switch e.Type {
			case domain.ErrorTypeConflict:
				return response.Error(c, fiber.StatusConflict, e.Message)
			case domain.ErrorTypeNotFound:
				return response.Error(c, fiber.StatusNotFound, e.Message)
			default:
				fmt.Println(e)
				return response.Error(c, fiber.StatusInternalServerError, "An unexpected error occurred")
			}
		default:
			return response.Error(c, fiber.StatusInternalServerError, "Failed to create university")
		}
	}

	return response.Success(c, "University created successfully", createdUniversity)
}

func (h *Handler) GetUniversity(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, domain.ErrInvalidID.Error())
	}

	university, err := h.useCase.GetUniversity(id)
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
			return response.Error(c, fiber.StatusInternalServerError, "Failed to retrieve university")
		}
	}

	return response.Success(c, "University retrieved successfully", university)
}

func (h *Handler) GetAllUniversities(c *fiber.Ctx) error {
	universities, err := h.useCase.GetAllUniversities()
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to retrieve universities")
	}

	return response.Success(c, "Universities retrieved successfully", universities)
}

func (h *Handler) UpdateUniversity(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, domain.ErrInvalidID.Error())
	}

	university := new(domain.UpdateUniversityDTO)
	if err := c.BodyParser(university); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid request body format")
	}

	if err := h.useCase.UpdateUniversity(id, university); err != nil {
		switch e := err.(type) {
		case *domain.ValidationErrors:
			return response.Error(c, fiber.StatusBadRequest, e.GetErrors())
		case *domain.DomainError:
			switch e.Type {
			case domain.ErrorTypeNotFound:
				return response.Error(c, fiber.StatusNotFound, e.Message)
			default:
				return response.Error(c, fiber.StatusInternalServerError, "An unexpected error occurred")
			}
		default:
			return response.Error(c, fiber.StatusInternalServerError, "Failed to update university")
		}
	}

	return response.Success(c, "University updated successfully", nil)
}

func (h *Handler) DeleteUniversity(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, domain.ErrInvalidID.Error())
	}

	if err := h.useCase.DeleteUniversity(id); err != nil {
		switch e := err.(type) {
		case *domain.DomainError:
			switch e.Type {
			case domain.ErrorTypeNotFound:
				return response.Error(c, fiber.StatusNotFound, e.Message)
			default:
				return response.Error(c, fiber.StatusInternalServerError, "An unexpected error occurred")
			}
		default:
			return response.Error(c, fiber.StatusInternalServerError, "Failed to delete university")
		}
	}

	return response.Success(c, "University deleted successfully", nil)
} 