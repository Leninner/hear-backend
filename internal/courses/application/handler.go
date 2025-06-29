package application

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/leninner/hear-backend/internal/courses/domain"
	"github.com/leninner/hear-backend/internal/shared/response"
)

type Handler struct {
	useCase *UseCase
}

func NewHandler(useCase *UseCase) *Handler {
	return &Handler{
		useCase: useCase,
	}
}

func (h *Handler) CreateCourse(c *fiber.Ctx) error {
	course := new(domain.CreateCourseDTO)
	if err := c.BodyParser(course); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid request body format")
	}

	createdCourse, err := h.useCase.CreateCourse(course)
	if err != nil {
		switch e := err.(type) {
		case *domain.ValidationErrors:
			return response.Error(c, fiber.StatusBadRequest, e.GetErrors())
		case *domain.DomainError:
			switch e.Type {
			case domain.ErrorTypeConflict:
				return response.Error(c, fiber.StatusConflict, e.Message)
			case domain.ErrorTypeNotFound:
				return response.Error(c, fiber.StatusNotFound, e.Message)
			default:
				fmt.Println(e)
				return response.Error(c, fiber.StatusInternalServerError, "An unexpected error occurred")
			}
		default:
			return response.Error(c, fiber.StatusInternalServerError, "Failed to create course")
		}
	}

	return response.Success(c, "Course created successfully", createdCourse)
}

func (h *Handler) GetCourse(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, domain.ErrInvalidID.Error())
	}

	course, err := h.useCase.GetCourse(id)
	if err != nil {
		switch e := err.(type) {
		case *domain.DomainError:
			switch e.Type {
			case domain.ErrorTypeNotFound:
				return response.Error(c, fiber.StatusNotFound, e.Message)
			default:
				return response.Error(c, fiber.StatusInternalServerError, "An unexpected error occurred")
			}
		default:
			return response.Error(c, fiber.StatusInternalServerError, "Failed to retrieve course")
		}
	}

	return response.Success(c, "Course retrieved successfully", course)
}

func (h *Handler) GetAllCourses(c *fiber.Ctx) error {
	courses, err := h.useCase.GetAllCourses()
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to retrieve courses")
	}

	return response.Success(c, "Courses retrieved successfully", courses)
}

func (h *Handler) UpdateCourse(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, domain.ErrInvalidID.Error())
	}

	course := new(domain.UpdateCourseDTO)
	if err := c.BodyParser(course); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid request body format")
	}

	if err := h.useCase.UpdateCourse(id, course); err != nil {
		switch e := err.(type) {
		case *domain.ValidationErrors:
			return response.Error(c, fiber.StatusBadRequest, e.GetErrors())
		case *domain.DomainError:
			switch e.Type {
			case domain.ErrorTypeNotFound:
				return response.Error(c, fiber.StatusNotFound, e.Message)
			default:
				return response.Error(c, fiber.StatusInternalServerError, "An unexpected error occurred")
			}
		default:
			return response.Error(c, fiber.StatusInternalServerError, "Failed to update course")
		}
	}

	return response.Success(c, "Course updated successfully", nil)
}

func (h *Handler) DeleteCourse(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, domain.ErrInvalidID.Error())
	}

	if err := h.useCase.DeleteCourse(id); err != nil {
		switch e := err.(type) {
		case *domain.DomainError:
			switch e.Type {
			case domain.ErrorTypeNotFound:
				return response.Error(c, fiber.StatusNotFound, e.Message)
			default:
				return response.Error(c, fiber.StatusInternalServerError, "An unexpected error occurred")
			}
		default:
			return response.Error(c, fiber.StatusInternalServerError, "Failed to delete course")
		}
	}

	return response.Success(c, "Course deleted successfully", nil)
}

