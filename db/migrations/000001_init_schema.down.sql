DROP TABLE IF EXISTS qr_codes;
DROP TABLE IF EXISTS attendance;
DROP TABLE IF EXISTS section_enrollments;
DROP TABLE IF EXISTS class_schedules;
DROP TABLE IF EXISTS course_sections;
DROP TABLE IF EXISTS courses;
DROP TABLE IF EXISTS classrooms;
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS faculties;
DROP TABLE IF EXISTS universities;

DROP TYPE IF EXISTS attendance_status;
DROP TYPE IF EXISTS user_role;

DROP EXTENSION IF EXISTS earthdistance;
DROP EXTENSION IF EXISTS cube;
DROP EXTENSION IF EXISTS "uuid-ossp"; 