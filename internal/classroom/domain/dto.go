package domain

import (
	"strings"
)

type CreateClassroomDTO struct {
	Name        string   `json:"name"`
	Building    string   `json:"building"`
	Floor       *int     `json:"floor"`
	Capacity    *int     `json:"capacity"`
	LocationLat *float64 `json:"locationLat"`
	LocationLng *float64 `json:"locationLng"`
}

func (dto *CreateClassroomDTO) Validate() error {
	validationErrors := NewValidationErrors()

	if strings.TrimSpace(dto.Name) == "" {
		validationErrors.AddError(ErrNameRequired)
	}

	if strings.TrimSpace(dto.Building) == "" {
		validationErrors.AddError(ErrBuildingRequired)
	}

	if dto.Floor == nil {
		validationErrors.AddError(ErrFloorRequired)
	} else if *dto.Floor < 0 {
		validationErrors.AddError(ErrFloorInvalid)
	}

	if dto.Capacity == nil {
		validationErrors.AddError(ErrCapacityRequired)
	} else if *dto.Capacity <= 0 {
		validationErrors.AddError(ErrCapacityTooSmall)
	} else if *dto.Capacity > 1000 {
		validationErrors.AddError(ErrCapacityTooLarge)
	}

	if dto.LocationLat == nil {
		validationErrors.AddError(ErrLocationLatRequired)
	} else if *dto.LocationLat < -90 || *dto.LocationLat > 90 {
		validationErrors.AddError(ErrLocationLatInvalid)
	}

	if dto.LocationLng == nil {
		validationErrors.AddError(ErrLocationLngRequired)
	} else if *dto.LocationLng < -180 || *dto.LocationLng > 180 {
		validationErrors.AddError(ErrLocationLngInvalid)
	}

	if validationErrors.HasErrors() {
		return validationErrors
	}

	return nil
}

type UpdateClassroomDTO struct {
	Name        string   `json:"name,omitempty"`
	Building    string   `json:"building,omitempty"`
	Floor       *int     `json:"floor,omitempty"`
	Capacity    *int     `json:"capacity,omitempty"`
	LocationLat *float64 `json:"locationLat,omitempty"`
	LocationLng *float64 `json:"locationLng,omitempty"`
}

func (dto *UpdateClassroomDTO) Validate() error {
	validationErrors := NewValidationErrors()

	if dto.Name != "" && strings.TrimSpace(dto.Name) == "" {
		validationErrors.AddError(ErrNameRequired)
	}

	if dto.Building != "" && strings.TrimSpace(dto.Building) == "" {
		validationErrors.AddError(ErrBuildingRequired)
	}

	if dto.Floor != nil && *dto.Floor < 0 {
		validationErrors.AddError(ErrFloorInvalid)
	}

	if dto.Capacity != nil {
		if *dto.Capacity <= 0 {
			validationErrors.AddError(ErrCapacityTooSmall)
		} else if *dto.Capacity > 1000 {
			validationErrors.AddError(ErrCapacityTooLarge)
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