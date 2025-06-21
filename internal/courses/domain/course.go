package domain

import (
	"time"

	"github.com/google/uuid"
)

type Course struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	FacultyID uuid.UUID `json:"facultyId"`
	Semester  string    `json:"semester"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func NewCourse(name string, facultyID uuid.UUID, semester string) *Course {
	now := time.Now()
	return &Course{
		ID:        uuid.New(),
		Name:      name,
		FacultyID: facultyID,
		Semester:  semester,
		CreatedAt: now,
		UpdatedAt: now,
	}
}

type Repository interface {
	Create(course *Course) error
	GetByID(id uuid.UUID) (*Course, error)
	GetAll() ([]*Course, error)
	GetByFacultyID(facultyID uuid.UUID) ([]*Course, error)
	GetBySemester(semester string) ([]*Course, error)
	Update(course *Course) error
	Delete(id uuid.UUID) error
} 