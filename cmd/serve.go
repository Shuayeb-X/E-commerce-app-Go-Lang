package cmd

import (
	"ecommerce/config"
	"ecommerce/rest"
	"ecommerce/rest/handlers/product"
)

func Serve() {

	cnf := config.GetConfig()

	productHandler := product.NewHandler()

	userHandler := user.NewHandler()

	server := rest.NewServer(
		cnf,
		productHandler,
		userHandler,
	)

	server.Start()

}
