-- +goose Up
CREATE TABLE users (
    id UUID PRIMARY KEY,
    username TEXT NOT NULL,
    email TEXT UNIQUE NOT NULL,
    password_hash BYTEA,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

CREATE TABLE transactions (
    id SERIAL PRIMARY KEY,
    user_id UUID NOT NULL,
    amount DOUBLE PRECISION NOT NULL,
    incoming BOOLEAN NOT NULL,
    recurring TEXT NOT NULL,
    start_date TIMESTAMP,
    end_date TIMESTAMP,
    CONSTRAINT fk_user_id
        FOREIGN KEY(user_id)
        REFERENCES users(id) 
        ON DELETE CASCADE
);

CREATE INDEX idx_users_id ON users(id);
CREATE INDEX idx_transactions_user_id ON transactions(user_id);

-- +goose Down
DROP TABLE transactions;
DROP TABLE users;
