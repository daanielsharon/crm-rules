INSERT INTO rules (name, condition, schedule)
VALUES
    ('Notify Dormant Users', 'last_active > 90', 'hourly'),
    ('Annual Review Email', 'created_at > ''2021-01-01''', 'every_5_minutes'),
    ('Detect Malicious Activity', 'failed_logins > 30', 'every_10_minutes'),
    ('Send Weekly Summary', 'plan > ''premium''', 'hourly'),
    ('Detect Inactive Users', 'last_login < 30', 'every_5_minutes'),
    ('Send Birthday Greeting', 'dob > NOW() - INTERVAL 1 DAY', 'every_10_minutes');
