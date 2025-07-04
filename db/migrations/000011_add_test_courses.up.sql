-- Insert test courses for different faculties
-- Universidad Nacional de Colombia - Facultad de Ingeniería
INSERT INTO courses (id, faculty_id, name, semester, created_at, updated_at) VALUES
    ('990e8400-e29b-41d4-a716-446655440001', '660e8400-e29b-41d4-a716-446655440001', 'Cálculo Diferencial', '1', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('990e8400-e29b-41d4-a716-446655440002', '660e8400-e29b-41d4-a716-446655440001', 'Programación I', '1', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('990e8400-e29b-41d4-a716-446655440003', '660e8400-e29b-41d4-a716-446655440001', 'Física Mecánica', '2', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('990e8400-e29b-41d4-a716-446655440004', '660e8400-e29b-41d4-a716-446655440001', 'Estructuras de Datos', '3', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('990e8400-e29b-41d4-a716-446655440005', '660e8400-e29b-41d4-a716-446655440001', 'Bases de Datos', '4', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('990e8400-e29b-41d4-a716-446655440006', '660e8400-e29b-41d4-a716-446655440001', 'Redes de Computadores', '5', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

-- Universidad Nacional de Colombia - Facultad de Ciencias
INSERT INTO courses (id, faculty_id, name, semester, created_at, updated_at) VALUES
    ('990e8400-e29b-41d4-a716-446655440007', '660e8400-e29b-41d4-a716-446655440002', 'Química General', '1', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('990e8400-e29b-41d4-a716-446655440008', '660e8400-e29b-41d4-a716-446655440002', 'Biología Celular', '2', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('990e8400-e29b-41d4-a716-446655440009', '660e8400-e29b-41d4-a716-446655440002', 'Estadística Básica', '3', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('990e8400-e29b-41d4-a716-446655440010', '660e8400-e29b-41d4-a716-446655440002', 'Genética Molecular', '4', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

-- Universidad Nacional de Colombia - Facultad de Medicina
INSERT INTO courses (id, faculty_id, name, semester, created_at, updated_at) VALUES
    ('990e8400-e29b-41d4-a716-446655440011', '660e8400-e29b-41d4-a716-446655440003', 'Anatomía Humana', '1', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('990e8400-e29b-41d4-a716-446655440012', '660e8400-e29b-41d4-a716-446655440003', 'Fisiología', '2', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('990e8400-e29b-41d4-a716-446655440013', '660e8400-e29b-41d4-a716-446655440003', 'Patología General', '3', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('990e8400-e29b-41d4-a716-446655440014', '660e8400-e29b-41d4-a716-446655440003', 'Farmacología', '4', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

-- Universidad de los Andes - Facultad de Ingeniería
INSERT INTO courses (id, faculty_id, name, semester, created_at, updated_at) VALUES
    ('990e8400-e29b-41d4-a716-446655440015', '660e8400-e29b-41d4-a716-446655440005', 'Ingeniería de Software', '5', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('990e8400-e29b-41d4-a716-446655440016', '660e8400-e29b-41d4-a716-446655440005', 'Inteligencia Artificial', '6', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('990e8400-e29b-41d4-a716-446655440017', '660e8400-e29b-41d4-a716-446655440005', 'Sistemas Distribuidos', '7', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('990e8400-e29b-41d4-a716-446655440018', '660e8400-e29b-41d4-a716-446655440005', 'Machine Learning', '8', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('990e8400-e29b-41d4-a716-446655440019', '660e8400-e29b-41d4-a716-446655440005', 'Arquitectura de Software', '9', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

-- Universidad de los Andes - Facultad de Administración
INSERT INTO courses (id, faculty_id, name, semester, created_at, updated_at) VALUES
    ('990e8400-e29b-41d4-a716-446655440020', '660e8400-e29b-41d4-a716-446655440006', 'Contabilidad Financiera', '1', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('990e8400-e29b-41d4-a716-446655440021', '660e8400-e29b-41d4-a716-446655440006', 'Marketing Estratégico', '2', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('990e8400-e29b-41d4-a716-446655440022', '660e8400-e29b-41d4-a716-446655440006', 'Finanzas Corporativas', '3', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('990e8400-e29b-41d4-a716-446655440023', '660e8400-e29b-41d4-a716-446655440006', 'Gestión de Operaciones', '4', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

-- Universidad Javeriana - Facultad de Ingeniería
INSERT INTO courses (id, faculty_id, name, semester, created_at, updated_at) VALUES
    ('990e8400-e29b-41d4-a716-446655440024', '660e8400-e29b-41d4-a716-446655440009', 'Desarrollo Web', '5', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('990e8400-e29b-41d4-a716-446655440025', '660e8400-e29b-41d4-a716-446655440009', 'Seguridad Informática', '6', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('990e8400-e29b-41d4-a716-446655440026', '660e8400-e29b-41d4-a716-446655440009', 'Desarrollo Móvil', '7', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('990e8400-e29b-41d4-a716-446655440027', '660e8400-e29b-41d4-a716-446655440009', 'Cloud Computing', '8', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

-- Universidad de Antioquia - Facultad de Ingeniería
INSERT INTO courses (id, faculty_id, name, semester, created_at, updated_at) VALUES
    ('990e8400-e29b-41d4-a716-446655440028', '660e8400-e29b-41d4-a716-446655440013', 'Sistemas Embebidos', '7', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('990e8400-e29b-41d4-a716-446655440029', '660e8400-e29b-41d4-a716-446655440013', 'Robótica', '8', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('990e8400-e29b-41d4-a716-446655440030', '660e8400-e29b-41d4-a716-446655440013', 'Control Automático', '9', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('990e8400-e29b-41d4-a716-446655440031', '660e8400-e29b-41d4-a716-446655440013', 'Procesamiento de Señales', '10', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

-- Universidad del Valle - Facultad de Ingeniería
INSERT INTO courses (id, faculty_id, name, semester, created_at, updated_at) VALUES
    ('990e8400-e29b-41d4-a716-446655440032', '660e8400-e29b-41d4-a716-446655440017', 'Ingeniería de Sistemas', '1', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('990e8400-e29b-41d4-a716-446655440033', '660e8400-e29b-41d4-a716-446655440017', 'Teoría de la Computación', '2', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('990e8400-e29b-41d4-a716-446655440034', '660e8400-e29b-41d4-a716-446655440017', 'Compiladores', '3', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('990e8400-e29b-41d4-a716-446655440035', '660e8400-e29b-41d4-a716-446655440017', 'Sistemas Operativos', '4', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP); 