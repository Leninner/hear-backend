-- Add location fields to attendance table for distance validation
ALTER TABLE attendance 
ADD COLUMN user_latitude DECIMAL(10, 8),
ADD COLUMN user_longitude DECIMAL(11, 8),
ADD COLUMN distance_meters DECIMAL(10, 2),
ADD COLUMN max_distance_meters INTEGER DEFAULT 100;

-- Add index for location-based queries
CREATE INDEX idx_attendance_location ON attendance(user_latitude, user_longitude);
CREATE INDEX idx_attendance_distance ON attendance(distance_meters); 