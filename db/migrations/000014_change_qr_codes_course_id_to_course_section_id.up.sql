-- Delete existing QR codes since they reference courses, not course sections
DELETE FROM qr_codes;

-- Change course_id to course_section_id in qr_codes table
ALTER TABLE qr_codes RENAME COLUMN course_id TO course_section_id;

-- Update the foreign key constraint
ALTER TABLE qr_codes DROP CONSTRAINT IF EXISTS qr_codes_course_id_fkey;
ALTER TABLE qr_codes ADD CONSTRAINT qr_codes_course_section_id_fkey 
    FOREIGN KEY (course_section_id) REFERENCES course_sections(id);

-- Update the index
DROP INDEX IF EXISTS idx_qr_codes_course;
CREATE INDEX idx_qr_codes_course_section ON qr_codes(course_section_id); 