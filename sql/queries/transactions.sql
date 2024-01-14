-- name: CreateTransaction :one
INSERT INTO transactions (user_id, amount, incoming, description, recurring, start_date, end_date)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;

-- name: GetUserTransactions :many
SELECT * FROM transactions WHERE user_id = $1;

-- name: GetUserIncomingTransactions :many
SELECT * FROM transactions WHERE user_id = $1 AND incoming = 1;

-- name: GetUserOutgoingTransactions :many
SELECT * FROM transactions WHERE user_id = $1 AND incoming = 0;

-- name: GetUserTransactionsByMonth :many
SELECT * FROM transactions 
WHERE user_id = $1 AND start_date <= $2 AND end_date >= $3;
