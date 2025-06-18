package domain

import (
	"time"

	"github.com/google/uuid"
)

type Course struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	TeacherID   uuid.UUID `json:"teacherId"`
	FacultyID   uuid.UUID `json:"facultyId"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func NewCourse(name, description string, teacherID, facultyID uuid.UUID) *Course {
	now := time.Now()
	return &Course{
		ID:          uuid.New(),
		Name:        name,
		Description: description,
		TeacherID:   teacherID,
		FacultyID:   facultyID,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
}

type Repository interface {
	Create(course *Course) error
	GetByID(id uuid.UUID) (*Course, error)
	GetAll() ([]*Course, error)
	GetByTeacherID(teacherID uuid.UUID) ([]*Course, error)
	Update(course *Course) error
	Delete(id uuid.UUID) error
} 