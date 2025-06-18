package infrastructure

import (
	"context"

	"github.com/google/uuid"
	"github.com/leninner/hear-backend/internal/infrastructure/db"
	"github.com/leninner/hear-backend/internal/universities/domain"
)

type PostgresRepository struct {
	queries *db.Queries
}

func NewPostgresRepository(queries *db.Queries) *PostgresRepository {
	return &PostgresRepository{
		queries: queries,
	}
}

func (r *PostgresRepository) Create(university *domain.University) error {
	ctx := context.Background()

	result, err := r.queries.CreateUniversity(ctx, university.Name)
	if err != nil {
		return domain.NewInternalError("failed to create university in database", err)
	}

	university.ID = result.ID
	if result.CreatedAt.Valid {
		university.CreatedAt = result.CreatedAt.Time
	}
	if result.UpdatedAt.Valid {
		university.UpdatedAt = result.UpdatedAt.Time
	}
	return nil
}

func (r *PostgresRepository) GetByID(id uuid.UUID) (*domain.University, error) {
	ctx := context.Background()

	result, err := r.queries.GetUniversityByID(ctx, id)
	if err != nil {
		return nil, domain.ErrUniversityNotFound
	}

	university := &domain.University{
		ID:   result.ID,
		Name: result.Name,
	}
	if result.CreatedAt.Valid {
		university.CreatedAt = result.CreatedAt.Time
	}
	if result.UpdatedAt.Valid {
		university.UpdatedAt = result.UpdatedAt.Time
	}

	return university, nil
}

func (r *PostgresRepository) GetByName(name string) (*domain.University, error) {
	ctx := context.Background()

	result, err := r.queries.GetUniversityByName(ctx, name)
	if err != nil {
		return nil, domain.ErrUniversityNotFound
	}

	university := &domain.University{
		ID:   result.ID,
		Name: result.Name,
	}
	if result.CreatedAt.Valid {
		university.CreatedAt = result.CreatedAt.Time
	}
	if result.UpdatedAt.Valid {
		university.UpdatedAt = result.UpdatedAt.Time
	}

	return university, nil
}

func (r *PostgresRepository) GetAll() ([]*domain.University, error) {
	ctx := context.Background()

	results, err := r.queries.ListUniversities(ctx)
	if err != nil {
		return nil, domain.NewInternalError("failed to retrieve universities from database", err)
	}

	universities := make([]*domain.University, 0, len(results))
	for _, result := range results {
		university := &domain.University{
			ID:   result.ID,
			Name: result.Name,
		}
		if result.CreatedAt.Valid {
			university.CreatedAt = result.CreatedAt.Time
		}
		if result.UpdatedAt.Valid {
			university.UpdatedAt = result.UpdatedAt.Time
		}
		universities = append(universities, university)
	}
	return universities, nil
}

func (r *PostgresRepository) Update(university *domain.University) error {
	ctx := context.Background()

	params := db.UpdateUniversityParams{
		ID:   university.ID,
		Name: university.Name,
	}

	result, err := r.queries.UpdateUniversity(ctx, params)
	if err != nil {
		return domain.NewInternalError("failed to update university in database", err)
	}

	if result.UpdatedAt.Valid {
		university.UpdatedAt = result.UpdatedAt.Time
	}
	return nil
}

func (r *PostgresRepository) Delete(id uuid.UUID) error {
	ctx := context.Background()

	err := r.queries.DeleteUniversity(ctx, id)
	if err != nil {
		return domain.ErrUniversityNotFound
	}

	return nil
} 