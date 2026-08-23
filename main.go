package main

import (
	"ecommerce/util"
	"fmt"
)

func main() {

	

	jwt, err := util.CreateJwt("my-secret-key", util.Payload{
		Sub:         45,
		FirstName:   "Shuayeb",
		LastName:    "Zet",
		Email:       "zt@gmail.com",
		IsShopowner: false,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(jwt)
}
