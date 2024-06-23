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

func (as *AccoutServiceTestSuite) TestListAccountService() {
	as.mockRepo.On("ListAccounts").Return([]db.Account{}, nil)

	accounts, err := as.aservice.ListAccountService()

	assert.NoError(as.T(), err)
	assert.NotNil(as.T(), accounts)
	as.mockRepo.AssertExpectations(as.T())
}

func (as *AccoutServiceTestSuite) TestUpdateAccountService() {

	testParams := db.UpdateAccountParams{
		ID:      1,
		Balance: 20000000,
	}

	var expectedOutput db.Account

	as.mockRepo.On("UpdateAccount", testParams).Return(expectedOutput, nil)

	account, err := as.aservice.UpdateAccountService(testParams)

	assert.NoError(as.T(), err)
	assert.NotNil(as.T(), account)

	as.mockRepo.AssertExpectations(as.T())
}

func (as *AccoutServiceTestSuite) TestDeleteAccountService() {
	id := 1

	as.mockRepo.On("DeleteAccount", int32(id)).Return(nil)

	err := as.aservice.DeleteAccountService(int32(id))

	assert.NoError(as.T(), err)

	as.mockRepo.AssertExpectations(as.T())
}

func TestRunAccountService(t *testing.T) {
	as := new(AccoutServiceTestSuite)
	suite.Run(t, as)
}
