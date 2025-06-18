package infrastructure

import (
	"context"
	"strconv"

	"github.com/google/uuid"
	"github.com/leninner/hear-backend/internal/classroom/domain"
	"github.com/leninner/hear-backend/internal/infrastructure/db"
)

type PostgresRepository struct {
	db *db.Queries
}

func NewPostgresRepository(db *db.Queries) *PostgresRepository {
	return &PostgresRepository{
		db: db,
	}
}

func (r *PostgresRepository) Create(classroom *domain.Classroom) error {
	ctx := context.Background()
	
	params := db.CreateClassroomParams{
		Name:        classroom.Name,
		Building:    classroom.Building,
		Floor:       int32(classroom.Floor),
		Capacity:    int32(classroom.Capacity),
		LocationLat: strconv.FormatFloat(classroom.LocationLat, 'f', -1, 64),
		LocationLng: strconv.FormatFloat(classroom.LocationLng, 'f', -1, 64),
	}
	
	result, err := r.db.CreateClassroom(ctx, params)
	if err != nil {
		return domain.NewInternalError("failed to create classroom in database", err)
	}

	classroom.ID = result.ID
	if result.CreatedAt.Valid {
		classroom.CreatedAt = result.CreatedAt.Time
	}
	if result.UpdatedAt.Valid {
		classroom.UpdatedAt = result.UpdatedAt.Time
	}
	
	return nil
}

func (r *PostgresRepository) GetByID(id uuid.UUID) (*domain.Classroom, error) {
	ctx := context.Background()
	
	dbClassroom, err := r.db.GetClassroomByID(ctx, id)
	if err != nil {
		return nil, domain.ErrClassroomNotFound
	}
	
	return r.mapToDomain(&dbClassroom), nil
}

func (r *PostgresRepository) GetByName(name string) (*domain.Classroom, error) {
	ctx := context.Background()
	
	dbClassroom, err := r.db.GetClassroomByName(ctx, name)
	if err != nil {
		return nil, domain.ErrClassroomNotFound
	}
	
	return r.mapToDomain(&dbClassroom), nil
}

func (r *PostgresRepository) GetAll() ([]*domain.Classroom, error) {
	ctx := context.Background()
	
	dbClassrooms, err := r.db.GetAll(ctx, db.GetAllParams{
		Limit:  100,
		Offset: 0,
	})
	if err != nil {
		return nil, domain.NewInternalError("failed to retrieve classrooms from database", err)
	}
	
	var classrooms []*domain.Classroom
	for _, dbClassroom := range dbClassrooms {
		classrooms = append(classrooms, r.mapToDomain(&dbClassroom))
	}
	
	return classrooms, nil
}

func (r *PostgresRepository) GetByLocation(lat, lng, radius float64) ([]*domain.Classroom, error) {
	ctx := context.Background()
	
	params := db.GetNearbyClassroomsParams{
		LlToEarth:   lat,
		LlToEarth_2: lng,
		EarthBox:    radius,
	}
	
	dbClassrooms, err := r.db.GetNearbyClassrooms(ctx, params)
	if err != nil {
		return nil, domain.NewInternalError("failed to retrieve classrooms by location from database", err)
	}
	
	var classrooms []*domain.Classroom
	for _, dbClassroom := range dbClassrooms {
		classroom := &domain.Classroom{
			ID:          dbClassroom.ID,
			Name:        dbClassroom.Name,
			Building:    dbClassroom.Building,
			Floor:       int(dbClassroom.Floor),
			Capacity:    int(dbClassroom.Capacity),
			LocationLat: r.parseFloat(dbClassroom.LocationLat),
			LocationLng: r.parseFloat(dbClassroom.LocationLng),
		}
		
		if dbClassroom.CreatedAt.Valid {
			classroom.CreatedAt = dbClassroom.CreatedAt.Time
		}
		if dbClassroom.UpdatedAt.Valid {
			classroom.UpdatedAt = dbClassroom.UpdatedAt.Time
		}
		
		classrooms = append(classrooms, classroom)
	}
	
	return classrooms, nil
}

func (r *PostgresRepository) Update(classroom *domain.Classroom) error {
	ctx := context.Background()
	
	params := db.UpdateClassroomParams{
		ID:          classroom.ID,
		Name:        classroom.Name,
		Building:    classroom.Building,
		Floor:       int32(classroom.Floor),
		Capacity:    int32(classroom.Capacity),
		LocationLat: strconv.FormatFloat(classroom.LocationLat, 'f', -1, 64),
		LocationLng: strconv.FormatFloat(classroom.LocationLng, 'f', -1, 64),
	}
	
	_, err := r.db.UpdateClassroom(ctx, params)
	if err != nil {
		return domain.NewInternalError("failed to update classroom in database", err)
	}
	
	return nil
}

func (r *PostgresRepository) Delete(id uuid.UUID) error {
	ctx := context.Background()
	
	err := r.db.DeleteClassroom(ctx, id)
	if err != nil {
		return domain.ErrClassroomNotFound
	}
	
	return nil
}

func (r *PostgresRepository) mapToDomain(dbClassroom *db.Classroom) *domain.Classroom {
	classroom := &domain.Classroom{
		ID:          dbClassroom.ID,
		Name:        dbClassroom.Name,
		Building:    dbClassroom.Building,
		Floor:       int(dbClassroom.Floor),
		Capacity:    int(dbClassroom.Capacity),
		LocationLat: r.parseFloat(dbClassroom.LocationLat),
		LocationLng: r.parseFloat(dbClassroom.LocationLng),
	}
	
	if dbClassroom.CreatedAt.Valid {
		classroom.CreatedAt = dbClassroom.CreatedAt.Time
	}
	if dbClassroom.UpdatedAt.Valid {
		classroom.UpdatedAt = dbClassroom.UpdatedAt.Time
	}
	
	return classroom
}

func (r *PostgresRepository) parseFloat(s string) float64 {
	f, _ := strconv.ParseFloat(s, 64)
	return f
} 