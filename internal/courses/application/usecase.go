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

// Section methods
func (uc *UseCase) CreateSection(dto *domain.CreateCourseSectionDTO) (*domain.CourseSection, error) {
	if err := dto.Validate(); err != nil {
		return nil, err
	}

	section := domain.NewCourseSection(dto.CourseID, dto.TeacherID, dto.Name, dto.MaxStudents)
	if err := uc.repository.CreateSection(section); err != nil {
		return nil, domain.NewInternalError("failed to create section in database", err)
	}

	return section, nil
}

func (uc *UseCase) GetSection(id uuid.UUID) (*domain.CourseSection, error) {
	section, err := uc.repository.GetSectionByID(id)
	if err != nil {
		return nil, domain.NewNotFoundError("section not found")
	}
	return section, nil
}

func (uc *UseCase) GetSectionsByCourse(courseID uuid.UUID) ([]*domain.CourseSection, error) {
	sections, err := uc.repository.GetSectionsByCourseID(courseID)
	if err != nil {
		return nil, domain.NewInternalError("failed to retrieve course sections from database", err)
	}
	return sections, nil
}

func (uc *UseCase) GetSectionsByTeacher(teacherID uuid.UUID) ([]*domain.CourseSection, error) {
	sections, err := uc.repository.GetSectionsByTeacherID(teacherID)
	if err != nil {
		return nil, domain.NewInternalError("failed to retrieve teacher sections from database", err)
	}
	return sections, nil
}

func (uc *UseCase) UpdateSection(id uuid.UUID, dto *domain.UpdateCourseSectionDTO) error {
	if err := dto.Validate(); err != nil {
		return err
	}

	section, err := uc.repository.GetSectionByID(id)
	if err != nil {
		return domain.NewNotFoundError("section not found")
	}

	if dto.Name != "" {
		section.Name = dto.Name
	}
	if dto.MaxStudents != nil {
		section.MaxStudents = *dto.MaxStudents
	}

	if err := uc.repository.UpdateSection(section); err != nil {
		return domain.NewInternalError("failed to update section in database", err)
	}

	return nil
}

func (uc *UseCase) DeleteSection(id uuid.UUID) error {
	if _, err := uc.repository.GetSectionByID(id); err != nil {
		return domain.NewNotFoundError("section not found")
	}

	if err := uc.repository.DeleteSection(id); err != nil {
		return domain.NewInternalError("failed to delete section from database", err)
	}

	return nil
}

func (uc *UseCase) EnrollInSection(sectionID uuid.UUID, dto *domain.EnrollInSectionDTO) (*domain.CourseSection, error) {
	if err := dto.Validate(); err != nil {
		return nil, err
	}

	// Get the section to check max students
	section, err := uc.repository.GetSectionByID(sectionID)
	if err != nil {
		return nil, domain.NewNotFoundError("section not found")
	}

	// Check if student is already enrolled
	alreadyEnrolled, err := uc.repository.IsStudentEnrolled(sectionID, dto.StudentID)
	if err != nil {
		return nil, domain.NewInternalError("failed to check enrollment status", err)
	}
	if alreadyEnrolled {
		return nil, domain.NewConflictError(domain.ErrAlreadyEnrolled)
	}

	// Check if section is full
	currentEnrollmentCount, err := uc.repository.GetEnrollmentCount(sectionID)
	if err != nil {
		return nil, domain.NewInternalError("failed to get enrollment count", err)
	}
	if currentEnrollmentCount >= section.MaxStudents {
		return nil, domain.NewConflictError(domain.ErrSectionFull)
	}

	// Enroll the student
	if err := uc.repository.EnrollStudent(sectionID, dto.StudentID); err != nil {
		return nil, domain.NewInternalError("failed to enroll student", err)
	}

	return section, nil
}