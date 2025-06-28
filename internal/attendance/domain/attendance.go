package domain

import (
	"fmt"
	"math"
	"time"

	"github.com/google/uuid"
)

type AttendanceStatus string

const (
	StatusPresent AttendanceStatus = "present"
	StatusAbsent  AttendanceStatus = "absent"
	StatusLate    AttendanceStatus = "late"
)

// Default maximum distance in meters for attendance validation
const DefaultMaxDistanceMeters = 50

type Attendance struct {
	ID                uuid.UUID       `json:"id"`
	StudentID         uuid.UUID       `json:"studentId"`
	ScheduleID        uuid.UUID       `json:"scheduleId"`
	Status            AttendanceStatus `json:"status"`
	Date              time.Time       `json:"date"`
	UserLatitude      *float64        `json:"userLatitude,omitempty"`
	UserLongitude     *float64        `json:"userLongitude,omitempty"`
	DistanceMeters    *float64        `json:"distanceMeters,omitempty"`
	MaxDistanceMeters int             `json:"maxDistanceMeters"`
	CreatedAt         time.Time       `json:"createdAt"`
	UpdatedAt         time.Time       `json:"updatedAt"`
}

func NewAttendance(studentID, scheduleID uuid.UUID, status AttendanceStatus, date time.Time) *Attendance {
	now := time.Now()
	return &Attendance{
		ID:                uuid.New(),
		StudentID:         studentID,
		ScheduleID:        scheduleID,
		Status:            status,
		Date:              date,
		MaxDistanceMeters: DefaultMaxDistanceMeters,
		CreatedAt:         now,
		UpdatedAt:         now,
	}
}

// CalculateDistance calculates the distance between two points using Haversine formula
func CalculateDistance(lat1, lon1, lat2, lon2 float64) float64 {
	const R = 6371000 // Earth's radius in meters

	lat1Rad := lat1 * math.Pi / 180
	lon1Rad := lon1 * math.Pi / 180
	lat2Rad := lat2 * math.Pi / 180
	lon2Rad := lon2 * math.Pi / 180

	dlat := lat2Rad - lat1Rad
	dlon := lon2Rad - lon1Rad

	a := math.Sin(dlat/2)*math.Sin(dlat/2) + math.Cos(lat1Rad)*math.Cos(lat2Rad)*math.Sin(dlon/2)*math.Sin(dlon/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return R * c
}

// IsWithinDistance checks if the calculated distance is within the maximum allowed distance
func (a *Attendance) IsWithinDistance() bool {
	if a.DistanceMeters == nil {
		return false
	}
	fmt.Println(*a.DistanceMeters, a.MaxDistanceMeters)
	return *a.DistanceMeters <= float64(a.MaxDistanceMeters)
}

type Repository interface {
	Create(attendance *Attendance) error
	GetByID(id uuid.UUID) (*Attendance, error)
	GetByStudentID(studentID uuid.UUID) ([]*Attendance, error)
	GetByScheduleID(scheduleID uuid.UUID) ([]*Attendance, error)
	GetByDate(date time.Time) ([]*Attendance, error)
	Update(attendance *Attendance) error
	Delete(id uuid.UUID) error
	GetByStudentScheduleAndDate(studentID, scheduleID uuid.UUID, date time.Time) (*Attendance, error)
}

// ScheduleRepository interface for getting schedule information
type ScheduleRepository interface {
	GetByID(id uuid.UUID) (*Schedule, error)
}

// ClassroomRepository interface for getting classroom information
type ClassroomRepository interface {
	GetByID(id uuid.UUID) (*Classroom, error)
}

// Schedule represents a class schedule with classroom information
type Schedule struct {
	ID          uuid.UUID `json:"id"`
	ClassroomID uuid.UUID `json:"classroomId"`
}

// Classroom represents a classroom with location information
type Classroom struct {
	ID          uuid.UUID `json:"id"`
	LocationLat float64   `json:"locationLat"`
	LocationLng float64   `json:"locationLng"`
} 