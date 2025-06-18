package domain

import (
	"strings"

	"github.com/google/uuid"
)

type CreateFacultyDTO struct {
	UniversityID uuid.UUID `json:"universityId"`
	Name         string    `json:"name"`
	LocationLat  float64   `json:"locationLat"`
	LocationLng  float64   `json:"locationLng"`
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

	if dto.LocationLat < -90 || dto.LocationLat > 90 {
		validationErrors.AddError(ErrLocationLatInvalid)
	}

	if dto.LocationLng < -180 || dto.LocationLng > 180 {
		validationErrors.AddError(ErrLocationLngInvalid)
	}

	if validationErrors.HasErrors() {
		return validationErrors
	}

	return nil
}

type UpdateFacultyDTO struct {
	UniversityID *uuid.UUID `json:"universityId,omitempty"`
	Name         string     `json:"name,omitempty"`
	LocationLat  *float64   `json:"locationLat,omitempty"`
	LocationLng  *float64   `json:"locationLng,omitempty"`
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

	if dto.LocationLat != nil && (*dto.LocationLat < -90 || *dto.LocationLat > 90) {
		validationErrors.AddError(ErrLocationLatInvalid)
	}

	if dto.LocationLng != nil && (*dto.LocationLng < -180 || *dto.LocationLng > 180) {
		validationErrors.AddError(ErrLocationLngInvalid)
	}

	if validationErrors.HasErrors() {
		return validationErrors
	}

	return nil
} 