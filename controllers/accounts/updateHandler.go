package accounts

import (
	db "bank/db/sqlc"
	"encoding/json"
	"net/http"
)

func (ac *AccountController) UpdateAccountsHandler(w http.ResponseWriter, req *http.Request) {

	var payload = &db.UpdateAccountParams{}

	if err := json.NewDecoder(req.Body).Decode(payload); err != nil {
		if err != nil {
			err := ErrReponse{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			}

			err.Serve(w)
		}
	}

	account, err := ac.accounts.UpdateAccountService(*payload)
	if err != nil {
		err := ErrReponse{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}

		err.Serve(w)
	}

	res := SuccessDataResponse{
		Status: http.StatusOK,
		Data:   account,
	}

	res.ServeR(w)
}
