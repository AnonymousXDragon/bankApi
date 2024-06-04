package accounts

import (
	db "bank/db/sqlc"
	"context"
	"database/sql"
	"fmt"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type TestAccountTxSuite struct {
	suite.Suite
	conn   *sql.DB
	driver string
	dbUrl  string
	q      *db.Queries
	ctx    context.Context
}

func (ta *TestAccountTxSuite) SetupSuite() {
	ta.driver = "postgres"
	ta.dbUrl = "postgresql://postgres:password@localhost:5432/bankdb?sslmode=disable"
	var err error
	ta.conn, err = sql.Open(ta.driver, ta.dbUrl)
	if err != nil {
		fmt.Println("error", err)
		log.Fatal(err.Error())
	}

	ta.q = db.New(ta.conn)
	ta.ctx = context.Background()
	ta.Truncate()
}

func (ta *TestAccountTxSuite) Truncate() {

	query := "TRUNCATE TABLE Accounts RESTART IDENTITY CASCADE"
	tx, _ := ta.conn.BeginTx(context.Background(), &sql.TxOptions{
		ReadOnly: false,
	})
	_, err := tx.Exec(query)

	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	tx.Commit()

}

func (ta *TestAccountTxSuite) TestCreateAccountTx() {
	acc := NewAccountRepo(ta.conn)

	payload := db.CreateAccountParams{
		Owner:    "arnold shazni",
		Balance:  12873192,
		Currency: "USD",
	}

	account, err := acc.AccountSideEffectsTx(payload, func() (db.Account, error) {
		return ta.q.CreateAccount(ta.ctx, payload)
	})

	assert.NoError(ta.T(), err)
	assert.NotEmpty(ta.T(), account)
	assert.NotNil(ta.T(), account)
	assert.IsType(ta.T(), db.Account{}, *account)
}

func (ta *TestAccountTxSuite) TestCreateAccountTx2() {
	acc := NewAccountRepo(ta.conn)

	payload := db.CreateAccountParams{
		Owner:    "arnold shazni",
		Balance:  12873192,
		Currency: "USD",
	}

	account, err := acc.CreateAccountTx(payload)

	assert.NoError(ta.T(), err)
	assert.NotEmpty(ta.T(), account)
	assert.NotNil(ta.T(), account)
	assert.IsType(ta.T(), db.Account{}, account)
}

func (ta *TestAccountTxSuite) TestUpdateAccountTx() {
	acc := NewAccountRepo(ta.conn)

	payload := db.UpdateAccountParams{
		ID:      1,
		Balance: 12873192,
	}

	account, err := acc.AccountSideEffectsTx(payload, func() (db.Account, error) {
		return ta.q.UpdateAccount(ta.ctx, payload)
	})

	assert.NoError(ta.T(), err)
	assert.NotEmpty(ta.T(), account)
	assert.NotNil(ta.T(), account)
	assert.IsType(ta.T(), db.Account{}, *account)
}

func (ta *TestAccountTxSuite) TestGetAccountTx() {
	acc := NewAccountRepo(ta.conn)

	account, err := acc.AccountSideEffectsTx(1, func() (db.Account, error) {
		return ta.q.GetAccount(ta.ctx, 1)
	})

	assert.NoError(ta.T(), err)
	assert.NotEmpty(ta.T(), account)
	assert.NotNil(ta.T(), account)
	assert.IsType(ta.T(), db.Account{}, *account)
}

func (ta *TestAccountTxSuite) TestDeleteAccountTx() {
	acc := NewAccountRepo(ta.conn)

	err := acc.AccountExecTx(1, func() error {
		return ta.q.DeleteAccount(ta.ctx, 1)
	})

	assert.NoError(ta.T(), err)
}

func (ta *TestAccountTxSuite) TestListAccountTx() {
	acc := NewAccountRepo(ta.conn)

	account, err := acc.AccountRetrieveTx(1, func() ([]db.Account, error) {
		return ta.q.ListAccounts(ta.ctx)
	})

	assert.NoError(ta.T(), err)
	assert.NotEmpty(ta.T(), account)
	assert.NotNil(ta.T(), account)
	assert.IsType(ta.T(), []db.Account{}, *account)
}

func (ta *TestAccountTxSuite) TearDownSuite() {
	ta.Truncate()
	if err := ta.conn.Close(); err != nil {
		log.Fatal("error in closing db connection ", err)
	}
}

func TestAccountTxRun(t *testing.T) {
	accountTx := new(TestAccountTxSuite)
	accountTx.SetT(t)

	accountTx.SetupSuite()
	defer accountTx.TearDownSuite()

	accountTx.Run("createAccountTx", accountTx.TestCreateAccountTx)
	accountTx.Run("craeteAccountTx2", accountTx.TestCreateAccountTx2)
	accountTx.Run("updateAccountTx", accountTx.TestUpdateAccountTx)
	accountTx.Run("getAccountTx", accountTx.TestGetAccountTx)
	accountTx.Run("listAccountTx", accountTx.TestListAccountTx)
	accountTx.Run("deleteAccountTx", accountTx.TestDeleteAccountTx)
}
