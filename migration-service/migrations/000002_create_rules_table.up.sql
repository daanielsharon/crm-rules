CREATE TABLE rules (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    condition TEXT NOT NULL,
    schedule TEXT NOT NULL,  
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

