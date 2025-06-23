-- name: CreateQRCode :one
INSERT INTO qr_codes (
    course_section_id,
    code,
    expires_at
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: GetQRCodeByID :one
SELECT * FROM qr_codes
WHERE id = $1 LIMIT 1;

-- name: GetQRCodeByCode :one
SELECT * FROM qr_codes
WHERE code = $1 LIMIT 1;

-- name: GetQRCodesByCourseSectionID :many
SELECT * FROM qr_codes
WHERE course_section_id = $1;

-- name: GetActiveQRCodeByCourseSectionID :one
SELECT * FROM qr_codes
WHERE course_section_id = $1 
AND expires_at > NOW()
ORDER BY created_at DESC
LIMIT 1;

-- name: UpdateQRCode :one
UPDATE qr_codes
SET code = $2, expires_at = $3, updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: DeleteQRCode :exec
DELETE FROM qr_codes
WHERE id = $1; 