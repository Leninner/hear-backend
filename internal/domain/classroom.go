package domain

import (
	"time"

	"github.com/google/uuid"
)

type Classroom struct {
	ID          uuid.UUID
	Name        string
	Building    string
	Floor       int
	Capacity    int
	LocationLat float64
	LocationLng float64
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewClassroom(name, building string, floor, capacity int, locationLat, locationLng float64) *Classroom {
	now := time.Now()
	return &Classroom{
		ID:          uuid.New(),
		Name:        name,
		Building:    building,
		Floor:       floor,
		Capacity:    capacity,
		LocationLat: locationLat,
		LocationLng: locationLng,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
}

func (c *Classroom) FullName() string {
	return c.Building + " - " + c.Name
}

func (c *Classroom) ValidateLocation(lat, lng float64, maxDistanceMeters float64) bool {
	distance := calculateDistance(c.LocationLat, c.LocationLng, lat, lng)
	return distance <= maxDistanceMeters
} 