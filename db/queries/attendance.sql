-- name: CreateAttendance :one
INSERT INTO attendance (
    student_id,
    class_schedule_id,
    status,
    date,
    user_latitude,
    user_longitude,
    distance_meters,
    max_distance_meters
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8
) RETURNING *;

-- name: GetAttendanceByID :one
SELECT * FROM attendance
WHERE id = $1 LIMIT 1;

-- name: GetAttendanceByStudentID :many
SELECT * FROM attendance
WHERE student_id = $1
AND date BETWEEN $2 AND $3;

-- name: GetAttendanceByClassScheduleID :many
SELECT * FROM attendance
WHERE class_schedule_id = $1
AND date = $2;

-- name: UpdateAttendance :one
UPDATE attendance
SET
    status = COALESCE($2, status),
    user_latitude = COALESCE($3, user_latitude),
    user_longitude = COALESCE($4, user_longitude),
    distance_meters = COALESCE($5, distance_meters),
    max_distance_meters = COALESCE($6, max_distance_meters),
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1
RETURNING *; 