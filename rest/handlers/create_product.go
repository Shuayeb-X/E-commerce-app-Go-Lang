package handlers

import (
	"ecommerce/database"
	"ecommerce/util"
	"encoding/json"
	"fmt"
	"net/http"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {

	var newProduct database.Product

	err := json.NewDecoder(r.Body).Decode(&newProduct)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Invalid JSON data", http.StatusBadRequest)
		return
	}

	createdProduct := database.Store(newProduct)

	util.SendData(
		w,
		createdProduct,
		http.StatusCreated,
	)
}
