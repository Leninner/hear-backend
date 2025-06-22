package domain

import (
	"time"

	"github.com/google/uuid"
)

type Faculty struct {
	ID           uuid.UUID `json:"id"`
	UniversityID uuid.UUID `json:"universityId"`
	Name         string    `json:"name"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

func NewFaculty(universityID uuid.UUID, name string) *Faculty {
	now := time.Now()
	return &Faculty{
		ID:           uuid.New(),
		UniversityID: universityID,
		Name:         name,
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