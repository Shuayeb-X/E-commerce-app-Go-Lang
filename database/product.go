package database

var ProductList []Product

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgUrl      string  `json:"img_url"`
}

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

	ProductList = append(ProductList, prd1, prd2)
}
