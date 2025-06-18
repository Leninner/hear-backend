package application

import (
	"github.com/gofiber/fiber/v2"
	"github.com/leninner/hear-backend/internal/auth/domain"
	"github.com/leninner/hear-backend/internal/shared/response"
	userDomain "github.com/leninner/hear-backend/internal/users/domain"
)

type Handler struct {
	useCase *UseCase
}

func NewHandler(useCase *UseCase) *Handler {
	return &Handler{
		useCase: useCase,
	}
}

func (h *Handler) GetUseCase() *UseCase {
	return h.useCase
}

func (h *Handler) Login(c *fiber.Ctx) error {
	loginDTO := new(domain.LoginDTO)
	if err := c.BodyParser(loginDTO); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid request body format")
	}

	authResponse, err := h.useCase.Login(loginDTO)
	if err != nil {
		switch e := err.(type) {
		case *userDomain.ValidationErrors:
			return response.Error(c, fiber.StatusBadRequest, e.GetErrors())
		case *userDomain.DomainError:
			switch e.Type {
			case userDomain.ErrorTypeValidation:
				return response.Error(c, fiber.StatusBadRequest, e.Message)
			case userDomain.ErrorTypeNotFound:
				return response.Error(c, fiber.StatusNotFound, e.Message)
			default:
				return response.Error(c, fiber.StatusInternalServerError, "An unexpected error occurred")
			}
		default:
			return response.Error(c, fiber.StatusInternalServerError, "Failed to authenticate user")
		}
	}

	return response.Success(c, "Login successful", authResponse)
}

func (h *Handler) Register(c *fiber.Ctx) error {
	registerDTO := new(domain.RegisterDTO)
	if err := c.BodyParser(registerDTO); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid request body format")
	}

	authResponse, err := h.useCase.Register(registerDTO)
	if err != nil {
		switch e := err.(type) {
		case *userDomain.ValidationErrors:
			return response.Error(c, fiber.StatusBadRequest, e.GetErrors())
		case *userDomain.DomainError:
			switch e.Type {
			case userDomain.ErrorTypeValidation:
				return response.Error(c, fiber.StatusBadRequest, e.Message)
			case userDomain.ErrorTypeConflict:
				return response.Error(c, fiber.StatusConflict, e.Message)
			default:
				return response.Error(c, fiber.StatusInternalServerError, "An unexpected error occurred")
			}
		default:
			return response.Error(c, fiber.StatusInternalServerError, "Failed to register user")
		}
	}

	return response.Success(c, "Registration successful", authResponse)
}

func (h *Handler) RefreshToken(c *fiber.Ctx) error {
	refreshDTO := new(domain.RefreshTokenDTO)
	if err := c.BodyParser(refreshDTO); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid request body format")
	}

	tokenPair, err := h.useCase.RefreshToken(refreshDTO)
	if err != nil {
		switch e := err.(type) {
		case *userDomain.DomainError:
			switch e.Type {
			case userDomain.ErrorTypeValidation:
				return response.Error(c, fiber.StatusBadRequest, e.Message)
			case userDomain.ErrorTypeNotFound:
				return response.Error(c, fiber.StatusNotFound, e.Message)
			default:
				return response.Error(c, fiber.StatusInternalServerError, "An unexpected error occurred")
			}
		default:
			return response.Error(c, fiber.StatusInternalServerError, "Failed to refresh token")
		}
	}

	return response.Success(c, "Token refreshed successfully", tokenPair)
}

func (h *Handler) Logout(c *fiber.Ctx) error {
	refreshDTO := new(domain.RefreshTokenDTO)
	if err := c.BodyParser(refreshDTO); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid request body format")
	}

	if err := h.useCase.Logout(refreshDTO.RefreshToken); err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to logout")
	}

	return response.Success(c, "Logout successful", nil)
} 