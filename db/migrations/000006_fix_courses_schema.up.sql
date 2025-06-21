-- Remove location fields from courses table and simplify structure
ALTER TABLE courses DROP COLUMN IF EXISTS location_lat;
ALTER TABLE courses DROP COLUMN IF EXISTS location_lng;
ALTER TABLE courses DROP COLUMN IF EXISTS code;
ALTER TABLE courses DROP COLUMN IF EXISTS description;
ALTER TABLE courses DROP COLUMN IF EXISTS credits;
ALTER TABLE courses DROP COLUMN IF EXISTS academic_year;
ALTER TABLE courses DROP COLUMN IF EXISTS is_active;

-- Add semester field if it doesn't exist
ALTER TABLE courses ADD COLUMN IF NOT EXISTS semester VARCHAR(20) NOT NULL DEFAULT '2024-1';

-- Update course_sections table to have a name instead of just section_number
ALTER TABLE course_sections ADD COLUMN IF NOT EXISTS name VARCHAR(100) NOT NULL DEFAULT 'Section 1';
ALTER TABLE course_sections DROP COLUMN IF EXISTS section_number;

-- Create schedules table if it doesn't exist
CREATE TABLE IF NOT EXISTS schedules (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    course_id UUID NOT NULL REFERENCES courses(id) ON DELETE CASCADE,
    section_id UUID NOT NULL REFERENCES course_sections(id) ON DELETE CASCADE,
    classroom_id UUID NOT NULL REFERENCES classrooms(id) ON DELETE CASCADE,
    day_of_week INTEGER NOT NULL CHECK (day_of_week BETWEEN 0 AND 6),
    start_time TIME NOT NULL,
    end_time TIME NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(course_id, section_id, classroom_id, day_of_week, start_time)
);

-- Drop the old class_schedules table if it exists
DROP TABLE IF EXISTS class_schedules CASCADE;

-- Update attendance table to reference schedules instead of class_schedules
ALTER TABLE attendance DROP CONSTRAINT IF EXISTS attendance_class_schedule_id_fkey;
ALTER TABLE attendance RENAME COLUMN class_schedule_id TO schedule_id;
ALTER TABLE attendance ADD CONSTRAINT attendance_schedule_id_fkey 
    FOREIGN KEY (schedule_id) REFERENCES schedules(id) ON DELETE CASCADE;

-- Update indexes
DROP INDEX IF EXISTS idx_class_schedules_course;
DROP INDEX IF EXISTS idx_class_schedules_classroom;
CREATE INDEX idx_schedules_course ON schedules(course_id);
CREATE INDEX idx_schedules_section ON schedules(section_id);
CREATE INDEX idx_schedules_classroom ON schedules(classroom_id);
CREATE INDEX idx_schedules_time ON schedules(day_of_week, start_time, end_time);
CREATE INDEX idx_attendance_schedule ON attendance(schedule_id); 