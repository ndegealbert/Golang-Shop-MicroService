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
