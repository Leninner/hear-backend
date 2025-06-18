package domain

import (
	"time"

	"github.com/google/uuid"
)

type AttendanceStatus string

const (
	StatusPresent AttendanceStatus = "present"
	StatusAbsent  AttendanceStatus = "absent"
	StatusLate    AttendanceStatus = "late"
)

type Attendance struct {
	ID             uuid.UUID       `json:"id"`
	StudentID      uuid.UUID       `json:"studentId"`
	ClassScheduleID uuid.UUID      `json:"classScheduleId"`
	Status         AttendanceStatus `json:"status"`
	Date           time.Time       `json:"date"`
	CreatedAt      time.Time       `json:"createdAt"`
	UpdatedAt      time.Time       `json:"updatedAt"`
}

func NewAttendance(studentID, classScheduleID uuid.UUID, status AttendanceStatus, date time.Time) *Attendance {
	now := time.Now()
	return &Attendance{
		ID:             uuid.New(),
		StudentID:      studentID,
		ClassScheduleID: classScheduleID,
		Status:         status,
		Date:           date,
		CreatedAt:      now,
		UpdatedAt:      now,
	}
}

type Repository interface {
	Create(attendance *Attendance) error
	GetByID(id uuid.UUID) (*Attendance, error)
	GetByStudentID(studentID uuid.UUID) ([]*Attendance, error)
	GetByClassScheduleID(classScheduleID uuid.UUID) ([]*Attendance, error)
	GetByDate(date time.Time) ([]*Attendance, error)
	Update(attendance *Attendance) error
	Delete(id uuid.UUID) error
} 