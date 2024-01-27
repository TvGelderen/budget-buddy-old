-- name: CreateTransaction :one
INSERT INTO transactions (user_id, amount, incoming, description, recurring, start_date, end_date)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;

-- name: UpdateTransaction :exec
UPDATE transactions
SET amount = $3, incoming = $4, description = $5, recurring = $6, start_date = $7, end_date = $8
WHERE id = $1 AND user_id = $2;

-- name: DeleteTransaction :exec
DELETE FROM transactions
WHERE id = $1 AND user_id = $2;

-- name: GetTransaction :one
SELECT * FROM transactions WHERE id = $1;

-- name: GetUserTransactions :many
SELECT * FROM transactions WHERE user_id = $1;

-- name: GetUserIncomingTransactions :many
SELECT * FROM transactions WHERE user_id = $1 AND incoming = 1;

-- name: GetUserOutgoingTransactions :many
SELECT * FROM transactions WHERE user_id = $1 AND incoming = 0;

-- name: GetUserTransactionsByMonth :many
SELECT * FROM transactions 
WHERE user_id = $1 AND start_date < $2 AND end_date >= $3;
