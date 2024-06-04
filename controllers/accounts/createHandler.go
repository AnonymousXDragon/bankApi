package accounts

import (
	db "bank/db/sqlc"
	"encoding/json"
	"net/http"
)

type SampleResponse interface {
	Serve(w http.ResponseWriter)
	ServeR(w http.ResponseWriter)
}

type ErrReponse struct {
	Status  int
	Message string
}

type SuccessResponse struct {
	Status int
	Data   any
}

func (sr *SuccessResponse) ServeR(w http.ResponseWriter) {
	w.WriteHeader(sr.Status)
	w.Header().Set("Content-Type", "application/json")
	data, err := json.Marshal(sr)
	if err != nil {
		panic(err)
	}
	w.Write(data)
}

func (err *ErrReponse) Serve(w http.ResponseWriter) {
	w.WriteHeader(err.Status)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(err)
}

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

	res := &SuccessResponse{
		Status: http.StatusOK,
		Data:   *account,
	}

	res.ServeR(w)
}
