package accounts

import db "bank/db/sqlc"

func (as *AccountService) CreateAccountService(payload db.CreateAccountParams) (*db.Account, error) {
	account, err := as.accrepo.CreateAccountTx(payload)
	if err != nil {
		return nil, err
	}

	return &account, nil
}

func (as *AccountService) ListAccountService() ([]db.Account, error) {
	return as.accrepo.ListAccounts()
}

func (as *AccountService) UpdateAccountService(params db.UpdateAccountParams) (db.Account, error) {
	return as.accrepo.UpdateAccount(params)
}

func (as *AccountService) DeleteAccountService(id int32) error {
	return as.accrepo.DeleteAccount(id)
}
