package domain

import (
	"fmt"
)

type ErrorType string

const (
	ErrorTypeValidation ErrorType = "VALIDATION_ERROR"
	ErrorTypeNotFound   ErrorType = "NOT_FOUND_ERROR"
	ErrorTypeConflict   ErrorType = "CONFLICT_ERROR"
	ErrorTypeInternal   ErrorType = "INTERNAL_ERROR"
)

type ValidationErrors struct {
	Errors []string
}

func (e *ValidationErrors) Error() string {
	return fmt.Sprintf("validation errors: %v", e.Errors)
}

func (e *ValidationErrors) AddError(err string) {
	e.Errors = append(e.Errors, err)
}

func (e *ValidationErrors) HasErrors() bool {
	return len(e.Errors) > 0
}

func (e *ValidationErrors) GetErrors() []string {
	return e.Errors
}

type DomainError struct {
	Type    ErrorType
	Message string
	Err     error
}

func (e *DomainError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %s (%v)", e.Type, e.Message, e.Err)
	}
	return fmt.Sprintf("%s: %s", e.Type, e.Message)
}

func (e *DomainError) Unwrap() error {
	return e.Err
}

func NewValidationError(message string) error {
	return &DomainError{
		Type:    ErrorTypeValidation,
		Message: message,
	}
}

func NewValidationErrors() *ValidationErrors {
	return &ValidationErrors{
		Errors: make([]string, 0),
	}
}

func NewNotFoundError(message string) error {
	return &DomainError{
		Type:    ErrorTypeNotFound,
		Message: message,
	}
}

func NewConflictError(message string) error {
	return &DomainError{
		Type:    ErrorTypeConflict,
		Message: message,
	}
}

func NewInternalError(message string, err error) error {
	return &DomainError{
		Type:    ErrorTypeInternal,
		Message: message,
		Err:     err,
	}
}

var (
	ErrNameRequired     = "name field is required and cannot be empty"
	ErrNameTooShort     = "name must be at least 2 characters long"
	ErrNameTooLong      = "name cannot exceed 255 characters"
	ErrUniversityExists = NewConflictError("a university with this name already exists")
	ErrUniversityNotFound = NewNotFoundError("university not found")
	ErrInvalidID        = NewValidationError("invalid university ID format")
) 