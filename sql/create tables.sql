CREATE TABLE wallets (
    id UUID PRIMARY KEY,
    owner_id VARCHAR(255) NOT NULL,
    balance NUMERIC(15, 2) NOT NULL DEFAULT 0,
    currency VARCHAR(10) NOT NULL,
    status VARCHAR(50) NOT NULL DEFAULT 'Active',
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

CREATE TABLE transactions (
    id UUID PRIMARY KEY,
    wallet_id UUID REFERENCES wallets(id),
    amount NUMERIC(15, 2) NOT NULL,
    type VARCHAR(50) NOT NULL, -- Credit or Debit
    status VARCHAR(50) NOT NULL, -- Pending, Completed
    timestamp TIMESTAMP NOT NULL
);
