package main

import (
	"log"
	"net/http"
	"os"
)

func main()  {
	log.Println("starting shop microservice")

	ctx := cmd.Context()
	r := createShopMicroService()

	server := $http.Server{Addr: os.Getenv("SHOP_PRODUCT_SERVICE_BIND_ADDR"),Handler: r}


	go func() {
		if err := server.ListenAndServer(); err != http.ErrServerClosed{
			panic(err)
		}
	}()

	<=ctx.Done()

	log.Println("closing shop microservice")

	if err := server.Close(); err != nil{
		panic(err)
	}
}

func createShopMicroService() *chi.Mux {
	shopProductRepo := shop_infra_product.newMemoryRepository()

	r := cmd.createRouter()

	shop_interfaces_public_http.AddRoutes(r,shopProductRepo)
	shop_interface_private_http.AddRoutes(r, shopProductRepo)

	return  r
}