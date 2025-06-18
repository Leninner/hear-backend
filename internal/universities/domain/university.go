package domain

import (
	"time"

	"github.com/google/uuid"
)

type University struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func NewUniversity(name string) *University {
	now := time.Now()
	return &University{
		ID:        uuid.New(),
		Name:      name,
		CreatedAt: now,
		UpdatedAt: now,
	}
}

type Repository interface {
	Create(university *University) error
	GetByID(id uuid.UUID) (*University, error)
	GetByName(name string) (*University, error)
	GetAll() ([]*University, error)
	Update(university *University) error
	Delete(id uuid.UUID) error
} 