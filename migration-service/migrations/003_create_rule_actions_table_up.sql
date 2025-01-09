CREATE TABLE rule_actions (
    id SERIAL PRIMARY KEY,
    rule_id INT NOT NULL,
    action TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    FOREIGN KEY (rule_id) REFERENCES rules (id) ON DELETE CASCADE
);


