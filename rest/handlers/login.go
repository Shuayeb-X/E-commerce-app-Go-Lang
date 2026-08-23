package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	
	"ecommerce/database"
	"ecommerce/util"
)

type ReqLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(w http.ResponseWriter, r *http.Request) {

	var reqLogin ReqLogin

	err := json.NewDecoder(r.Body).Decode(&reqLogin)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Ivalid Request data", http.StatusBadRequest)
		return
	}

	usr := database.Find(reqLogin.Email, reqLogin.Password)

	if usr == nil {
		http.Error(w, "Invalid Credentials", http.StatusBadRequest)
		return
	}

	util.SendData(w,usr,http.StatusOK)
}
