-- name: CreateQRCode :one
INSERT INTO qr_codes (
    course_id,
    code,
    expires_at
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: GetQRCodeByCode :one
SELECT * FROM qr_codes
WHERE code = $1 LIMIT 1;

-- name: GetQRCodesByCourseID :many
SELECT * FROM qr_codes
WHERE course_id = $1;

-- name: DeleteQRCode :exec
DELETE FROM qr_codes
WHERE id = $1; 