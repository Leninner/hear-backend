package application

import (
	"github.com/google/uuid"
	classroomDomain "github.com/leninner/hear-backend/internal/classroom/domain"
	facultiesDomain "github.com/leninner/hear-backend/internal/faculties/domain"
)

type UseCase struct {
	repository        classroomDomain.Repository
	facultyRepository facultiesDomain.Repository
}

func NewUseCase(repository classroomDomain.Repository, facultyRepository facultiesDomain.Repository) *UseCase {
	return &UseCase{
		repository:        repository,
		facultyRepository: facultyRepository,
	}
}

func (uc *UseCase) CreateClassroom(dto *classroomDomain.CreateClassroomDTO) (*classroomDomain.Classroom, error) {
	if err := dto.Validate(); err != nil {
		return nil, err
	}

	// Validate that faculty exists
	_, err := uc.facultyRepository.GetByID(dto.FacultyID)
	if err != nil {
		return nil, classroomDomain.ErrFacultyNotFound
	}

	existingClassroom, err := uc.repository.GetByName(dto.Name)
	if err == nil && existingClassroom != nil {
		return nil, classroomDomain.ErrClassroomNameExists
	}

	classroom := classroomDomain.NewClassroom(
		dto.Name, 
		dto.Building, 
		*dto.Floor, 
		*dto.Capacity,
		dto.FacultyID,
		*dto.LocationLat, 
		*dto.LocationLng,
	)
	if err := uc.repository.Create(classroom); err != nil {
		return nil, classroomDomain.NewInternalError("failed to create classroom in database", err)
	}

	return classroom, nil
}

func (uc *UseCase) UpdateClassroom(id uuid.UUID, dto *classroomDomain.UpdateClassroomDTO) error {
	if err := dto.Validate(); err != nil {
		return err
	}

	classroom, err := uc.repository.GetByID(id)
	if err != nil {
		return classroomDomain.ErrClassroomNotFound
	}

	// Validate that faculty exists if it's being updated
	if dto.FacultyID != nil {
		_, err := uc.facultyRepository.GetByID(*dto.FacultyID)
		if err != nil {
			return classroomDomain.ErrFacultyNotFound
		}
	}

	if dto.Name != "" {
		// Check if the new name conflicts with another classroom
		if dto.Name != classroom.Name {
			existingClassroom, err := uc.repository.GetByName(dto.Name)
			if err == nil && existingClassroom != nil {
				return classroomDomain.ErrClassroomNameExists
			}
		}
		classroom.Name = dto.Name
	}

	if dto.Building != "" {
		classroom.Building = dto.Building
	}

	if dto.Floor != nil {
		classroom.Floor = *dto.Floor
	}

	if dto.Capacity != nil {
		classroom.Capacity = *dto.Capacity
	}

	if dto.FacultyID != nil {
		classroom.FacultyID = *dto.FacultyID
	}

	if dto.LocationLat != nil {
		classroom.LocationLat = *dto.LocationLat
	}

	if dto.LocationLng != nil {
		classroom.LocationLng = *dto.LocationLng
	}

	if err := uc.repository.Update(classroom); err != nil {
		return classroomDomain.NewInternalError("failed to update classroom in database", err)
	}

	return nil
}

func (uc *UseCase) GetClassroom(id uuid.UUID) (*classroomDomain.Classroom, error) {
	classroom, err := uc.repository.GetByID(id)
	if err != nil {
		return nil, classroomDomain.ErrClassroomNotFound
	}
	return classroom, nil
}

func (uc *UseCase) GetAllClassrooms() ([]*classroomDomain.Classroom, error) {
	classrooms, err := uc.repository.GetAll()
	if err != nil {
		return nil, classroomDomain.NewInternalError("failed to retrieve classrooms from database", err)
	}
	return classrooms, nil
}

func (uc *UseCase) GetClassroomsByFaculty(facultyID uuid.UUID) ([]*classroomDomain.Classroom, error) {
	classrooms, err := uc.repository.GetByFacultyID(facultyID)
	if err != nil {
		return nil, classroomDomain.NewInternalError("failed to retrieve classrooms by faculty from database", err)
	}
	return classrooms, nil
}

func (uc *UseCase) GetClassroomsByLocation(lat, lng, radius float64) ([]*classroomDomain.Classroom, error) {
	classrooms, err := uc.repository.GetByLocation(lat, lng, radius)
	if err != nil {
		return nil, classroomDomain.NewInternalError("failed to retrieve classrooms by location from database", err)
	}
	return classrooms, nil
}

func (uc *UseCase) DeleteClassroom(id uuid.UUID) error {
	// Check if classroom exists before deleting
	_, err := uc.repository.GetByID(id)
	if err != nil {
		return classroomDomain.ErrClassroomNotFound
	}

	if err := uc.repository.Delete(id); err != nil {
		return classroomDomain.NewInternalError("failed to delete classroom from database", err)
	}

	return nil
} 