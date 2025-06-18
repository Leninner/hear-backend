package infrastructure

import (
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/leninner/hear-backend/internal/qrcode/domain"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) *PostgresRepository {
	return &PostgresRepository{
		db: db,
	}
}

func (r *PostgresRepository) Create(qrcode *domain.QRCode) error {
	query := `
		INSERT INTO qrcodes (id, course_id, code, expires_at, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6)
	`
	_, err := r.db.Exec(query,
		qrcode.ID,
		qrcode.CourseID,
		qrcode.Code,
		qrcode.ExpiresAt,
		qrcode.CreatedAt,
		qrcode.UpdatedAt,
	)
	return err
}

func (r *PostgresRepository) GetByID(id uuid.UUID) (*domain.QRCode, error) {
	query := `
		SELECT id, course_id, code, expires_at, created_at, updated_at
		FROM qrcodes
		WHERE id = $1
	`
	qrcode := &domain.QRCode{}
	err := r.db.QueryRow(query, id).Scan(
		&qrcode.ID,
		&qrcode.CourseID,
		&qrcode.Code,
		&qrcode.ExpiresAt,
		&qrcode.CreatedAt,
		&qrcode.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, errors.New("QR code not found")
	}
	return qrcode, err
}

func (r *PostgresRepository) GetByCode(code string) (*domain.QRCode, error) {
	query := `
		SELECT id, course_id, code, expires_at, created_at, updated_at
		FROM qrcodes
		WHERE code = $1
	`
	qrcode := &domain.QRCode{}
	err := r.db.QueryRow(query, code).Scan(
		&qrcode.ID,
		&qrcode.CourseID,
		&qrcode.Code,
		&qrcode.ExpiresAt,
		&qrcode.CreatedAt,
		&qrcode.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, errors.New("QR code not found")
	}
	return qrcode, err
}

func (r *PostgresRepository) GetByCourseID(courseID uuid.UUID) ([]*domain.QRCode, error) {
	query := `
		SELECT id, course_id, code, expires_at, created_at, updated_at
		FROM qrcodes
		WHERE course_id = $1
	`
	rows, err := r.db.Query(query, courseID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var qrcodes []*domain.QRCode
	for rows.Next() {
		qrcode := &domain.QRCode{}
		err := rows.Scan(
			&qrcode.ID,
			&qrcode.CourseID,
			&qrcode.Code,
			&qrcode.ExpiresAt,
			&qrcode.CreatedAt,
			&qrcode.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		qrcodes = append(qrcodes, qrcode)
	}
	return qrcodes, nil
}

func (r *PostgresRepository) GetActiveByCourseID(courseID uuid.UUID) (*domain.QRCode, error) {
	query := `
		SELECT id, course_id, code, expires_at, created_at, updated_at
		FROM qrcodes
		WHERE course_id = $1 AND expires_at > $2
		ORDER BY created_at DESC
		LIMIT 1
	`
	qrcode := &domain.QRCode{}
	err := r.db.QueryRow(query, courseID, time.Now()).Scan(
		&qrcode.ID,
		&qrcode.CourseID,
		&qrcode.Code,
		&qrcode.ExpiresAt,
		&qrcode.CreatedAt,
		&qrcode.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, errors.New("no active QR code found")
	}
	return qrcode, err
}

func (r *PostgresRepository) Update(qrcode *domain.QRCode) error {
	query := `
		UPDATE qrcodes
		SET course_id = $1, code = $2, expires_at = $3, updated_at = $4
		WHERE id = $5
	`
	result, err := r.db.Exec(query,
		qrcode.CourseID,
		qrcode.Code,
		qrcode.ExpiresAt,
		qrcode.UpdatedAt,
		qrcode.ID,
	)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("QR code not found")
	}
	return nil
}

func (r *PostgresRepository) Delete(id uuid.UUID) error {
	query := `DELETE FROM qrcodes WHERE id = $1`
	result, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("QR code not found")
	}
	return nil
} 