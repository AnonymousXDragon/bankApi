package accounts

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (ac *AccountController) DeleteAccountsHandler(w http.ResponseWriter, req *http.Request) {

	params := mux.Vars(req)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		err := ErrReponse{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}

		err.Serve(w)
	}

	err = ac.accounts.DeleteAccountService(int32(id))
	if err != nil {
		err := ErrReponse{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}

		err.Serve(w)
	}

	res := SuccessMessageResponse{
		Status:  http.StatusOK,
		Message: "successfull",
	}

	res.ServeR(w)
}
