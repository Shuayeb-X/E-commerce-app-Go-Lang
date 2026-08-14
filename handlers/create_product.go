package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"ecommerce/database"
	"ecommerce/util"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {

	var newProduct database.Product

	err := json.NewDecoder(r.Body).Decode(&newProduct)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Invalid JSON data", http.StatusBadRequest)
		return
	}

	newProduct.ID = len(database.ProductList) + 1

	database.ProductList = append(
		database.ProductList,
		newProduct,
	)

	util.SendData(
		w,
		newProduct,
		http.StatusCreated,
	)
}