// Section handlers
func (h *Handler) CreateSection(c *fiber.Ctx) error {
	section := new(domain.CreateCourseSectionDTO)
	if err := c.BodyParser(section); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid request body format")
	}

	createdSection, err := h.useCase.CreateSection(section)
	if err != nil {
		switch e := err.(type) {
		case *domain.ValidationErrors:
			return response.Error(c, fiber.StatusBadRequest, e.GetErrors())
		case *domain.DomainError:
			switch e.Type {
			case domain.ErrorTypeConflict:
				return response.Error(c, fiber.StatusConflict, e.Message)
			case domain.ErrorTypeNotFound:
				return response.Error(c, fiber.StatusNotFound, e.Message)
			default:
				return response.Error(c, fiber.StatusInternalServerError, "An unexpected error occurred")
			}
		default:
			return response.Error(c, fiber.StatusInternalServerError, "Failed to create section")
		}
	}

	return response.Success(c, "Section created successfully", createdSection)
}

func (h *Handler) GetSection(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid section ID format")
	}

	section, err := h.useCase.GetSection(id)
	if err != nil {
		switch e := err.(type) {
		case *domain.DomainError:
			switch e.Type {
			case domain.ErrorTypeNotFound:
				return response.Error(c, fiber.StatusNotFound, e.Message)
			default:
				return response.Error(c, fiber.StatusInternalServerError, "An unexpected error occurred")
			}
		default:
			return response.Error(c, fiber.StatusInternalServerError, "Failed to retrieve section")
		}
	}

	return response.Success(c, "Section retrieved successfully", section)
}

func (h *Handler) GetSectionsByCourse(c *fiber.Ctx) error {
	courseID, err := uuid.Parse(c.Params("courseId"))
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid course ID format")
	}

	sections, err := h.useCase.GetSectionsByCourse(courseID)
	if err != nil {
		switch e := err.(type) {
		case *domain.DomainError:
			switch e.Type {
			case domain.ErrorTypeNotFound:
				return response.Error(c, fiber.StatusNotFound, e.Message)
			default:
				return response.Error(c, fiber.StatusInternalServerError, "An unexpected error occurred")
			}
		default:
			return response.Error(c, fiber.StatusInternalServerError, "Failed to retrieve course sections")
		}
	}

	return response.Success(c, "Course sections retrieved successfully", sections)
}

func (h *Handler) GetSectionsByTeacher(c *fiber.Ctx) error {
	teacherID, err := uuid.Parse(c.Params("teacherId"))
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid teacher ID format")
	}

	sections, err := h.useCase.GetSectionsByTeacher(teacherID)
	if err != nil {
		switch e := err.(type) {
		case *domain.DomainError:
			switch e.Type {
			case domain.ErrorTypeNotFound:
				return response.Error(c, fiber.StatusNotFound, e.Message)
			default:
				return response.Error(c, fiber.StatusInternalServerError, "An unexpected error occurred")
			}
		default:
			return response.Error(c, fiber.StatusInternalServerError, "Failed to retrieve teacher sections")
		}
	}

	return response.Success(c, "Teacher sections retrieved successfully", sections)
}

func (h *Handler) GetSectionsByStudent(c *fiber.Ctx) error {
	studentID, err := uuid.Parse(c.Params("studentId"))
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid student ID format")
	}

	sections, err := h.useCase.GetSectionsByStudent(studentID)
	if err != nil {
		switch e := err.(type) {
		case *domain.DomainError:
			switch e.Type {
			case domain.ErrorTypeNotFound:
				return response.Error(c, fiber.StatusNotFound, e.Message)
			default:
				return response.Error(c, fiber.StatusInternalServerError, "An unexpected error occurred")
			}
		default:
			return response.Error(c, fiber.StatusInternalServerError, "Failed to retrieve student sections")
		}
	}

	return response.Success(c, "Student sections retrieved successfully", sections)
}

