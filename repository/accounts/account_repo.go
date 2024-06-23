package accounts

import (
	db "bank/db/sqlc"
	"context"
	"database/sql"
	"fmt"
)

type AccountRepo struct {
	db *sql.DB
	q  *db.Queries
}

func NewAccountRepo(d *sql.DB) *AccountRepo {
	return &AccountRepo{
		db: d,
		q:  db.New(d),
	}
}

func (a *AccountRepo) AccountSideEffectsTx(args any, fn func() (db.Account, error)) (*db.Account, error) {

	tx, err := a.db.BeginTx(context.Background(), &sql.TxOptions{
		ReadOnly: false,
	})

	if err != nil {
		return nil, err
	}

	account, err := fn()
	// fmt.Println("creating", account)
	// fmt.Println("error", err)
	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("execution failed %v", err.Error())
	}

	err = tx.Commit()
	if err != nil {
		return nil, fmt.Errorf("failed to commit the transaction: %v", err.Error())
	}
	return &account, nil
}

func (a *AccountRepo) CreateAccountTx(args db.CreateAccountParams) (db.Account, error) {
	// account, err := acc.AccountSideEffectsTx(payload, func() (db.Account, error) {
	// 	return ta.q.CreateAccount(ta.ctx, payload)
	// })
	if a.q == nil {
		return db.Account{}, fmt.Errorf("query executor is nil")
	}

	account, err := a.AccountSideEffectsTx(args, func() (db.Account, error) {
		return a.q.CreateAccount(context.Background(), args)
	})
	fmt.Println(account)
	return *account, err
}

// preventing from dirty-reads and for getting consistent data reads
// only see the committed changes
func (a *AccountRepo) AccountRetrieveTx(args any, fn func() ([]db.Account, error)) (*[]db.Account, error) {
	tx, err := a.db.BeginTx(context.Background(), &sql.TxOptions{
		ReadOnly: false,
	})

	if err != nil {
		return nil, err
	}

	account, err := fn()
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	return &account, nil
}

func (a *AccountRepo) AccountExecTx(args any, fn func() error) error {
	tx, err := a.db.BeginTx(context.Background(), &sql.TxOptions{
		ReadOnly: false,
	})

	if err != nil {
		return err
	}

	err = fn()
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}

func (a *AccountRepo) ListAccounts() ([]db.Account, error) {
	accounts, err := a.q.ListAccounts(context.Background())
	if err != nil {
		return []db.Account{}, err
	}

	return accounts, nil
}

func (a *AccountRepo) UpdateAccount(params db.UpdateAccountParams) (db.Account, error) {
	return a.q.UpdateAccount(context.Background(), params)
}

func (a *AccountRepo) DeleteAccount(id int32) error {
	return a.q.DeleteAccount(context.Background(), id)
}
