package domain

import (
	"time"

	"github.com/google/uuid"
)

type Classroom struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Building    string    `json:"building"`
	Floor       int       `json:"floor"`
	Capacity    int       `json:"capacity"`
	LocationLat float64   `json:"locationLat"`
	LocationLng float64   `json:"locationLng"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
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

type Repository interface {
	Create(classroom *Classroom) error
	GetByID(id uuid.UUID) (*Classroom, error)
	GetByName(name string) (*Classroom, error)
	GetAll() ([]*Classroom, error)
	GetByLocation(lat, lng float64, radius float64) ([]*Classroom, error)
	Update(classroom *Classroom) error
	Delete(id uuid.UUID) error
} 