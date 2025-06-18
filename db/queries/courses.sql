-- name: CreateCourse :one
INSERT INTO courses (
    faculty_id,
    name,
    code
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: GetCourseByID :one
SELECT * FROM courses
WHERE id = $1 LIMIT 1;

-- name: GetCoursesByFacultyID :many
SELECT * FROM courses
WHERE faculty_id = $1;

-- name: UpdateCourse :one
UPDATE courses
SET
    name = COALESCE($2, name),
    code = COALESCE($3, code),
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1
RETURNING *;

-- name: DeleteCourse :exec
DELETE FROM courses
WHERE id = $1;

-- name: CreateClassSchedule :one
INSERT INTO class_schedules (
    course_id,
    day_of_week,
    start_time,
    end_time,
    classroom_id
) VALUES (
    $1, $2, $3, $4, $5
) RETURNING *;

-- name: GetClassScheduleByID :one
SELECT * FROM class_schedules
WHERE id = $1 LIMIT 1;

-- name: GetClassSchedulesByCourseID :many
SELECT * FROM class_schedules
WHERE course_id = $1;

-- name: GetClassSchedulesByClassroomAndTime :many
SELECT * FROM class_schedules
WHERE classroom_id = $1
AND day_of_week = $2
AND (
    (start_time <= $3 AND end_time > $3) OR
    (start_time < $4 AND end_time >= $4) OR
    (start_time >= $3 AND end_time <= $4)
);

-- name: UpdateClassSchedule :one
UPDATE class_schedules
SET
    day_of_week = COALESCE($2, day_of_week),
    start_time = COALESCE($3, start_time),
    end_time = COALESCE($4, end_time),
    classroom_id = COALESCE($5, classroom_id),
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1
RETURNING *;

-- name: DeleteClassSchedule :exec
DELETE FROM class_schedules
WHERE id = $1; 