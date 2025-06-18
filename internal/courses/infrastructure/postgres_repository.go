package infrastructure

import (
	"context"

	"github.com/google/uuid"
	"github.com/leninner/hear-backend/internal/courses/domain"
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

func (r *PostgresRepository) Create(course *domain.Course) error {
	ctx := context.Background()

	params := db.CreateCourseParams{
		FacultyID: course.FacultyID,
		Name:      course.Name,
		Code:      course.Description,
	}

	createdCourse, err := r.queries.CreateCourse(ctx, params)
	if err != nil {
		return err
	}

	course.ID = createdCourse.ID
	if createdCourse.CreatedAt.Valid {
		course.CreatedAt = createdCourse.CreatedAt.Time
	}
	if createdCourse.UpdatedAt.Valid {
		course.UpdatedAt = createdCourse.UpdatedAt.Time
	}

	return nil
}

func (r *PostgresRepository) GetByID(id uuid.UUID) (*domain.Course, error) {
	ctx := context.Background()

	course, err := r.queries.GetCourseByID(ctx, id)
	if err != nil {
		return nil, domain.ErrCourseNotFound
	}

	domainCourse := &domain.Course{
		ID:          course.ID,
		Name:        course.Name,
		Description: course.Code,
		FacultyID:   course.FacultyID,
	}
	if course.CreatedAt.Valid {
		domainCourse.CreatedAt = course.CreatedAt.Time
	}
	if course.UpdatedAt.Valid {
		domainCourse.UpdatedAt = course.UpdatedAt.Time
	}

	return domainCourse, nil
}

func (r *PostgresRepository) GetAll() ([]*domain.Course, error) {
	ctx := context.Background()

	courses, err := r.queries.GetCoursesByFacultyID(ctx, uuid.Nil)
	if err != nil {
		return nil, err
	}

	var domainCourses []*domain.Course
	for _, course := range courses {
		domainCourse := &domain.Course{
			ID:          course.ID,
			Name:        course.Name,
			Description: course.Code,
			FacultyID:   course.FacultyID,
		}
		if course.CreatedAt.Valid {
			domainCourse.CreatedAt = course.CreatedAt.Time
		}
		if course.UpdatedAt.Valid {
			domainCourse.UpdatedAt = course.UpdatedAt.Time
		}
		domainCourses = append(domainCourses, domainCourse)
	}

	return domainCourses, nil
}

func (r *PostgresRepository) GetByTeacherID(teacherID uuid.UUID) ([]*domain.Course, error) {
	ctx := context.Background()

	courses, err := r.queries.GetCoursesByFacultyID(ctx, teacherID)
	if err != nil {
		return nil, err
	}

	var domainCourses []*domain.Course
	for _, course := range courses {
		domainCourse := &domain.Course{
			ID:          course.ID,
			Name:        course.Name,
			Description: course.Code,
			FacultyID:   course.FacultyID,
		}
		if course.CreatedAt.Valid {
			domainCourse.CreatedAt = course.CreatedAt.Time
		}
		if course.UpdatedAt.Valid {
			domainCourse.UpdatedAt = course.UpdatedAt.Time
		}
		domainCourses = append(domainCourses, domainCourse)
	}

	return domainCourses, nil
}

func (r *PostgresRepository) Update(course *domain.Course) error {
	ctx := context.Background()

	params := db.UpdateCourseParams{
		ID:   course.ID,
		Name: course.Name,
		Code: course.Description,
	}

	updatedCourse, err := r.queries.UpdateCourse(ctx, params)
	if err != nil {
		return err
	}

	if updatedCourse.UpdatedAt.Valid {
		course.UpdatedAt = updatedCourse.UpdatedAt.Time
	}
	return nil
}

func (r *PostgresRepository) Delete(id uuid.UUID) error {
	ctx := context.Background()

	err := r.queries.DeleteCourse(ctx, id)
	if err != nil {
		return domain.ErrCourseNotFound
	}

	return nil
} 