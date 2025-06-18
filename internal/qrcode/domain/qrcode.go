package domain

import (
	"time"

	"github.com/google/uuid"
)

type QRCode struct {
	ID        uuid.UUID `json:"id"`
	CourseID  uuid.UUID `json:"courseId"`
	Code      string    `json:"code"`
	ExpiresAt time.Time `json:"expiresAt"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func NewQRCode(courseID uuid.UUID, code string, expiresAt time.Time) *QRCode {
	now := time.Now()
	return &QRCode{
		ID:        uuid.New(),
		CourseID:  courseID,
		Code:      code,
		ExpiresAt: expiresAt,
		CreatedAt: now,
		UpdatedAt: now,
	}
}

func (q *QRCode) IsExpired() bool {
	return time.Now().After(q.ExpiresAt)
}

type Repository interface {
	Create(qrcode *QRCode) error
	GetByID(id uuid.UUID) (*QRCode, error)
	GetByCode(code string) (*QRCode, error)
	GetByCourseID(courseID uuid.UUID) ([]*QRCode, error)
	GetActiveByCourseID(courseID uuid.UUID) (*QRCode, error)
	Update(qrcode *QRCode) error
	Delete(id uuid.UUID) error
} 