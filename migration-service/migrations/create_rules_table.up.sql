CREATE TABLE rules (
    id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    name TEXT NOT NULL,
    condition TEXT NOT NULL,
    schedule TEXT NOT NULL,  
    action TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS execution_logs (
    id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    rule_id UUID NOT NULL,
    status TEXT NOT NULL,
    executed_at TIMESTAMP NOT NULL DEFAULT NOW(),
    FOREIGN KEY (rule_id) REFERENCES rules(id)
);
