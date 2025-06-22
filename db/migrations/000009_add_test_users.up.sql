-- Insert test users with different roles
-- Password hash for all users: "password123" (bcrypt hash)
-- Admin user
INSERT INTO users (id, email, password_hash, first_name, last_name, role, created_at, updated_at) VALUES
    ('770e8400-e29b-41d4-a716-446655440001', '1@admin', '$2a$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi', 'Admin', 'System', 'admin', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

-- Teacher users (10 teachers)
INSERT INTO users (id, email, password_hash, first_name, last_name, role, created_at, updated_at) VALUES
    ('770e8400-e29b-41d4-a716-446655440002', '1@teacher', '$2a$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi', 'María', 'González', 'teacher', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('770e8400-e29b-41d4-a716-446655440003', '2@teacher', '$2a$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi', 'Carlos', 'Rodríguez', 'teacher', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('770e8400-e29b-41d4-a716-446655440004', '3@teacher', '$2a$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi', 'Ana', 'Martínez', 'teacher', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('770e8400-e29b-41d4-a716-446655440005', '4@teacher', '$2a$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi', 'Luis', 'Hernández', 'teacher', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('770e8400-e29b-41d4-a716-446655440006', '5@teacher', '$2a$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi', 'Patricia', 'López', 'teacher', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('770e8400-e29b-41d4-a716-446655440007', '6@teacher', '$2a$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi', 'Roberto', 'García', 'teacher', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('770e8400-e29b-41d4-a716-446655440008', '7@teacher', '$2a$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi', 'Carmen', 'Pérez', 'teacher', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('770e8400-e29b-41d4-a716-446655440009', '8@teacher', '$2a$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi', 'Fernando', 'Sánchez', 'teacher', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('770e8400-e29b-41d4-a716-446655440010', '9@teacher', '$2a$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi', 'Isabel', 'Ramírez', 'teacher', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('770e8400-e29b-41d4-a716-446655440011', '10@teacher', '$2a$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi', 'Miguel', 'Torres', 'teacher', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

-- Student users (20 students)
INSERT INTO users (id, email, password_hash, first_name, last_name, role, created_at, updated_at) VALUES
    ('770e8400-e29b-41d4-a716-446655440012', '1@student', '$2a$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi', 'Juan', 'Díaz', 'student', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('770e8400-e29b-41d4-a716-446655440013', '2@student', '$2a$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi', 'Sofia', 'Moreno', 'student', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('770e8400-e29b-41d4-a716-446655440014', '3@student', '$2a$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi', 'Andrés', 'Jiménez', 'student', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('770e8400-e29b-41d4-a716-446655440015', '4@student', '$2a$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi', 'Valentina', 'Ruiz', 'student', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('770e8400-e29b-41d4-a716-446655440016', '5@student', '$2a$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi', 'Diego', 'Flores', 'student', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('770e8400-e29b-41d4-a716-446655440017', '6@student', '$2a$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi', 'Camila', 'Vargas', 'student', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('770e8400-e29b-41d4-a716-446655440018', '7@student', '$2a$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi', 'Sebastián', 'Reyes', 'student', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('770e8400-e29b-41d4-a716-446655440019', '8@student', '$2a$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi', 'Mariana', 'Castro', 'student', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('770e8400-e29b-41d4-a716-446655440020', '9@student', '$2a$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi', 'Nicolás', 'Morales', 'student', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('770e8400-e29b-41d4-a716-446655440021', '10@student', '$2a$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi', 'Gabriela', 'Ortiz', 'student', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('770e8400-e29b-41d4-a716-446655440022', '11@student', '$2a$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi', 'Daniel', 'Silva', 'student', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('770e8400-e29b-41d4-a716-446655440023', '12@student', '$2a$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi', 'Laura', 'Cruz', 'student', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('770e8400-e29b-41d4-a716-446655440024', '13@student', '$2a$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi', 'Felipe', 'Mendoza', 'student', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('770e8400-e29b-41d4-a716-446655440025', '14@student', '$2a$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi', 'Alejandra', 'Herrera', 'student', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('770e8400-e29b-41d4-a716-446655440026', '15@student', '$2a$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi', 'Ricardo', 'Rojas', 'student', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('770e8400-e29b-41d4-a716-446655440027', '16@student', '$2a$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi', 'Natalia', 'Acosta', 'student', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('770e8400-e29b-41d4-a716-446655440028', '17@student', '$2a$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi', 'Carlos', 'Medina', 'student', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('770e8400-e29b-41d4-a716-446655440029', '18@student', '$2a$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi', 'Paula', 'Guerrero', 'student', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('770e8400-e29b-41d4-a716-446655440030', '19@student', '$2a$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi', 'Alejandro', 'Romero', 'student', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('770e8400-e29b-41d4-a716-446655440031', '20@student', '$2a$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi', 'Isabella', 'Suárez', 'student', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP); 