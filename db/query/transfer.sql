-- name: CreateTransfer :one
INSERT INTO transfers (
    "from_account_id",
    "to_account_id",
    "amount") VALUES ($1, $2, $3) RETURNING *;

-- name: GetTransferFromAccount :one
SELECT * FROM transfers
WHERE from_account_id = $1 LIMIT 1;

-- name: ListTransfersFromAccount :many
SELECT * FROM transfers
ORDER BY from_account_id
LIMIT $1
OFFSET $2;

-- name: UpdateTransfer :one
UPDATE transfers 
SET amount = $2
WHERE from_account_id = $1 AND to_account_id = $2
RETURNING *;

-- name: DeleteTransfer :exec
DELETE FROM transfers WHERE from_account_id = $1 AND to_account_id = $2;