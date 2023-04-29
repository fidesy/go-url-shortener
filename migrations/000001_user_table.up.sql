CREATE TABLE IF NOT EXISTS users(
    id SERIAL PRIMARY KEY,
    name VARCHAR(100),
    username VARCHAR(50),
    password_hash VARCHAR(100)
);

CREATE TABLE IF NOT EXISTS urls(
    id SERIAL PRIMARY KEY,
    user_id INT,
    hash VARCHAR(6) UNIQUE,
    original_url TEXT,
    creation_date TIMESTAMP,
    expiration_date TIMESTAMP
);

