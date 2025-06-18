package domain

import "github.com/leninner/hear-backend/internal/users/domain"

var (
	ErrInvalidCredentials = domain.NewValidationError("invalid email or password")
	ErrTokenExpired       = domain.NewValidationError("token has expired")
	ErrInvalidToken       = domain.NewValidationError("invalid token")
	ErrTokenNotFound      = domain.NewNotFoundError("refresh token not found")
	ErrEmailExists        = domain.ErrEmailExists
) 