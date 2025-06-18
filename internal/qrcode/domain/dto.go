package domain

import (
	"time"

	"github.com/google/uuid"
)

type CreateQRCodeDTO struct {
	CourseID  uuid.UUID `json:"courseId"`
	Code      string    `json:"code"`
	ExpiresAt time.Time `json:"expiresAt"`
}

func (dto *CreateQRCodeDTO) Validate() error {
	validationErrors := NewValidationErrors()

	if dto.CourseID == uuid.Nil {
		validationErrors.AddError(ErrCourseIDRequired)
	}

	if dto.Code == "" {
		validationErrors.AddError(ErrCodeRequired)
	} else if len(dto.Code) < 6 {
		validationErrors.AddError(ErrCodeTooShort)
	} else if len(dto.Code) > 255 {
		validationErrors.AddError(ErrCodeTooLong)
	}

	if dto.ExpiresAt.IsZero() {
		validationErrors.AddError(ErrExpiresAtRequired)
	} else if dto.ExpiresAt.Before(time.Now()) {
		validationErrors.AddError(ErrExpiresAtCannotBePast)
	}

	if validationErrors.HasErrors() {
		return validationErrors
	}

	return nil
}

type UpdateQRCodeDTO struct {
	Code      string     `json:"code,omitempty"`
	ExpiresAt *time.Time `json:"expiresAt,omitempty"`
}

func (dto *UpdateQRCodeDTO) Validate() error {
	validationErrors := NewValidationErrors()

	if dto.Code != "" {
		if len(dto.Code) < 6 {
			validationErrors.AddError(ErrCodeTooShort)
		} else if len(dto.Code) > 255 {
			validationErrors.AddError(ErrCodeTooLong)
		}
	}

	if dto.ExpiresAt != nil {
		if dto.ExpiresAt.IsZero() {
			validationErrors.AddError(ErrExpiresAtRequired)
		} else if dto.ExpiresAt.Before(time.Now()) {
			validationErrors.AddError(ErrExpiresAtCannotBePast)
		}
	}

	if validationErrors.HasErrors() {
		return validationErrors
	}

	return nil
} 