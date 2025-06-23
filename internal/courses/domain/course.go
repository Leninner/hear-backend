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
	// Course methods
	Create(course *Course) error
	GetByID(id uuid.UUID) (*Course, error)
	GetAll() ([]*Course, error)
	GetByFacultyID(facultyID uuid.UUID) ([]*Course, error)
	GetBySemester(semester string) ([]*Course, error)
	Update(course *Course) error
	Delete(id uuid.UUID) error
	
	// Section methods
	CreateSection(section *CourseSection) error
	GetSectionByID(id uuid.UUID) (*CourseSection, error)
	GetSectionsByCourseID(courseID uuid.UUID) ([]*CourseSection, error)
	GetSectionsByTeacherID(teacherID uuid.UUID) ([]*CourseSection, error)
	UpdateSection(section *CourseSection) error
	DeleteSection(id uuid.UUID) error
	
	// Enrollment methods
	EnrollStudent(sectionID, studentID uuid.UUID) error
	GetEnrollmentCount(sectionID uuid.UUID) (int, error)
	IsStudentEnrolled(sectionID, studentID uuid.UUID) (bool, error)
	UnenrollStudent(sectionID, studentID uuid.UUID) error
} 