package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	log.Println("Starting payments microservice")

	defer log.Println("Closing payment microservice")

	ctx := cmd.Context()
	paymentInterface := createPaymentMicroService()

	if err := paymentInterface.Run(ctx); err != nil {

		panic(err)
	}
}

func createPaymentMicroService() amqp.paymentsInterface {
	cmd.waitForService(os.Getenv("SHO_RABBITMQ_ADDR"))
	paymentsService := payments_app.newPaymentService(
		payments_infra_orders.NewHTTPClient(os.Getenv("SHOP_ORDER_SERVICE_ADD")),
	)

	paymentInterface, err := amqp.NewPaymentsInterface(

		fmt.Sprintf("amqp://%s",os.Getenv("SHOP_RABBITMQ_ADDR"))
		os.Getenv("SHOP_RABBITMQ_ORDERS_TO_PAY_QUE")
		paymentsService,
		)

	if err!=nil{
		panic(err)
	}

	return paymentsService
}
