package accounts

import (
	db "bank/db/sqlc"
	"encoding/json"
	"net/http"
)

func (ac *AccountController) CreateAccountHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	payload := db.CreateAccountParams{}
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		errRes := &ErrReponse{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
		errRes.Serve(w)
	}

	account, err := ac.accounts.CreateAccountService(payload)

	if err != nil {
		errRes := &ErrReponse{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
		errRes.Serve(w)
	}

	res := &SuccessDataResponse{
		Status: http.StatusOK,
		Data:   *account,
	}

	res.ServeR(w)
}
