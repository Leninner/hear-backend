package infrastructure

import (
	"context"

	"github.com/google/uuid"
	"github.com/leninner/hear-backend/internal/faculties/domain"
	"github.com/leninner/hear-backend/internal/infrastructure/db"
)

type PostgresRepository struct {
	queries *db.Queries
}

func NewPostgresRepository(queries *db.Queries) *PostgresRepository {
	return &PostgresRepository{
		queries: queries,
	}
}

func (r *PostgresRepository) Create(faculty *domain.Faculty) error {
	ctx := context.Background()

	params := db.CreateFacultyParams{
		UniversityID: faculty.UniversityID,
		Name:         faculty.Name,
	}

	result, err := r.queries.CreateFaculty(ctx, params)
	if err != nil {
		return domain.NewInternalError("failed to create faculty in database", err)
	}

	faculty.ID = result.ID
	if result.CreatedAt.Valid {
		faculty.CreatedAt = result.CreatedAt.Time
	}
	if result.UpdatedAt.Valid {
		faculty.UpdatedAt = result.UpdatedAt.Time
	}
	return nil
}

func (r *PostgresRepository) GetByID(id uuid.UUID) (*domain.Faculty, error) {
	ctx := context.Background()

	result, err := r.queries.GetFacultyByID(ctx, id)
	if err != nil {
		return nil, domain.ErrFacultyNotFound
	}

	faculty := &domain.Faculty{
		ID:           result.ID,
		UniversityID: result.UniversityID,
		Name:         result.Name,
	}
	if result.CreatedAt.Valid {
		faculty.CreatedAt = result.CreatedAt.Time
	}
	if result.UpdatedAt.Valid {
		faculty.UpdatedAt = result.UpdatedAt.Time
	}

	return faculty, nil
}

func (r *PostgresRepository) GetByName(name string) (*domain.Faculty, error) {
	ctx := context.Background()

	result, err := r.queries.GetFacultyByName(ctx, name)
	if err != nil {
		return nil, domain.ErrFacultyNotFound
	}

	faculty := &domain.Faculty{
		ID:           result.ID,
		UniversityID: result.UniversityID,
		Name:         result.Name,
	}
	if result.CreatedAt.Valid {
		faculty.CreatedAt = result.CreatedAt.Time
	}
	if result.UpdatedAt.Valid {
		faculty.UpdatedAt = result.UpdatedAt.Time
	}

	return faculty, nil
}

func (r *PostgresRepository) GetAll() ([]*domain.Faculty, error) {
	ctx := context.Background()

	results, err := r.queries.ListFaculties(ctx)
	if err != nil {
		return nil, domain.NewInternalError("failed to retrieve faculties from database", err)
	}

	faculties := make([]*domain.Faculty, 0, len(results))
	for _, result := range results {
		faculty := &domain.Faculty{
			ID:           result.ID,
			UniversityID: result.UniversityID,
			Name:         result.Name,
		}
		if result.CreatedAt.Valid {
			faculty.CreatedAt = result.CreatedAt.Time
		}
		if result.UpdatedAt.Valid {
			faculty.UpdatedAt = result.UpdatedAt.Time
		}
		faculties = append(faculties, faculty)
	}
	return faculties, nil
}

func (r *PostgresRepository) GetByUniversityID(universityID uuid.UUID) ([]*domain.Faculty, error) {
	ctx := context.Background()

	results, err := r.queries.GetFacultiesByUniversityID(ctx, universityID)
	if err != nil {
		return nil, domain.NewInternalError("failed to retrieve faculties from database", err)
	}

	faculties := make([]*domain.Faculty, 0, len(results))
	for _, result := range results {
		faculty := &domain.Faculty{
			ID:           result.ID,
			UniversityID: result.UniversityID,
			Name:         result.Name,
		}
		if result.CreatedAt.Valid {
			faculty.CreatedAt = result.CreatedAt.Time
		}
		if result.UpdatedAt.Valid {
			faculty.UpdatedAt = result.UpdatedAt.Time
		}
		faculties = append(faculties, faculty)
	}
	return faculties, nil
}

func (r *PostgresRepository) Update(faculty *domain.Faculty) error {
	ctx := context.Background()

	params := db.UpdateFacultyParams{
		ID:           faculty.ID,
		UniversityID: faculty.UniversityID,
		Name:         faculty.Name,
	}

	result, err := r.queries.UpdateFaculty(ctx, params)
	if err != nil {
		return domain.NewInternalError("failed to update faculty in database", err)
	}

	if result.UpdatedAt.Valid {
		faculty.UpdatedAt = result.UpdatedAt.Time
	}
	return nil
}

func (r *PostgresRepository) Delete(id uuid.UUID) error {
	ctx := context.Background()

	err := r.queries.DeleteFaculty(ctx, id)
	if err != nil {
		return domain.ErrFacultyNotFound
	}

	return nil
} 