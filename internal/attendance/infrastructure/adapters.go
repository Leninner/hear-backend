package infrastructure

import (
	"github.com/google/uuid"
	"github.com/leninner/hear-backend/internal/attendance/domain"
	classroomDomain "github.com/leninner/hear-backend/internal/classroom/domain"
	schedulesDomain "github.com/leninner/hear-backend/internal/schedules/domain"
)

// ScheduleAdapter adapts the schedules repository to the attendance domain interface
type ScheduleAdapter struct {
	scheduleRepository schedulesDomain.Repository
}

func NewScheduleAdapter(scheduleRepository schedulesDomain.Repository) *ScheduleAdapter {
	return &ScheduleAdapter{
		scheduleRepository: scheduleRepository,
	}
}

func (a *ScheduleAdapter) GetByID(id uuid.UUID) (*domain.Schedule, error) {
	schedule, err := a.scheduleRepository.GetByID(id)
	if err != nil {
		return nil, err
	}
	
	return &domain.Schedule{
		ID:          schedule.ID,
		ClassroomID: schedule.ClassroomID,
	}, nil
}

// ClassroomAdapter adapts the classroom repository to the attendance domain interface
type ClassroomAdapter struct {
	classroomRepository classroomDomain.Repository
}

func NewClassroomAdapter(classroomRepository classroomDomain.Repository) *ClassroomAdapter {
	return &ClassroomAdapter{
		classroomRepository: classroomRepository,
	}
}

func (a *ClassroomAdapter) GetByID(id uuid.UUID) (*domain.Classroom, error) {
	classroom, err := a.classroomRepository.GetByID(id)
	if err != nil {
		return nil, err
	}
	
	return &domain.Classroom{
		ID:          classroom.ID,
		LocationLat: classroom.LocationLat,
		LocationLng: classroom.LocationLng,
	}, nil
} 