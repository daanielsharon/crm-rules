INSERT INTO rules (name, condition, action)
VALUES
    ('Notify Dormant Users', 'last_active > 90', 'send_notification'),
    ('Annual Review Email', 'registration_date < ''2021-01-01''', 'send_email');
    ('Detect Malicious Activity', 'failed_logins > 30', 'alert_admin'),
    ('Send Weekly Summary', 'last_login > 30', 'send_summary'),
    ('Detect Inactive Users', 'last_login < 30', 'alert_admin'),
    ('Send Birthday Greeting', 'dob > NOW() - INTERVAL 1 DAY', 'send_birthday_greeting');
