package infrastructure

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/leninner/hear-backend/internal/infrastructure/db"
	"github.com/leninner/hear-backend/internal/qrcode/domain"
)

type PostgresRepository struct {
	db *db.Queries
}

func NewPostgresRepository(db *db.Queries) *PostgresRepository {
	return &PostgresRepository{
		db: db,
	}
}

func (r *PostgresRepository) Create(qrcode *domain.QRCode) error {
	ctx := context.Background()
	
	params := db.CreateQRCodeParams{
		CourseID:  qrcode.CourseID,
		Code:      qrcode.Code,
		ExpiresAt: qrcode.ExpiresAt,
	}
	
	_, err := r.db.CreateQRCode(ctx, params)
	return err
}

func (r *PostgresRepository) GetByID(id uuid.UUID) (*domain.QRCode, error) {
	ctx := context.Background()
	
	dbQRCodes, err := r.db.GetQRCodesByCourseID(ctx, id)
	if err != nil {
		return nil, errors.New("QR code not found")
	}
	
	if len(dbQRCodes) == 0 {
		return nil, errors.New("QR code not found")
	}
	
	return r.mapToDomain(&dbQRCodes[0]), nil
}

func (r *PostgresRepository) GetByCode(code string) (*domain.QRCode, error) {
	ctx := context.Background()
	
	dbQRCode, err := r.db.GetQRCodeByCode(ctx, code)
	if err != nil {
		return nil, errors.New("QR code not found")
	}
	
	return r.mapToDomain(&dbQRCode), nil
}

func (r *PostgresRepository) GetByCourseID(courseID uuid.UUID) ([]*domain.QRCode, error) {
	ctx := context.Background()
	
	dbQRCodes, err := r.db.GetQRCodesByCourseID(ctx, courseID)
	if err != nil {
		return nil, err
	}
	
	var qrcodes []*domain.QRCode
	for _, dbQRCode := range dbQRCodes {
		qrcodes = append(qrcodes, r.mapToDomain(&dbQRCode))
	}
	
	return qrcodes, nil
}

func (r *PostgresRepository) GetActiveByCourseID(courseID uuid.UUID) (*domain.QRCode, error) {
	ctx := context.Background()
	
	dbQRCodes, err := r.db.GetQRCodesByCourseID(ctx, courseID)
	if err != nil {
		return nil, err
	}
	
	now := time.Now()
	for _, dbQRCode := range dbQRCodes {
		if dbQRCode.ExpiresAt.After(now) {
			return r.mapToDomain(&dbQRCode), nil
		}
	}
	
	return nil, errors.New("no active QR code found")
}

func (r *PostgresRepository) Update(qrcode *domain.QRCode) error {
	return errors.New("update operation not implemented in generated queries")
}

func (r *PostgresRepository) Delete(id uuid.UUID) error {
	ctx := context.Background()
	
	err := r.db.DeleteQRCode(ctx, id)
	if err != nil {
		return errors.New("QR code not found")
	}
	
	return nil
}

func (r *PostgresRepository) mapToDomain(dbQRCode *db.QrCode) *domain.QRCode {
	qrcode := &domain.QRCode{
		ID:        dbQRCode.ID,
		CourseID:  dbQRCode.CourseID,
		Code:      dbQRCode.Code,
		ExpiresAt: dbQRCode.ExpiresAt,
	}
	
	if dbQRCode.CreatedAt.Valid {
		qrcode.CreatedAt = dbQRCode.CreatedAt.Time
	}
	
	return qrcode
} 