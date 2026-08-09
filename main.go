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

	if r.Method != "GET" {
		http.Error(w, "Please give me Get Request", 400)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(productList)
}

func createProduct(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		http.Error(w, "Please give me Post Request", 400)
		return
	}

	var newProduct Product

	err := json.NewDecoder(r.Body).Decode(&newProduct)
	if err != nil {
		http.Error(w, "Invalid JSON data", http.StatusBadRequest)
		return
	}

	newProduct.ID = len(productList) + 1
	productList = append(productList, newProduct)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newProduct)
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
