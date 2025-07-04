// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0

package db

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type AttendanceStatus string

const (
	AttendanceStatusPresent AttendanceStatus = "present"
	AttendanceStatusAbsent  AttendanceStatus = "absent"
	AttendanceStatusLate    AttendanceStatus = "late"
)

func (e *AttendanceStatus) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = AttendanceStatus(s)
	case string:
		*e = AttendanceStatus(s)
	default:
		return fmt.Errorf("unsupported scan type for AttendanceStatus: %T", src)
	}
	return nil
}

type NullAttendanceStatus struct {
	AttendanceStatus AttendanceStatus `json:"attendance_status"`
	Valid            bool             `json:"valid"` // Valid is true if AttendanceStatus is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullAttendanceStatus) Scan(value interface{}) error {
	if value == nil {
		ns.AttendanceStatus, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.AttendanceStatus.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullAttendanceStatus) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.AttendanceStatus), nil
}

type UserRole string

const (
	UserRoleAdmin   UserRole = "admin"
	UserRoleTeacher UserRole = "teacher"
	UserRoleStudent UserRole = "student"
)

func (e *UserRole) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = UserRole(s)
	case string:
		*e = UserRole(s)
	default:
		return fmt.Errorf("unsupported scan type for UserRole: %T", src)
	}
	return nil
}

type NullUserRole struct {
	UserRole UserRole `json:"user_role"`
	Valid    bool     `json:"valid"` // Valid is true if UserRole is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullUserRole) Scan(value interface{}) error {
	if value == nil {
		ns.UserRole, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.UserRole.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullUserRole) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.UserRole), nil
}

type Attendance struct {
	ID                uuid.UUID        `json:"id"`
	StudentID         uuid.UUID        `json:"student_id"`
	ScheduleID        uuid.UUID        `json:"schedule_id"`
	Status            AttendanceStatus `json:"status"`
	Date              time.Time        `json:"date"`
	CreatedAt         sql.NullTime     `json:"created_at"`
	UpdatedAt         sql.NullTime     `json:"updated_at"`
	UserLatitude      sql.NullString   `json:"user_latitude"`
	UserLongitude     sql.NullString   `json:"user_longitude"`
	DistanceMeters    sql.NullString   `json:"distance_meters"`
	MaxDistanceMeters sql.NullInt32    `json:"max_distance_meters"`
}

type Classroom struct {
	ID          uuid.UUID     `json:"id"`
	Name        string        `json:"name"`
	Building    string        `json:"building"`
	Floor       int32         `json:"floor"`
	Capacity    int32         `json:"capacity"`
	LocationLat string        `json:"location_lat"`
	LocationLng string        `json:"location_lng"`
	CreatedAt   sql.NullTime  `json:"created_at"`
	UpdatedAt   sql.NullTime  `json:"updated_at"`
	FacultyID   uuid.NullUUID `json:"faculty_id"`
}

type Course struct {
	ID        uuid.UUID      `json:"id"`
	FacultyID uuid.UUID      `json:"faculty_id"`
	Name      string         `json:"name"`
	Semester  sql.NullString `json:"semester"`
	CreatedAt sql.NullTime   `json:"created_at"`
	UpdatedAt sql.NullTime   `json:"updated_at"`
}

type CourseSection struct {
	ID          uuid.UUID    `json:"id"`
	CourseID    uuid.UUID    `json:"course_id"`
	TeacherID   uuid.UUID    `json:"teacher_id"`
	MaxStudents int32        `json:"max_students"`
	CreatedAt   sql.NullTime `json:"created_at"`
	UpdatedAt   sql.NullTime `json:"updated_at"`
	Name        string       `json:"name"`
}

type Faculty struct {
	ID           uuid.UUID    `json:"id"`
	UniversityID uuid.UUID    `json:"university_id"`
	Name         string       `json:"name"`
	CreatedAt    sql.NullTime `json:"created_at"`
	UpdatedAt    sql.NullTime `json:"updated_at"`
}

type QrCode struct {
	ID              uuid.UUID    `json:"id"`
	CourseSectionID uuid.UUID    `json:"course_section_id"`
	Code            string       `json:"code"`
	ExpiresAt       time.Time    `json:"expires_at"`
	CreatedAt       sql.NullTime `json:"created_at"`
}

type RefreshToken struct {
	ID        uuid.UUID    `json:"id"`
	UserID    uuid.UUID    `json:"user_id"`
	Token     string       `json:"token"`
	ExpiresAt time.Time    `json:"expires_at"`
	CreatedAt sql.NullTime `json:"created_at"`
}

type Schedule struct {
	ID          uuid.UUID    `json:"id"`
	CourseID    uuid.UUID    `json:"course_id"`
	SectionID   uuid.UUID    `json:"section_id"`
	ClassroomID uuid.UUID    `json:"classroom_id"`
	DayOfWeek   int32        `json:"day_of_week"`
	StartTime   time.Time    `json:"start_time"`
	EndTime     time.Time    `json:"end_time"`
	CreatedAt   sql.NullTime `json:"created_at"`
	UpdatedAt   sql.NullTime `json:"updated_at"`
}

type SectionEnrollment struct {
	ID        uuid.UUID    `json:"id"`
	SectionID uuid.UUID    `json:"section_id"`
	StudentID uuid.UUID    `json:"student_id"`
	CreatedAt sql.NullTime `json:"created_at"`
	UpdatedAt sql.NullTime `json:"updated_at"`
}

type University struct {
	ID        uuid.UUID    `json:"id"`
	Name      string       `json:"name"`
	CreatedAt sql.NullTime `json:"created_at"`
	UpdatedAt sql.NullTime `json:"updated_at"`
}

type User struct {
	ID           uuid.UUID    `json:"id"`
	Email        string       `json:"email"`
	PasswordHash string       `json:"password_hash"`
	FirstName    string       `json:"first_name"`
	LastName     string       `json:"last_name"`
	Role         UserRole     `json:"role"`
	CreatedAt    sql.NullTime `json:"created_at"`
	UpdatedAt    sql.NullTime `json:"updated_at"`
}
