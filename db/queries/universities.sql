-- name: CreateUniversity :one
INSERT INTO universities (
    name
) VALUES (
    $1
) RETURNING *;

-- name: GetUniversityByID :one
SELECT * FROM universities
WHERE id = $1 LIMIT 1;

-- name: GetUniversityByName :one
SELECT * FROM universities
WHERE name = $1 LIMIT 1;

-- name: ListUniversities :many
SELECT * FROM universities
ORDER BY name;

-- name: UpdateUniversity :one
UPDATE universities
SET
    name = COALESCE($2, name),
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1
RETURNING *;

-- name: DeleteUniversity :exec
DELETE FROM universities
WHERE id = $1; 