package accounts

import (
	db "bank/db/sqlc"
)

type AccountRepoType interface {
	AccountSideEffectsTx(args any, fn func() (db.Account, error)) (*db.Account, error)
	AccountRetrieveTx(args any, fn func() ([]db.Account, error)) (*[]db.Account, error)
	AccountExecTx(args any, fn func() error) error
	CreateAccountTx(args db.CreateAccountParams) (db.Account, error)
	ListAccounts() ([]db.Account, error)
	UpdateAccount(params db.UpdateAccountParams) (db.Account, error)
	DeleteAccount(id int32) error
}
