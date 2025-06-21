package domain

import (
	"strings"

	"github.com/google/uuid"
)

type CreateCourseDTO struct {
	Name      string    `json:"name"`
	FacultyID uuid.UUID `json:"facultyId"`
	Semester  string    `json:"semester"`
}

func (dto *CreateCourseDTO) Validate() error {
	validationErrors := NewValidationErrors()

	if strings.TrimSpace(dto.Name) == "" {
		validationErrors.AddError(ErrNameRequired)
	}

	if dto.FacultyID == uuid.Nil {
		validationErrors.AddError(ErrFacultyIDRequired)
	}

	if strings.TrimSpace(dto.Semester) == "" {
		validationErrors.AddError(ErrSemesterRequired)
	}

	if validationErrors.HasErrors() {
		return validationErrors
	}

	return nil
}

type UpdateCourseDTO struct {
	Name     string `json:"name,omitempty"`
	Semester string `json:"semester,omitempty"`
}

func (dto *UpdateCourseDTO) Validate() error {
	validationErrors := NewValidationErrors()

	if dto.Name != "" && strings.TrimSpace(dto.Name) == "" {
		validationErrors.AddError(ErrNameRequired)
	}

	if dto.Semester != "" && strings.TrimSpace(dto.Semester) == "" {
		validationErrors.AddError(ErrSemesterRequired)
	}

	if validationErrors.HasErrors() {
		return validationErrors
	}

	return nil
} 