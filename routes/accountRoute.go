package routes

import (
	"net/http"

	"bank/db"

	"github.com/gorilla/mux"

	accCont "bank/controllers/accounts"
	accRepo "bank/repository/accounts"
	accServ "bank/services/accounts"
)

func AccountRouter() *mux.Router {

	accountRepo := accRepo.NewAccountRepo(db.MakeConnection())
	accountService := accServ.NewAccountService(accountRepo)
	accountController := accCont.NewAccountController(accountService)
	accountR := mux.NewRouter().PathPrefix("/accounts").Subrouter()

	accountR.HandleFunc("/", AccountHomeHandler).Methods(http.MethodGet).Name("AccountHome")
	accountR.HandleFunc("/create", accountController.CreateAccountHandler).Methods(http.MethodPost).Name("CreateAccount")

	return accountR
}

func AccountHomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("welcome to accounts api"))
}
