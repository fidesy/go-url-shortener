CREATE TABLE IF NOT EXISTS urls(
    id SERIAL PRIMARY KEY,
    hash VARCHAR(6) UNIQUE,
    original_url TEXT,
    creation_date TIMESTAMP,
    expiration_date TIMESTAMP
);

