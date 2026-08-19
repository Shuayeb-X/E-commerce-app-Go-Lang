package handlers

import (
	"ecommerce/database"
	"ecommerce/util"
	"net/http"
	"strconv"
)

func GetProduct(w http.ResponseWriter, r *http.Request) {

	productId := r.PathValue("id")

	pID, err := strconv.Atoi(productId)

	if err != nil {
		http.Error(w, "Please give me valid id", http.StatusBadRequest)
		return
	}

	product := database.Get(pID)

	if product == nil {
		util.SendError(w, 404, "Prodcut not found")
		return
	}
	util.SendData(w,product,200)
}
