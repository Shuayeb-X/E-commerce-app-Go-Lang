package main

import (
	"ecommerce/cmd"
)

func main() {

	// jwt, err := util.CreateJwt("my-secret-key", util.Payload{
	// 	Sub:         45,
	// 	FirstName:   "Shuayeb",
	// 	LastName:    "Zet",
	// 	Email:       "zt@gmail.com",
	// 	IsShopowner: false,
	// })
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(jwt)

	cmd.Serve()

}

// eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVC
// J9.eyJzdWIiOjQ1LCJmaXJ
// zdF9uYW1lIjoiU2h1YXllYiIsImxhc3RfbmFtZSI6IlpldCIsImVtYWlsIjoienRAZ21haWwuY29tIiwiaXNfc2hvc
// F9vd25lciI6ZmFsc2V9.4tEsPcR5LhwqjEJn3
// Wy9zHi4hbGw80nGs_ba-1sIzhw
