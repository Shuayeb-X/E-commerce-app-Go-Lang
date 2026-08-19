package database

var productList []Product

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgUrl      string  `json:"img_url"`
}

func Store(p Product) Product{

	p.ID =len(productList) + 1
	productList = append(productList, p)
	return p
}

func List() []Product {
	return productList
}

func Get(productID int) *Product {
	for _, product := range productList {

		if product.ID == productID {
			return &product
		}
	}
	return nil
}

func Update(product Product) {

	for idx, p := range productList {

		if p.ID == product.ID {
			productList[idx] = product

		}
	}

}

func Delete(productID int) {

	var tempList []Product = make([]Product, 0)

	for _, p := range productList {

		if p.ID != productID {
			tempList =append(tempList,p)
		}
	}

	productList = tempList
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

	productList = append(productList, prd1, prd2)
}
