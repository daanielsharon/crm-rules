INSERT INTO users (name, email, last_active, plan, email_verified, failed_logins)
VALUES
('John Doe', 'john.doe@example.com', NOW() - INTERVAL '100 days', 'premium', TRUE, 0),
('Jane Smith', 'jane.smith@example.com', NOW() - INTERVAL '50 days', 'free', FALSE, 0),
('Mike Brown', 'mike.brown@example.com', NOW() - INTERVAL '200 days', 'basic', TRUE, 0);

