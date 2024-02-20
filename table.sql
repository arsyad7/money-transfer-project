CREATE TABLE user_banks (
    id SERIAL PRIMARY KEY,
    account_name VARCHAR(255) NOT NULL,
    account_number VARCHAR(20) NOT NULL,
    bank_code VARCHAR(50) NOT NULL,
    account_balance DECIMAL(18, 2),
    created_at TIMESTAMP DEFAULT current_timestamp,
    created_by VARCHAR(255),
    modified_at TIMESTAMP,
    modified_by VARCHAR(255),
    deleted_at TIMESTAMP,
    deleted_by VARCHAR(255),
    is_active BOOLEAN DEFAULT true
);

CREATE TABLE money_transfer.transactions (
    id SERIAL PRIMARY KEY,
    transaction_id VARCHAR(50),
    user_bank_id INTEGER REFERENCES user_banks(id),
    amount DECIMAL(18, 2) NOT NULL,
    status VARCHAR(50),
    transaction_type VARCHAR(50),
    created_at TIMESTAMP DEFAULT current_timestamp,
    created_by VARCHAR(255),
    modified_at TIMESTAMP,
    modified_by VARCHAR(255),
    deleted_at TIMESTAMP,
    deleted_by VARCHAR(255),
    is_active BOOLEAN DEFAULT true
);

