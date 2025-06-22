package domain

import (
	"strconv"
	"strings"

	"github.com/google/uuid"
)

// Semester represents valid semester values from 1 to 10
type Semester int

const (
	Semester1  Semester = 1
	Semester2  Semester = 2
	Semester3  Semester = 3
	Semester4  Semester = 4
	Semester5  Semester = 5
	Semester6  Semester = 6
	Semester7  Semester = 7
	Semester8  Semester = 8
	Semester9  Semester = 9
	Semester10 Semester = 10
)

// IsValidSemester checks if a semester value is valid (1-10)
func IsValidSemester(semester string) bool {
	if strings.TrimSpace(semester) == "" {
		return false
	}
	
	value, err := strconv.Atoi(semester)
	if err != nil {
		return false
	}
	
	return value >= 1 && value <= 10
}

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
	} else if !IsValidSemester(dto.Semester) {
		validationErrors.AddError(ErrSemesterInvalid)
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

	if dto.Semester != "" {
		if strings.TrimSpace(dto.Semester) == "" {
			validationErrors.AddError(ErrSemesterRequired)
		} else if !IsValidSemester(dto.Semester) {
			validationErrors.AddError(ErrSemesterInvalid)
		}
	}

	if validationErrors.HasErrors() {
		return validationErrors
	}

	return nil
} 