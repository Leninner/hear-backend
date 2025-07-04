-- Test schedules for one section per faculty
-- Universidad Nacional de Colombia - Facultad de Ingeniería
INSERT INTO schedules (id, course_id, section_id, classroom_id, day_of_week, start_time, end_time, created_at, updated_at) VALUES
  ('bb0e8400-e29b-41d4-a716-446655440001', '990e8400-e29b-41d4-a716-446655440001', 'aa0e8400-e29b-41d4-a716-446655440001', '880e8400-e29b-41d4-a716-446655440001', 1, '08:00', '09:00', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
  ('bb0e8400-e29b-41d4-a716-446655440002', '990e8400-e29b-41d4-a716-446655440001', 'aa0e8400-e29b-41d4-a716-446655440001', '880e8400-e29b-41d4-a716-446655440002', 3, '09:00', '10:00', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
  ('bb0e8400-e29b-41d4-a716-446655440003', '990e8400-e29b-41d4-a716-446655440001', 'aa0e8400-e29b-41d4-a716-446655440001', '880e8400-e29b-41d4-a716-446655440003', 5, '10:00', '11:00', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

-- Universidad Nacional de Colombia - Facultad de Ciencias
INSERT INTO schedules (id, course_id, section_id, classroom_id, day_of_week, start_time, end_time, created_at, updated_at) VALUES
  ('bb0e8400-e29b-41d4-a716-446655440004', '990e8400-e29b-41d4-a716-446655440007', 'aa0e8400-e29b-41d4-a716-446655440015', '880e8400-e29b-41d4-a716-446655440008', 1, '08:00', '09:00', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
  ('bb0e8400-e29b-41d4-a716-446655440005', '990e8400-e29b-41d4-a716-446655440007', 'aa0e8400-e29b-41d4-a716-446655440015', '880e8400-e29b-41d4-a716-446655440009', 3, '09:00', '10:00', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
  ('bb0e8400-e29b-41d4-a716-446655440006', '990e8400-e29b-41d4-a716-446655440007', 'aa0e8400-e29b-41d4-a716-446655440015', '880e8400-e29b-41d4-a716-446655440010', 5, '10:00', '11:00', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

-- Universidad Nacional de Colombia - Facultad de Medicina
INSERT INTO schedules (id, course_id, section_id, classroom_id, day_of_week, start_time, end_time, created_at, updated_at) VALUES
  ('bb0e8400-e29b-41d4-a716-446655440007', '990e8400-e29b-41d4-a716-446655440011', 'aa0e8400-e29b-41d4-a716-446655440024', '880e8400-e29b-41d4-a716-446655440012', 1, '08:00', '09:00', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
  ('bb0e8400-e29b-41d4-a716-446655440008', '990e8400-e29b-41d4-a716-446655440011', 'aa0e8400-e29b-41d4-a716-446655440024', '880e8400-e29b-41d4-a716-446655440013', 3, '09:00', '10:00', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
  ('bb0e8400-e29b-41d4-a716-446655440009', '990e8400-e29b-41d4-a716-446655440011', 'aa0e8400-e29b-41d4-a716-446655440024', '880e8400-e29b-41d4-a716-446655440014', 5, '10:00', '11:00', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

-- Universidad de los Andes - Facultad de Ingeniería
INSERT INTO schedules (id, course_id, section_id, classroom_id, day_of_week, start_time, end_time, created_at, updated_at) VALUES
  ('bb0e8400-e29b-41d4-a716-446655440010', '990e8400-e29b-41d4-a716-446655440015', 'aa0e8400-e29b-41d4-a716-446655440032', '880e8400-e29b-41d4-a716-446655440015', 1, '08:00', '09:00', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
  ('bb0e8400-e29b-41d4-a716-446655440011', '990e8400-e29b-41d4-a716-446655440015', 'aa0e8400-e29b-41d4-a716-446655440032', '880e8400-e29b-41d4-a716-446655440016', 3, '09:00', '10:00', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
  ('bb0e8400-e29b-41d4-a716-446655440012', '990e8400-e29b-41d4-a716-446655440015', 'aa0e8400-e29b-41d4-a716-446655440032', '880e8400-e29b-41d4-a716-446655440017', 5, '10:00', '11:00', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

-- Universidad de los Andes - Facultad de Administración
INSERT INTO schedules (id, course_id, section_id, classroom_id, day_of_week, start_time, end_time, created_at, updated_at) VALUES
  ('bb0e8400-e29b-41d4-a716-446655440013', '990e8400-e29b-41d4-a716-446655440020', 'aa0e8400-e29b-41d4-a716-446655440042', '880e8400-e29b-41d4-a716-446655440020', 1, '08:00', '09:00', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
  ('bb0e8400-e29b-41d4-a716-446655440014', '990e8400-e29b-41d4-a716-446655440020', 'aa0e8400-e29b-41d4-a716-446655440042', '880e8400-e29b-41d4-a716-446655440021', 3, '09:00', '10:00', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
  ('bb0e8400-e29b-41d4-a716-446655440015', '990e8400-e29b-41d4-a716-446655440020', 'aa0e8400-e29b-41d4-a716-446655440042', '880e8400-e29b-41d4-a716-446655440022', 5, '10:00', '11:00', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

-- Universidad Javeriana - Facultad de Ingeniería
INSERT INTO schedules (id, course_id, section_id, classroom_id, day_of_week, start_time, end_time, created_at, updated_at) VALUES
  ('bb0e8400-e29b-41d4-a716-446655440016', '990e8400-e29b-41d4-a716-446655440024', 'aa0e8400-e29b-41d4-a716-446655440050', '880e8400-e29b-41d4-a716-446655440023', 1, '08:00', '09:00', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
  ('bb0e8400-e29b-41d4-a716-446655440017', '990e8400-e29b-41d4-a716-446655440024', 'aa0e8400-e29b-41d4-a716-446655440050', '880e8400-e29b-41d4-a716-446655440024', 3, '09:00', '10:00', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
  ('bb0e8400-e29b-41d4-a716-446655440018', '990e8400-e29b-41d4-a716-446655440024', 'aa0e8400-e29b-41d4-a716-446655440050', '880e8400-e29b-41d4-a716-446655440025', 5, '10:00', '11:00', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

-- Universidad de Antioquia - Facultad de Ingeniería
INSERT INTO schedules (id, course_id, section_id, classroom_id, day_of_week, start_time, end_time, created_at, updated_at) VALUES
  ('bb0e8400-e29b-41d4-a716-446655440019', '990e8400-e29b-41d4-a716-446655440028', 'aa0e8400-e29b-41d4-a716-446655440058', '880e8400-e29b-41d4-a716-446655440027', 1, '08:00', '09:00', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
  ('bb0e8400-e29b-41d4-a716-446655440020', '990e8400-e29b-41d4-a716-446655440028', 'aa0e8400-e29b-41d4-a716-446655440058', '880e8400-e29b-41d4-a716-446655440028', 3, '09:00', '10:00', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
  ('bb0e8400-e29b-41d4-a716-446655440021', '990e8400-e29b-41d4-a716-446655440028', 'aa0e8400-e29b-41d4-a716-446655440058', '880e8400-e29b-41d4-a716-446655440029', 5, '10:00', '11:00', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

-- Universidad del Valle - Facultad de Ingeniería
INSERT INTO schedules (id, course_id, section_id, classroom_id, day_of_week, start_time, end_time, created_at, updated_at) VALUES
  ('bb0e8400-e29b-41d4-a716-446655440022', '990e8400-e29b-41d4-a716-446655440032', 'aa0e8400-e29b-41d4-a716-446655440066', '880e8400-e29b-41d4-a716-446655440031', 1, '08:00', '09:00', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
  ('bb0e8400-e29b-41d4-a716-446655440023', '990e8400-e29b-41d4-a716-446655440032', 'aa0e8400-e29b-41d4-a716-446655440066', '880e8400-e29b-41d4-a716-446655440032', 3, '09:00', '10:00', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
  ('bb0e8400-e29b-41d4-a716-446655440024', '990e8400-e29b-41d4-a716-446655440032', 'aa0e8400-e29b-41d4-a716-446655440066', '880e8400-e29b-41d4-a716-446655440033', 5, '10:00', '11:00', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP); 