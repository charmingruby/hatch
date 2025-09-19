CREATE TABLE IF NOT EXISTS notes (
    id VARCHAR PRIMARY KEY,
    title VARCHAR NOT NULL,
    content VARCHAR NOT NULL,
    archived BOOLEAN NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP
);