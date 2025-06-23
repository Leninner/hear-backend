package application

import (
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

func (uc *UseCase) CreateQRCode(dto *domain.CreateQRCodeDTO) (*domain.QRCode, error) {
	if err := dto.Validate(); err != nil {
		return nil, err
	}

	qrcode := domain.NewQRCode(dto.CourseSectionID, dto.Code, dto.ExpiresAt)
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

func (uc *UseCase) GetCourseSectionQRCodes(courseSectionID uuid.UUID) ([]*domain.QRCode, error) {
	return uc.repository.GetByCourseSectionID(courseSectionID)
}

func (uc *UseCase) GetActiveQRCode(courseSectionID uuid.UUID) (*domain.QRCode, error) {
	return uc.repository.GetActiveByCourseSectionID(courseSectionID)
}

func (uc *UseCase) UpdateQRCode(id uuid.UUID, dto *domain.UpdateQRCodeDTO) (*domain.QRCode, error) {
	if err := dto.Validate(); err != nil {
		return nil, err
	}

	qrcode, err := uc.repository.GetByID(id)
	if err != nil {
		return nil, err
	}

	if dto.Code != "" {
		qrcode.Code = dto.Code
	}
	if dto.ExpiresAt != nil {
		qrcode.ExpiresAt = *dto.ExpiresAt
	}

	err = uc.repository.Update(qrcode)
	if err != nil {
		return nil, err
	}

	return qrcode, nil
}

func (uc *UseCase) DeleteQRCode(id uuid.UUID) error {
	return uc.repository.Delete(id)
} 