func (h *Handler) GetSectionsWithSchedulesByStudent(c *fiber.Ctx) error {
	studentID, err := uuid.Parse(c.Params("studentId"))
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid student ID format")
	}

	sections, err := h.useCase.GetSectionsWithSchedulesByStudent(studentID)
	if err != nil {
		switch e := err.(type) {
		case *domain.DomainError:
			switch e.Type {
			case domain.ErrorTypeNotFound:
				return response.Error(c, fiber.StatusNotFound, e.Message)
			default:
				return response.Error(c, fiber.StatusInternalServerError, "An unexpected error occurred")
			}
		default:
			return response.Error(c, fiber.StatusInternalServerError, "Failed to retrieve student sections with schedules")
		}
	}

	return response.Success(c, "Student sections with schedules retrieved successfully", sections)
}

func (h *Handler) UpdateSection(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid section ID format")
	}

	section := new(domain.UpdateCourseSectionDTO)
	if err := c.BodyParser(section); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid request body format")
	}

	if err := h.useCase.UpdateSection(id, section); err != nil {
		switch e := err.(type) {
		case *domain.ValidationErrors:
			return response.Error(c, fiber.StatusBadRequest, e.GetErrors())
		case *domain.DomainError:
			switch e.Type {
			case domain.ErrorTypeNotFound:
				return response.Error(c, fiber.StatusNotFound, e.Message)
			default:
				return response.Error(c, fiber.StatusInternalServerError, "An unexpected error occurred")
			}
		default:
			return response.Error(c, fiber.StatusInternalServerError, "Failed to update section")
		}
	}

	return response.Success(c, "Section updated successfully", nil)
}

func (h *Handler) DeleteSection(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid section ID format")
	}

	if err := h.useCase.DeleteSection(id); err != nil {
		switch e := err.(type) {
		case *domain.DomainError:
			switch e.Type {
			case domain.ErrorTypeNotFound:
				return response.Error(c, fiber.StatusNotFound, e.Message)
			default:
				return response.Error(c, fiber.StatusInternalServerError, "An unexpected error occurred")
			}
		default:
			return response.Error(c, fiber.StatusInternalServerError, "Failed to delete section")
		}
	}

	return response.Success(c, "Section deleted successfully", nil)
} 

func (h *Handler) EnrollInSection(c *fiber.Ctx) error {
	sectionID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid section ID format")
	}

	enrollment := new(domain.EnrollInSectionDTO)
	if err := c.BodyParser(enrollment); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid request body format")
	}
	
	section, err := h.useCase.EnrollInSection(sectionID, enrollment)
	if err != nil {
		switch e := err.(type) {
		case *domain.ValidationErrors:
			return response.Error(c, fiber.StatusBadRequest, e.GetErrors())
		case *domain.DomainError:
			switch e.Type {
			case domain.ErrorTypeNotFound:
				return response.Error(c, fiber.StatusNotFound, e.Message)
			case domain.ErrorTypeConflict:
				return response.Error(c, fiber.StatusConflict, e.Message)
			default:
				return response.Error(c, fiber.StatusInternalServerError, "An unexpected error occurred")
			}
		default:
			return response.Error(c, fiber.StatusInternalServerError, "Failed to enroll in section")
		}
	}

	return response.Success(c, "Enrollment successful", section)
}

func (h *Handler) GetEnrollmentsWithDetailsBySection(c *fiber.Ctx) error {
	sectionID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid section ID format")
	}

	sectionEnrollments, err := h.useCase.GetEnrollmentsWithDetailsBySection(sectionID)
	if err != nil {
		switch e := err.(type) {
		case *domain.DomainError:
			switch e.Type {
			case domain.ErrorTypeNotFound:
				return response.Error(c, fiber.StatusNotFound, e.Message)
			default:
				return response.Error(c, fiber.StatusInternalServerError, "An unexpected error occurred")
			}
		default:
			return response.Error(c, fiber.StatusInternalServerError, "Failed to retrieve enrollments with details")
		}
	}

	return response.Success(c, "Section enrollments with details retrieved successfully", sectionEnrollments)
}