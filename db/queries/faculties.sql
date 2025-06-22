-- name: CreateFaculty :one
INSERT INTO faculties (
    university_id,
    name
) VALUES (
    $1, $2
) RETURNING *;

-- name: GetFacultyByID :one
SELECT * FROM faculties
WHERE id = $1 LIMIT 1;

-- name: GetFacultyByName :one
SELECT * FROM faculties
WHERE name = $1 LIMIT 1;

-- name: ListFaculties :many
SELECT * FROM faculties
ORDER BY name;

-- name: GetFacultiesByUniversityID :many
SELECT * FROM faculties
WHERE university_id = $1
ORDER BY name;

-- name: UpdateFaculty :one
UPDATE faculties
SET
    university_id = COALESCE($2, university_id),
    name = COALESCE($3, name),
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1
RETURNING *;

-- name: DeleteFaculty :exec
DELETE FROM faculties
WHERE id = $1; 