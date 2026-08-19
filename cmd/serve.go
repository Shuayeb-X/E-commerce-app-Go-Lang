package cmd

import (
	"ecommerce/config"
	"ecommerce/middleware"
	"fmt"
	"net/http"
)

func Serve() {

	config.GetConfig()
	manager := middleware.NewManager()

	manager.Use(
		middleware.Preflight,
		middleware.Cors,
		middleware.Logger,
	)

	mux := http.NewServeMux()

	wrappedMux := manager.WrapMux(mux)

	initRoutes(mux, manager)

	fmt.Println("Server Running on :3000")

	err := http.ListenAndServe(":3000", wrappedMux)

	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
