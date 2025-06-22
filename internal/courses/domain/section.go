package domain

import (
	"time"

	"github.com/google/uuid"
)

type CourseSection struct {
	ID          uuid.UUID `json:"id"`
	CourseID    uuid.UUID `json:"courseId"`
	Name        string    `json:"name"`
	TeacherID   uuid.UUID `json:"teacherId"`
	MaxStudents int       `json:"maxStudents"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func NewCourseSection(courseID, teacherID uuid.UUID, name string, maxStudents int) *CourseSection {
	now := time.Now()
	return &CourseSection{
		ID:          uuid.New(),
		CourseID:    courseID,
		TeacherID:   teacherID,
		Name:        name,
		MaxStudents: maxStudents,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
} 