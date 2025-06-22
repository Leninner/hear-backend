package application

import (
	"time"

	"github.com/google/uuid"
	"github.com/leninner/hear-backend/internal/schedules/domain"
)

type UseCase struct {
	repository domain.Repository
	coursesRepository domain.CoursesRepository
}

func NewUseCase(repository domain.Repository, coursesRepository domain.CoursesRepository) *UseCase {
	return &UseCase{
		repository: repository,
		coursesRepository: coursesRepository,
	}
}

func (uc *UseCase) CreateSchedule(dto *domain.CreateScheduleDTO) (*domain.Schedule, error) {
	if err := dto.Validate(); err != nil {
		return nil, err
	}

	section, err := uc.coursesRepository.GetSectionByID(dto.SectionID)
	if err != nil {
		return nil, domain.NewNotFoundError("section not found")
	}

	startTime, err := time.Parse("15:04", dto.StartTime)
	if err != nil {
		return nil, domain.NewValidationError("invalid start time format")
	}

	endTime, err := time.Parse("15:04", dto.EndTime)
	if err != nil {
		return nil, domain.NewValidationError("invalid end time format")
	}

	if startTime.After(endTime) || startTime.Equal(endTime) {
		return nil, domain.NewValidationError("start time must be before end time")
	}

	// Convert DayOfWeek enum to int for storage
	dayOfWeekInt := dto.DayOfWeek.ToInt()
	
	schedule := domain.NewSchedule(section.CourseID, dto.SectionID, dto.ClassroomID, dayOfWeekInt, startTime, endTime)
	
	if err := uc.repository.Create(schedule); err != nil {
		return nil, domain.NewInternalError("failed to create schedule in database", err)
	}

	return schedule, nil
}

func (uc *UseCase) GetSchedule(id uuid.UUID) (*domain.Schedule, error) {
	schedule, err := uc.repository.GetByID(id)
	if err != nil {
		return nil, domain.ErrScheduleNotFound
	}
	return schedule, nil
}

func (uc *UseCase) GetSchedulesByCourse(courseID uuid.UUID) ([]*domain.Schedule, error) {
	schedules, err := uc.repository.GetByCourseID(courseID)
	if err != nil {
		return nil, domain.NewInternalError("failed to retrieve course schedules from database", err)
	}
	return schedules, nil
}

func (uc *UseCase) GetSchedulesBySection(sectionID uuid.UUID) ([]*domain.Schedule, error) {
	schedules, err := uc.repository.GetBySectionID(sectionID)
	if err != nil {
		return nil, domain.NewInternalError("failed to retrieve section schedules from database", err)
	}
	return schedules, nil
}

func (uc *UseCase) GetSchedulesByClassroom(classroomID uuid.UUID) ([]*domain.Schedule, error) {
	schedules, err := uc.repository.GetByClassroomID(classroomID)
	if err != nil {
		return nil, domain.NewInternalError("failed to retrieve classroom schedules from database", err)
	}
	return schedules, nil
}

func (uc *UseCase) GetSchedulesByClassroomAndTime(classroomID uuid.UUID, dayOfWeek int, startTime, endTime time.Time) ([]*domain.Schedule, error) {
	schedules, err := uc.repository.GetByClassroomAndTime(classroomID, dayOfWeek, startTime, endTime)
	if err != nil {
		return nil, domain.NewInternalError("failed to retrieve classroom schedules from database", err)
	}
	return schedules, nil
}

func (uc *UseCase) UpdateSchedule(id uuid.UUID, dto *domain.UpdateScheduleDTO) error {
	if err := dto.Validate(); err != nil {
		return err
	}

	schedule, err := uc.repository.GetByID(id)
	if err != nil {
		return domain.ErrScheduleNotFound
	}

	if dto.SectionID != nil {
		schedule.SectionID = *dto.SectionID
	}
	if dto.ClassroomID != nil {
		schedule.ClassroomID = *dto.ClassroomID
	}
	if dto.DayOfWeek != nil {
		schedule.DayOfWeek = *dto.DayOfWeek
	}
	if dto.StartTime != "" {
		startTime, err := time.Parse("15:04", dto.StartTime)
		if err != nil {
			return domain.NewValidationError("invalid start time format")
		}
		schedule.StartTime = startTime
	}
	if dto.EndTime != "" {
		endTime, err := time.Parse("15:04", dto.EndTime)
		if err != nil {
			return domain.NewValidationError("invalid end time format")
		}
		schedule.EndTime = endTime
	}

	// Validate that start time is before end time
	if schedule.StartTime.After(schedule.EndTime) || schedule.StartTime.Equal(schedule.EndTime) {
		return domain.NewValidationError("start time must be before end time")
	}

	if err := uc.repository.Update(schedule); err != nil {
		return domain.NewInternalError("failed to update schedule in database", err)
	}

	return nil
}

func (uc *UseCase) DeleteSchedule(id uuid.UUID) error {
	if _, err := uc.repository.GetByID(id); err != nil {
		return domain.ErrScheduleNotFound
	}

	if err := uc.repository.Delete(id); err != nil {
		return domain.NewInternalError("failed to delete schedule from database", err)
	}

	return nil
} 