-- name: CreateAttendance :one
INSERT INTO attendance (
    student_id,
    class_schedule_id,
    status,
    date
) VALUES (
    $1, $2, $3, $4
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
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1
RETURNING *; 