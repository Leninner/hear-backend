-- name: CreateSchedule :one
INSERT INTO schedules (
    course_id,
    section_id,
    classroom_id,
    day_of_week,
    start_time,
    end_time
) VALUES (
    $1, $2, $3, $4, $5, $6
) RETURNING *;

-- name: GetScheduleByID :one
SELECT * FROM schedules
WHERE id = $1 LIMIT 1;

-- name: GetSchedulesByCourseID :many
SELECT * FROM schedules
WHERE course_id = $1;

-- name: GetSchedulesBySectionID :many
SELECT * FROM schedules
WHERE section_id = $1;

-- name: GetSchedulesByClassroomID :many
SELECT * FROM schedules
WHERE classroom_id = $1;

-- name: GetSchedulesByClassroomAndTime :many
SELECT * FROM schedules
WHERE classroom_id = $1
AND day_of_week = $2
AND (
    (start_time <= $3 AND end_time > $3) OR
    (start_time < $4 AND end_time >= $4) OR
    (start_time >= $3 AND end_time <= $4)
);

-- name: UpdateSchedule :one
UPDATE schedules
SET
    course_id = COALESCE($2, course_id),
    section_id = COALESCE($3, section_id),
    classroom_id = COALESCE($4, classroom_id),
    day_of_week = COALESCE($5, day_of_week),
    start_time = COALESCE($6, start_time),
    end_time = COALESCE($7, end_time),
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1
RETURNING *;

-- name: DeleteSchedule :exec
DELETE FROM schedules
WHERE id = $1; 