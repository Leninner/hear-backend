package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/leninner/hear-backend/internal/courses/domain"
)

type DayOfWeek string

const (
	Sunday    DayOfWeek = "0"
	Monday    DayOfWeek = "1"
	Tuesday   DayOfWeek = "2"
	Wednesday DayOfWeek = "3"
	Thursday  DayOfWeek = "4"
	Friday    DayOfWeek = "5"
	Saturday  DayOfWeek = "6"
)

func (d DayOfWeek) IsValid() bool {
	switch d {
	case Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday:
		return true
	default:
		return false
	}
}

func (d DayOfWeek) ToInt() int {
	switch d {
	case Sunday:
		return 0
	case Monday:
		return 1
	case Tuesday:
		return 2
	case Wednesday:
		return 3
	case Thursday:
		return 4
	case Friday:
		return 5
	case Saturday:
		return 6
	default:
		return -1
	}
}

type Schedule struct {
	ID          uuid.UUID `json:"id"`
	CourseID    uuid.UUID `json:"courseId"`
	SectionID   uuid.UUID `json:"sectionId"`
	ClassroomID uuid.UUID `json:"classroomId"`
	DayOfWeek   int       `json:"dayOfWeek"`
	StartTime   time.Time `json:"startTime"`
	EndTime     time.Time `json:"endTime"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func NewSchedule(courseID, sectionID, classroomID uuid.UUID, dayOfWeek int, startTime, endTime time.Time) *Schedule {
	now := time.Now()
	return &Schedule{
		ID:          uuid.New(),
		CourseID:    courseID,
		SectionID:   sectionID,
		ClassroomID: classroomID,
		DayOfWeek:   dayOfWeek,
		StartTime:   startTime,
		EndTime:     endTime,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
}

type Repository interface {
	Create(schedule *Schedule) error
	GetByID(id uuid.UUID) (*Schedule, error)
	GetByCourseID(courseID uuid.UUID) ([]*Schedule, error)
	GetBySectionID(sectionID uuid.UUID) ([]*Schedule, error)
	GetByClassroomID(classroomID uuid.UUID) ([]*Schedule, error)
	GetByClassroomAndTime(classroomID uuid.UUID, dayOfWeek int, startTime, endTime time.Time) ([]*Schedule, error)
	Update(schedule *Schedule) error
	Delete(id uuid.UUID) error
}

type CoursesRepository interface {
	GetSectionByID(id uuid.UUID) (*domain.CourseSection, error)
} 