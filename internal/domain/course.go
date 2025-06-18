package domain

import (
	"time"

	"github.com/google/uuid"
)

type Course struct {
	ID          uuid.UUID
	FacultyID   uuid.UUID
	Name        string
	Code        string
	Description string
	Credits     int
	Semester    string
	AcademicYear string
	IsActive    bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type CourseSection struct {
	ID           uuid.UUID
	CourseID     uuid.UUID
	SectionNumber int
	TeacherID    uuid.UUID
	MaxStudents  int
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type ClassSchedule struct {
	ID          uuid.UUID
	CourseID    uuid.UUID
	SectionID   uuid.UUID
	DayOfWeek   int
	StartTime   time.Time
	EndTime     time.Time
	ClassroomID uuid.UUID
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewCourse(facultyID uuid.UUID, name, code, description, semester, academicYear string, credits int) *Course {
	now := time.Now()
	return &Course{
		ID:           uuid.New(),
		FacultyID:    facultyID,
		Name:         name,
		Code:         code,
		Description:  description,
		Credits:      credits,
		Semester:     semester,
		AcademicYear: academicYear,
		IsActive:     true,
		CreatedAt:    now,
		UpdatedAt:    now,
	}
}

func NewCourseSection(courseID uuid.UUID, sectionNumber int, teacherID uuid.UUID, maxStudents int) *CourseSection {
	now := time.Now()
	return &CourseSection{
		ID:           uuid.New(),
		CourseID:     courseID,
		SectionNumber: sectionNumber,
		TeacherID:    teacherID,
		MaxStudents:  maxStudents,
		CreatedAt:    now,
		UpdatedAt:    now,
	}
}

func NewClassSchedule(courseID, sectionID uuid.UUID, dayOfWeek int, startTime, endTime time.Time, classroomID uuid.UUID) *ClassSchedule {
	now := time.Now()
	return &ClassSchedule{
		ID:          uuid.New(),
		CourseID:    courseID,
		SectionID:   sectionID,
		DayOfWeek:   dayOfWeek,
		StartTime:   startTime,
		EndTime:     endTime,
		ClassroomID: classroomID,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
}

func (cs *ClassSchedule) IsValidAttendanceTime(attendanceTime time.Time) bool {
	attendanceDay := int(attendanceTime.Weekday())
	if attendanceDay != cs.DayOfWeek {
		return false
	}

	attendanceHour := attendanceTime.Hour()
	attendanceMinute := attendanceTime.Minute()
	startHour := cs.StartTime.Hour()
	startMinute := cs.StartTime.Minute()
	endHour := cs.EndTime.Hour()
	endMinute := cs.EndTime.Minute()

	attendanceMinutes := attendanceHour*60 + attendanceMinute
	startMinutes := startHour*60 + startMinute
	endMinutes := endHour*60 + endMinute

	return attendanceMinutes >= startMinutes && attendanceMinutes <= endMinutes
} 