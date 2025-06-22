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
	ErrNameRequired        = "name field is required and cannot be empty"
	ErrBuildingRequired    = "building field is required and cannot be empty"
	ErrFacultyIDRequired   = "faculty ID field is required and cannot be empty"
	ErrFloorRequired       = "floor field is required"
	ErrCapacityRequired    = "capacity field is required"
	ErrLocationLatRequired = "locationLat field is required"
	ErrLocationLngRequired = "locationLng field is required"
	ErrFloorInvalid        = "floor must be a non-negative integer"
	ErrCapacityInvalid     = "capacity must be a positive integer"
	ErrCapacityTooSmall    = "capacity must be at least 1"
	ErrCapacityTooLarge    = "capacity cannot exceed 1000"
	ErrLocationLatInvalid  = "locationLat must be between -90 and 90"
	ErrLocationLngInvalid  = "locationLng must be between -180 and 180"
	ErrClassroomNotFound   = NewNotFoundError("classroom not found")
	ErrFacultyNotFound     = NewNotFoundError("faculty not found")
	ErrInvalidID           = NewValidationError("invalid classroom ID format")
	ErrClassroomNameExists = NewConflictError("a classroom with this name already exists")
) 