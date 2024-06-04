-- name: createEntires :one
INSERT INTO Entries (account_id, amount)
VALUES ($1, $2) 
RETURNING *;

-- name: getEntrie :one
SELECT * FROM Entries
WHERE id=$1 LIMIT 1;

-- name: updateEntrie :one
UPDATE Entries
SET account_id=$2, amount=$3
WHERE id=$1 RETURNING *;

-- name: listEntries :many
SELECT * FROM Entries
WHERE account_id=$1
ORDER BY created_at;

-- name: deleteEntrie :exec
DELETE FROM Entries
WHERE id=$1;