// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0
// source: universities.sql

package db

import (
	"context"

	"github.com/google/uuid"
)

const createUniversity = `-- name: CreateUniversity :one
INSERT INTO universities (
    name
) VALUES (
    $1
) RETURNING id, name, created_at, updated_at
`

func (q *Queries) CreateUniversity(ctx context.Context, name string) (University, error) {
	row := q.queryRow(ctx, q.createUniversityStmt, createUniversity, name)
	var i University
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteUniversity = `-- name: DeleteUniversity :exec
DELETE FROM universities
WHERE id = $1
`

func (q *Queries) DeleteUniversity(ctx context.Context, id uuid.UUID) error {
	_, err := q.exec(ctx, q.deleteUniversityStmt, deleteUniversity, id)
	return err
}

const getUniversityByID = `-- name: GetUniversityByID :one
SELECT id, name, created_at, updated_at FROM universities
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetUniversityByID(ctx context.Context, id uuid.UUID) (University, error) {
	row := q.queryRow(ctx, q.getUniversityByIDStmt, getUniversityByID, id)
	var i University
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUniversityByName = `-- name: GetUniversityByName :one
SELECT id, name, created_at, updated_at FROM universities
WHERE name = $1 LIMIT 1
`

func (q *Queries) GetUniversityByName(ctx context.Context, name string) (University, error) {
	row := q.queryRow(ctx, q.getUniversityByNameStmt, getUniversityByName, name)
	var i University
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listUniversities = `-- name: ListUniversities :many
SELECT id, name, created_at, updated_at FROM universities
ORDER BY name
`

func (q *Queries) ListUniversities(ctx context.Context) ([]University, error) {
	rows, err := q.query(ctx, q.listUniversitiesStmt, listUniversities)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []University{}
	for rows.Next() {
		var i University
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateUniversity = `-- name: UpdateUniversity :one
UPDATE universities
SET
    name = COALESCE($2, name),
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1
RETURNING id, name, created_at, updated_at
`

type UpdateUniversityParams struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

func (q *Queries) UpdateUniversity(ctx context.Context, arg UpdateUniversityParams) (University, error) {
	row := q.queryRow(ctx, q.updateUniversityStmt, updateUniversity, arg.ID, arg.Name)
	var i University
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
