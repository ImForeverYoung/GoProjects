


CREATE TABLE groups (
    group_id SERIAL PRIMARY KEY,
    group_name VARCHAR(50) NOT NULL,
    faculty VARCHAR(50) CHECK (faculty IN ('Engineers', 'Humanities'))
);

CREATE TABLE students (
    student_id SERIAL PRIMARY KEY,
    student_name VARCHAR(80) NOT NULL,
    gender VARCHAR(5) CHECK (gender IN ('M', 'F')),
    birth_date DATE,
    group_id INT REFERENCES groups(group_id)
);
CREATE TABLE subjects (
    subject_id SERIAL PRIMARY KEY,
    subject_name VARCHAR(100) NOT NULL UNIQUE
);

CREATE TABLE schedules (
    schedule_id SERIAL PRIMARY KEY,
    group_id INT REFERENCES groups(group_id),
    subject_id INT REFERENCES subjects(subject_id),
    time_slot VARCHAR(50)
);


CREATE TABLE attendance {
    id SERIAL PRIMARY KEY,
    student_id INT REFERENCES students(student_id),
    subject_id INT REFERENCES subjects(subject_id),
    visit_day DATE,
    visited BOOLEAN,
};



INSERT INTO groups (group_name, faculty) VALUES 
('Eng-101', 'Engineers'),
('Hum-202', 'Humanities'),
('Eng-102', 'Engineers');


INSERT INTO students (student_name, gender, birth_date, group_id) VALUES 
('Amre Jumadiyev', 'M', '2003-05-15', 1),
('Dina Kairatkyzy', 'F', '2004-02-20', 1),    
('Sanzhar Aliyev', 'M', '2003-11-10', 2),
('Aidana Aiyasheva', 'F', '2005-01-05', 2),   
('Uliyana Mazurenko', 'F', '2002-08-30', 3);    


INSERT INTO subjects (subject_name) VALUES 
('Math'), 
('Physics'), 
('Philosophy'), 
('History');


INSERT INTO schedules (group_id, subject_id, time_slot) VALUES 
(1, 1, '09:00 - 10:30'), -- Eng-101: Math
(1, 2, '10:45 - 12:15'), -- Eng-101: Physics
(2, 3, '09:00 - 10:30'), -- Hum-202: Philosophy
(2, 4, '13:00 - 14:30'); -- Hum-202: History

-- Отмечаем посещаемость (пример)
INSERT INTO attendance (student_id, subject_id, visit_day, visited) VALUES 
(1, 1, '2026-01-11', TRUE),  -- Amre был на математике
(2, 1, '2026-01-11', FALSE), -- Dina пропустила математику
(1, 2, '2026-01-11', TRUE);  -- Amre был на физике


