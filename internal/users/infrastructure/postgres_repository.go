package infrastructure

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/leninner/hear-backend/internal/infrastructure/db"
	"github.com/leninner/hear-backend/internal/users/domain"
)

type PostgresRepository struct {
	queries *db.Queries
}

func NewPostgresRepository(sqlDB *sql.DB) *PostgresRepository {
	return &PostgresRepository{
		queries: db.New(sqlDB),
	}
}

func (r *PostgresRepository) Create(user *domain.User) error {
	params := db.CreateUserParams{
		Email:    user.Email,
		Password: user.Password,
		Name:     user.Name,
	}

	result, err := r.queries.CreateUser(context.Background(), params)
	if err != nil {
		return err
	}

	user.ID = result.ID
	return nil
}

func (r *PostgresRepository) GetByID(id uuid.UUID) (*domain.User, error) {
	result, err := r.queries.GetUserByID(context.Background(), id)
	if err != nil {
		return nil, err
	}

	return &domain.User{
		ID:       result.ID,			
		Email:    result.Email,
		Password: result.Password,
		Name:     result.Name,
	}, nil
}

func (r *PostgresRepository) GetByEmail(email string) (*domain.User, error) {
	result, err := r.queries.GetUserByEmail(context.Background(), email)
	if err != nil {
		return nil, err
	}

	return &domain.User{
		ID:       result.ID,
		Email:    result.Email,
		Password: result.Password,
		Name:     result.Name,
	}, nil
}

func (r *PostgresRepository) Update(user *domain.User) error {
	params := db.UpdateUserParams{
		ID:       user.ID,
		Email:    user.Email,
		Password: user.Password,
		Name:     user.Name,
	}

	_, err := r.queries.UpdateUser(context.Background(), params)
	return err
}

func (r *PostgresRepository) Delete(id uuid.UUID) error {
	return r.queries.DeleteUser(context.Background(), id)
}

func (r *PostgresRepository) GetAll() ([]*domain.User, error) {
	results, err := r.queries.ListUsers(context.Background())
	if err != nil {
		return nil, err
	}

	users := make([]*domain.User, len(results))
	for i, result := range results {
		users[i] = &domain.User{
			ID:       result.ID,
			Email:    result.Email,
			Password: result.Password,
			Name:     result.Name,
		}
	}

	return users, nil
} 