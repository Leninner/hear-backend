-- Insert test users with different roles
-- Password hash for all users: "password123" (bcrypt hash)
-- Admin user
INSERT INTO users (id, email, password_hash, first_name, last_name, role, created_at, updated_at) VALUES
    ('770e8400-e29b-41d4-a716-446655440001', '1@admin.com', '$2a$10$DtCQWRxGmo/9ctlhCJt7rud/5pkooLtxbzW.Cwys3/jRz.T5kfNoy', 'Admin', 'System', 'admin', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

-- Teacher users (10 teachers)
INSERT INTO users (id, email, password_hash, first_name, last_name, role, created_at, updated_at) VALUES
    ('770e8400-e29b-41d4-a716-446655440002', '1@teacher.com', '$2a$10$DtCQWRxGmo/9ctlhCJt7rud/5pkooLtxbzW.Cwys3/jRz.T5kfNoy', 'María', 'González', 'teacher', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('770e8400-e29b-41d4-a716-446655440003', '2@teacher.com', '$2a$10$DtCQWRxGmo/9ctlhCJt7rud/5pkooLtxbzW.Cwys3/jRz.T5kfNoy', 'Carlos', 'Rodríguez', 'teacher', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('770e8400-e29b-41d4-a716-446655440004', '3@teacher.com', '$2a$10$DtCQWRxGmo/9ctlhCJt7rud/5pkooLtxbzW.Cwys3/jRz.T5kfNoy', 'Ana', 'Martínez', 'teacher', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('770e8400-e29b-41d4-a716-446655440005', '4@teacher.com', '$2a$10$DtCQWRxGmo/9ctlhCJt7rud/5pkooLtxbzW.Cwys3/jRz.T5kfNoy', 'Luis', 'Hernández', 'teacher', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('770e8400-e29b-41d4-a716-446655440006', '5@teacher.com', '$2a$10$DtCQWRxGmo/9ctlhCJt7rud/5pkooLtxbzW.Cwys3/jRz.T5kfNoy', 'Patricia', 'López', 'teacher', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('770e8400-e29b-41d4-a716-446655440007', '6@teacher.com', '$2a$10$DtCQWRxGmo/9ctlhCJt7rud/5pkooLtxbzW.Cwys3/jRz.T5kfNoy', 'Roberto', 'García', 'teacher', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('770e8400-e29b-41d4-a716-446655440008', '7@teacher.com', '$2a$10$DtCQWRxGmo/9ctlhCJt7rud/5pkooLtxbzW.Cwys3/jRz.T5kfNoy', 'Carmen', 'Pérez', 'teacher', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('770e8400-e29b-41d4-a716-446655440009', '8@teacher.com', '$2a$10$DtCQWRxGmo/9ctlhCJt7rud/5pkooLtxbzW.Cwys3/jRz.T5kfNoy', 'Fernando', 'Sánchez', 'teacher', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('770e8400-e29b-41d4-a716-446655440010', '9@teacher.com', '$2a$10$DtCQWRxGmo/9ctlhCJt7rud/5pkooLtxbzW.Cwys3/jRz.T5kfNoy', 'Isabel', 'Ramírez', 'teacher', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('770e8400-e29b-41d4-a716-446655440011', '10@teacher.com', '$2a$10$DtCQWRxGmo/9ctlhCJt7rud/5pkooLtxbzW.Cwys3/jRz.T5kfNoy', 'Miguel', 'Torres', 'teacher', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

-- Student users (20 students)
INSERT INTO users (id, email, password_hash, first_name, last_name, role, created_at, updated_at) VALUES
    ('770e8400-e29b-41d4-a716-446655440012', '1@student.com', '$2a$10$DtCQWRxGmo/9ctlhCJt7rud/5pkooLtxbzW.Cwys3/jRz.T5kfNoy', 'Juan', 'Díaz', 'student', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('770e8400-e29b-41d4-a716-446655440013', '2@student.com', '$2a$10$DtCQWRxGmo/9ctlhCJt7rud/5pkooLtxbzW.Cwys3/jRz.T5kfNoy', 'Sofia', 'Moreno', 'student', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('770e8400-e29b-41d4-a716-446655440014', '3@student.com', '$2a$10$DtCQWRxGmo/9ctlhCJt7rud/5pkooLtxbzW.Cwys3/jRz.T5kfNoy', 'Andrés', 'Jiménez', 'student', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('770e8400-e29b-41d4-a716-446655440015', '4@student.com', '$2a$10$DtCQWRxGmo/9ctlhCJt7rud/5pkooLtxbzW.Cwys3/jRz.T5kfNoy', 'Valentina', 'Ruiz', 'student', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('770e8400-e29b-41d4-a716-446655440016', '5@student.com', '$2a$10$DtCQWRxGmo/9ctlhCJt7rud/5pkooLtxbzW.Cwys3/jRz.T5kfNoy', 'Diego', 'Flores', 'student', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('770e8400-e29b-41d4-a716-446655440017', '6@student.com', '$2a$10$DtCQWRxGmo/9ctlhCJt7rud/5pkooLtxbzW.Cwys3/jRz.T5kfNoy', 'Camila', 'Vargas', 'student', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('770e8400-e29b-41d4-a716-446655440018', '7@student.com', '$2a$10$DtCQWRxGmo/9ctlhCJt7rud/5pkooLtxbzW.Cwys3/jRz.T5kfNoy', 'Sebastián', 'Reyes', 'student', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('770e8400-e29b-41d4-a716-446655440019', '8@student.com', '$2a$10$DtCQWRxGmo/9ctlhCJt7rud/5pkooLtxbzW.Cwys3/jRz.T5kfNoy', 'Mariana', 'Castro', 'student', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('770e8400-e29b-41d4-a716-446655440020', '9@student.com', '$2a$10$DtCQWRxGmo/9ctlhCJt7rud/5pkooLtxbzW.Cwys3/jRz.T5kfNoy', 'Nicolás', 'Morales', 'student', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('770e8400-e29b-41d4-a716-446655440021', '10@student.com', '$2a$10$DtCQWRxGmo/9ctlhCJt7rud/5pkooLtxbzW.Cwys3/jRz.T5kfNoy', 'Gabriela', 'Ortiz', 'student', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('770e8400-e29b-41d4-a716-446655440022', '11@student.com', '$2a$10$DtCQWRxGmo/9ctlhCJt7rud/5pkooLtxbzW.Cwys3/jRz.T5kfNoy', 'Daniel', 'Silva', 'student', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('770e8400-e29b-41d4-a716-446655440023', '12@student.com', '$2a$10$DtCQWRxGmo/9ctlhCJt7rud/5pkooLtxbzW.Cwys3/jRz.T5kfNoy', 'Laura', 'Cruz', 'student', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('770e8400-e29b-41d4-a716-446655440024', '13@student.com', '$2a$10$DtCQWRxGmo/9ctlhCJt7rud/5pkooLtxbzW.Cwys3/jRz.T5kfNoy', 'Felipe', 'Mendoza', 'student', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('770e8400-e29b-41d4-a716-446655440025', '14@student.com', '$2a$10$DtCQWRxGmo/9ctlhCJt7rud/5pkooLtxbzW.Cwys3/jRz.T5kfNoy', 'Alejandra', 'Herrera', 'student', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('770e8400-e29b-41d4-a716-446655440026', '15@student.com', '$2a$10$DtCQWRxGmo/9ctlhCJt7rud/5pkooLtxbzW.Cwys3/jRz.T5kfNoy', 'Ricardo', 'Rojas', 'student', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('770e8400-e29b-41d4-a716-446655440027', '16@student.com', '$2a$10$DtCQWRxGmo/9ctlhCJt7rud/5pkooLtxbzW.Cwys3/jRz.T5kfNoy', 'Natalia', 'Acosta', 'student', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('770e8400-e29b-41d4-a716-446655440028', '17@student.com', '$2a$10$DtCQWRxGmo/9ctlhCJt7rud/5pkooLtxbzW.Cwys3/jRz.T5kfNoy', 'Carlos', 'Medina', 'student', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('770e8400-e29b-41d4-a716-446655440029', '18@student.com', '$2a$10$DtCQWRxGmo/9ctlhCJt7rud/5pkooLtxbzW.Cwys3/jRz.T5kfNoy', 'Paula', 'Guerrero', 'student', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('770e8400-e29b-41d4-a716-446655440030', '19@student.com', '$2a$10$DtCQWRxGmo/9ctlhCJt7rud/5pkooLtxbzW.Cwys3/jRz.T5kfNoy', 'Alejandro', 'Romero', 'student', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('770e8400-e29b-41d4-a716-446655440031', '20@student.com', '$2a$10$DtCQWRxGmo/9ctlhCJt7rud/5pkooLtxbzW.Cwys3/jRz.T5kfNoy', 'Isabella', 'Suárez', 'student', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP); 