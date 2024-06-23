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

	accountR.HandleFunc("/home", AccountHomeHandler).Methods(http.MethodGet).Name("AccountHome")
	accountR.HandleFunc("/create", accountController.CreateAccountHandler).Methods(http.MethodPost).Name("CreateAccount")
	accountR.HandleFunc("/update", accountController.UpdateAccountsHandler).Methods(http.MethodPut).Name("CreateAccount")
	accountR.HandleFunc("/{id}", accountController.DeleteAccountsHandler).Methods(http.MethodPost).Name("CreateAccount")
	accountR.HandleFunc("/", accountController.ListAccountsHandler).Methods(http.MethodGet).Name("ListAccounts")

	return accountR
}

func AccountHomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("welcome to accounts api"))
}
