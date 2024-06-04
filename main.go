package main

import (
	"fmt"
	"log"
	"net/http"

	appRoutes "bank/routes"

	"github.com/gorilla/mux"
	"github.com/subosito/gotenv"
)

func main() {
	err := gotenv.Load()
	if err != nil {
		log.Fatal(err.Error())
	}

	app := mux.NewRouter()
	app.HandleFunc("/", HomeHandler).Methods(http.MethodGet)

	accountRouter := appRoutes.AccountRouter()

	app.PathPrefix("/accounts").Handler(accountRouter)

	server := &http.Server{
		Handler: app,
		Addr:    "0.0.0.0:8080",
	}

	log.Fatal(server.ListenAndServe())
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello,World")
}
