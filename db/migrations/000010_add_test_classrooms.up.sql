-- Insert test classrooms for different faculties
-- Universidad Nacional de Colombia - Facultad de Ingeniería
INSERT INTO classrooms (id, name, building, floor, capacity, faculty_id, location_lat, location_lng, created_at, updated_at) VALUES
    ('880e8400-e29b-41d4-a716-446655440001', 'A-101', 'Edificio de Ingeniería', 1, 30, '660e8400-e29b-41d4-a716-446655440001', 4.6379, -74.0847, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('880e8400-e29b-41d4-a716-446655440002', 'A-102', 'Edificio de Ingeniería', 1, 25, '660e8400-e29b-41d4-a716-446655440001', 4.6379, -74.0847, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('880e8400-e29b-41d4-a716-446655440003', 'A-201', 'Edificio de Ingeniería', 2, 40, '660e8400-e29b-41d4-a716-446655440001', 4.6379, -74.0847, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('880e8400-e29b-41d4-a716-446655440004', 'A-202', 'Edificio de Ingeniería', 2, 35, '660e8400-e29b-41d4-a716-446655440001', 4.6379, -74.0847, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('880e8400-e29b-41d4-a716-446655440005', 'B-101', 'Edificio de Sistemas', 1, 50, '660e8400-e29b-41d4-a716-446655440001', 4.6379, -74.0847, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('880e8400-e29b-41d4-a716-446655440006', 'B-102', 'Edificio de Sistemas', 1, 45, '660e8400-e29b-41d4-a716-446655440001', 4.6379, -74.0847, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('880e8400-e29b-41d4-a716-446655440007', 'C-101', 'Auditorio Principal', 1, 120, '660e8400-e29b-41d4-a716-446655440001', 4.6379, -74.0847, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

-- Universidad Nacional de Colombia - Facultad de Ciencias
INSERT INTO classrooms (id, name, building, floor, capacity, faculty_id, location_lat, location_lng, created_at, updated_at) VALUES
    ('880e8400-e29b-41d4-a716-446655440008', 'Lab-101', 'Laboratorio de Física', 1, 20, '660e8400-e29b-41d4-a716-446655440002', 4.6379, -74.0847, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('880e8400-e29b-41d4-a716-446655440009', 'Lab-102', 'Laboratorio de Química', 1, 25, '660e8400-e29b-41d4-a716-446655440002', 4.6379, -74.0847, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('880e8400-e29b-41d4-a716-446655440010', 'D-201', 'Edificio de Ciencias', 2, 35, '660e8400-e29b-41d4-a716-446655440002', 4.6379, -74.0847, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('880e8400-e29b-41d4-a716-446655440011', 'D-202', 'Edificio de Ciencias', 2, 30, '660e8400-e29b-41d4-a716-446655440002', 4.6379, -74.0847, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

-- Universidad Nacional de Colombia - Facultad de Medicina
INSERT INTO classrooms (id, name, building, floor, capacity, faculty_id, location_lat, location_lng, created_at, updated_at) VALUES
    ('880e8400-e29b-41d4-a716-446655440012', 'Med-101', 'Hospital Universitario', 1, 40, '660e8400-e29b-41d4-a716-446655440003', 4.6379, -74.0847, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('880e8400-e29b-41d4-a716-446655440013', 'Med-102', 'Hospital Universitario', 1, 35, '660e8400-e29b-41d4-a716-446655440003', 4.6379, -74.0847, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('880e8400-e29b-41d4-a716-446655440014', 'E-201', 'Edificio de Medicina', 2, 50, '660e8400-e29b-41d4-a716-446655440003', 4.6379, -74.0847, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

-- Universidad de los Andes - Facultad de Ingeniería
INSERT INTO classrooms (id, name, building, floor, capacity, faculty_id, location_lat, location_lng, created_at, updated_at) VALUES
    ('880e8400-e29b-41d4-a716-446655440015', 'ML-101', 'Mario Laserna', 1, 30, '660e8400-e29b-41d4-a716-446655440005', 4.6014, -74.0661, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('880e8400-e29b-41d4-a716-446655440016', 'ML-102', 'Mario Laserna', 1, 25, '660e8400-e29b-41d4-a716-446655440005', 4.6014, -74.0661, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('880e8400-e29b-41d4-a716-446655440017', 'ML-201', 'Mario Laserna', 2, 40, '660e8400-e29b-41d4-a716-446655440005', 4.6014, -74.0661, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('880e8400-e29b-41d4-a716-446655440018', 'ML-202', 'Mario Laserna', 2, 35, '660e8400-e29b-41d4-a716-446655440005', 4.6014, -74.0661, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('880e8400-e29b-41d4-a716-446655440019', 'Aud-101', 'Auditorio Lleras', 1, 150, '660e8400-e29b-41d4-a716-446655440005', 4.6014, -74.0661, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

-- Universidad de los Andes - Facultad de Administración
INSERT INTO classrooms (id, name, building, floor, capacity, faculty_id, location_lat, location_lng, created_at, updated_at) VALUES
    ('880e8400-e29b-41d4-a716-446655440020', 'AD-101', 'Edificio de Administración', 1, 45, '660e8400-e29b-41d4-a716-446655440006', 4.6014, -74.0661, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('880e8400-e29b-41d4-a716-446655440021', 'AD-102', 'Edificio de Administración', 1, 40, '660e8400-e29b-41d4-a716-446655440006', 4.6014, -74.0661, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('880e8400-e29b-41d4-a716-446655440022', 'AD-201', 'Edificio de Administración', 2, 55, '660e8400-e29b-41d4-a716-446655440006', 4.6014, -74.0661, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

-- Universidad Javeriana - Facultad de Ingeniería
INSERT INTO classrooms (id, name, building, floor, capacity, faculty_id, location_lat, location_lng, created_at, updated_at) VALUES
    ('880e8400-e29b-41d4-a716-446655440023', 'Jav-101', 'Edificio José Gabriel Maldonado', 1, 30, '660e8400-e29b-41d4-a716-446655440009', 4.6286, -74.0647, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('880e8400-e29b-41d4-a716-446655440024', 'Jav-102', 'Edificio José Gabriel Maldonado', 1, 25, '660e8400-e29b-41d4-a716-446655440009', 4.6286, -74.0647, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('880e8400-e29b-41d4-a716-446655440025', 'Jav-201', 'Edificio José Gabriel Maldonado', 2, 40, '660e8400-e29b-41d4-a716-446655440009', 4.6286, -74.0647, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('880e8400-e29b-41d4-a716-446655440026', 'Lab-Jav-101', 'Laboratorio de Ingeniería', 1, 20, '660e8400-e29b-41d4-a716-446655440009', 4.6286, -74.0647, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

-- Universidad de Antioquia - Facultad de Ingeniería
INSERT INTO classrooms (id, name, building, floor, capacity, faculty_id, location_lat, location_lng, created_at, updated_at) VALUES
    ('880e8400-e29b-41d4-a716-446655440027', 'Ant-101', 'Bloque 21', 1, 35, '660e8400-e29b-41d4-a716-446655440013', 6.2442, -75.5812, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('880e8400-e29b-41d4-a716-446655440028', 'Ant-102', 'Bloque 21', 1, 30, '660e8400-e29b-41d4-a716-446655440013', 6.2442, -75.5812, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('880e8400-e29b-41d4-a716-446655440029', 'Ant-201', 'Bloque 21', 2, 45, '660e8400-e29b-41d4-a716-446655440013', 6.2442, -75.5812, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('880e8400-e29b-41d4-a716-446655440030', 'Aud-Ant-101', 'Auditorio Principal', 1, 200, '660e8400-e29b-41d4-a716-446655440013', 6.2442, -75.5812, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

-- Universidad del Valle - Facultad de Ingeniería
INSERT INTO classrooms (id, name, building, floor, capacity, faculty_id, location_lat, location_lng, created_at, updated_at) VALUES
    ('880e8400-e29b-41d4-a716-446655440031', 'Val-101', 'Edificio 350', 1, 40, '660e8400-e29b-41d4-a716-446655440017', 3.4516, -76.5320, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('880e8400-e29b-41d4-a716-446655440032', 'Val-102', 'Edificio 350', 1, 35, '660e8400-e29b-41d4-a716-446655440017', 3.4516, -76.5320, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('880e8400-e29b-41d4-a716-446655440033', 'Val-201', 'Edificio 350', 2, 50, '660e8400-e29b-41d4-a716-446655440017', 3.4516, -76.5320, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('880e8400-e29b-41d4-a716-446655440034', 'Lab-Val-101', 'Laboratorio de Sistemas', 1, 25, '660e8400-e29b-41d4-a716-446655440017', 3.4516, -76.5320, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP); 