SET search_path TO public;

-- Drop the users table if it already exists
DROP TABLE IF EXISTS books;

-- Create the users table
CREATE TABLE books (
    id SERIAL PRIMARY KEY,
    title VARCHAR(50) NOT NULL
);

-- Insert 10 records into the books table
INSERT INTO
    books (title)
VALUES
    ('Meditations'),
    ('The Art of War'),
    ('The Prince'),
    ('The Republic'),
    ('The Wealth of Nations');

-- Select all records to verify the insertion
SELECT
    *
FROM
    books;
