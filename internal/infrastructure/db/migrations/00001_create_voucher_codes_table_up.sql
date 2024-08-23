CREATE TABLE IF NOT EXISTS voucher_codes (
    id SERIAL PRIMARY KEY,
    code VARCHAR(255) UNIQUE NOT NULL,
    amount INT NOT NULL DEFAULT 0,
    state VARCHAR(50) NOT NULL,
    usage_limit INT NOT NULL DEFAULT 0,
    user_limit INT NOT NULL DEFAULT 0,
    current_usage INT NOT NULL DEFAULT 0,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);