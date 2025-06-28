package application

import (
	"time"

	"github.com/google/uuid"
	"github.com/leninner/hear-backend/internal/attendance/domain"
)

type UseCase struct {
	repository         domain.Repository
	scheduleRepository domain.ScheduleRepository
	classroomRepository domain.ClassroomRepository
}

func NewUseCase(repository domain.Repository, scheduleRepository domain.ScheduleRepository, classroomRepository domain.ClassroomRepository) *UseCase {
	return &UseCase{
		repository:         repository,
		scheduleRepository: scheduleRepository,
		classroomRepository: classroomRepository,
	}
}

func (uc *UseCase) CreateAttendance(studentID, classScheduleID uuid.UUID, status domain.AttendanceStatus, date time.Time) (*domain.Attendance, error) {
	attendance := domain.NewAttendance(studentID, classScheduleID, status, date)
	err := uc.repository.Create(attendance)
	if err != nil {
		return nil, err
	}
	return attendance, nil
}

func (uc *UseCase) CreateAttendanceWithLocation(
	studentID, classScheduleID uuid.UUID, 
	status domain.AttendanceStatus, 
	date time.Time,
	userLatitude, userLongitude *float64,
) (*domain.Attendance, error) {
	// Check if attendance already exists
	existing, err := uc.repository.GetByStudentScheduleAndDate(studentID, classScheduleID, date)
	if err == nil && existing != nil {
		// Update existing attendance
		existing.Status = status
		if userLatitude != nil && userLongitude != nil {
			existing.UserLatitude = userLatitude
			existing.UserLongitude = userLongitude
			// Get schedule to find classroom ID
			schedule, err := uc.scheduleRepository.GetByID(classScheduleID)
			if err != nil {
				return nil, err
			}
			// Get classroom to get location coordinates
			classroom, err := uc.classroomRepository.GetByID(schedule.ClassroomID)
			if err != nil {
				return nil, err
			}
			// Calculate distance between user and classroom
			distance := domain.CalculateDistance(*userLatitude, *userLongitude, classroom.LocationLat, classroom.LocationLng)
			existing.DistanceMeters = &distance
			// Check if user is within allowed distance
			if !existing.IsWithinDistance() {
				return nil, domain.NewValidationError(domain.ErrDistanceTooFar)
			}
		}
		err = uc.repository.Update(existing)
		if err != nil {
			return nil, err
		}
		return existing, nil
	}
	// If not found, create new attendance
	attendance := domain.NewAttendance(studentID, classScheduleID, status, date)
	if userLatitude != nil && userLongitude != nil {
		attendance.UserLatitude = userLatitude
		attendance.UserLongitude = userLongitude
		// Get schedule to find classroom ID
		schedule, err := uc.scheduleRepository.GetByID(classScheduleID)
		if err != nil {
			return nil, err
		}
		// Get classroom to get location coordinates
		classroom, err := uc.classroomRepository.GetByID(schedule.ClassroomID)
		if err != nil {
			return nil, err
		}
		// Calculate distance between user and classroom
		distance := domain.CalculateDistance(*userLatitude, *userLongitude, classroom.LocationLat, classroom.LocationLng)
		attendance.DistanceMeters = &distance
		// Check if user is within allowed distance
		if !attendance.IsWithinDistance() {
			return nil, domain.NewValidationError(domain.ErrDistanceTooFar)
		}
	}
	err = uc.repository.Create(attendance)
	if err != nil {
		return nil, err
	}
	return attendance, nil
}

func (uc *UseCase) GetAttendance(id uuid.UUID) (*domain.Attendance, error) {
	return uc.repository.GetByID(id)
}

func (uc *UseCase) GetStudentAttendance(studentID uuid.UUID) ([]*domain.Attendance, error) {
	return uc.repository.GetByStudentID(studentID)
}

func (uc *UseCase) GetClassScheduleAttendance(classScheduleID uuid.UUID) ([]*domain.Attendance, error) {
	return uc.repository.GetByScheduleID(classScheduleID)
}

func (uc *UseCase) GetAttendanceByDate(date time.Time) ([]*domain.Attendance, error) {
	return uc.repository.GetByDate(date)
}

func (uc *UseCase) UpdateAttendance(attendance *domain.Attendance) error {
	return uc.repository.Update(attendance)
}

func (uc *UseCase) DeleteAttendance(id uuid.UUID) error {
	return uc.repository.Delete(id)
} 