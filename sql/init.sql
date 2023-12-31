-- USERS TABLE CONFIGURAT
CREATE TABLE IF NOT EXISTS person
(
    id SERIAL PRIMARY KEY,
    email VARCHAR(320) UNIQUE NOT NULL,
    password VARCHAR(128) NOT NULL, 
    activation_link VARCHAR(128),
    is_activated BOOLEAN NOT NULL,
);

ALTER TABLE person
ALTER COLUMN is_activated
SET DEFAULT false;


-- TOKEN TABLE CONFIGURAT
CREATE TABLE IF NOT EXISTS tokens
(
    person_id INT REFERENCES person(id),
    refresh_token TEXT UNIQUE NOT NULL
);