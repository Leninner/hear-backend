-- name: CreateAttendance :one
INSERT INTO attendance (
    student_id,
    class_schedule_id,
    attendance_date,
    image_url,
    image_metadata,
    is_valid
) VALUES (
    $1, $2, $3, $4, $5, $6
) RETURNING *;

-- name: GetAttendanceByID :one
SELECT * FROM attendance
WHERE id = $1 LIMIT 1;

-- name: GetAttendanceByStudentID :many
SELECT * FROM attendance
WHERE student_id = $1
AND attendance_date BETWEEN $2 AND $3;

-- name: GetAttendanceByClassScheduleID :many
SELECT * FROM attendance
WHERE class_schedule_id = $1
AND attendance_date = $2;

-- name: UpdateAttendance :one
UPDATE attendance
SET
    is_valid = COALESCE($2, is_valid),
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1
RETURNING *; 