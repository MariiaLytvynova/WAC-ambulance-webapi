 CREATE TABLE IF NOT EXISTS doctors (
            id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
            name VARCHAR(255) NOT NULL
        );

CREATE TABLE IF NOT EXISTS patients (
            id_patient INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
            name VARCHAR(255) NOT NULL,
            fullname VARCHAR(255) NOT NULL
);


CREATE TABLE IF NOT EXISTS examinations (
            id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
            id_patient INT NOT NULL,
            day VARCHAR(50) NOT NULL,
            start_time VARCHAR(50) NOT NULL,
            end_time VARCHAR(50) NOT NULL,
            name_examination VARCHAR(255) NOT NULL,

            CONSTRAINT fk_examination_patient
            FOREIGN KEY (id_patient)
            REFERENCES patients(id_patient)
);

INSERT INTO doctors (name)
VALUES
    ('Dr.Bobulova'),
    ('Dr.Max'),
    ('Dr.Semenko')
ON CONFLICT DO NOTHING;


INSERT INTO patients (name, fullname)
VALUES
    ('Alex', 'Paradajka'),
    ('Mrkvicka', 'Sliva')
ON CONFLICT DO NOTHING;


INSERT INTO examinations (id_patient, day, start_time, end_time, name_examination)
VALUES
    ('1', 'Pondelok', '10:00', '11:00', 'MR'),
    ('2', 'Pondelok', '9:00', '10:00', 'MR'),
    ('1', 'Streda', '10:00', '11:00', 'Prechadzka'),
    ('2', 'Utorok', '11:00', '12:00', 'Prechadzka'),
    ('1', 'Stvrtok', '8:00', '10:00', 'Odober krvi'),
    ('2', 'Piatok', '11:00', '12:00', 'Masaz')
ON CONFLICT DO NOTHING;

