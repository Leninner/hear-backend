package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type UserRepository interface {
	Create(ctx context.Context, user *User) error
	GetByID(ctx context.Context, id uuid.UUID) (*User, error)
	GetByEmail(ctx context.Context, email string) (*User, error)
	Update(ctx context.Context, user *User) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type CourseRepository interface {
	Create(ctx context.Context, course *Course) error
	GetByID(ctx context.Context, id uuid.UUID) (*Course, error)
	GetByFacultyID(ctx context.Context, facultyID uuid.UUID) ([]*Course, error)
	GetActive(ctx context.Context) ([]*Course, error)
	GetBySemester(ctx context.Context, semester, academicYear string) ([]*Course, error)
	Update(ctx context.Context, course *Course) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type CourseSectionRepository interface {
	Create(ctx context.Context, section *CourseSection) error
	GetByID(ctx context.Context, id uuid.UUID) (*CourseSection, error)
	GetByCourseID(ctx context.Context, courseID uuid.UUID) ([]*CourseSection, error)
	GetByTeacherID(ctx context.Context, teacherID uuid.UUID) ([]*CourseSection, error)
	Update(ctx context.Context, section *CourseSection) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type ClassScheduleRepository interface {
	Create(ctx context.Context, schedule *ClassSchedule) error
	GetByID(ctx context.Context, id uuid.UUID) (*ClassSchedule, error)
	GetByCourseID(ctx context.Context, courseID uuid.UUID) ([]*ClassSchedule, error)
	GetByClassroomAndTime(ctx context.Context, classroomID uuid.UUID, dayOfWeek int, startTime, endTime time.Time) ([]*ClassSchedule, error)
	Update(ctx context.Context, schedule *ClassSchedule) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type ClassroomRepository interface {
	Create(ctx context.Context, classroom *Classroom) error
	GetByID(ctx context.Context, id uuid.UUID) (*Classroom, error)
	GetByBuilding(ctx context.Context, building string) ([]*Classroom, error)
	GetByCapacity(ctx context.Context, minCapacity int) ([]*Classroom, error)
	GetNearby(ctx context.Context, lat, lng float64, maxDistanceMeters float64) ([]*Classroom, error)
	Update(ctx context.Context, classroom *Classroom) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type AttendanceRepository interface {
	Create(ctx context.Context, attendance *Attendance) error
	GetByID(ctx context.Context, id uuid.UUID) (*Attendance, error)
	GetByStudentID(ctx context.Context, studentID uuid.UUID, startDate, endDate time.Time) ([]*Attendance, error)
	GetByClassScheduleID(ctx context.Context, classScheduleID uuid.UUID, date time.Time) ([]*Attendance, error)
	Update(ctx context.Context, attendance *Attendance) error
}

type QRCodeRepository interface {
	Create(ctx context.Context, qrCode *QRCode) error
	GetByCode(ctx context.Context, code string) (*QRCode, error)
	GetByCourseID(ctx context.Context, courseID uuid.UUID) ([]*QRCode, error)
	Delete(ctx context.Context, id uuid.UUID) error
} 