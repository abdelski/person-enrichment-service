CREATE TABLE people (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    surname VARCHAR(255) NOT NULL,
    age INTEGER,
    gender VARCHAR(50),
    nationality VARCHAR(50),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE INDEX idx_people_name ON people(name);
CREATE INDEX idx_people_surname ON people(surname);
CREATE INDEX idx_people_age ON people(age);
CREATE INDEX idx_people_gender ON people(gender);
CREATE INDEX idx_people_nationality ON people(nationality);