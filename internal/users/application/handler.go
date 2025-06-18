package application

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/leninner/hear-backend/internal/shared/response"
	"github.com/leninner/hear-backend/internal/users/domain"
)

type Handler struct {
	useCase *UseCase
}

func NewHandler(useCase *UseCase) *Handler {
	return &Handler{
		useCase: useCase,
	}
}

func (h *Handler) Create(c *fiber.Ctx) error {
	user := new(domain.CreateUserDTO)
	if err := c.BodyParser(user); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid request body format")
	}

	createdUser, err := h.useCase.CreateUser(user)
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
			return response.Error(c, fiber.StatusInternalServerError, "Failed to create user")
		}
	}

	return response.Success(c, "User created successfully", createdUser)
}

func (h *Handler) GetByID(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, domain.ErrInvalidID.Error())
	}

	user, err := h.useCase.GetUser(id)
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
			return response.Error(c, fiber.StatusInternalServerError, "Failed to retrieve user")
		}
	}

	return response.Success(c, "User retrieved successfully", user)
}

func (h *Handler) GetAll(c *fiber.Ctx) error {
	users, err := h.useCase.GetAllUsers()
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to retrieve users")
	}

	return response.Success(c, "Users retrieved successfully", users)
} 