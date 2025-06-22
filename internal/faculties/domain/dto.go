package domain

import (
	"strings"

	"github.com/google/uuid"
)

type CreateFacultyDTO struct {
	UniversityID uuid.UUID `json:"universityId"`
	Name         string    `json:"name"`
}

func (dto *CreateFacultyDTO) Validate() error {
	validationErrors := NewValidationErrors()

	if dto.UniversityID == uuid.Nil {
		validationErrors.AddError(ErrUniversityIDRequired)
	}

	if strings.TrimSpace(dto.Name) == "" {
		validationErrors.AddError(ErrNameRequired)
	} else if len(strings.TrimSpace(dto.Name)) < 2 {
		validationErrors.AddError(ErrNameTooShort)
	} else if len(strings.TrimSpace(dto.Name)) > 255 {
		validationErrors.AddError(ErrNameTooLong)
	}

	if validationErrors.HasErrors() {
		return validationErrors
	}

	return nil
}

type UpdateFacultyDTO struct {
	UniversityID *uuid.UUID `json:"universityId,omitempty"`
	Name         string     `json:"name,omitempty"`
}

func (dto *UpdateFacultyDTO) Validate() error {
	validationErrors := NewValidationErrors()

	if dto.UniversityID != nil && *dto.UniversityID == uuid.Nil {
		validationErrors.AddError(ErrUniversityIDRequired)
	}

	if dto.Name != "" {
		if strings.TrimSpace(dto.Name) == "" {
			validationErrors.AddError(ErrNameRequired)
		} else if len(strings.TrimSpace(dto.Name)) < 2 {
			validationErrors.AddError(ErrNameTooShort)
		} else if len(strings.TrimSpace(dto.Name)) > 255 {
			validationErrors.AddError(ErrNameTooLong)
		}
	}

	if validationErrors.HasErrors() {
		return validationErrors
	}

	return nil
} 