package accounts

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	db "bank/db/sqlc"
	"bank/repository/accounts/mocks"
	aService "bank/services/accounts"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type AccountControllerTestSuite struct {
	suite.Suite
	aservice    aService.AccountService
	mockRepo    *mocks.AccountRepoType
	aController *AccountController
}

func (as *AccountControllerTestSuite) SetupSuite() {
	as.mockRepo = mocks.NewAccountRepoType(as.T())
	as.aservice = *aService.NewAccountService(as.mockRepo)
	as.aController = NewAccountController(&as.aservice)
}

func (as *AccountControllerTestSuite) TestCreateAccountHandler() {

	testParams := db.CreateAccountParams{
		Owner:    "test01",
		Balance:  7000,
		Currency: "USD",
	}

	expectedOutput := db.Account{
		ID:       1,
		Owner:    "test01",
		Balance:  7000,
		Currency: "USD",
	}

	as.mockRepo.On("CreateAccountTx", testParams).Return(expectedOutput, nil)

	output, err := as.aservice.CreateAccountService(testParams)

	assert.NoError(as.T(), err)
	assert.Equal(as.T(), expectedOutput, *output)

	data, err := json.Marshal(testParams)
	assert.NoError(as.T(), err)

	req, err := http.NewRequest(http.MethodPost, "/create", bytes.NewBuffer(data))
	assert.NoError(as.T(), err)

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(as.aController.CreateAccountHandler)
	handler.ServeHTTP(rr, req)

	assert.Equal(as.T(), http.StatusOK, rr.Code)

	var response SuccessDataResponse
	var responseData SuccessDataResponse

	err = json.NewDecoder(rr.Body).Decode(&response)
	assert.NoError(as.T(), err)

	resData, err := json.Marshal(response)
	assert.NoError(as.T(), err)

	err = json.Unmarshal(resData, &responseData)
	assert.NoError(as.T(), err)

	assert.NotEmpty(as.T(), response.Data)
	assert.Equal(as.T(), http.StatusOK, response.Status)
	assert.Equal(as.T(), responseData.Data, response.Data)

	as.mockRepo.AssertExpectations(as.T())
}

func TestRunControllerService(t *testing.T) {
	as := new(AccountControllerTestSuite)
	suite.Run(t, as)
}
