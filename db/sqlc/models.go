// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package db

import (
	"database/sql"
	"time"
)

type Account struct {
	ID        int32        `json:"id"`
	Owner     string       `json:"owner"`
	Balance   int64        `json:"balance"`
	Currency  string       `json:"currency"`
	CreatedAt sql.NullTime `json:"created_at"`
}

type Entry struct {
	ID        int32     `json:"id"`
	AccountID int32     `json:"account_id"`
	Amount    int64     `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}

type Transfer struct {
	ID            int32     `json:"id"`
	FromAccountID int32     `json:"from_account_id"`
	ToAccountID   int32     `json:"to_account_id"`
	Amount        int64     `json:"amount"`
	CreatedAt     time.Time `json:"created_at"`
}
