-- name: CreateTransfer :one
INSERT INTO transfers(
    from_account_id,
    to_account_id,
    amount
)VALUES(
    $1,$2,$3
)RETURNING *
;

-- name: DeleteTransfer :exec
DELETE FROM transfers
WHERE id =$1;


-- name: UpdateTransfer :exec
UPDATE transfers
SET amount=$2
WHERE id = $1;

-- name: GetTransfer :one
SELECT * FROM transfers
WHERE id = $1;

-- name: FromListTransfer :many
SELECT * FROM transfers
WHERE from_account_id =$1
ORDER BY created_at
LIMIT $2
OFFSET $3;


-- name: ToListTransfer :many
SELECT * FROM transfers
WHERE to_account_id =$1
ORDER BY created_at
LIMIT $2
OFFSET $3;