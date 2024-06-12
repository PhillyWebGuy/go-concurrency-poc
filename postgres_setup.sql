SET search_path TO public;

-- Drop the users table if it already exists
DROP TABLE IF EXISTS users;

-- Create the users table
CREATE TABLE users
(
    id         SERIAL PRIMARY KEY,
    first_name VARCHAR(50)                                          NOT NULL
);

-- Insert 10 records into the users table
INSERT INTO users (first_name)
VALUES ('Joe'),
       ('John'),
       ('Jane'),
       ('Robert'),
       ('Emily'),
       ('Michael'),
       ('Linda'),
       ('David'),
       ('Elizabeth'),
       ('Richard'),
       ('Susan');

-- Select all records to verify the insertion
SELECT *
FROM users;