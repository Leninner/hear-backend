package domain

import (
	"crypto/rand"
	"encoding/base64"
	"time"

	"github.com/google/uuid"
)

type QRCode struct {
	ID        uuid.UUID
	CourseID  uuid.UUID
	Code      string
	ExpiresAt time.Time
	CreatedAt time.Time
}

func NewQRCode(courseID uuid.UUID, duration time.Duration) (*QRCode, error) {
	code, err := generateRandomCode()
	if err != nil {
		return nil, err
	}

	now := time.Now()
	return &QRCode{
		ID:        uuid.New(),
		CourseID:  courseID,
		Code:      code,
		ExpiresAt: now.Add(duration),
		CreatedAt: now,
	}, nil
}

func (qr *QRCode) IsExpired() bool {
	return time.Now().After(qr.ExpiresAt)
}

func generateRandomCode() (string, error) {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
} 