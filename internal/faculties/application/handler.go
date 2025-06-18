package application

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/leninner/hear-backend/internal/faculties/domain"
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

func (h *Handler) CreateFaculty(c *fiber.Ctx) error {
	faculty := new(domain.CreateFacultyDTO)
	if err := c.BodyParser(faculty); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid request body format")
	}

	createdFaculty, err := h.useCase.CreateFaculty(faculty)
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
			return response.Error(c, fiber.StatusInternalServerError, "Failed to create faculty")
		}
	}

	return response.Success(c, "Faculty created successfully", createdFaculty)
}

func (h *Handler) GetFaculty(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, domain.ErrInvalidID.Error())
	}

	faculty, err := h.useCase.GetFaculty(id)
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
			return response.Error(c, fiber.StatusInternalServerError, "Failed to retrieve faculty")
		}
	}

	return response.Success(c, "Faculty retrieved successfully", faculty)
}

func (h *Handler) GetAllFaculties(c *fiber.Ctx) error {
	faculties, err := h.useCase.GetAllFaculties()
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to retrieve faculties")
	}

	return response.Success(c, "Faculties retrieved successfully", faculties)
}

func (h *Handler) GetFacultiesByUniversity(c *fiber.Ctx) error {
	universityID, err := uuid.Parse(c.Params("universityId"))
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid university ID format")
	}

	faculties, err := h.useCase.GetFacultiesByUniversity(universityID)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to retrieve faculties")
	}

	return response.Success(c, "Faculties retrieved successfully", faculties)
}

func (h *Handler) UpdateFaculty(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, domain.ErrInvalidID.Error())
	}

	faculty := new(domain.UpdateFacultyDTO)
	if err := c.BodyParser(faculty); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid request body format")
	}

	if err := h.useCase.UpdateFaculty(id, faculty); err != nil {
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
			return response.Error(c, fiber.StatusInternalServerError, "Failed to update faculty")
		}
	}

	return response.Success(c, "Faculty updated successfully", nil)
}

func (h *Handler) DeleteFaculty(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, domain.ErrInvalidID.Error())
	}

	if err := h.useCase.DeleteFaculty(id); err != nil {
		switch e := err.(type) {
		case *domain.DomainError:
			switch e.Type {
			case domain.ErrorTypeNotFound:
				return response.Error(c, fiber.StatusNotFound, e.Message)
			default:
				return response.Error(c, fiber.StatusInternalServerError, "An unexpected error occurred")
			}
		default:
			return response.Error(c, fiber.StatusInternalServerError, "Failed to delete faculty")
		}
	}

	return response.Success(c, "Faculty deleted successfully", nil)
} 