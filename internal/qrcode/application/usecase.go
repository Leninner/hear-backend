package application

import (
	"time"

	"github.com/google/uuid"
	"github.com/leninner/hear-backend/internal/qrcode/domain"
)

type UseCase struct {
	repository domain.Repository
}

func NewUseCase(repository domain.Repository) *UseCase {
	return &UseCase{
		repository: repository,
	}
}

func (uc *UseCase) CreateQRCode(courseID uuid.UUID, code string, expiresAt time.Time) (*domain.QRCode, error) {
	qrcode := domain.NewQRCode(courseID, code, expiresAt)
	err := uc.repository.Create(qrcode)
	if err != nil {
		return nil, err
	}
	return qrcode, nil
}

func (uc *UseCase) GetQRCode(id uuid.UUID) (*domain.QRCode, error) {
	return uc.repository.GetByID(id)
}

func (uc *UseCase) GetQRCodeByCode(code string) (*domain.QRCode, error) {
	return uc.repository.GetByCode(code)
}

func (uc *UseCase) GetCourseQRCodes(courseID uuid.UUID) ([]*domain.QRCode, error) {
	return uc.repository.GetByCourseID(courseID)
}

func (uc *UseCase) GetActiveQRCode(courseID uuid.UUID) (*domain.QRCode, error) {
	return uc.repository.GetActiveByCourseID(courseID)
}

func (uc *UseCase) UpdateQRCode(qrcode *domain.QRCode) error {
	return uc.repository.Update(qrcode)
}

func (uc *UseCase) DeleteQRCode(id uuid.UUID) error {
	return uc.repository.Delete(id)
} 