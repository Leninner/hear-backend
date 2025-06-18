package infrastructure

import (
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/leninner/hear-backend/internal/attendance/domain"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) *PostgresRepository {
	return &PostgresRepository{
		db: db,
	}
}

func (r *PostgresRepository) Create(attendance *domain.Attendance) error {
	query := `
		INSERT INTO attendances (id, student_id, course_id, status, date, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`
	_, err := r.db.Exec(query,
		attendance.ID,
		attendance.StudentID,
		attendance.CourseID,
		attendance.Status,
		attendance.Date,
		attendance.CreatedAt,
		attendance.UpdatedAt,
	)
	return err
}

func (r *PostgresRepository) GetByID(id uuid.UUID) (*domain.Attendance, error) {
	query := `
		SELECT id, student_id, course_id, status, date, created_at, updated_at
		FROM attendances
		WHERE id = $1
	`
	attendance := &domain.Attendance{}
	err := r.db.QueryRow(query, id).Scan(
		&attendance.ID,
		&attendance.StudentID,
		&attendance.CourseID,
		&attendance.Status,
		&attendance.Date,
		&attendance.CreatedAt,
		&attendance.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, errors.New("attendance not found")
	}
	return attendance, err
}

func (r *PostgresRepository) GetByStudentID(studentID uuid.UUID) ([]*domain.Attendance, error) {
	query := `
		SELECT id, student_id, course_id, status, date, created_at, updated_at
		FROM attendances
		WHERE student_id = $1
	`
	rows, err := r.db.Query(query, studentID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var attendances []*domain.Attendance
	for rows.Next() {
		attendance := &domain.Attendance{}
		err := rows.Scan(
			&attendance.ID,
			&attendance.StudentID,
			&attendance.CourseID,
			&attendance.Status,
			&attendance.Date,
			&attendance.CreatedAt,
			&attendance.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		attendances = append(attendances, attendance)
	}
	return attendances, nil
}

func (r *PostgresRepository) GetByCourseID(courseID uuid.UUID) ([]*domain.Attendance, error) {
	query := `
		SELECT id, student_id, course_id, status, date, created_at, updated_at
		FROM attendances
		WHERE course_id = $1
	`
	rows, err := r.db.Query(query, courseID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var attendances []*domain.Attendance
	for rows.Next() {
		attendance := &domain.Attendance{}
		err := rows.Scan(
			&attendance.ID,
			&attendance.StudentID,
			&attendance.CourseID,
			&attendance.Status,
			&attendance.Date,
			&attendance.CreatedAt,
			&attendance.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		attendances = append(attendances, attendance)
	}
	return attendances, nil
}

func (r *PostgresRepository) GetByDate(date time.Time) ([]*domain.Attendance, error) {
	query := `
		SELECT id, student_id, course_id, status, date, created_at, updated_at
		FROM attendances
		WHERE date::date = $1::date
	`
	rows, err := r.db.Query(query, date)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var attendances []*domain.Attendance
	for rows.Next() {
		attendance := &domain.Attendance{}
		err := rows.Scan(
			&attendance.ID,
			&attendance.StudentID,
			&attendance.CourseID,
			&attendance.Status,
			&attendance.Date,
			&attendance.CreatedAt,
			&attendance.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		attendances = append(attendances, attendance)
	}
	return attendances, nil
}

func (r *PostgresRepository) Update(attendance *domain.Attendance) error {
	query := `
		UPDATE attendances
		SET student_id = $1, course_id = $2, status = $3, date = $4, updated_at = $5
		WHERE id = $6
	`
	result, err := r.db.Exec(query,
		attendance.StudentID,
		attendance.CourseID,
		attendance.Status,
		attendance.Date,
		attendance.UpdatedAt,
		attendance.ID,
	)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("attendance not found")
	}
	return nil
}

func (r *PostgresRepository) Delete(id uuid.UUID) error {
	query := `DELETE FROM attendances WHERE id = $1`
	result, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("attendance not found")
	}
	return nil
} 