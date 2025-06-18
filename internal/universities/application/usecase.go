package application

import (
	"github.com/google/uuid"
	"github.com/leninner/hear-backend/internal/universities/domain"
)

type UseCase struct {
	repository domain.Repository
}

func NewUseCase(repository domain.Repository) *UseCase {
	return &UseCase{
		repository: repository,
	}
}

func (uc *UseCase) CreateUniversity(dto *domain.CreateUniversityDTO) (*domain.University, error) {
	if err := dto.Validate(); err != nil {
		return nil, err
	}

	university := domain.NewUniversity(dto.Name)
	if err := uc.repository.Create(university); err != nil {
		return nil, domain.NewInternalError("failed to create university in database", err)
	}

	return university, nil
}

func (uc *UseCase) GetUniversity(id uuid.UUID) (*domain.University, error) {
	university, err := uc.repository.GetByID(id)
	if err != nil {
		return nil, domain.ErrUniversityNotFound
	}
	return university, nil
}

func (uc *UseCase) GetAllUniversities() ([]*domain.University, error) {
	universities, err := uc.repository.GetAll()
	if err != nil {
		return nil, domain.NewInternalError("failed to retrieve universities from database", err)
	}
	return universities, nil
}

func (uc *UseCase) UpdateUniversity(id uuid.UUID, dto *domain.UpdateUniversityDTO) error {
	if err := dto.Validate(); err != nil {
		return err
	}

	university, err := uc.repository.GetByID(id)
	if err != nil {
		return domain.ErrUniversityNotFound
	}

	if dto.Name != "" {
		university.Name = dto.Name
	}

	if err := uc.repository.Update(university); err != nil {
		return domain.NewInternalError("failed to update university in database", err)
	}

	return nil
}

func (uc *UseCase) DeleteUniversity(id uuid.UUID) error {
	if err := uc.repository.Delete(id); err != nil {
		return domain.ErrUniversityNotFound
	}
	return nil
} 