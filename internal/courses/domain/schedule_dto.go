package domain

import (
	"time"

	"github.com/google/uuid"
)

type CreateScheduleDTO struct {
	CourseID    uuid.UUID `json:"courseId"`
	SectionID   uuid.UUID `json:"sectionId"`
	ClassroomID uuid.UUID `json:"classroomId"`
	DayOfWeek   int       `json:"dayOfWeek"`
	StartTime   string    `json:"startTime"`
	EndTime     string    `json:"endTime"`
}

func (dto *CreateScheduleDTO) Validate() error {
	validationErrors := NewValidationErrors()

	if dto.CourseID == uuid.Nil {
		validationErrors.AddError(ErrCourseIDRequired)
	}

	if dto.SectionID == uuid.Nil {
		validationErrors.AddError(ErrSectionIDRequired)
	}

	if dto.ClassroomID == uuid.Nil {
		validationErrors.AddError(ErrClassroomIDRequired)
	}

	if dto.DayOfWeek < 0 || dto.DayOfWeek > 6 {
		validationErrors.AddError(ErrDayOfWeekInvalid)
	}

	if dto.StartTime == "" {
		validationErrors.AddError(ErrStartTimeRequired)
	}

	if dto.EndTime == "" {
		validationErrors.AddError(ErrEndTimeRequired)
	}

	// Validate time format
	if dto.StartTime != "" {
		if _, err := time.Parse("15:04", dto.StartTime); err != nil {
			validationErrors.AddError(ErrStartTimeInvalid)
		}
	}

	if dto.EndTime != "" {
		if _, err := time.Parse("15:04", dto.EndTime); err != nil {
			validationErrors.AddError(ErrEndTimeInvalid)
		}
	}

	if validationErrors.HasErrors() {
		return validationErrors
	}

	return nil
}

type UpdateScheduleDTO struct {
	CourseID    *uuid.UUID `json:"courseId,omitempty"`
	SectionID   *uuid.UUID `json:"sectionId,omitempty"`
	ClassroomID *uuid.UUID `json:"classroomId,omitempty"`
	DayOfWeek   *int       `json:"dayOfWeek,omitempty"`
	StartTime   string     `json:"startTime,omitempty"`
	EndTime     string     `json:"endTime,omitempty"`
}

func (dto *UpdateScheduleDTO) Validate() error {
	validationErrors := NewValidationErrors()

	if dto.DayOfWeek != nil && (*dto.DayOfWeek < 0 || *dto.DayOfWeek > 6) {
		validationErrors.AddError(ErrDayOfWeekInvalid)
	}

	if dto.StartTime != "" {
		if _, err := time.Parse("15:04", dto.StartTime); err != nil {
			validationErrors.AddError(ErrStartTimeInvalid)
		}
	}

	if dto.EndTime != "" {
		if _, err := time.Parse("15:04", dto.EndTime); err != nil {
			validationErrors.AddError(ErrEndTimeInvalid)
		}
	}

	if validationErrors.HasErrors() {
		return validationErrors
	}

	return nil
} 