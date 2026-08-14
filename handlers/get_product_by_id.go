package handlers

import (
	"ecommerce/database"
	"ecommerce/util"
	"net/http"
	"strconv"
)

func GetProductByID(w http.ResponseWriter, r *http.Request) {

	productId := r.PathValue("productId")

	id, err := strconv.Atoi(productId)

	if err != nil {
		http.Error(w, "Please give me valid id", http.StatusBadRequest)
		return
	}

	for _, product := range database.ProductList {

		if product.ID == id {
			util.SendData(w, product, http.StatusOK)
			return
		}
	}

	http.Error(w, "Product not found", http.StatusNotFound)
}
