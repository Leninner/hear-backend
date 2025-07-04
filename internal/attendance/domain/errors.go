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
	ErrStudentIDRequired      = "student ID field is required and cannot be empty"
	ErrClassScheduleIDRequired = "class schedule ID field is required and cannot be empty"
	ErrStatusRequired         = "status field is required and cannot be empty"
	ErrStatusInvalid          = "status must be one of: present, absent, late"
	ErrDateRequired           = "date field is required and cannot be empty"
	ErrDateCannotBeFuture     = "date cannot be in the future"
	ErrLatitudeRequired       = "latitude is required when longitude is provided"
	ErrLatitudeInvalid        = "latitude must be between -90 and 90 degrees"
	ErrLongitudeRequired      = "longitude is required when latitude is provided"
	ErrLongitudeInvalid       = "longitude must be between -180 and 180 degrees"
	ErrMaxDistanceInvalid     = "max distance must be greater than 0"
	ErrDistanceTooFar         = "user is too far from the classroom location"
	ErrAttendanceExists       = NewConflictError("attendance record already exists for this student, class schedule, and date")
	ErrAttendanceNotFound     = NewNotFoundError("attendance record not found")
	ErrInvalidID              = NewValidationError("invalid attendance ID format")
) 