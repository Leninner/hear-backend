package application

import (
	"github.com/google/uuid"
	"github.com/leninner/hear-backend/internal/classroom/domain"
)

type UseCase struct {
	repository domain.Repository
}

func NewUseCase(repository domain.Repository) *UseCase {
	return &UseCase{
		repository: repository,
	}
}

func (uc *UseCase) CreateClassroom(dto *domain.CreateClassroomDTO) (*domain.Classroom, error) {
	if err := dto.Validate(); err != nil {
		return nil, err
	}

	existingClassroom, err := uc.repository.GetByName(dto.Name)
	if err == nil && existingClassroom != nil {
		return nil, domain.ErrClassroomNameExists
	}

	classroom := domain.NewClassroom(
		dto.Name, 
		dto.Building, 
		*dto.Floor, 
		*dto.Capacity, 
		*dto.LocationLat, 
		*dto.LocationLng,
	)
	if err := uc.repository.Create(classroom); err != nil {
		return nil, domain.NewInternalError("failed to create classroom in database", err)
	}

	return classroom, nil
}

func (uc *UseCase) UpdateClassroom(id uuid.UUID, dto *domain.UpdateClassroomDTO) error {
	if err := dto.Validate(); err != nil {
		return err
	}

	classroom, err := uc.repository.GetByID(id)
	if err != nil {
		return domain.ErrClassroomNotFound
	}

	if dto.Name != "" {
		// Check if the new name conflicts with another classroom
		if dto.Name != classroom.Name {
			existingClassroom, err := uc.repository.GetByName(dto.Name)
			if err == nil && existingClassroom != nil {
				return domain.ErrClassroomNameExists
			}
		}
		classroom.Name = dto.Name
	}

	if dto.Building != "" {
		classroom.Building = dto.Building
	}

	if dto.Floor != nil {
		classroom.Floor = *dto.Floor
	}

	if dto.Capacity != nil {
		classroom.Capacity = *dto.Capacity
	}

	if dto.LocationLat != nil {
		classroom.LocationLat = *dto.LocationLat
	}

	if dto.LocationLng != nil {
		classroom.LocationLng = *dto.LocationLng
	}

	if err := uc.repository.Update(classroom); err != nil {
		return domain.NewInternalError("failed to update classroom in database", err)
	}

	return nil
}

func (uc *UseCase) GetClassroom(id uuid.UUID) (*domain.Classroom, error) {
	classroom, err := uc.repository.GetByID(id)
	if err != nil {
		return nil, domain.ErrClassroomNotFound
	}
	return classroom, nil
}

func (uc *UseCase) GetAllClassrooms() ([]*domain.Classroom, error) {
	classrooms, err := uc.repository.GetAll()
	if err != nil {
		return nil, domain.NewInternalError("failed to retrieve classrooms from database", err)
	}
	return classrooms, nil
}

func (uc *UseCase) GetClassroomsByLocation(lat, lng, radius float64) ([]*domain.Classroom, error) {
	classrooms, err := uc.repository.GetByLocation(lat, lng, radius)
	if err != nil {
		return nil, domain.NewInternalError("failed to retrieve classrooms by location from database", err)
	}
	return classrooms, nil
}

func (uc *UseCase) DeleteClassroom(id uuid.UUID) error {
	// Check if classroom exists before deleting
	_, err := uc.repository.GetByID(id)
	if err != nil {
		return domain.ErrClassroomNotFound
	}

	if err := uc.repository.Delete(id); err != nil {
		return domain.NewInternalError("failed to delete classroom from database", err)
	}

	return nil
} 