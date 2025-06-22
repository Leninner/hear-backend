-- Remove faculty_id from classrooms table
DROP INDEX IF EXISTS idx_classrooms_faculty;
ALTER TABLE classrooms DROP COLUMN IF EXISTS faculty_id; 