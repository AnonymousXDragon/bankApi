package accounts

import (
	"bank/repository/accounts"
)

type AccountService struct {
	accrepo accounts.AccountRepoType
}

func NewAccountService(accrepo accounts.AccountRepoType) *AccountService {
	return &AccountService{
		accrepo,
	}
}
