-- Remove location fields from attendance table
DROP INDEX IF EXISTS idx_attendance_distance;
DROP INDEX IF EXISTS idx_attendance_location;

ALTER TABLE attendance 
DROP COLUMN IF EXISTS max_distance_meters,
DROP COLUMN IF EXISTS distance_meters,
DROP COLUMN IF EXISTS user_longitude,
DROP COLUMN IF EXISTS user_latitude; 