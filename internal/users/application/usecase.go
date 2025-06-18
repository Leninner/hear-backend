package application

import (
	"github.com/google/uuid"
	"github.com/leninner/hear-backend/internal/users/domain"
	"golang.org/x/crypto/bcrypt"
)

type UseCase struct {
	repository domain.Repository
}

func NewUseCase(repository domain.Repository) *UseCase {
	return &UseCase{
		repository: repository,
	}
}

func (uc *UseCase) CreateUser(dto *domain.CreateUserDTO) (*domain.User, error) {
	if err := dto.Validate(); err != nil {
		return nil, err
	}

	existingUser, err := uc.repository.GetByEmail(dto.Email)
	if err == nil && existingUser != nil {
		return nil, domain.ErrEmailExists
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(dto.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, domain.NewInternalError("failed to hash password", err)
	}

	user := domain.NewUser(dto.Email, string(hashedPassword), dto.FirstName, dto.LastName, dto.Role)
	if err := uc.repository.Create(user); err != nil {
		return nil, domain.NewInternalError("failed to create user in database", err)
	}

	return user, nil
}

func (uc *UseCase) UpdateUser(id uuid.UUID, dto *domain.UpdateUserDTO) error {
	if err := dto.Validate(); err != nil {
		return err
	}

	user, err := uc.repository.GetByID(id)
	if err != nil {
		return domain.ErrUserNotFound
	}

	if dto.FirstName != "" {
		user.FirstName = dto.FirstName
	}
	if dto.LastName != "" {
		user.LastName = dto.LastName
	}
	if dto.Role != "" {
		user.Role = dto.Role
	}

	if err := uc.repository.Update(user); err != nil {
		return domain.NewInternalError("failed to update user in database", err)
	}

	return nil
}

func (uc *UseCase) GetUser(id uuid.UUID) (*domain.User, error) {
	user, err := uc.repository.GetByID(id)
	if err != nil {
		return nil, domain.ErrUserNotFound
	}
	return user, nil
}

func (uc *UseCase) GetAllUsers() ([]*domain.User, error) {
	users, err := uc.repository.GetAll()
	if err != nil {
		return nil, domain.NewInternalError("failed to retrieve users from database", err)
	}
	return users, nil
}

func (uc *UseCase) GetUserByEmail(email string) (*domain.User, error) {
	user, err := uc.repository.GetByEmail(email)
	if err != nil {
		return nil, domain.ErrUserNotFound
	}
	return user, nil
} 