package domain

import (
	"math"
	"time"

	"github.com/google/uuid"
)

type ImageMetadata struct {
	Timestamp time.Time `json:"timestamp"`
	Latitude  float64   `json:"latitude"`
	Longitude float64   `json:"longitude"`
}

type Attendance struct {
	ID              uuid.UUID
	StudentID       uuid.UUID
	ClassScheduleID uuid.UUID
	AttendanceDate  time.Time
	ImageURL        string
	ImageMetadata   ImageMetadata
	IsValid         bool
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

func NewAttendance(studentID, classScheduleID uuid.UUID, imageURL string, metadata ImageMetadata) *Attendance {
	now := time.Now()
	return &Attendance{
		ID:              uuid.New(),
		StudentID:       studentID,
		ClassScheduleID: classScheduleID,
		AttendanceDate:  now,
		ImageURL:        imageURL,
		ImageMetadata:   metadata,
		IsValid:         false,
		CreatedAt:       now,
		UpdatedAt:       now,
	}
}

func (a *Attendance) ValidateLocation(facultyLat, facultyLng float64, maxDistanceMeters float64) bool {
	distance := calculateDistance(a.ImageMetadata.Latitude, a.ImageMetadata.Longitude, facultyLat, facultyLng)
	return distance <= maxDistanceMeters
}

func calculateDistance(lat1, lon1, lat2, lon2 float64) float64 {
	const R = 6371000
	φ1 := lat1 * (3.141592653589793 / 180)
	φ2 := lat2 * (3.141592653589793 / 180)
	Δφ := (lat2 - lat1) * (3.141592653589793 / 180)
	Δλ := (lon2 - lon1) * (3.141592653589793 / 180)

	a := (Δφ/2)*(Δφ/2) + φ1*φ2*(Δλ/2)*(Δλ/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return R * c
} 