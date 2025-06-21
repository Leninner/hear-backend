package domain

import (
	"strings"

	"github.com/google/uuid"
)

type CreateCourseSectionDTO struct {
	CourseID    uuid.UUID `json:"courseId"`
	Name        string    `json:"name"`
	TeacherID   uuid.UUID `json:"teacherId"`
	MaxStudents int       `json:"maxStudents"`
}

func (dto *CreateCourseSectionDTO) Validate() error {
	validationErrors := NewValidationErrors()

	if dto.CourseID == uuid.Nil {
		validationErrors.AddError(ErrCourseIDRequired)
	}

	if strings.TrimSpace(dto.Name) == "" {
		validationErrors.AddError(ErrSectionNameRequired)
	}

	if dto.TeacherID == uuid.Nil {
		validationErrors.AddError(ErrTeacherIDRequired)
	}

	if dto.MaxStudents <= 0 {
		validationErrors.AddError(ErrMaxStudentsInvalid)
	}

	if validationErrors.HasErrors() {
		return validationErrors
	}

	return nil
}

type UpdateCourseSectionDTO struct {
	Name        string `json:"name,omitempty"`
	MaxStudents *int   `json:"maxStudents,omitempty"`
}

func (dto *UpdateCourseSectionDTO) Validate() error {
	validationErrors := NewValidationErrors()

	if dto.Name != "" && strings.TrimSpace(dto.Name) == "" {
		validationErrors.AddError(ErrSectionNameRequired)
	}

	if dto.MaxStudents != nil && *dto.MaxStudents <= 0 {
		validationErrors.AddError(ErrMaxStudentsInvalid)
	}

	if validationErrors.HasErrors() {
		return validationErrors
	}

	return nil
} 