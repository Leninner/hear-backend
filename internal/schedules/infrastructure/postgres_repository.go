package infrastructure

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/leninner/hear-backend/internal/infrastructure/db"
	"github.com/leninner/hear-backend/internal/schedules/domain"
)

type PostgresRepository struct {
	db *db.Queries
}

func NewPostgresRepository(db *db.Queries) *PostgresRepository {
	return &PostgresRepository{
		db: db,
	}
}

func (r *PostgresRepository) Create(schedule *domain.Schedule) error {
	ctx := context.Background()
	
	startTime, err := time.Parse("15:04", schedule.StartTime.Format("15:04"))
	if err != nil {
		return err
	}

	endTime, err := time.Parse("15:04", schedule.EndTime.Format("15:04"))
	if err != nil {
		return err
	}

	params := db.CreateScheduleParams{
		CourseID:    schedule.CourseID,
		SectionID:   schedule.SectionID,
		ClassroomID: schedule.ClassroomID,
		DayOfWeek:   int32(schedule.DayOfWeek),
		StartTime:   startTime,
		EndTime:     endTime,
	}

	dbSchedule, err := r.db.CreateSchedule(ctx, params)
	if err != nil {
		return err
	}

	schedule.ID = dbSchedule.ID
	schedule.CreatedAt = dbSchedule.CreatedAt.Time
	schedule.UpdatedAt = dbSchedule.UpdatedAt.Time

	return nil
}

func (r *PostgresRepository) GetByID(id uuid.UUID) (*domain.Schedule, error) {
	ctx := context.Background()
	
	dbSchedule, err := r.db.GetScheduleByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &domain.Schedule{
		ID:          dbSchedule.ID,
		CourseID:    dbSchedule.CourseID,
		SectionID:   dbSchedule.SectionID,
		ClassroomID: dbSchedule.ClassroomID,
		DayOfWeek:   int(dbSchedule.DayOfWeek),
		StartTime:   dbSchedule.StartTime,
		EndTime:     dbSchedule.EndTime,
		CreatedAt:   dbSchedule.CreatedAt.Time,
		UpdatedAt:   dbSchedule.UpdatedAt.Time,
	}, nil
}

func (r *PostgresRepository) GetByCourseID(courseID uuid.UUID) ([]*domain.Schedule, error) {
	ctx := context.Background()
	
	dbSchedules, err := r.db.GetSchedulesByCourseID(ctx, courseID)
	if err != nil {
		return nil, err
	}

	schedules := make([]*domain.Schedule, len(dbSchedules))
	for i, dbSchedule := range dbSchedules {
		schedules[i] = &domain.Schedule{
			ID:          dbSchedule.ID,
			CourseID:    dbSchedule.CourseID,
			SectionID:   dbSchedule.SectionID,
			ClassroomID: dbSchedule.ClassroomID,
			DayOfWeek:   int(dbSchedule.DayOfWeek),
			StartTime:   dbSchedule.StartTime,
			EndTime:     dbSchedule.EndTime,
			CreatedAt:   dbSchedule.CreatedAt.Time,
			UpdatedAt:   dbSchedule.UpdatedAt.Time,
		}
	}

	return schedules, nil
}

func (r *PostgresRepository) GetBySectionID(sectionID uuid.UUID) ([]*domain.Schedule, error) {
	ctx := context.Background()
	
	dbSchedules, err := r.db.GetSchedulesBySectionID(ctx, sectionID)
	if err != nil {
		return nil, err
	}

	schedules := make([]*domain.Schedule, len(dbSchedules))
	for i, dbSchedule := range dbSchedules {
		schedules[i] = &domain.Schedule{
			ID:          dbSchedule.ID,
			CourseID:    dbSchedule.CourseID,
			SectionID:   dbSchedule.SectionID,
			ClassroomID: dbSchedule.ClassroomID,
			DayOfWeek:   int(dbSchedule.DayOfWeek),
			StartTime:   dbSchedule.StartTime,
			EndTime:     dbSchedule.EndTime,
			CreatedAt:   dbSchedule.CreatedAt.Time,
			UpdatedAt:   dbSchedule.UpdatedAt.Time,
		}
	}

	return schedules, nil
}

func (r *PostgresRepository) GetByClassroomID(classroomID uuid.UUID) ([]*domain.Schedule, error) {
	ctx := context.Background()
	
	dbSchedules, err := r.db.GetSchedulesByClassroomID(ctx, classroomID)
	if err != nil {
		return nil, err
	}

	schedules := make([]*domain.Schedule, len(dbSchedules))
	for i, dbSchedule := range dbSchedules {
		schedules[i] = &domain.Schedule{
			ID:          dbSchedule.ID,
			CourseID:    dbSchedule.CourseID,
			SectionID:   dbSchedule.SectionID,
			ClassroomID: dbSchedule.ClassroomID,
			DayOfWeek:   int(dbSchedule.DayOfWeek),
			StartTime:   dbSchedule.StartTime,
			EndTime:     dbSchedule.EndTime,
			CreatedAt:   dbSchedule.CreatedAt.Time,
			UpdatedAt:   dbSchedule.UpdatedAt.Time,
		}
	}

	return schedules, nil
}

func (r *PostgresRepository) GetByClassroomAndTime(classroomID uuid.UUID, dayOfWeek int, startTime, endTime time.Time) ([]*domain.Schedule, error) {
	ctx := context.Background()
	
	params := db.GetSchedulesByClassroomAndTimeParams{
		ClassroomID: classroomID,
		DayOfWeek:   int32(dayOfWeek),
		StartTime:   startTime,
		StartTime_2: endTime,
	}

	dbSchedules, err := r.db.GetSchedulesByClassroomAndTime(ctx, params)
	if err != nil {
		return nil, err
	}

	schedules := make([]*domain.Schedule, len(dbSchedules))
	for i, dbSchedule := range dbSchedules {
		schedules[i] = &domain.Schedule{
			ID:          dbSchedule.ID,
			CourseID:    dbSchedule.CourseID,
			SectionID:   dbSchedule.SectionID,
			ClassroomID: dbSchedule.ClassroomID,
			DayOfWeek:   int(dbSchedule.DayOfWeek),
			StartTime:   dbSchedule.StartTime,
			EndTime:     dbSchedule.EndTime,
			CreatedAt:   dbSchedule.CreatedAt.Time,
			UpdatedAt:   dbSchedule.UpdatedAt.Time,
		}
	}

	return schedules, nil
}

func (r *PostgresRepository) Update(schedule *domain.Schedule) error {
	ctx := context.Background()
	
	startTime, err := time.Parse("15:04", schedule.StartTime.Format("15:04"))
	if err != nil {
		return err
	}

	endTime, err := time.Parse("15:04", schedule.EndTime.Format("15:04"))
	if err != nil {
		return err
	}

	params := db.UpdateScheduleParams{
		ID:          schedule.ID,
		CourseID:    schedule.CourseID,
		SectionID:   schedule.SectionID,
		ClassroomID: schedule.ClassroomID,
		DayOfWeek:   int32(schedule.DayOfWeek),
		StartTime:   startTime,
		EndTime:     endTime,
	}

	dbSchedule, err := r.db.UpdateSchedule(ctx, params)
	if err != nil {
		return err
	}

	schedule.UpdatedAt = dbSchedule.UpdatedAt.Time
	return nil
}

func (r *PostgresRepository) Delete(id uuid.UUID) error {
	ctx := context.Background()
	
	return r.db.DeleteSchedule(ctx, id)
} 