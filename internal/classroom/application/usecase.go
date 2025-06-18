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

func (u *UseCase) CreateClassroom(name, building string, floor, capacity int, locationLat, locationLng float64) (*domain.Classroom, error) {
	classroom := domain.NewClassroom(name, building, floor, capacity, locationLat, locationLng)
	if err := u.repository.Create(classroom); err != nil {
		return nil, err
	}
	return classroom, nil
}

func (u *UseCase) GetClassroom(id uuid.UUID) (*domain.Classroom, error) {
	return u.repository.GetByID(id)
}

func (u *UseCase) GetAllClassrooms() ([]*domain.Classroom, error) {
	return u.repository.GetAll()
}

func (u *UseCase) GetClassroomsByLocation(lat, lng, radius float64) ([]*domain.Classroom, error) {
	return u.repository.GetByLocation(lat, lng, radius)
}

func (u *UseCase) UpdateClassroom(classroom *domain.Classroom) error {
	return u.repository.Update(classroom)
}

func (u *UseCase) DeleteClassroom(id uuid.UUID) error {
	return u.repository.Delete(id)
} 