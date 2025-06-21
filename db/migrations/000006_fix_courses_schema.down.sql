-- Revert attendance table changes
ALTER TABLE attendance DROP CONSTRAINT IF EXISTS attendance_schedule_id_fkey;
ALTER TABLE attendance RENAME COLUMN schedule_id TO class_schedule_id;

-- Recreate class_schedules table
CREATE TABLE class_schedules (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    course_id UUID NOT NULL REFERENCES courses(id),
    section_id UUID NOT NULL REFERENCES course_sections(id),
    day_of_week INTEGER NOT NULL CHECK (day_of_week BETWEEN 0 AND 6),
    start_time TIME NOT NULL,
    end_time TIME NOT NULL,
    classroom_id UUID NOT NULL REFERENCES classrooms(id),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Re-add attendance foreign key
ALTER TABLE attendance ADD CONSTRAINT attendance_class_schedule_id_fkey 
    FOREIGN KEY (class_schedule_id) REFERENCES class_schedules(id);

-- Drop schedules table
DROP TABLE IF EXISTS schedules;

-- Revert course_sections table changes
ALTER TABLE course_sections DROP COLUMN IF EXISTS name;
ALTER TABLE course_sections ADD COLUMN IF NOT EXISTS section_number INTEGER NOT NULL DEFAULT 1;

-- Revert courses table changes
ALTER TABLE courses DROP COLUMN IF EXISTS semester;
ALTER TABLE courses ADD COLUMN IF NOT EXISTS code VARCHAR(50) NOT NULL DEFAULT '';
ALTER TABLE courses ADD COLUMN IF NOT EXISTS description TEXT;
ALTER TABLE courses ADD COLUMN IF NOT EXISTS credits INTEGER NOT NULL DEFAULT 3;
ALTER TABLE courses ADD COLUMN IF NOT EXISTS academic_year VARCHAR(9);
ALTER TABLE courses ADD COLUMN IF NOT EXISTS is_active BOOLEAN DEFAULT true;
ALTER TABLE courses ADD COLUMN IF NOT EXISTS location_lat DECIMAL(10, 8);
ALTER TABLE courses ADD COLUMN IF NOT EXISTS location_lng DECIMAL(11, 8);

-- Recreate indexes
DROP INDEX IF EXISTS idx_schedules_course;
DROP INDEX IF EXISTS idx_schedules_section;
DROP INDEX IF EXISTS idx_schedules_classroom;
DROP INDEX IF EXISTS idx_schedules_time;
DROP INDEX IF EXISTS idx_attendance_schedule;
CREATE INDEX idx_class_schedules_course ON class_schedules(course_id);
CREATE INDEX idx_class_schedules_classroom ON class_schedules(classroom_id);
CREATE INDEX idx_attendance_class_schedule ON attendance(class_schedule_id); 