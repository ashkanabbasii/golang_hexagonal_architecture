CREATE TABLE IF NOT EXISTS voucher_redemption_history (
    id SERIAL PRIMARY KEY,
    voucher_id INT REFERENCES voucher_codes(id) ON DELETE CASCADE,
    amount INT NOT NULL DEFAULT 0,
    code VARCHAR(255) NOT NULL,
    redeemed_at TIMESTAMP DEFAULT NOW(),
    user_id INT,
);

CREATE INDEX idx_redemption_code ON voucher_redemption_history(code);
CREATE INDEX idx_redemption_user_id ON voucher_redemption_history(user_id);