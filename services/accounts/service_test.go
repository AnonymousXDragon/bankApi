package accounts

import (
	"testing"

	db "bank/db/sqlc"
	"bank/repository/accounts/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type AccoutServiceTestSuite struct {
	suite.Suite
	aservice AccountService
	mockRepo *mocks.AccountRepoType
}

func (as *AccoutServiceTestSuite) SetupSuite() {
	as.mockRepo = mocks.NewAccountRepoType(as.T())
	as.aservice = *NewAccountService(as.mockRepo)
}

func (as *AccoutServiceTestSuite) TestCreateAccountService() {
	testParams := db.CreateAccountParams{
		Owner:    "test01",
		Balance:  70000,
		Currency: "USD",
	}

	expectedAccount := db.Account{
		ID:       1,
		Owner:    "test01",
		Balance:  70000,
		Currency: "USD",
	}

	as.mockRepo.On("CreateAccountTx", testParams).Return(expectedAccount, nil)

	account, err := as.aservice.accrepo.CreateAccountTx(testParams)

	assert.NoError(as.T(), err)
	assert.Equal(as.T(), expectedAccount, account)

	as.mockRepo.AssertExpectations(as.T())
}

func TestRunAccountService(t *testing.T) {
	as := new(AccoutServiceTestSuite)
	suite.Run(t, as)
}
