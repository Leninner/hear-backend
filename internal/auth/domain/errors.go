package domain

import "github.com/leninner/hear-backend/internal/users/domain"

type ValidationErrors = domain.ValidationErrors
type DomainError = domain.DomainError
type ErrorType = domain.ErrorType

const (
	ErrorTypeValidation = domain.ErrorTypeValidation
	ErrorTypeNotFound   = domain.ErrorTypeNotFound
	ErrorTypeConflict   = domain.ErrorTypeConflict
	ErrorTypeInternal   = domain.ErrorTypeInternal
)

func NewValidationErrors() *ValidationErrors {
	return domain.NewValidationErrors()
}

func NewValidationError(message string) error {
	return domain.NewValidationError(message)
}

func NewNotFoundError(message string) error {
	return domain.NewNotFoundError(message)
}

func NewConflictError(message string) error {
	return domain.NewConflictError(message)
}

func NewInternalError(message string, err error) error {
	return domain.NewInternalError(message, err)
}

var (
	ErrInvalidCredentials = domain.NewValidationError("invalid email or password")
	ErrTokenExpired       = domain.NewValidationError("token has expired")
	ErrInvalidToken       = domain.NewValidationError("invalid token")
	ErrTokenNotFound      = domain.NewNotFoundError("refresh token not found")
	ErrEmailExists        = domain.ErrEmailExists
	ErrEmailRequired      = domain.ErrEmailRequired
	ErrEmailInvalid       = domain.ErrEmailInvalid
	ErrPasswordRequired   = domain.ErrPasswordRequired
	ErrPasswordTooShort   = domain.ErrPasswordTooShort
	ErrFirstNameRequired  = domain.ErrFirstNameRequired
	ErrLastNameRequired   = domain.ErrLastNameRequired
	ErrRoleRequired       = domain.ErrRoleRequired
	ErrRoleInvalid        = domain.ErrRoleInvalid
	ErrRefreshTokenRequired = "refresh token field is required and cannot be empty"
) 