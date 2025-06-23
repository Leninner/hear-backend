-- name: CreateCourseSection :one
INSERT INTO course_sections (
    course_id,
    name,
    teacher_id,
    max_students
) VALUES (
    $1, $2, $3, $4
) RETURNING *;

-- name: GetCourseSectionByID :one
SELECT * FROM course_sections
WHERE id = $1 LIMIT 1;

-- name: GetCourseSectionsByCourseID :many
SELECT * FROM course_sections
WHERE course_id = $1;

-- name: GetCourseSectionsByTeacherID :many
SELECT * FROM course_sections
WHERE teacher_id = $1;

-- name: UpdateCourseSection :one
UPDATE course_sections
SET
    name = COALESCE($2, name),
    teacher_id = COALESCE($3, teacher_id),
    max_students = COALESCE($4, max_students),
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1
RETURNING *;

-- name: DeleteCourseSection :exec
DELETE FROM course_sections
WHERE id = $1;

-- name: EnrollStudent :one
INSERT INTO section_enrollments (
    section_id,
    student_id
) VALUES (
    $1, $2
) RETURNING *;

-- name: GetEnrollmentCount :one
SELECT COUNT(*) FROM section_enrollments
WHERE section_id = $1;

-- name: IsStudentEnrolled :one
SELECT EXISTS(
    SELECT 1 FROM section_enrollments
    WHERE section_id = $1 AND student_id = $2
);

-- name: UnenrollStudent :exec
DELETE FROM section_enrollments
WHERE section_id = $1 AND student_id = $2; 