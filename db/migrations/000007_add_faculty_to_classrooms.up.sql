-- Add faculty_id to classrooms table
ALTER TABLE classrooms ADD COLUMN faculty_id UUID REFERENCES faculties(id);

-- Create index for better query performance
CREATE INDEX idx_classrooms_faculty ON classrooms(faculty_id); 