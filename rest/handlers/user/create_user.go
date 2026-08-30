package user

import (
	"encoding/json"
	"fmt"
	"net/http"

	"ecommerce/database"
	"ecommerce/util"
)

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {

	var newUser database.User

	err := json.NewDecoder(r.Body).Decode(&newUser)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Ivalid Request data", http.StatusBadRequest)
		return
	}

	createdUser := newUser.Store()

	util.SendData(
		w,
		createdUser,
		http.StatusCreated,
	)
}
