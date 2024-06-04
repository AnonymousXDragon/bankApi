package accounts

import db "bank/db/sqlc"

func (as *AccountService) CreateAccountService(payload db.CreateAccountParams) (*db.Account, error) {
	account, err := as.accrepo.CreateAccountTx(payload)
	if err != nil {
		return nil, err
	}

	return &account, nil
}
