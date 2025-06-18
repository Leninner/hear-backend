package infrastructure

import (
	"context"
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
		ClassScheduleID: attendance.ClassScheduleID,
		Status:          db.AttendanceStatus(attendance.Status),
		Date:            attendance.Date,
	}
	
	result, err := r.db.CreateAttendance(ctx, params)
	if err != nil {
		return err
	}
	
	attendance.ID = result.ID
	if result.CreatedAt.Valid {
		attendance.CreatedAt = result.CreatedAt.Time
	}
	if result.UpdatedAt.Valid {
		attendance.UpdatedAt = result.UpdatedAt.Time
	}
	
	return nil
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
		StudentID: studentID,
		Date:      startDate,
		Date_2:    endDate,
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

func (r *PostgresRepository) GetByClassScheduleID(classScheduleID uuid.UUID) ([]*domain.Attendance, error) {
	ctx := context.Background()
	
	params := db.GetAttendanceByClassScheduleIDParams{
		ClassScheduleID: classScheduleID,
		Date:            time.Now(),
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
		Date:            date,
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
		ID:     attendance.ID,
		Status: db.AttendanceStatus(attendance.Status),
	}
	
	result, err := r.db.UpdateAttendance(ctx, params)
	if err != nil {
		return err
	}
	
	if result.UpdatedAt.Valid {
		attendance.UpdatedAt = result.UpdatedAt.Time
	}
	
	return nil
}

func (r *PostgresRepository) Delete(id uuid.UUID) error {
	return errors.New("delete operation not implemented in generated queries")
}

func (r *PostgresRepository) mapToDomain(dbAttendance *db.Attendance) *domain.Attendance {
	attendance := &domain.Attendance{
		ID:             dbAttendance.ID,
		StudentID:      dbAttendance.StudentID,
		ClassScheduleID: dbAttendance.ClassScheduleID,
		Status:         domain.AttendanceStatus(dbAttendance.Status),
		Date:           dbAttendance.Date,
	}
	
	if dbAttendance.CreatedAt.Valid {
		attendance.CreatedAt = dbAttendance.CreatedAt.Time
	}
	if dbAttendance.UpdatedAt.Valid {
		attendance.UpdatedAt = dbAttendance.UpdatedAt.Time
	}
	
	return attendance
} 