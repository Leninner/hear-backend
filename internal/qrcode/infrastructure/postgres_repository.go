package infrastructure

import (
	"context"
	"errors"

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
		CourseSectionID: qrcode.CourseSectionID,
		Code:            qrcode.Code,
		ExpiresAt:       qrcode.ExpiresAt,
	}
	
	_, err := r.db.CreateQRCode(ctx, params)
	return err
}

func (r *PostgresRepository) GetByID(id uuid.UUID) (*domain.QRCode, error) {
	ctx := context.Background()
	
	dbQRCode, err := r.db.GetQRCodeByID(ctx, id)
	if err != nil {
		return nil, errors.New("QR code not found")
	}
	
	return r.mapToDomain(&dbQRCode), nil
}

func (r *PostgresRepository) GetByCode(code string) (*domain.QRCode, error) {
	ctx := context.Background()
	
	dbQRCode, err := r.db.GetQRCodeByCode(ctx, code)
	if err != nil {
		return nil, errors.New("QR code not found")
	}
	
	return r.mapToDomain(&dbQRCode), nil
}

func (r *PostgresRepository) GetByCourseSectionID(courseSectionID uuid.UUID) ([]*domain.QRCode, error) {
	ctx := context.Background()
	
	dbQRCodes, err := r.db.GetQRCodesByCourseSectionID(ctx, courseSectionID)
	if err != nil {
		return nil, err
	}
	
	var qrcodes []*domain.QRCode
	for _, dbQRCode := range dbQRCodes {
		qrcodes = append(qrcodes, r.mapToDomain(&dbQRCode))
	}
	
	return qrcodes, nil
}

func (r *PostgresRepository) GetActiveByCourseSectionID(courseSectionID uuid.UUID) (*domain.QRCode, error) {
	ctx := context.Background()
	
	dbQRCode, err := r.db.GetActiveQRCodeByCourseSectionID(ctx, courseSectionID)
	if err != nil {
		return nil, errors.New("no active QR code found")
	}
	
	return r.mapToDomain(&dbQRCode), nil
}

func (r *PostgresRepository) Update(qrcode *domain.QRCode) error {
	ctx := context.Background()
	
	params := db.UpdateQRCodeParams{
		ID:        qrcode.ID,
		Code:      qrcode.Code,
		ExpiresAt: qrcode.ExpiresAt,
	}
	
	_, err := r.db.UpdateQRCode(ctx, params)
	return err
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
		ID:              dbQRCode.ID,
		CourseSectionID: dbQRCode.CourseSectionID,
		Code:            dbQRCode.Code,
		ExpiresAt:       dbQRCode.ExpiresAt,
	}
	
	if dbQRCode.CreatedAt.Valid {
		qrcode.CreatedAt = dbQRCode.CreatedAt.Time
	}
	
	return qrcode
} 