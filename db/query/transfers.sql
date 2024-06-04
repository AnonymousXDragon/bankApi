-- name: createTransfer :one
INSERT INTO Transfers ( from_account_id , to_account_id , amount )
VALUES ($1,$2,$3) RETURNING *;


-- name: getTransfer :one
SELECT * FROM Transfers
WHERE id=$1 LIMIT 1;

-- name: listTransfers :many
SELECT * FROM Transfers
WHERE from_account_id=$1
ORDER BY created_at;