CREATE SCHEMA checklist;

CREATE TABLE checklist.tasks (
    id SERIAL PRIMARY KEY,
    title VARCHAR(50) NOT NULL CHECK (char_length(title) BETWEEN 1 AND 50),
    description VARCHAR(1000) CHECK (char_length(description) BETWEEN 1 AND 1000),
    completed BOOLEAN NOT NULL
)