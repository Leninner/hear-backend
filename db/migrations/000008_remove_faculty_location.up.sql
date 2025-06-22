-- Remove location columns from faculties table
ALTER TABLE faculties DROP COLUMN IF EXISTS location_lat;
ALTER TABLE faculties DROP COLUMN IF EXISTS location_lng; 