CREATE TABLE IF NOT EXISTS urls (
    hash VARCHAR(6) PRIMARY KEY,
    original_url VARCHAR,
    creation_date DATE,
    expiration_date DATE
);