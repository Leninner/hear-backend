-- Add back location columns to faculties table
ALTER TABLE faculties ADD COLUMN location_lat DECIMAL(10, 8);
ALTER TABLE faculties ADD COLUMN location_lng DECIMAL(11, 8); 