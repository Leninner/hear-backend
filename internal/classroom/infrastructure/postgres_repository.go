package infrastructure

import (
	"database/sql"
	"errors"

	"github.com/google/uuid"
	"github.com/leninner/hear-backend/internal/classroom/domain"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) *PostgresRepository {
	return &PostgresRepository{
		db: db,
	}
}

func (r *PostgresRepository) Create(classroom *domain.Classroom) error {
	query := `
		INSERT INTO classrooms (id, name, building, floor, capacity, location_lat, location_lng, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	`
	_, err := r.db.Exec(query,
		classroom.ID,
		classroom.Name,
		classroom.Building,
		classroom.Floor,
		classroom.Capacity,
		classroom.LocationLat,
		classroom.LocationLng,
		classroom.CreatedAt,
		classroom.UpdatedAt,
	)
	return err
}

func (r *PostgresRepository) GetByID(id uuid.UUID) (*domain.Classroom, error) {
	query := `
		SELECT id, name, building, floor, capacity, location_lat, location_lng, created_at, updated_at
		FROM classrooms
		WHERE id = $1
	`
	classroom := &domain.Classroom{}
	err := r.db.QueryRow(query, id).Scan(
		&classroom.ID,
		&classroom.Name,
		&classroom.Building,
		&classroom.Floor,
		&classroom.Capacity,
		&classroom.LocationLat,
		&classroom.LocationLng,
		&classroom.CreatedAt,
		&classroom.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, errors.New("classroom not found")
	}
	return classroom, err
}

func (r *PostgresRepository) GetAll() ([]*domain.Classroom, error) {
	query := `
		SELECT id, name, building, floor, capacity, location_lat, location_lng, created_at, updated_at
		FROM classrooms
		ORDER BY name
	`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var classrooms []*domain.Classroom
	for rows.Next() {
		classroom := &domain.Classroom{}
		err := rows.Scan(
			&classroom.ID,
			&classroom.Name,
			&classroom.Building,
			&classroom.Floor,
			&classroom.Capacity,
			&classroom.LocationLat,
			&classroom.LocationLng,
			&classroom.CreatedAt,
			&classroom.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		classrooms = append(classrooms, classroom)
	}
	return classrooms, nil
}

func (r *PostgresRepository) GetByLocation(lat, lng, radius float64) ([]*domain.Classroom, error) {
	query := `
		SELECT id, name, building, floor, capacity, location_lat, location_lng, created_at, updated_at
		FROM classrooms
		WHERE (
			(6371 * acos(
				cos(radians($1)) * cos(radians(location_lat)) *
				cos(radians(location_lng) - radians($2)) +
				sin(radians($1)) * sin(radians(location_lat))
			)) <= $3
		)
		ORDER BY (
			(6371 * acos(
				cos(radians($1)) * cos(radians(location_lat)) *
				cos(radians(location_lng) - radians($2)) +
				sin(radians($1)) * sin(radians(location_lat))
			))
		)
	`
	rows, err := r.db.Query(query, lat, lng, radius)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var classrooms []*domain.Classroom
	for rows.Next() {
		classroom := &domain.Classroom{}
		err := rows.Scan(
			&classroom.ID,
			&classroom.Name,
			&classroom.Building,
			&classroom.Floor,
			&classroom.Capacity,
			&classroom.LocationLat,
			&classroom.LocationLng,
			&classroom.CreatedAt,
			&classroom.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		classrooms = append(classrooms, classroom)
	}
	return classrooms, nil
}

func (r *PostgresRepository) Update(classroom *domain.Classroom) error {
	query := `
		UPDATE classrooms
		SET name = $1, building = $2, floor = $3, capacity = $4,
			location_lat = $5, location_lng = $6, updated_at = $7
		WHERE id = $8
	`
	result, err := r.db.Exec(query,
		classroom.Name,
		classroom.Building,
		classroom.Floor,
		classroom.Capacity,
		classroom.LocationLat,
		classroom.LocationLng,
		classroom.UpdatedAt,
		classroom.ID,
	)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("classroom not found")
	}
	return nil
}

func (r *PostgresRepository) Delete(id uuid.UUID) error {
	query := `DELETE FROM classrooms WHERE id = $1`
	result, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("classroom not found")
	}
	return nil
} 