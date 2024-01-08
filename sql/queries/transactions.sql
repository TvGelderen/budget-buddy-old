-- name: CreateTransaction :one
INSERT INTO transactions (user_id, amount, incoming, recurring, start_date, end_date)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetTransactionsByUserId :many
SELECT * FROM transactions WHERE user_id = $1;

-- name: GetIncomingTransactionslByUserId :many
SELECT * FROM transactions WHERE user_id = $1 AND incoming = 1;

-- name: GetOutgoingTransactionslByUserId :many
SELECT * FROM transactions WHERE user_id = $1 AND incoming = 0;
