INSERT INTO users (name, email, last_active, plan, email_verified)
VALUES
('John Doe', 'john.doe@example.com', NOW() - INTERVAL '100 days', 'premium', TRUE),
('Jane Smith', 'jane.smith@example.com', NOW() - INTERVAL '50 days', 'free', FALSE),
('Mike Brown', 'mike.brown@example.com', NOW() - INTERVAL '200 days', 'basic', TRUE);

