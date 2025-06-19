package domain

import (
	"time"

	"github.com/google/uuid"
)

type CreateAttendanceDTO struct {
	StudentID         uuid.UUID       `json:"studentId"`
	ClassScheduleID   uuid.UUID       `json:"classScheduleId"`
	Status            AttendanceStatus `json:"status"`
	Date              time.Time       `json:"date"`
	UserLatitude      *float64        `json:"userLatitude,omitempty"`
	UserLongitude     *float64        `json:"userLongitude,omitempty"`
	MaxDistanceMeters *int            `json:"maxDistanceMeters,omitempty"`
}

func (dto *CreateAttendanceDTO) Validate() error {
	validationErrors := NewValidationErrors()

	if dto.StudentID == uuid.Nil {
		validationErrors.AddError(ErrStudentIDRequired)
	}

	if dto.ClassScheduleID == uuid.Nil {
		validationErrors.AddError(ErrClassScheduleIDRequired)
	}

	if dto.Status == "" {
		validationErrors.AddError(ErrStatusRequired)
	} else {
		switch dto.Status {
		case StatusPresent, StatusAbsent, StatusLate:
		default:
			validationErrors.AddError(ErrStatusInvalid)
		}
	}

	if dto.Date.IsZero() {
		validationErrors.AddError(ErrDateRequired)
	} else if dto.Date.After(time.Now()) {
		validationErrors.AddError(ErrDateCannotBeFuture)
	}

	// Validate location coordinates if provided
	if dto.UserLatitude != nil || dto.UserLongitude != nil {
		if dto.UserLatitude == nil {
			validationErrors.AddError(ErrLatitudeRequired)
		} else if *dto.UserLatitude < -90 || *dto.UserLatitude > 90 {
			validationErrors.AddError(ErrLatitudeInvalid)
		}

		if dto.UserLongitude == nil {
			validationErrors.AddError(ErrLongitudeRequired)
		} else if *dto.UserLongitude < -180 || *dto.UserLongitude > 180 {
			validationErrors.AddError(ErrLongitudeInvalid)
		}
	}

	// Validate max distance if provided
	if dto.MaxDistanceMeters != nil {
		if *dto.MaxDistanceMeters <= 0 {
			validationErrors.AddError(ErrMaxDistanceInvalid)
		}
	}

	if validationErrors.HasErrors() {
		return validationErrors
	}

	return nil
}

type UpdateAttendanceDTO struct {
	Status            AttendanceStatus `json:"status,omitempty"`
	Date              *time.Time       `json:"date,omitempty"`
	UserLatitude      *float64         `json:"userLatitude,omitempty"`
	UserLongitude     *float64         `json:"userLongitude,omitempty"`
	MaxDistanceMeters *int             `json:"maxDistanceMeters,omitempty"`
}

func (dto *UpdateAttendanceDTO) Validate() error {
	validationErrors := NewValidationErrors()

	if dto.Status != "" {
		switch dto.Status {
		case StatusPresent, StatusAbsent, StatusLate:
		default:
			validationErrors.AddError(ErrStatusInvalid)
		}
	}

	if dto.Date != nil {
		if dto.Date.IsZero() {
			validationErrors.AddError(ErrDateRequired)
		} else if dto.Date.After(time.Now()) {
			validationErrors.AddError(ErrDateCannotBeFuture)
		}
	}

	// Validate location coordinates if provided
	if dto.UserLatitude != nil {
		if *dto.UserLatitude < -90 || *dto.UserLatitude > 90 {
			validationErrors.AddError(ErrLatitudeInvalid)
		}
	}

	if dto.UserLongitude != nil {
		if *dto.UserLongitude < -180 || *dto.UserLongitude > 180 {
			validationErrors.AddError(ErrLongitudeInvalid)
		}
	}

	// Validate max distance if provided
	if dto.MaxDistanceMeters != nil {
		if *dto.MaxDistanceMeters <= 0 {
			validationErrors.AddError(ErrMaxDistanceInvalid)
		}
	}

	if validationErrors.HasErrors() {
		return validationErrors
	}

	return nil
} 