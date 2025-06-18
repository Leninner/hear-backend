-- Insert test universities
INSERT INTO universities (id, name, created_at, updated_at) VALUES
    ('550e8400-e29b-41d4-a716-446655440001', 'Universidad Nacional de Colombia', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('550e8400-e29b-41d4-a716-446655440002', 'Universidad de los Andes', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('550e8400-e29b-41d4-a716-446655440003', 'Universidad Javeriana', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('550e8400-e29b-41d4-a716-446655440004', 'Universidad de Antioquia', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('550e8400-e29b-41d4-a716-446655440005', 'Universidad del Valle', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

-- Insert test faculties for Universidad Nacional de Colombia
INSERT INTO faculties (id, university_id, name, location_lat, location_lng, created_at, updated_at) VALUES
    ('660e8400-e29b-41d4-a716-446655440001', '550e8400-e29b-41d4-a716-446655440001', 'Facultad de Ingeniería', 4.6379, -74.0847, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('660e8400-e29b-41d4-a716-446655440002', '550e8400-e29b-41d4-a716-446655440001', 'Facultad de Ciencias', 4.6379, -74.0847, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('660e8400-e29b-41d4-a716-446655440003', '550e8400-e29b-41d4-a716-446655440001', 'Facultad de Medicina', 4.6379, -74.0847, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('660e8400-e29b-41d4-a716-446655440004', '550e8400-e29b-41d4-a716-446655440001', 'Facultad de Ciencias Económicas', 4.6379, -74.0847, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

-- Insert test faculties for Universidad de los Andes
INSERT INTO faculties (id, university_id, name, location_lat, location_lng, created_at, updated_at) VALUES
    ('660e8400-e29b-41d4-a716-446655440005', '550e8400-e29b-41d4-a716-446655440002', 'Facultad de Ingeniería', 4.6014, -74.0661, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('660e8400-e29b-41d4-a716-446655440006', '550e8400-e29b-41d4-a716-446655440002', 'Facultad de Administración', 4.6014, -74.0661, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('660e8400-e29b-41d4-a716-446655440007', '550e8400-e29b-41d4-a716-446655440002', 'Facultad de Derecho', 4.6014, -74.0661, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('660e8400-e29b-41d4-a716-446655440008', '550e8400-e29b-41d4-a716-446655440002', 'Facultad de Medicina', 4.6014, -74.0661, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

-- Insert test faculties for Universidad Javeriana
INSERT INTO faculties (id, university_id, name, location_lat, location_lng, created_at, updated_at) VALUES
    ('660e8400-e29b-41d4-a716-446655440009', '550e8400-e29b-41d4-a716-446655440003', 'Facultad de Ingeniería', 4.6286, -74.0647, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('660e8400-e29b-41d4-a716-446655440010', '550e8400-e29b-41d4-a716-446655440003', 'Facultad de Ciencias Económicas y Administrativas', 4.6286, -74.0647, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('660e8400-e29b-41d4-a716-446655440011', '550e8400-e29b-41d4-a716-446655440003', 'Facultad de Medicina', 4.6286, -74.0647, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('660e8400-e29b-41d4-a716-446655440012', '550e8400-e29b-41d4-a716-446655440003', 'Facultad de Ciencias', 4.6286, -74.0647, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

-- Insert test faculties for Universidad de Antioquia
INSERT INTO faculties (id, university_id, name, location_lat, location_lng, created_at, updated_at) VALUES
    ('660e8400-e29b-41d4-a716-446655440013', '550e8400-e29b-41d4-a716-446655440004', 'Facultad de Ingeniería', 6.2442, -75.5812, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('660e8400-e29b-41d4-a716-446655440014', '550e8400-e29b-41d4-a716-446655440004', 'Facultad de Ciencias Exactas y Naturales', 6.2442, -75.5812, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('660e8400-e29b-41d4-a716-446655440015', '550e8400-e29b-41d4-a716-446655440004', 'Facultad de Medicina', 6.2442, -75.5812, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('660e8400-e29b-41d4-a716-446655440016', '550e8400-e29b-41d4-a716-446655440004', 'Facultad de Ciencias Económicas', 6.2442, -75.5812, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

-- Insert test faculties for Universidad del Valle
INSERT INTO faculties (id, university_id, name, location_lat, location_lng, created_at, updated_at) VALUES
    ('660e8400-e29b-41d4-a716-446655440017', '550e8400-e29b-41d4-a716-446655440005', 'Facultad de Ingeniería', 3.4516, -76.5320, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('660e8400-e29b-41d4-a716-446655440018', '550e8400-e29b-41d4-a716-446655440005', 'Facultad de Ciencias Naturales y Exactas', 3.4516, -76.5320, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('660e8400-e29b-41d4-a716-446655440019', '550e8400-e29b-41d4-a716-446655440005', 'Facultad de Ciencias de la Administración', 3.4516, -76.5320, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('660e8400-e29b-41d4-a716-446655440020', '550e8400-e29b-41d4-a716-446655440005', 'Facultad de Salud', 3.4516, -76.5320, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP); 