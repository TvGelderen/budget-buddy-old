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
    date TIMESTAMP NOT NULL,
    next_date TIMESTAMP,
    end_date TIMESTAMP,
    CONSTRAINT fk_user_id
        FOREIGN KEY(user_id)
        REFERENCES users(id) 
        ON DELETE CASCADE
);

CREATE INDEX id_idx ON users(id);
CREATE INDEX user_id ON transactions(user_id);

-- +goose Down
DROP TABLE transactions;
DROP TABLE users;

DROP INDEX id_idx ON users(id);
DROP INDEX user_id ON transactions(user_id);
