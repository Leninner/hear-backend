package domain

import (
	"strings"
)

type CreateUniversityDTO struct {
	Name string `json:"name"`
}

func (dto *CreateUniversityDTO) Validate() error {
	validationErrors := NewValidationErrors()

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

type UpdateUniversityDTO struct {
	Name string `json:"name,omitempty"`
}

func (dto *UpdateUniversityDTO) Validate() error {
	validationErrors := NewValidationErrors()

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