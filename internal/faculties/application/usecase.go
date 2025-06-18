package application

import (
	"github.com/google/uuid"
	"github.com/leninner/hear-backend/internal/faculties/domain"
)

type UseCase struct {
	repository domain.Repository
}

func NewUseCase(repository domain.Repository) *UseCase {
	return &UseCase{
		repository: repository,
	}
}

func (uc *UseCase) CreateFaculty(dto *domain.CreateFacultyDTO) (*domain.Faculty, error) {
	if err := dto.Validate(); err != nil {
		return nil, err
	}

	faculty := domain.NewFaculty(dto.UniversityID, dto.Name, dto.LocationLat, dto.LocationLng)
	if err := uc.repository.Create(faculty); err != nil {
		return nil, domain.NewInternalError("failed to create faculty in database", err)
	}

	return faculty, nil
}

func (uc *UseCase) GetFaculty(id uuid.UUID) (*domain.Faculty, error) {
	faculty, err := uc.repository.GetByID(id)
	if err != nil {
		return nil, domain.ErrFacultyNotFound
	}
	return faculty, nil
}

func (uc *UseCase) GetAllFaculties() ([]*domain.Faculty, error) {
	faculties, err := uc.repository.GetAll()
	if err != nil {
		return nil, domain.NewInternalError("failed to retrieve faculties from database", err)
	}
	return faculties, nil
}

func (uc *UseCase) GetFacultiesByUniversity(universityID uuid.UUID) ([]*domain.Faculty, error) {
	faculties, err := uc.repository.GetByUniversityID(universityID)
	if err != nil {
		return nil, domain.NewInternalError("failed to retrieve faculties from database", err)
	}
	return faculties, nil
}

func (uc *UseCase) UpdateFaculty(id uuid.UUID, dto *domain.UpdateFacultyDTO) error {
	if err := dto.Validate(); err != nil {
		return err
	}

	faculty, err := uc.repository.GetByID(id)
	if err != nil {
		return domain.ErrFacultyNotFound
	}

	if dto.UniversityID != nil {
		faculty.UniversityID = *dto.UniversityID
	}
	if dto.Name != "" {
		faculty.Name = dto.Name
	}
	if dto.LocationLat != nil {
		faculty.LocationLat = *dto.LocationLat
	}
	if dto.LocationLng != nil {
		faculty.LocationLng = *dto.LocationLng
	}

	if err := uc.repository.Update(faculty); err != nil {
		return domain.NewInternalError("failed to update faculty in database", err)
	}

	return nil
}

func (uc *UseCase) DeleteFaculty(id uuid.UUID) error {
	if err := uc.repository.Delete(id); err != nil {
		return domain.ErrFacultyNotFound
	}
	return nil
} 