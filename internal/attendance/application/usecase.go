package application

import (
	"time"

	"github.com/google/uuid"
	"github.com/leninner/hear-backend/internal/attendance/domain"
)

type UseCase struct {
	repository domain.Repository
}

func NewUseCase(repository domain.Repository) *UseCase {
	return &UseCase{
		repository: repository,
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

func (uc *UseCase) GetAttendance(id uuid.UUID) (*domain.Attendance, error) {
	return uc.repository.GetByID(id)
}

func (uc *UseCase) GetStudentAttendance(studentID uuid.UUID) ([]*domain.Attendance, error) {
	return uc.repository.GetByStudentID(studentID)
}

func (uc *UseCase) GetClassScheduleAttendance(classScheduleID uuid.UUID) ([]*domain.Attendance, error) {
	return uc.repository.GetByClassScheduleID(classScheduleID)
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