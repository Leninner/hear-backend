package application

import (
	"github.com/google/uuid"
	"github.com/leninner/hear-backend/internal/courses/domain"
)

type UseCase struct {
	repository domain.Repository
}

func NewUseCase(repository domain.Repository) *UseCase {
	return &UseCase{
		repository: repository,
	}
}

func (uc *UseCase) CreateCourse(dto *domain.CreateCourseDTO) (*domain.Course, error) {
	if err := dto.Validate(); err != nil {
		return nil, err
	}

	course := domain.NewCourse(dto.Name, dto.FacultyID, dto.Semester)
	if err := uc.repository.Create(course); err != nil {
		return nil, domain.NewInternalError("failed to create course in database", err)
	}

	return course, nil
}

func (uc *UseCase) GetCourse(id uuid.UUID) (*domain.Course, error) {
	course, err := uc.repository.GetByID(id)
	if err != nil {
		return nil, domain.ErrCourseNotFound
	}
	return course, nil
}

func (uc *UseCase) GetAllCourses() ([]*domain.Course, error) {
	courses, err := uc.repository.GetAll()
	if err != nil {
		return nil, domain.NewInternalError("failed to retrieve courses from database", err)
	}
	return courses, nil
}

func (uc *UseCase) GetCoursesByFaculty(facultyID uuid.UUID) ([]*domain.Course, error) {
	courses, err := uc.repository.GetByFacultyID(facultyID)
	if err != nil {
		return nil, domain.NewInternalError("failed to retrieve faculty courses from database", err)
	}
	return courses, nil
}

func (uc *UseCase) GetCoursesBySemester(semester string) ([]*domain.Course, error) {
	courses, err := uc.repository.GetBySemester(semester)
	if err != nil {
		return nil, domain.NewInternalError("failed to retrieve semester courses from database", err)
	}
	return courses, nil
}

func (uc *UseCase) UpdateCourse(id uuid.UUID, dto *domain.UpdateCourseDTO) error {
	if err := dto.Validate(); err != nil {
		return err
	}

	course, err := uc.repository.GetByID(id)
	if err != nil {
		return domain.ErrCourseNotFound
	}

	if dto.Name != "" {
		course.Name = dto.Name
	}
	if dto.Semester != "" {
		course.Semester = dto.Semester
	}

	if err := uc.repository.Update(course); err != nil {
		return domain.NewInternalError("failed to update course in database", err)
	}

	return nil
}

func (uc *UseCase) DeleteCourse(id uuid.UUID) error {
	if _, err := uc.repository.GetByID(id); err != nil {
		return domain.ErrCourseNotFound
	}

	if err := uc.repository.Delete(id); err != nil {
		return domain.NewInternalError("failed to delete course from database", err)
	}

	return nil
} 