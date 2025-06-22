-- Remove test users
DELETE FROM users WHERE email IN (
    '1@admin',
    '1@teacher', '2@teacher', '3@teacher', '4@teacher', '5@teacher',
    '6@teacher', '7@teacher', '8@teacher', '9@teacher', '10@teacher',
    '1@student', '2@student', '3@student', '4@student', '5@student',
    '6@student', '7@student', '8@student', '9@student', '10@student',
    '11@student', '12@student', '13@student', '14@student', '15@student',
    '16@student', '17@student', '18@student', '19@student', '20@student'
); 