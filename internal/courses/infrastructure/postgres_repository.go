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

	courses, err := r.queries.GetAllCourses(ctx)
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

// Section methods
func (r *PostgresRepository) CreateSection(section *domain.CourseSection) error {
	ctx := context.Background()

	params := db.CreateCourseSectionParams{
		CourseID:    section.CourseID,
		Name:        section.Name,
		TeacherID:   section.TeacherID,
		MaxStudents: int32(section.MaxStudents),
	}

	createdSection, err := r.queries.CreateCourseSection(ctx, params)
	if err != nil {
		return err
	}

	section.ID = createdSection.ID
	if createdSection.CreatedAt.Valid {
		section.CreatedAt = createdSection.CreatedAt.Time
	}
	if createdSection.UpdatedAt.Valid {
		section.UpdatedAt = createdSection.UpdatedAt.Time
	}

	return nil
}

func (r *PostgresRepository) GetSectionByID(id uuid.UUID) (*domain.CourseSection, error) {
	ctx := context.Background()

	section, err := r.queries.GetCourseSectionByID(ctx, id)
	if err != nil {
		return nil, err
	}

	domainSection := &domain.CourseSection{
		ID:          section.ID,
		CourseID:    section.CourseID,
		Name:        section.Name,
		TeacherID:   section.TeacherID,
		MaxStudents: int(section.MaxStudents),
	}
	if section.CreatedAt.Valid {
		domainSection.CreatedAt = section.CreatedAt.Time
	}
	if section.UpdatedAt.Valid {
		domainSection.UpdatedAt = section.UpdatedAt.Time
	}

	return domainSection, nil
}

func (r *PostgresRepository) GetSectionsByCourseID(courseID uuid.UUID) ([]*domain.CourseSection, error) {
	ctx := context.Background()

	sections, err := r.queries.GetCourseSectionsByCourseID(ctx, courseID)
	if err != nil {
		return nil, err
	}

	var domainSections []*domain.CourseSection
	for _, section := range sections {
		domainSection := &domain.CourseSection{
			ID:          section.ID,
			CourseID:    section.CourseID,
			Name:        section.Name,
			TeacherID:   section.TeacherID,
			MaxStudents: int(section.MaxStudents),
		}
		if section.CreatedAt.Valid {
			domainSection.CreatedAt = section.CreatedAt.Time
		}
		if section.UpdatedAt.Valid {
			domainSection.UpdatedAt = section.UpdatedAt.Time
		}
		domainSections = append(domainSections, domainSection)
	}

	return domainSections, nil
}

func (r *PostgresRepository) GetSectionsByTeacherID(teacherID uuid.UUID) ([]*domain.CourseSection, error) {
	ctx := context.Background()

	sections, err := r.queries.GetCourseSectionsByTeacherID(ctx, teacherID)
	if err != nil {
		return nil, err
	}

	var domainSections []*domain.CourseSection
	for _, section := range sections {
		domainSection := &domain.CourseSection{
			ID:          section.ID,
			CourseID:    section.CourseID,
			Name:        section.Name,
			TeacherID:   section.TeacherID,
			MaxStudents: int(section.MaxStudents),
		}
		if section.CreatedAt.Valid {
			domainSection.CreatedAt = section.CreatedAt.Time
		}
		if section.UpdatedAt.Valid {
			domainSection.UpdatedAt = section.UpdatedAt.Time
		}
		domainSections = append(domainSections, domainSection)
	}

	return domainSections, nil
}

func (r *PostgresRepository) GetSectionsByStudentID(studentID uuid.UUID) ([]*domain.CourseSection, error) {
	ctx := context.Background()

	sections, err := r.queries.GetCourseSectionsByStudentID(ctx, studentID)
	if err != nil {
		return nil, err
	}

	var domainSections []*domain.CourseSection
	for _, section := range sections {
		domainSection := &domain.CourseSection{
			ID:          section.ID,
			CourseID:    section.CourseID,
			Name:        section.Name,
			TeacherID:   section.TeacherID,
			MaxStudents: int(section.MaxStudents),
		}
		if section.CreatedAt.Valid {
			domainSection.CreatedAt = section.CreatedAt.Time
		}
		if section.UpdatedAt.Valid {
			domainSection.UpdatedAt = section.UpdatedAt.Time
		}
		domainSections = append(domainSections, domainSection)
	}

	return domainSections, nil
}

