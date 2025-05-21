package application

import (
	"github.com/google/uuid"
	"github.com/leninner/hear-backend/internal/users/domain"
)

type UseCase struct {
	repository domain.Repository
}

func NewUseCase(repository domain.Repository) *UseCase {
	return &UseCase{
		repository: repository,
	}
}

func (uc *UseCase) CreateUser(user *domain.User) error {
	return uc.repository.Create(user)
}

func (uc *UseCase) GetUser(id uuid.UUID) (*domain.User, error) {
	return uc.repository.GetByID(id)
}

func (uc *UseCase) GetAllUsers() ([]*domain.User, error) {
	return uc.repository.GetAll()
}

func (uc *UseCase) GetUserByEmail(email string) (*domain.User, error) {
	return uc.repository.GetByEmail(email)
} 