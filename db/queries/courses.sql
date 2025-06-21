-- name: CreateCourse :one
INSERT INTO courses (
    faculty_id,
    name,
    semester
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: GetCourseByID :one
SELECT * FROM courses
WHERE id = $1 LIMIT 1;

-- name: GetCoursesByFacultyID :many
SELECT * FROM courses
WHERE faculty_id = $1;

-- name: GetCoursesBySemester :many
SELECT * FROM courses
WHERE semester = $1;

-- name: UpdateCourse :one
UPDATE courses
SET
    name = COALESCE($2, name),
    semester = COALESCE($3, semester),
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1
RETURNING *;

-- name: DeleteCourse :exec
DELETE FROM courses
WHERE id = $1; 