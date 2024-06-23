package accounts

import (
	as "bank/services/accounts"
	"encoding/json"
	"log"
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

func (err *ErrReponse) Serve(w http.ResponseWriter) {
	w.WriteHeader(err.Status)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(err)
}

type SuccessDataResponse struct {
	Status int
	Data   any
}

func (sr *SuccessDataResponse) ServeR(w http.ResponseWriter) {
	w.WriteHeader(sr.Status)
	w.Header().Set("Content-Type", "application/json")
	data, err := json.Marshal(sr)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(data)
}

type SuccessMessageResponse struct {
	Status  int
	Message string
}

func (sr *SuccessMessageResponse) ServeR(w http.ResponseWriter) {
	w.WriteHeader(sr.Status)
	w.Header().Set("Content-Type", "application/json")
	data, err := json.Marshal(sr)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(data)
}

type AccountController struct {
	accounts *as.AccountService
}

func NewAccountController(ac *as.AccountService) *AccountController {
	return &AccountController{
		accounts: ac,
	}
}
