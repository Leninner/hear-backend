package infrastructure

import (
	"context"

	"github.com/google/uuid"
	"github.com/leninner/hear-backend/internal/infrastructure/db"
	"github.com/leninner/hear-backend/internal/users/domain"
)

type PostgresRepository struct {
	queries *db.Queries
}

func NewPostgresRepository(queries *db.Queries) *PostgresRepository {
	return &PostgresRepository{
		queries: queries,
	}
}

func (r *PostgresRepository) Create(user *domain.User) error {
	ctx := context.Background()

	params := db.CreateUserParams{
		Email:        user.Email,
		PasswordHash: user.PasswordHash,
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		Role:         db.UserRole(user.Role),
	}

	result, err := r.queries.CreateUser(ctx, params)
	if err != nil {
		return domain.NewInternalError("failed to create user in database", err)
	}

	user.ID = result.ID
	if result.CreatedAt.Valid {
		user.CreatedAt = result.CreatedAt.Time
	}
	if result.UpdatedAt.Valid {
		user.UpdatedAt = result.UpdatedAt.Time
	}
	return nil
}

func (r *PostgresRepository) GetByID(id uuid.UUID) (*domain.User, error) {
	ctx := context.Background()

	result, err := r.queries.GetUserByID(ctx, id)
	if err != nil {
		return nil, domain.ErrUserNotFound
	}

	user := &domain.User{
		ID:           result.ID,
		Email:        result.Email,
		PasswordHash: result.PasswordHash,
		FirstName:    result.FirstName,
		LastName:     result.LastName,
		Role:         domain.UserRole(result.Role),
	}
	if result.CreatedAt.Valid {
		user.CreatedAt = result.CreatedAt.Time
	}
	if result.UpdatedAt.Valid {
		user.UpdatedAt = result.UpdatedAt.Time
	}

	return user, nil
}

func (r *PostgresRepository) GetByEmail(email string) (*domain.User, error) {
	ctx := context.Background()

	result, err := r.queries.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, domain.ErrUserNotFound
	}

	user := &domain.User{
		ID:           result.ID,
		Email:        result.Email,
		PasswordHash: result.PasswordHash,
		FirstName:    result.FirstName,
		LastName:     result.LastName,
		Role:         domain.UserRole(result.Role),
	}
	if result.CreatedAt.Valid {
		user.CreatedAt = result.CreatedAt.Time
	}
	if result.UpdatedAt.Valid {
		user.UpdatedAt = result.UpdatedAt.Time
	}

	return user, nil
}

func (r *PostgresRepository) GetAll() ([]*domain.User, error) {
	ctx := context.Background()

	results, err := r.queries.ListUsers(ctx)
	if err != nil {
		return nil, domain.NewInternalError("failed to retrieve users from database", err)
	}

	users := make([]*domain.User, 0, len(results))
	for _, result := range results {
		user := &domain.User{
			ID:           result.ID,
			Email:        result.Email,
			PasswordHash: result.PasswordHash,
			FirstName:    result.FirstName,
			LastName:     result.LastName,
			Role:         domain.UserRole(result.Role),
		}
		if result.CreatedAt.Valid {
			user.CreatedAt = result.CreatedAt.Time
		}
		if result.UpdatedAt.Valid {
			user.UpdatedAt = result.UpdatedAt.Time
		}
		users = append(users, user)
	}
	return users, nil
}

func (r *PostgresRepository) Update(user *domain.User) error {
	ctx := context.Background()

	params := db.UpdateUserParams{
		ID:           user.ID,
		Email:        user.Email,
		PasswordHash: user.PasswordHash,
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		Role:         db.UserRole(user.Role),
	}

	result, err := r.queries.UpdateUser(ctx, params)
	if err != nil {
		return domain.NewInternalError("failed to update user in database", err)
	}

	if result.UpdatedAt.Valid {
		user.UpdatedAt = result.UpdatedAt.Time
	}
	return nil
}

func (r *PostgresRepository) Delete(id uuid.UUID) error {
	ctx := context.Background()

	err := r.queries.DeleteUser(ctx, id)
	if err != nil {
		return domain.ErrUserNotFound
	}

	return nil
} 