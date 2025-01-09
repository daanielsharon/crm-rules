CREATE TABLE IF NOT EXISTS execution_logs (
    id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    rule_id INT NOT NULL,
    user_id UUID NOT NULL,
    action TEXT NOT NULL,
    status TEXT NOT NULL,
    executed_at TIMESTAMP NOT NULL DEFAULT NOW(),
    FOREIGN KEY (rule_id) REFERENCES rules(id)
);
