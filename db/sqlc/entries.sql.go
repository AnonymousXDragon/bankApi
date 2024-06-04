// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: entries.sql

package db

import (
	"context"
)

const createEntires = `-- name: createEntires :one
INSERT INTO Entries (account_id, amount)
VALUES ($1, $2) 
RETURNING id, account_id, amount, created_at
`

type createEntiresParams struct {
	AccountID int32 `json:"account_id"`
	Amount    int64 `json:"amount"`
}

func (q *Queries) createEntires(ctx context.Context, arg createEntiresParams) (Entry, error) {
	row := q.queryRow(ctx, q.createEntiresStmt, createEntires, arg.AccountID, arg.Amount)
	var i Entry
	err := row.Scan(
		&i.ID,
		&i.AccountID,
		&i.Amount,
		&i.CreatedAt,
	)
	return i, err
}

const deleteEntrie = `-- name: deleteEntrie :exec
DELETE FROM Entries
WHERE id=$1
`

func (q *Queries) deleteEntrie(ctx context.Context, id int32) error {
	_, err := q.exec(ctx, q.deleteEntrieStmt, deleteEntrie, id)
	return err
}

const getEntrie = `-- name: getEntrie :one
SELECT id, account_id, amount, created_at FROM Entries
WHERE id=$1 LIMIT 1
`

func (q *Queries) getEntrie(ctx context.Context, id int32) (Entry, error) {
	row := q.queryRow(ctx, q.getEntrieStmt, getEntrie, id)
	var i Entry
	err := row.Scan(
		&i.ID,
		&i.AccountID,
		&i.Amount,
		&i.CreatedAt,
	)
	return i, err
}

const listEntries = `-- name: listEntries :many
SELECT id, account_id, amount, created_at FROM Entries
WHERE account_id=$1
ORDER BY created_at
`

func (q *Queries) listEntries(ctx context.Context, accountID int32) ([]Entry, error) {
	rows, err := q.query(ctx, q.listEntriesStmt, listEntries, accountID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Entry
	for rows.Next() {
		var i Entry
		if err := rows.Scan(
			&i.ID,
			&i.AccountID,
			&i.Amount,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateEntrie = `-- name: updateEntrie :one
UPDATE Entries
SET account_id=$2, amount=$3
WHERE id=$1 RETURNING id, account_id, amount, created_at
`

type updateEntrieParams struct {
	ID        int32 `json:"id"`
	AccountID int32 `json:"account_id"`
	Amount    int64 `json:"amount"`
}

func (q *Queries) updateEntrie(ctx context.Context, arg updateEntrieParams) (Entry, error) {
	row := q.queryRow(ctx, q.updateEntrieStmt, updateEntrie, arg.ID, arg.AccountID, arg.Amount)
	var i Entry
	err := row.Scan(
		&i.ID,
		&i.AccountID,
		&i.Amount,
		&i.CreatedAt,
	)
	return i, err
}