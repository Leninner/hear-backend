package domain

import (
	"strings"

	"github.com/google/uuid"
)

type CreateCourseDTO struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	TeacherID   uuid.UUID `json:"teacherId"`
	FacultyID   uuid.UUID `json:"facultyId"`
}

func (dto *CreateCourseDTO) Validate() error {
	validationErrors := NewValidationErrors()

	if strings.TrimSpace(dto.Name) == "" {
		validationErrors.AddError(ErrNameRequired)
	}

	if strings.TrimSpace(dto.Description) == "" {
		validationErrors.AddError(ErrDescriptionRequired)
	}

	if dto.TeacherID == uuid.Nil {
		validationErrors.AddError(ErrTeacherIDRequired)
	}

	if dto.FacultyID == uuid.Nil {
		validationErrors.AddError(ErrFacultyIDRequired)
	}

	if validationErrors.HasErrors() {
		return validationErrors
	}

	return nil
}

type UpdateCourseDTO struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

func (dto *UpdateCourseDTO) Validate() error {
	validationErrors := NewValidationErrors()

	if dto.Name != "" && strings.TrimSpace(dto.Name) == "" {
		validationErrors.AddError(ErrNameRequired)
	}

	if dto.Description != "" && strings.TrimSpace(dto.Description) == "" {
		validationErrors.AddError(ErrDescriptionRequired)
	}

	if validationErrors.HasErrors() {
		return validationErrors
	}

	return nil
} 