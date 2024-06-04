package accounts

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"

	db "bank/db/sqlc"

	"context"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	_ "github.com/lib/pq"
)

type AccountTestSuite struct {
	suite.Suite
	conn   *sql.DB
	driver string
	dbUrl  string
	q      *db.Queries
	ctx    context.Context
}

func (at *AccountTestSuite) SetupSuite() {

	at.driver = "postgres"
	at.dbUrl = "postgresql://postgres:password@localhost:5432/bankdb?sslmode=disable"
	var err error
	at.conn, err = sql.Open(at.driver, at.dbUrl)
	if err != nil {
		fmt.Println("error", err)
		log.Fatal(err.Error())
	}

	at.q = db.New(at.conn)
	at.ctx = context.Background()
	at.Truncate()
}

func (at *AccountTestSuite) TestCreateAccount() {

	payload := db.CreateAccountParams{
		Owner:    "arnold shazni",
		Balance:  1002030102,
		Currency: "USD",
	}

	result, err := at.q.CreateAccount(at.ctx, payload)

	assert.NoError(at.T(), err)
	assert.NotEmpty(at.T(), result)

	assert.Equal(at.T(), payload.Owner, result.Owner)
	assert.Equal(at.T(), payload.Balance, result.Balance)
	assert.Equal(at.T(), payload.Currency, result.Currency)

	assert.NotZero(at.T(), result.ID)
	assert.NotZero(at.T(), result.CreatedAt)
}

func (at *AccountTestSuite) TestUpdateAccount() {
	payload := db.UpdateAccountParams{
		ID:      1,
		Balance: 13712319231023,
	}

	result, err := at.q.UpdateAccount(at.ctx, payload)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	assert.NoError(at.T(), err)
	assert.NotEmpty(at.T(), result)

	assert.Equal(at.T(), payload.ID, result.ID)
	assert.NotZero(at.T(), result.Balance)
}

func (at *AccountTestSuite) TestGetAccount() {
	result, err := at.q.GetAccount(at.ctx, 1)

	assert.NoError(at.T(), err)
	assert.NotEmpty(at.T(), result)

	assert.NotZero(at.T(), result.Balance)
	assert.NotZero(at.T(), result.ID)

	assert.NotNil(at.T(), result.Owner)
	assert.IsType(at.T(), string(""), result.Owner, "owner shhould be a string type")
	assert.NotNil(at.T(), result.CreatedAt)
}

func (at *AccountTestSuite) TestListAccount() {
	result, err := at.q.ListAccounts(at.ctx)

	assert.NoError(at.T(), err)
	assert.NotEmpty(at.T(), result)

	assert.IsType(at.T(), []db.Account{}, result, "result must by type of []Account")
}

func (at *AccountTestSuite) TestDeleteAccount() {
	err := at.q.DeleteAccount(at.ctx, 1)
	assert.NoError(at.T(), err)
}

func (at *AccountTestSuite) Truncate() {

	query := "TRUNCATE TABLE Accounts RESTART IDENTITY CASCADE"
	tx, _ := at.conn.BeginTx(context.Background(), &sql.TxOptions{
		ReadOnly: false,
	})
	_, err := tx.Exec(query)

	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	tx.Commit()

}

func (at *AccountTestSuite) TearDownTest() {
	fmt.Println("....tested")
}

func (at *AccountTestSuite) TearDownSuite() {
	at.Truncate()
	if err := at.conn.Close(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

func TestAccountRepo(t *testing.T) {
	suite := new(AccountTestSuite)
	suite.SetT(t)

	suite.SetupSuite()
	defer suite.TearDownSuite()

	suite.Run("createAccount", suite.TestCreateAccount)
	suite.Run("updateAccount", suite.TestUpdateAccount)
	suite.Run("getAccount", suite.TestGetAccount)
	suite.Run("listAccount", suite.TestListAccount)
	suite.Run("deleteAccount", suite.TestDeleteAccount)
}
