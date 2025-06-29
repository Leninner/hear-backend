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

-- name: GetEnrollmentsWithDetailsBySection :many
SELECT 
    se.id,
    se.section_id,
    se.student_id,
    se.created_at,
    se.updated_at,
    u.id as student_id,
    u.email,
    u.first_name,
    u.last_name,
    u.role,
    u.created_at as user_created_at,
    u.updated_at as user_updated_at
FROM section_enrollments se
JOIN users u ON se.student_id = u.id
WHERE se.section_id = $1
ORDER BY se.created_at;

-- name: GetCourseSectionsByStudentID :many
SELECT cs.* FROM course_sections cs
JOIN section_enrollments se ON cs.id = se.section_id
WHERE se.student_id = $1;

-- name: GetCourseSectionsWithSchedulesByStudentID :many
SELECT 
    cs.id,
    cs.course_id,
    cs.teacher_id,
    cs.max_students,
    cs.created_at,
    cs.updated_at,
    cs.name,
    s.id as schedule_id,
    s.course_id as schedule_course_id,
    s.section_id as schedule_section_id,
    s.classroom_id,
    s.day_of_week,
    s.start_time,
    s.end_time,
    s.created_at as schedule_created_at,
    s.updated_at as schedule_updated_at
FROM course_sections cs
JOIN section_enrollments se ON cs.id = se.section_id
LEFT JOIN schedules s ON cs.id = s.section_id
WHERE se.student_id = $1
ORDER BY cs.id, s.day_of_week, s.start_time; 