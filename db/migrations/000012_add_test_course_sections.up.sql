-- Insert test course sections for different courses
-- Each course will have 2-3 sections with different teachers

-- Universidad Nacional de Colombia - Facultad de Ingeniería
-- Cálculo Diferencial (2 sections)
INSERT INTO course_sections (id, course_id, name, teacher_id, max_students, created_at, updated_at) VALUES
    ('aa0e8400-e29b-41d4-a716-446655440001', '990e8400-e29b-41d4-a716-446655440001', 'Sección A', '770e8400-e29b-41d4-a716-446655440002', 30, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('aa0e8400-e29b-41d4-a716-446655440002', '990e8400-e29b-41d4-a716-446655440001', 'Sección B', '770e8400-e29b-41d4-a716-446655440003', 25, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

-- Programación I (3 sections)
INSERT INTO course_sections (id, course_id, name, teacher_id, max_students, created_at, updated_at) VALUES
    ('aa0e8400-e29b-41d4-a716-446655440003', '990e8400-e29b-41d4-a716-446655440002', 'Sección A', '770e8400-e29b-41d4-a716-446655440004', 25, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('aa0e8400-e29b-41d4-a716-446655440004', '990e8400-e29b-41d4-a716-446655440002', 'Sección B', '770e8400-e29b-41d4-a716-446655440005', 20, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('aa0e8400-e29b-41d4-a716-446655440005', '990e8400-e29b-41d4-a716-446655440002', 'Sección C', '770e8400-e29b-41d4-a716-446655440006', 30, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

-- Física Mecánica (2 sections)
INSERT INTO course_sections (id, course_id, name, teacher_id, max_students, created_at, updated_at) VALUES
    ('aa0e8400-e29b-41d4-a716-446655440006', '990e8400-e29b-41d4-a716-446655440003', 'Sección A', '770e8400-e29b-41d4-a716-446655440007', 35, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('aa0e8400-e29b-41d4-a716-446655440007', '990e8400-e29b-41d4-a716-446655440003', 'Sección B', '770e8400-e29b-41d4-a716-446655440008', 30, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

-- Estructuras de Datos (2 sections)
INSERT INTO course_sections (id, course_id, name, teacher_id, max_students, created_at, updated_at) VALUES
    ('aa0e8400-e29b-41d4-a716-446655440008', '990e8400-e29b-41d4-a716-446655440004', 'Sección A', '770e8400-e29b-41d4-a716-446655440009', 25, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('aa0e8400-e29b-41d4-a716-446655440009', '990e8400-e29b-41d4-a716-446655440004', 'Sección B', '770e8400-e29b-41d4-a716-446655440010', 20, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

-- Bases de Datos (3 sections)
INSERT INTO course_sections (id, course_id, name, teacher_id, max_students, created_at, updated_at) VALUES
    ('aa0e8400-e29b-41d4-a716-446655440010', '990e8400-e29b-41d4-a716-446655440005', 'Sección A', '770e8400-e29b-41d4-a716-446655440002', 30, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('aa0e8400-e29b-41d4-a716-446655440011', '990e8400-e29b-41d4-a716-446655440005', 'Sección B', '770e8400-e29b-41d4-a716-446655440003', 25, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('aa0e8400-e29b-41d4-a716-446655440012', '990e8400-e29b-41d4-a716-446655440005', 'Sección C', '770e8400-e29b-41d4-a716-446655440004', 20, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

-- Redes de Computadores (2 sections)
INSERT INTO course_sections (id, course_id, name, teacher_id, max_students, created_at, updated_at) VALUES
    ('aa0e8400-e29b-41d4-a716-446655440013', '990e8400-e29b-41d4-a716-446655440006', 'Sección A', '770e8400-e29b-41d4-a716-446655440005', 25, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('aa0e8400-e29b-41d4-a716-446655440014', '990e8400-e29b-41d4-a716-446655440006', 'Sección B', '770e8400-e29b-41d4-a716-446655440006', 20, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

-- Universidad Nacional de Colombia - Facultad de Ciencias
-- Química General (2 sections)
INSERT INTO course_sections (id, course_id, name, teacher_id, max_students, created_at, updated_at) VALUES
    ('aa0e8400-e29b-41d4-a716-446655440015', '990e8400-e29b-41d4-a716-446655440007', 'Sección A', '770e8400-e29b-41d4-a716-446655440007', 30, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('aa0e8400-e29b-41d4-a716-446655440016', '990e8400-e29b-41d4-a716-446655440007', 'Sección B', '770e8400-e29b-41d4-a716-446655440008', 25, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

-- Biología Celular (2 sections)
INSERT INTO course_sections (id, course_id, name, teacher_id, max_students, created_at, updated_at) VALUES
    ('aa0e8400-e29b-41d4-a716-446655440017', '990e8400-e29b-41d4-a716-446655440008', 'Sección A', '770e8400-e29b-41d4-a716-446655440009', 20, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('aa0e8400-e29b-41d4-a716-446655440018', '990e8400-e29b-41d4-a716-446655440008', 'Sección B', '770e8400-e29b-41d4-a716-446655440010', 25, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

-- Estadística Básica (3 sections)
INSERT INTO course_sections (id, course_id, name, teacher_id, max_students, created_at, updated_at) VALUES
    ('aa0e8400-e29b-41d4-a716-446655440019', '990e8400-e29b-41d4-a716-446655440009', 'Sección A', '770e8400-e29b-41d4-a716-446655440002', 35, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('aa0e8400-e29b-41d4-a716-446655440020', '990e8400-e29b-41d4-a716-446655440009', 'Sección B', '770e8400-e29b-41d4-a716-446655440003', 30, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('aa0e8400-e29b-41d4-a716-446655440021', '990e8400-e29b-41d4-a716-446655440009', 'Sección C', '770e8400-e29b-41d4-a716-446655440004', 25, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

-- Genética Molecular (2 sections)
INSERT INTO course_sections (id, course_id, name, teacher_id, max_students, created_at, updated_at) VALUES
    ('aa0e8400-e29b-41d4-a716-446655440022', '990e8400-e29b-41d4-a716-446655440010', 'Sección A', '770e8400-e29b-41d4-a716-446655440005', 20, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('aa0e8400-e29b-41d4-a716-446655440023', '990e8400-e29b-41d4-a716-446655440010', 'Sección B', '770e8400-e29b-41d4-a716-446655440006', 25, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

-- Universidad Nacional de Colombia - Facultad de Medicina
-- Anatomía Humana (2 sections)
INSERT INTO course_sections (id, course_id, name, teacher_id, max_students, created_at, updated_at) VALUES
    ('aa0e8400-e29b-41d4-a716-446655440024', '990e8400-e29b-41d4-a716-446655440011', 'Sección A', '770e8400-e29b-41d4-a716-446655440007', 30, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('aa0e8400-e29b-41d4-a716-446655440025', '990e8400-e29b-41d4-a716-446655440011', 'Sección B', '770e8400-e29b-41d4-a716-446655440008', 25, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

-- Fisiología (2 sections)
INSERT INTO course_sections (id, course_id, name, teacher_id, max_students, created_at, updated_at) VALUES
    ('aa0e8400-e29b-41d4-a716-446655440026', '990e8400-e29b-41d4-a716-446655440012', 'Sección A', '770e8400-e29b-41d4-a716-446655440009', 25, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('aa0e8400-e29b-41d4-a716-446655440027', '990e8400-e29b-41d4-a716-446655440012', 'Sección B', '770e8400-e29b-41d4-a716-446655440010', 30, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

-- Patología General (2 sections)
INSERT INTO course_sections (id, course_id, name, teacher_id, max_students, created_at, updated_at) VALUES
    ('aa0e8400-e29b-41d4-a716-446655440028', '990e8400-e29b-41d4-a716-446655440013', 'Sección A', '770e8400-e29b-41d4-a716-446655440002', 20, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('aa0e8400-e29b-41d4-a716-446655440029', '990e8400-e29b-41d4-a716-446655440013', 'Sección B', '770e8400-e29b-41d4-a716-446655440003', 25, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

-- Farmacología (2 sections)
INSERT INTO course_sections (id, course_id, name, teacher_id, max_students, created_at, updated_at) VALUES
    ('aa0e8400-e29b-41d4-a716-446655440030', '990e8400-e29b-41d4-a716-446655440014', 'Sección A', '770e8400-e29b-41d4-a716-446655440004', 25, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('aa0e8400-e29b-41d4-a716-446655440031', '990e8400-e29b-41d4-a716-446655440014', 'Sección B', '770e8400-e29b-41d4-a716-446655440005', 20, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

-- Universidad de los Andes - Facultad de Ingeniería
-- Ingeniería de Software (2 sections)
INSERT INTO course_sections (id, course_id, name, teacher_id, max_students, created_at, updated_at) VALUES
    ('aa0e8400-e29b-41d4-a716-446655440032', '990e8400-e29b-41d4-a716-446655440015', 'Sección A', '770e8400-e29b-41d4-a716-446655440006', 25, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('aa0e8400-e29b-41d4-a716-446655440033', '990e8400-e29b-41d4-a716-446655440015', 'Sección B', '770e8400-e29b-41d4-a716-446655440007', 20, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

-- Inteligencia Artificial (2 sections)
INSERT INTO course_sections (id, course_id, name, teacher_id, max_students, created_at, updated_at) VALUES
    ('aa0e8400-e29b-41d4-a716-446655440034', '990e8400-e29b-41d4-a716-446655440016', 'Sección A', '770e8400-e29b-41d4-a716-446655440008', 20, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('aa0e8400-e29b-41d4-a716-446655440035', '990e8400-e29b-41d4-a716-446655440016', 'Sección B', '770e8400-e29b-41d4-a716-446655440009', 25, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

-- Sistemas Distribuidos (2 sections)
INSERT INTO course_sections (id, course_id, name, teacher_id, max_students, created_at, updated_at) VALUES
    ('aa0e8400-e29b-41d4-a716-446655440036', '990e8400-e29b-41d4-a716-446655440017', 'Sección A', '770e8400-e29b-41d4-a716-446655440010', 20, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('aa0e8400-e29b-41d4-a716-446655440037', '990e8400-e29b-41d4-a716-446655440017', 'Sección B', '770e8400-e29b-41d4-a716-446655440002', 15, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

-- Machine Learning (2 sections)
INSERT INTO course_sections (id, course_id, name, teacher_id, max_students, created_at, updated_at) VALUES
    ('aa0e8400-e29b-41d4-a716-446655440038', '990e8400-e29b-41d4-a716-446655440018', 'Sección A', '770e8400-e29b-41d4-a716-446655440003', 20, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('aa0e8400-e29b-41d4-a716-446655440039', '990e8400-e29b-41d4-a716-446655440018', 'Sección B', '770e8400-e29b-41d4-a716-446655440004', 25, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

-- Arquitectura de Software (2 sections)
INSERT INTO course_sections (id, course_id, name, teacher_id, max_students, created_at, updated_at) VALUES
    ('aa0e8400-e29b-41d4-a716-446655440040', '990e8400-e29b-41d4-a716-446655440019', 'Sección A', '770e8400-e29b-41d4-a716-446655440005', 20, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('aa0e8400-e29b-41d4-a716-446655440041', '990e8400-e29b-41d4-a716-446655440019', 'Sección B', '770e8400-e29b-41d4-a716-446655440006', 15, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

-- Universidad de los Andes - Facultad de Administración
-- Contabilidad Financiera (2 sections)
INSERT INTO course_sections (id, course_id, name, teacher_id, max_students, created_at, updated_at) VALUES
    ('aa0e8400-e29b-41d4-a716-446655440042', '990e8400-e29b-41d4-a716-446655440020', 'Sección A', '770e8400-e29b-41d4-a716-446655440007', 30, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('aa0e8400-e29b-41d4-a716-446655440043', '990e8400-e29b-41d4-a716-446655440020', 'Sección B', '770e8400-e29b-41d4-a716-446655440008', 25, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

-- Marketing Estratégico (2 sections)
INSERT INTO course_sections (id, course_id, name, teacher_id, max_students, created_at, updated_at) VALUES
    ('aa0e8400-e29b-41d4-a716-446655440044', '990e8400-e29b-41d4-a716-446655440021', 'Sección A', '770e8400-e29b-41d4-a716-446655440009', 25, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('aa0e8400-e29b-41d4-a716-446655440045', '990e8400-e29b-41d4-a716-446655440021', 'Sección B', '770e8400-e29b-41d4-a716-446655440010', 30, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

-- Finanzas Corporativas (2 sections)
INSERT INTO course_sections (id, course_id, name, teacher_id, max_students, created_at, updated_at) VALUES
    ('aa0e8400-e29b-41d4-a716-446655440046', '990e8400-e29b-41d4-a716-446655440022', 'Sección A', '770e8400-e29b-41d4-a716-446655440002', 25, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('aa0e8400-e29b-41d4-a716-446655440047', '990e8400-e29b-41d4-a716-446655440022', 'Sección B', '770e8400-e29b-41d4-a716-446655440003', 20, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

-- Gestión de Operaciones (2 sections)
INSERT INTO course_sections (id, course_id, name, teacher_id, max_students, created_at, updated_at) VALUES
    ('aa0e8400-e29b-41d4-a716-446655440048', '990e8400-e29b-41d4-a716-446655440023', 'Sección A', '770e8400-e29b-41d4-a716-446655440004', 30, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('aa0e8400-e29b-41d4-a716-446655440049', '990e8400-e29b-41d4-a716-446655440023', 'Sección B', '770e8400-e29b-41d4-a716-446655440005', 25, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

-- Universidad Javeriana - Facultad de Ingeniería
-- Desarrollo Web (2 sections)
INSERT INTO course_sections (id, course_id, name, teacher_id, max_students, created_at, updated_at) VALUES
    ('aa0e8400-e29b-41d4-a716-446655440050', '990e8400-e29b-41d4-a716-446655440024', 'Sección A', '770e8400-e29b-41d4-a716-446655440006', 25, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('aa0e8400-e29b-41d4-a716-446655440051', '990e8400-e29b-41d4-a716-446655440024', 'Sección B', '770e8400-e29b-41d4-a716-446655440007', 20, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

-- Seguridad Informática (2 sections)
INSERT INTO course_sections (id, course_id, name, teacher_id, max_students, created_at, updated_at) VALUES
    ('aa0e8400-e29b-41d4-a716-446655440052', '990e8400-e29b-41d4-a716-446655440025', 'Sección A', '770e8400-e29b-41d4-a716-446655440008', 20, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('aa0e8400-e29b-41d4-a716-446655440053', '990e8400-e29b-41d4-a716-446655440025', 'Sección B', '770e8400-e29b-41d4-a716-446655440009', 15, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

-- Desarrollo Móvil (2 sections)
INSERT INTO course_sections (id, course_id, name, teacher_id, max_students, created_at, updated_at) VALUES
    ('aa0e8400-e29b-41d4-a716-446655440054', '990e8400-e29b-41d4-a716-446655440026', 'Sección A', '770e8400-e29b-41d4-a716-446655440010', 20, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('aa0e8400-e29b-41d4-a716-446655440055', '990e8400-e29b-41d4-a716-446655440026', 'Sección B', '770e8400-e29b-41d4-a716-446655440002', 25, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

-- Cloud Computing (2 sections)
INSERT INTO course_sections (id, course_id, name, teacher_id, max_students, created_at, updated_at) VALUES
    ('aa0e8400-e29b-41d4-a716-446655440056', '990e8400-e29b-41d4-a716-446655440027', 'Sección A', '770e8400-e29b-41d4-a716-446655440003', 20, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('aa0e8400-e29b-41d4-a716-446655440057', '990e8400-e29b-41d4-a716-446655440027', 'Sección B', '770e8400-e29b-41d4-a716-446655440004', 15, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

-- Universidad de Antioquia - Facultad de Ingeniería
-- Sistemas Embebidos (2 sections)
INSERT INTO course_sections (id, course_id, name, teacher_id, max_students, created_at, updated_at) VALUES
    ('aa0e8400-e29b-41d4-a716-446655440058', '990e8400-e29b-41d4-a716-446655440028', 'Sección A', '770e8400-e29b-41d4-a716-446655440005', 20, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('aa0e8400-e29b-41d4-a716-446655440059', '990e8400-e29b-41d4-a716-446655440028', 'Sección B', '770e8400-e29b-41d4-a716-446655440006', 15, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

-- Robótica (2 sections)
INSERT INTO course_sections (id, course_id, name, teacher_id, max_students, created_at, updated_at) VALUES
    ('aa0e8400-e29b-41d4-a716-446655440060', '990e8400-e29b-41d4-a716-446655440029', 'Sección A', '770e8400-e29b-41d4-a716-446655440007', 15, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('aa0e8400-e29b-41d4-a716-446655440061', '990e8400-e29b-41d4-a716-446655440029', 'Sección B', '770e8400-e29b-41d4-a716-446655440008', 20, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

-- Control Automático (2 sections)
INSERT INTO course_sections (id, course_id, name, teacher_id, max_students, created_at, updated_at) VALUES
    ('aa0e8400-e29b-41d4-a716-446655440062', '990e8400-e29b-41d4-a716-446655440030', 'Sección A', '770e8400-e29b-41d4-a716-446655440009', 20, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('aa0e8400-e29b-41d4-a716-446655440063', '990e8400-e29b-41d4-a716-446655440030', 'Sección B', '770e8400-e29b-41d4-a716-446655440010', 15, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

-- Procesamiento de Señales (2 sections)
INSERT INTO course_sections (id, course_id, name, teacher_id, max_students, created_at, updated_at) VALUES
    ('aa0e8400-e29b-41d4-a716-446655440064', '990e8400-e29b-41d4-a716-446655440031', 'Sección A', '770e8400-e29b-41d4-a716-446655440002', 15, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('aa0e8400-e29b-41d4-a716-446655440065', '990e8400-e29b-41d4-a716-446655440031', 'Sección B', '770e8400-e29b-41d4-a716-446655440003', 20, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

-- Universidad del Valle - Facultad de Ingeniería
-- Ingeniería de Sistemas (2 sections)
INSERT INTO course_sections (id, course_id, name, teacher_id, max_students, created_at, updated_at) VALUES
    ('aa0e8400-e29b-41d4-a716-446655440066', '990e8400-e29b-41d4-a716-446655440032', 'Sección A', '770e8400-e29b-41d4-a716-446655440004', 30, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('aa0e8400-e29b-41d4-a716-446655440067', '990e8400-e29b-41d4-a716-446655440032', 'Sección B', '770e8400-e29b-41d4-a716-446655440005', 25, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

-- Teoría de la Computación (2 sections)
INSERT INTO course_sections (id, course_id, name, teacher_id, max_students, created_at, updated_at) VALUES
    ('aa0e8400-e29b-41d4-a716-446655440068', '990e8400-e29b-41d4-a716-446655440033', 'Sección A', '770e8400-e29b-41d4-a716-446655440006', 20, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('aa0e8400-e29b-41d4-a716-446655440069', '990e8400-e29b-41d4-a716-446655440033', 'Sección B', '770e8400-e29b-41d4-a716-446655440007', 25, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

-- Compiladores (2 sections)
INSERT INTO course_sections (id, course_id, name, teacher_id, max_students, created_at, updated_at) VALUES
    ('aa0e8400-e29b-41d4-a716-446655440070', '990e8400-e29b-41d4-a716-446655440034', 'Sección A', '770e8400-e29b-41d4-a716-446655440008', 20, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('aa0e8400-e29b-41d4-a716-446655440071', '990e8400-e29b-41d4-a716-446655440034', 'Sección B', '770e8400-e29b-41d4-a716-446655440009', 15, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

-- Sistemas Operativos (2 sections)
INSERT INTO course_sections (id, course_id, name, teacher_id, max_students, created_at, updated_at) VALUES
    ('aa0e8400-e29b-41d4-a716-446655440072', '990e8400-e29b-41d4-a716-446655440035', 'Sección A', '770e8400-e29b-41d4-a716-446655440010', 25, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('aa0e8400-e29b-41d4-a716-446655440073', '990e8400-e29b-41d4-a716-446655440035', 'Sección B', '770e8400-e29b-41d4-a716-446655440002', 20, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP); 