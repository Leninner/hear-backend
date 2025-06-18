package application

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/leninner/hear-backend/internal/classroom/domain"
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

func (h *Handler) CreateClassroom(c *fiber.Ctx) error {
	classroom := new(domain.CreateClassroomDTO)
	if err := c.BodyParser(classroom); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid request body format")
	}

	createdClassroom, err := h.useCase.CreateClassroom(classroom)
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
				return response.Error(c, fiber.StatusInternalServerError, "An unexpected error occurred")
			}
		default:
			return response.Error(c, fiber.StatusInternalServerError, "Failed to create classroom")
		}
	}

	return response.Success(c, "Classroom created successfully", createdClassroom)
}

func (h *Handler) GetClassroom(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, domain.ErrInvalidID.Error())
	}

	classroom, err := h.useCase.GetClassroom(id)
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
			return response.Error(c, fiber.StatusInternalServerError, "Failed to retrieve classroom")
		}
	}

	return response.Success(c, "Classroom retrieved successfully", classroom)
}

func (h *Handler) GetAllClassrooms(c *fiber.Ctx) error {
	classrooms, err := h.useCase.GetAllClassrooms()
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to retrieve classrooms")
	}

	return response.Success(c, "Classrooms retrieved successfully", classrooms)
}

func (h *Handler) GetClassroomsByLocation(c *fiber.Ctx) error {
	var input struct {
		Lat    float64 `query:"lat"`
		Lng    float64 `query:"lng"`
		Radius float64 `query:"radius"`
	}

	if err := c.QueryParser(&input); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid location parameters")
	}

	classrooms, err := h.useCase.GetClassroomsByLocation(input.Lat, input.Lng, input.Radius)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to fetch classrooms by location")
	}

	return response.Success(c, "Classrooms by location retrieved successfully", classrooms)
}

func (h *Handler) UpdateClassroom(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, domain.ErrInvalidID.Error())
	}

	classroom := new(domain.UpdateClassroomDTO)
	if err := c.BodyParser(classroom); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid request body format")
	}

	err = h.useCase.UpdateClassroom(id, classroom)
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
				return response.Error(c, fiber.StatusInternalServerError, "An unexpected error occurred")
			}
		default:
			return response.Error(c, fiber.StatusInternalServerError, "Failed to update classroom")
		}
	}

	return response.Success(c, "Classroom updated successfully", nil)
}

func (h *Handler) DeleteClassroom(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, domain.ErrInvalidID.Error())
	}

	err = h.useCase.DeleteClassroom(id)
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
			return response.Error(c, fiber.StatusInternalServerError, "Failed to delete classroom")
		}
	}

	return response.Success(c, "Classroom deleted successfully", nil)
} 