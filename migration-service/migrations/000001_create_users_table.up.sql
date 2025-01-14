CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users ( 
    id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    name TEXT NOT NULL,
    email TEXT NOT NULL UNIQUE,
    last_active TIMESTAMP NOT NULL,
    plan TEXT DEFAULT 'free',
    failed_logins INTEGER DEFAULT 0,
    email_verified BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);