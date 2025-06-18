package domain

import (
	"time"

	"github.com/google/uuid"
)

type Faculty struct {
	ID           uuid.UUID `json:"id"`
	UniversityID uuid.UUID `json:"universityId"`
	Name         string    `json:"name"`
	LocationLat  float64   `json:"locationLat"`
	LocationLng  float64   `json:"locationLng"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

func NewFaculty(universityID uuid.UUID, name string, locationLat, locationLng float64) *Faculty {
	now := time.Now()
	return &Faculty{
		ID:           uuid.New(),
		UniversityID: universityID,
		Name:         name,
		LocationLat:  locationLat,
		LocationLng:  locationLng,
		CreatedAt:    now,
		UpdatedAt:    now,
	}
}

type Repository interface {
	Create(faculty *Faculty) error
	GetByID(id uuid.UUID) (*Faculty, error)
	GetByName(name string) (*Faculty, error)
	GetAll() ([]*Faculty, error)
	GetByUniversityID(universityID uuid.UUID) ([]*Faculty, error)
	Update(faculty *Faculty) error
	Delete(id uuid.UUID) error
} 