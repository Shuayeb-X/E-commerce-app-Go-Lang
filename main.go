package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgUrl      string  `json:"img_url"`
}

var productList []Product

func getProducts(w http.ResponseWriter, r *http.Request) {

	sendData(w,productList,http.StatusOK)

}

func createProduct(w http.ResponseWriter, r *http.Request) {

	var newProduct Product

	err := json.NewDecoder(r.Body).Decode(&newProduct)
	if err != nil {
		http.Error(w, "Invalid JSON data", http.StatusBadRequest)
		return
	}

	newProduct.ID = len(productList) + 1
	productList = append(productList, newProduct)

	sendata(w,newProduct,http.StatusCreated)
}

func sendData(w http.ResponseWriter, data interface{}, status int) {

	w.WriteHeader(status)

	json.NewEncoder(w).Encode(data)
}

// Initial products
func init() {

	prd1 := Product{
		ID:          1,
		Name:        "Cherry",
		Description: "Cherry color is red",
		Price:       60,
		ImgUrl:      "https://images.contentstack.io/v3/assets/bltcedd8dbd5891265b/blt2a5be8abcac1a15f/667081fd5014f14c2a033ce6/types-of-cherries-on-branch.jpg",
	}

	prd2 := Product{
		ID:          2,
		Name:        "Carambola",
		Description: "Carambola color is green",
		Price:       20,
		ImgUrl:      "https://example.com/carambola.jpg",
	}

	productList = append(productList, prd1, prd2)
}

func main() {

	mux := http.NewServeMux()

	mux.Handle(
		"GET /products",
		corsMiddleware(http.HandlerFunc(getProducts)),
	)
	mux.Handle("POST /create-product",
		corsMiddleware(http.HandlerFunc(createProduct)),
	)

	fmt.Println("Server Running on :3000")

	globalRouter := globalRouter(mux)

	err := http.ListenAndServe(":3000", globalRouter)

	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

func corsMiddleware(next http.Handler) http.Handler {

	handleCors := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST,PUT,PATCH,DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Sayem")
		w.Header().Set("Content-Type", "application/json")

		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(handleCors)

}

func globalRouter(mux *http.ServeMux) http.Handler {

	handleAllReq := func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST,PUT,PATCH,DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Sayem")
		w.Header().Set("Content-Type", "application/json")
		if r.Method == "OPTIONS" {
			w.WriteHeader(200)
			return
		} else {
			mux.ServeHTTP(w, r)
		}

	}

	return http.HandlerFunc(handleAllReq)
}
