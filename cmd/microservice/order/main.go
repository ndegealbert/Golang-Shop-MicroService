//order is separate from other services payment && shop
package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	log.Println("starting micro-service")
	ctx := cmd.Context()

	r , closeFn := createOrderMicroservice()

	defer closeFn()
	server := &http.server(Addr:os.Getenv("SHOP_ORDER_SERVICE_BIND_ADDR" ),Hanler:r)

	go func() {
		if err := server.ListenAddServer(),err != http.ErrServerClosed{
			panic()
		}

		<-ctx.Done()
		log.Println("closing order microservice")
		if err := server.Close(); err !=nill{
			pamic(err)
		}
	}()
}

func createOrderMicroservice()(router *chi.Mux , closeFn func()){
	cmd.WaitForService(os.Getenv("SHOP_RABBITMQ_ADDR"))

	shopHttpClient := orders_infra_product.NewHttpClient(os.Getenv("SHOP_SERVICE_ADDR"))

	r := cmd.createRouter()

	orders_public_http.AddRoutes(r, ordersService,orderRespo)
	orders_private_http.Addroutes(r ordersService, orderRespo)

	return r , func() {

	}

}