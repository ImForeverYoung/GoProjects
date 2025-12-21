


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

CREATE TABLE schedules (
    schedule_id SERIAL PRIMARY KEY,
    group_id INT REFERENCES groups(group_id),
    subject VARCHAR(80),
    time_slot VARCHAR(50)
);


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


INSERT INTO schedules (group_id, subject, time_slot) VALUES 
(1, 'Math', '09:00 - 10:30'),         
(1, 'Physics', '10:45 - 12:15'),      
(2, 'Philosophy', '09:00 - 10:30'),   
(2, 'History', '13:00 - 14:30');


