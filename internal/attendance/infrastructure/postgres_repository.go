package infrastructure

import (
	"context"
	"database/sql"
	"errors"
	"strconv"
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
	
	// Convert domain types to SQL types
	var userLat, userLng, distanceMeters sql.NullString
	var maxDistance sql.NullInt32
	
	if attendance.UserLatitude != nil {
		userLat.String = strconv.FormatFloat(*attendance.UserLatitude, 'f', 8, 64)
		userLat.Valid = true
	}
	
	if attendance.UserLongitude != nil {
		userLng.String = strconv.FormatFloat(*attendance.UserLongitude, 'f', 8, 64)
		userLng.Valid = true
	}
	
	if attendance.DistanceMeters != nil {
		distanceMeters.String = strconv.FormatFloat(*attendance.DistanceMeters, 'f', 2, 64)
		distanceMeters.Valid = true
	}
	
	maxDistance.Int32 = int32(attendance.MaxDistanceMeters)
	maxDistance.Valid = true
	
	params := db.CreateAttendanceParams{
		StudentID:         attendance.StudentID,
		ClassScheduleID:   attendance.ClassScheduleID,
		Status:            db.AttendanceStatus(attendance.Status),
		Date:              attendance.Date,
		UserLatitude:      userLat,
		UserLongitude:     userLng,
		DistanceMeters:    distanceMeters,
		MaxDistanceMeters: maxDistance,
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
	
	// Convert domain types to SQL types
	var userLat, userLng, distanceMeters sql.NullString
	var maxDistance sql.NullInt32
	
	if attendance.UserLatitude != nil {
		userLat.String = strconv.FormatFloat(*attendance.UserLatitude, 'f', 8, 64)
		userLat.Valid = true
	}
	
	if attendance.UserLongitude != nil {
		userLng.String = strconv.FormatFloat(*attendance.UserLongitude, 'f', 8, 64)
		userLng.Valid = true
	}
	
	if attendance.DistanceMeters != nil {
		distanceMeters.String = strconv.FormatFloat(*attendance.DistanceMeters, 'f', 2, 64)
		distanceMeters.Valid = true
	}
	
	maxDistance.Int32 = int32(attendance.MaxDistanceMeters)
	maxDistance.Valid = true
	
	params := db.UpdateAttendanceParams{
		ID:                attendance.ID,
		Status:            db.AttendanceStatus(attendance.Status),
		UserLatitude:      userLat,
		UserLongitude:     userLng,
		DistanceMeters:    distanceMeters,
		MaxDistanceMeters: maxDistance,
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
		ID:                dbAttendance.ID,
		StudentID:         dbAttendance.StudentID,
		ClassScheduleID:   dbAttendance.ClassScheduleID,
		Status:            domain.AttendanceStatus(dbAttendance.Status),
		Date:              dbAttendance.Date,
		MaxDistanceMeters: int(dbAttendance.MaxDistanceMeters.Int32),
	}
	
	// Convert location fields from SQL types to domain types
	if dbAttendance.UserLatitude.Valid {
		if lat, err := strconv.ParseFloat(dbAttendance.UserLatitude.String, 64); err == nil {
			attendance.UserLatitude = &lat
		}
	}
	
	if dbAttendance.UserLongitude.Valid {
		if lng, err := strconv.ParseFloat(dbAttendance.UserLongitude.String, 64); err == nil {
			attendance.UserLongitude = &lng
		}
	}
	
	if dbAttendance.DistanceMeters.Valid {
		if distance, err := strconv.ParseFloat(dbAttendance.DistanceMeters.String, 64); err == nil {
			attendance.DistanceMeters = &distance
		}
	}
	
	if dbAttendance.CreatedAt.Valid {
		attendance.CreatedAt = dbAttendance.CreatedAt.Time
	}
	if dbAttendance.UpdatedAt.Valid {
		attendance.UpdatedAt = dbAttendance.UpdatedAt.Time
	}
	
	return attendance
} 