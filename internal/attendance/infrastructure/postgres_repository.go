package infrastructure

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/leninner/hear-backend/internal/attendance/domain"
	"github.com/leninner/hear-backend/internal/infrastructure/db"
)

type PostgresRepository struct {
	db *db.Queries
}

func NewPostgresRepository(db *db.Queries) *PostgresRepository {
	return &PostgresRepository{
		db: db,
	}
}

func (r *PostgresRepository) Create(attendance *domain.Attendance) error {
	ctx := context.Background()
	
	params := db.CreateAttendanceParams{
		StudentID:       attendance.StudentID,
		ClassScheduleID: attendance.CourseID,
		AttendanceDate:  attendance.Date,
		ImageUrl:        "",
		ImageMetadata:   json.RawMessage("{}"),
		IsValid:         attendance.Status == domain.StatusPresent,
	}
	
	_, err := r.db.CreateAttendance(ctx, params)
	return err
}

func (r *PostgresRepository) GetByID(id uuid.UUID) (*domain.Attendance, error) {
	ctx := context.Background()
	
	dbAttendance, err := r.db.GetAttendanceByID(ctx, id)
	if err != nil {
		return nil, errors.New("attendance not found")
	}
	
	return r.mapToDomain(&dbAttendance), nil
}

func (r *PostgresRepository) GetByStudentID(studentID uuid.UUID) ([]*domain.Attendance, error) {
	ctx := context.Background()
	
	startDate := time.Now().AddDate(0, -6, 0)
	endDate := time.Now().AddDate(0, 6, 0)
	
	params := db.GetAttendanceByStudentIDParams{
		StudentID:        studentID,
		AttendanceDate:   startDate,
		AttendanceDate_2: endDate,
	}
	
	dbAttendances, err := r.db.GetAttendanceByStudentID(ctx, params)
	if err != nil {
		return nil, err
	}
	
	var attendances []*domain.Attendance
	for _, dbAttendance := range dbAttendances {
		attendances = append(attendances, r.mapToDomain(&dbAttendance))
	}
	
	return attendances, nil
}

func (r *PostgresRepository) GetByCourseID(courseID uuid.UUID) ([]*domain.Attendance, error) {
	ctx := context.Background()
	
	params := db.GetAttendanceByClassScheduleIDParams{
		ClassScheduleID: courseID,
		AttendanceDate:  time.Now(),
	}
	
	dbAttendances, err := r.db.GetAttendanceByClassScheduleID(ctx, params)
	if err != nil {
		return nil, err
	}
	
	var attendances []*domain.Attendance
	for _, dbAttendance := range dbAttendances {
		attendances = append(attendances, r.mapToDomain(&dbAttendance))
	}
	
	return attendances, nil
}

func (r *PostgresRepository) GetByDate(date time.Time) ([]*domain.Attendance, error) {
	ctx := context.Background()
	
	params := db.GetAttendanceByClassScheduleIDParams{
		ClassScheduleID: uuid.Nil,
		AttendanceDate:  date,
	}
	
	dbAttendances, err := r.db.GetAttendanceByClassScheduleID(ctx, params)
	if err != nil {
		return nil, err
	}
	
	var attendances []*domain.Attendance
	for _, dbAttendance := range dbAttendances {
		attendances = append(attendances, r.mapToDomain(&dbAttendance))
	}
	
	return attendances, nil
}

func (r *PostgresRepository) Update(attendance *domain.Attendance) error {
	ctx := context.Background()
	
	params := db.UpdateAttendanceParams{
		ID:      attendance.ID,
		IsValid: attendance.Status == domain.StatusPresent,
	}
	
	_, err := r.db.UpdateAttendance(ctx, params)
	return err
}

func (r *PostgresRepository) Delete(id uuid.UUID) error {
	return errors.New("delete operation not implemented in generated queries")
}

func (r *PostgresRepository) mapToDomain(dbAttendance *db.Attendance) *domain.Attendance {
	status := domain.StatusAbsent
	if dbAttendance.IsValid {
		status = domain.StatusPresent
	}
	
	attendance := &domain.Attendance{
		ID:        dbAttendance.ID,
		StudentID: dbAttendance.StudentID,
		CourseID:  dbAttendance.ClassScheduleID,
		Status:    status,
		Date:      dbAttendance.AttendanceDate,
	}
	
	if dbAttendance.CreatedAt.Valid {
		attendance.CreatedAt = dbAttendance.CreatedAt.Time
	}
	if dbAttendance.UpdatedAt.Valid {
		attendance.UpdatedAt = dbAttendance.UpdatedAt.Time
	}
	
	return attendance
} 