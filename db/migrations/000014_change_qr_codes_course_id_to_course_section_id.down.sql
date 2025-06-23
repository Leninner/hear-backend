-- Delete existing QR codes since we're reverting back to course_id
DELETE FROM qr_codes;

-- Revert course_section_id back to course_id in qr_codes table
ALTER TABLE qr_codes RENAME COLUMN course_section_id TO course_id;

-- Update the foreign key constraint back to courses
ALTER TABLE qr_codes DROP CONSTRAINT IF EXISTS qr_codes_course_section_id_fkey;
ALTER TABLE qr_codes ADD CONSTRAINT qr_codes_course_id_fkey 
    FOREIGN KEY (course_id) REFERENCES courses(id);

-- Update the index back
DROP INDEX IF EXISTS idx_qr_codes_course_section;
CREATE INDEX idx_qr_codes_course ON qr_codes(course_id); 