func (r *PostgresRepository) GetSectionsWithSchedulesByStudentID(studentID uuid.UUID) ([]*domain.CourseSectionWithSchedules, error) {
	ctx := context.Background()

	// Get sections with schedules by student ID
	rows, err := r.queries.GetCourseSectionsWithSchedulesByStudentID(ctx, studentID)
	if err != nil {
		return nil, err
	}

	// Group sections and their schedules
	sectionsMap := make(map[uuid.UUID]*domain.CourseSectionWithSchedules)
	
	for _, row := range rows {
		sectionID := row.ID
		
		// Create or get existing section
		sectionWithSchedules, exists := sectionsMap[sectionID]
		if !exists {
			domainSection := &domain.CourseSection{
				ID:          row.ID,
				CourseID:    row.CourseID,
				Name:        row.Name,
				TeacherID:   row.TeacherID,
				MaxStudents: int(row.MaxStudents),
			}
			if row.CreatedAt.Valid {
				domainSection.CreatedAt = row.CreatedAt.Time
			}
			if row.UpdatedAt.Valid {
				domainSection.UpdatedAt = row.UpdatedAt.Time
			}
			
			sectionWithSchedules = &domain.CourseSectionWithSchedules{
				CourseSection: domainSection,
				Schedules:     []*domain.Schedule{},
			}
			sectionsMap[sectionID] = sectionWithSchedules
		}
		
		// Add schedule if it exists (schedule_id is not null)
		if row.ScheduleID.Valid {
			domainSchedule := &domain.Schedule{
				ID:          row.ScheduleID.UUID.String(),
				CourseID:    row.ScheduleCourseID.UUID.String(),
				SectionID:   row.ScheduleSectionID.UUID.String(),
				ClassroomID: row.ClassroomID.UUID.String(),
				DayOfWeek:   int(row.DayOfWeek.Int32),
				StartTime:   row.StartTime.Time.Format("15:04"),
				EndTime:     row.EndTime.Time.Format("15:04"),
				CreatedAt:   row.ScheduleCreatedAt.Time.Format("2006-01-02T15:04:05Z"),
				UpdatedAt:   row.ScheduleUpdatedAt.Time.Format("2006-01-02T15:04:05Z"),
			}
			sectionWithSchedules.Schedules = append(sectionWithSchedules.Schedules, domainSchedule)
		}
	}
	
	// Convert map to slice
	var result []*domain.CourseSectionWithSchedules
	for _, section := range sectionsMap {
		result = append(result, section)
	}
	
	return result, nil
}

func (r *PostgresRepository) UpdateSection(section *domain.CourseSection) error {
	ctx := context.Background()

	params := db.UpdateCourseSectionParams{
		ID:          section.ID,
		Name:        section.Name,
		TeacherID:   section.TeacherID,
		MaxStudents: int32(section.MaxStudents),
	}

	updatedSection, err := r.queries.UpdateCourseSection(ctx, params)
	if err != nil {
		return err
	}

	if updatedSection.UpdatedAt.Valid {
		section.UpdatedAt = updatedSection.UpdatedAt.Time
	}
	return nil
}

func (r *PostgresRepository) DeleteSection(id uuid.UUID) error {
	ctx := context.Background()

	return r.queries.DeleteCourseSection(ctx, id)
}

// Enrollment methods
func (r *PostgresRepository) EnrollStudent(sectionID, studentID uuid.UUID) error {
	ctx := context.Background()

	params := db.EnrollStudentParams{
		SectionID: sectionID,
		StudentID: studentID,
	}

	_, err := r.queries.EnrollStudent(ctx, params)
	return err
}

func (r *PostgresRepository) GetEnrollmentCount(sectionID uuid.UUID) (int, error) {
	ctx := context.Background()

	count, err := r.queries.GetEnrollmentCount(ctx, sectionID)
	if err != nil {
		return 0, err
	}

	return int(count), nil
}

func (r *PostgresRepository) IsStudentEnrolled(sectionID, studentID uuid.UUID) (bool, error) {
	ctx := context.Background()

	params := db.IsStudentEnrolledParams{
		SectionID: sectionID,
		StudentID: studentID,
	}

	enrolled, err := r.queries.IsStudentEnrolled(ctx, params)
	if err != nil {
		return false, err
	}

	return enrolled, nil
}

func (r *PostgresRepository) UnenrollStudent(sectionID, studentID uuid.UUID) error {
	ctx := context.Background()

	params := db.UnenrollStudentParams{
		SectionID: sectionID,
		StudentID: studentID,
	}

	return r.queries.UnenrollStudent(ctx, params)
}

func (r *PostgresRepository) GetEnrollmentsWithDetailsBySection(sectionID uuid.UUID) ([]*domain.EnrollmentWithStudentDetails, error) {
	ctx := context.Background()

	enrollments, err := r.queries.GetEnrollmentsWithDetailsBySection(ctx, sectionID)
	if err != nil {
		return nil, err
	}

	var domainEnrollments []*domain.EnrollmentWithStudentDetails
	for _, enrollment := range enrollments {
		// Create student object
		student := &domain.Student{
			ID:        enrollment.StudentID_2,
			Email:     enrollment.Email,
			FirstName: enrollment.FirstName,
			LastName:  enrollment.LastName,
			Role:      string(enrollment.Role),
		}
		if enrollment.UserCreatedAt.Valid {
			student.CreatedAt = enrollment.UserCreatedAt.Time
		}
		if enrollment.UserUpdatedAt.Valid {
			student.UpdatedAt = enrollment.UserUpdatedAt.Time
		}

		// Create enrollment with student details
		domainEnrollment := &domain.EnrollmentWithStudentDetails{
			ID:        enrollment.ID,
			SectionID: enrollment.SectionID,
			StudentID: enrollment.StudentID,
			Student:   student,
		}
		if enrollment.CreatedAt.Valid {
			domainEnrollment.CreatedAt = enrollment.CreatedAt.Time
		}
		if enrollment.UpdatedAt.Valid {
			domainEnrollment.UpdatedAt = enrollment.UpdatedAt.Time
		}
		domainEnrollments = append(domainEnrollments, domainEnrollment)
	}

	return domainEnrollments, nil
} 