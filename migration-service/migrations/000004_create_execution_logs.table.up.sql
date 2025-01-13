CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE execution_logs (
    id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    rule_id INT NOT NULL,
    user_id UUID NOT NULL,
    action TEXT NOT NULL,
    status TEXT NOT NULL CHECK (status IN ('success', 'failure', 'pending', 'skipped')),
    executed_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (rule_id) REFERENCES rules(id),
    FOREIGN KEY (user_id) REFERENCES users(id)
);
