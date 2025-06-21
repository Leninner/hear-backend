package infrastructure

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/leninner/hear-backend/internal/courses/domain"
	"github.com/leninner/hear-backend/internal/infrastructure/db"
)

// TODO: After running the migration (000006_fix_courses_schema.up.sql),
// regenerate SQLC code with: sqlc generate
// Then update this repository to use the proper field names:
// - Replace Code with Semester in CreateCourseParams and UpdateCourseParams
// - Use GetCoursesBySemester query instead of filtering in Go

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
		Semester:  sql.NullString{String: course.Semester, Valid: true},
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
		ID:        course.ID,
		Name:      course.Name,
		FacultyID: course.FacultyID,
		Semester:  course.Semester.String,
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
			ID:        course.ID,
			Name:      course.Name,
			FacultyID: course.FacultyID,
			Semester:  course.Semester.String,
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

func (r *PostgresRepository) GetByFacultyID(facultyID uuid.UUID) ([]*domain.Course, error) {
	ctx := context.Background()

	courses, err := r.queries.GetCoursesByFacultyID(ctx, facultyID)
	if err != nil {
		return nil, err
	}

	var domainCourses []*domain.Course
	for _, course := range courses {
		domainCourse := &domain.Course{
			ID:        course.ID,
			Name:      course.Name,
			FacultyID: course.FacultyID,
			Semester:  course.Semester.String,
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

func (r *PostgresRepository) GetBySemester(semester string) ([]*domain.Course, error) {
	ctx := context.Background()

	params := sql.NullString{String: semester, Valid: true}
	courses, err := r.queries.GetCoursesBySemester(ctx, params)
	if err != nil {
		return nil, err
	}

	var domainCourses []*domain.Course
	for _, course := range courses {
		domainCourse := &domain.Course{
			ID:        course.ID,
			Name:      course.Name,
			FacultyID: course.FacultyID,
			Semester:  course.Semester.String,
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
		ID:       course.ID,
		Name:     course.Name,
		Semester: sql.NullString{String: course.Semester, Valid: true},
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