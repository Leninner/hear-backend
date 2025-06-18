package domain

import (
	"time"

	"github.com/google/uuid"
)

type CreateAttendanceDTO struct {
	StudentID      uuid.UUID       `json:"studentId"`
	ClassScheduleID uuid.UUID      `json:"classScheduleId"`
	Status         AttendanceStatus `json:"status"`
	Date           time.Time       `json:"date"`
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

	if validationErrors.HasErrors() {
		return validationErrors
	}

	return nil
}

type UpdateAttendanceDTO struct {
	Status AttendanceStatus `json:"status,omitempty"`
	Date   *time.Time       `json:"date,omitempty"`
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

	if validationErrors.HasErrors() {
		return validationErrors
	}

	return nil
} 