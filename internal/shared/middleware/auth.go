package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/leninner/hear-backend/internal/auth/domain"
	"github.com/leninner/hear-backend/internal/shared/response"
)

type AuthMiddleware struct {
	validateToken func(string) (*domain.TokenClaims, error)
}

func NewAuthMiddleware(validateToken func(string) (*domain.TokenClaims, error)) *AuthMiddleware {
	return &AuthMiddleware{
		validateToken: validateToken,
	}
}

func (m *AuthMiddleware) Authenticate() fiber.Handler {
	return func(c *fiber.Ctx) error {
		path := c.Path()
		method := c.Method()
		
		if IsPublicRoute(path, method) {
			return c.Next()
		}

		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return response.Error(c, fiber.StatusUnauthorized, "Authorization header is required")
		}

		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			return response.Error(c, fiber.StatusUnauthorized, "Invalid authorization header format")
		}

		token := tokenParts[1]
		claims, err := m.validateToken(token)
		if err != nil {
			return response.Error(c, fiber.StatusUnauthorized, "Invalid or expired token")
		}

		c.Locals("user", claims)
		return c.Next()
	}
}

func (m *AuthMiddleware) RequireRole(allowedRoles ...string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := c.Locals("user").(*domain.TokenClaims)

		if user.Role == "admin" {
			return c.Next()
		}
		
		for _, role := range allowedRoles {
			if user.Role == role {
				return c.Next()
			}
		}

		return response.Error(c, fiber.StatusForbidden, "Insufficient permissions")
	}
}

func (m *AuthMiddleware) RequireAnyRole(allowedRoles ...string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := c.Locals("user").(*domain.TokenClaims)
		
		for _, role := range allowedRoles {
			if user.Role == role {
				return c.Next()
			}
		}

		return response.Error(c, fiber.StatusForbidden, "Insufficient permissions")
	}
}

func (m *AuthMiddleware) RequireAdmin() fiber.Handler {
	return m.RequireRole("admin")
}

func (m *AuthMiddleware) RequireTeacher() fiber.Handler {
	return m.RequireRole("teacher")
}

func (m *AuthMiddleware) RequireStudent() fiber.Handler {
	return m.RequireRole("student")
}

func (m *AuthMiddleware) RequireTeacherOrAdmin() fiber.Handler {
	return m.RequireAnyRole("teacher", "admin")
}

func (m *AuthMiddleware) RequireTeacherOrStudent() fiber.Handler {
	return m.RequireAnyRole("teacher", "student")
} 