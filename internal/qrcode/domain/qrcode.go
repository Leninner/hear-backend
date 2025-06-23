package domain

import (
	"time"

	"github.com/google/uuid"
)

type QRCode struct {
	ID              uuid.UUID `json:"id"`
	CourseSectionID uuid.UUID `json:"courseSectionId"`
	Code            string    `json:"code"`
	ExpiresAt       time.Time `json:"expiresAt"`
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt"`
}

func NewQRCode(courseSectionID uuid.UUID, code string, expiresAt time.Time) *QRCode {
	now := time.Now()
	return &QRCode{
		ID:              uuid.New(),
		CourseSectionID: courseSectionID,
		Code:            code,
		ExpiresAt:       expiresAt,
		CreatedAt:       now,
		UpdatedAt:       now,
	}
}

func (q *QRCode) IsExpired() bool {
	return time.Now().After(q.ExpiresAt)
}

type Repository interface {
	Create(qrcode *QRCode) error
	GetByID(id uuid.UUID) (*QRCode, error)
	GetByCode(code string) (*QRCode, error)
	GetByCourseSectionID(courseSectionID uuid.UUID) ([]*QRCode, error)
	GetActiveByCourseSectionID(courseSectionID uuid.UUID) (*QRCode, error)
	Update(qrcode *QRCode) error
	Delete(id uuid.UUID) error
} 