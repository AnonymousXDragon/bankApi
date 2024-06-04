-- name: CreateAccount :one
INSERT INTO Accounts ( owner , balance , currency) 
VALUES ($1,$2,$3) RETURNING *;

-- name: UpdateAccount :one
UPDATE Accounts
SET balance=$2
WHERE id=$1 RETURNING *;

-- name: GetAccount :one
SELECT * FROM Accounts
WHERE id=$1 LIMIT 1;

-- name: GetAccount :one
SELECT * FROM Accounts
WHERE id=$1 LIMIT 1;


-- name: ListAccounts :many
SELECT * FROM Accounts
ORDER BY id;

-- name: DeleteAccount :exec
DELETE FROM Accounts 
WHERE id=$1;