package accounts

import (
	as "bank/services/accounts"
)

type AccountController struct {
	accounts *as.AccountService
}

func NewAccountController(ac *as.AccountService) *AccountController {
	return &AccountController{
		accounts: ac,
	}
}
