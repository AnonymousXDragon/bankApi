package accounts

import (
	db "bank/db/sqlc"
	"encoding/json"
	"log"
	"net/http"
)

type ListSucessResponse struct {
	Status  int
	Message string
	Data    []db.Account
}

func (s *ListSucessResponse) ServerResponse(w http.ResponseWriter) {
	w.WriteHeader(s.Status)
	w.Header().Set("Content-Type", "application/json")

	data, err := json.Marshal(s)
	if err != nil {
		log.Fatal(err)
	}

	w.Write(data)
}

func (ac *AccountController) ListAccountsHandler(w http.ResponseWriter, req *http.Request) {
	accounts, err := ac.accounts.ListAccountService()
	if err != nil {
		err := ErrReponse{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}

		err.Serve(w)
	}

	res := ListSucessResponse{
		Status:  http.StatusOK,
		Message: "successfull",
		Data:    accounts,
	}

	res.ServerResponse(w)
}
