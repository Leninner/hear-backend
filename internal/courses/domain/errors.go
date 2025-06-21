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
	ErrDescriptionRequired = "description field is required and cannot be empty"
	ErrTeacherIDRequired   = "teacher ID field is required and cannot be empty"
	ErrFacultyIDRequired   = "faculty ID field is required and cannot be empty"
	ErrSemesterRequired    = "semester field is required and cannot be empty"
	ErrCourseIDRequired    = "course ID field is required and cannot be empty"
	ErrSectionIDRequired   = "section ID field is required and cannot be empty"
	ErrClassroomIDRequired = "classroom ID field is required and cannot be empty"
	ErrSectionNameRequired = "section name field is required and cannot be empty"
	ErrMaxStudentsInvalid  = "max students must be greater than 0"
	ErrDayOfWeekInvalid    = "day of week must be between 0 and 6"
	ErrStartTimeRequired   = "start time field is required and cannot be empty"
	ErrEndTimeRequired     = "end time field is required and cannot be empty"
	ErrStartTimeInvalid    = "start time must be in HH:MM format"
	ErrEndTimeInvalid      = "end time must be in HH:MM format"
	ErrCourseNotFound      = NewNotFoundError("course not found")
	ErrInvalidID          = NewValidationError("invalid course ID format")
) 