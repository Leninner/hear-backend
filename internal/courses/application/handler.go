package application

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/leninner/hear-backend/internal/courses/domain"
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

func (h *Handler) CreateCourse(c *fiber.Ctx) error {
	course := new(domain.CreateCourseDTO)
	if err := c.BodyParser(course); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid request body format")
	}

	createdCourse, err := h.useCase.CreateCourse(course)
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
			return response.Error(c, fiber.StatusInternalServerError, "Failed to create course")
		}
	}

	return response.Success(c, "Course created successfully", createdCourse)
}

func (h *Handler) GetCourse(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, domain.ErrInvalidID.Error())
	}

	course, err := h.useCase.GetCourse(id)
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
			return response.Error(c, fiber.StatusInternalServerError, "Failed to retrieve course")
		}
	}

	return response.Success(c, "Course retrieved successfully", course)
}

func (h *Handler) GetAllCourses(c *fiber.Ctx) error {
	courses, err := h.useCase.GetAllCourses()
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to retrieve courses")
	}

	return response.Success(c, "Courses retrieved successfully", courses)
}

func (h *Handler) UpdateCourse(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, domain.ErrInvalidID.Error())
	}

	course := new(domain.UpdateCourseDTO)
	if err := c.BodyParser(course); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid request body format")
	}

	if err := h.useCase.UpdateCourse(id, course); err != nil {
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
			return response.Error(c, fiber.StatusInternalServerError, "Failed to update course")
		}
	}

	return response.Success(c, "Course updated successfully", nil)
}

func (h *Handler) DeleteCourse(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, domain.ErrInvalidID.Error())
	}

	if err := h.useCase.DeleteCourse(id); err != nil {
		switch e := err.(type) {
		case *domain.DomainError:
			switch e.Type {
			case domain.ErrorTypeNotFound:
				return response.Error(c, fiber.StatusNotFound, e.Message)
			default:
				return response.Error(c, fiber.StatusInternalServerError, "An unexpected error occurred")
			}
		default:
			return response.Error(c, fiber.StatusInternalServerError, "Failed to delete course")
		}
	}

	return response.Success(c, "Course deleted successfully", nil)
} 