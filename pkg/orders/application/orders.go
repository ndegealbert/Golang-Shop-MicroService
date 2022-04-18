package application

type productServices interface {
}

type paymentService interface {
}

type OrderService struct {
}

func newOrderService() {

}

type PlaceOrderCommand struct {
}

func (s OrderService) PlaceOrder(cmd PlaceOrderCommand) error {

}

type makeOrderAsPaidCommand struct {
}

func (s OrderService) MarkOrderAsPaid(cmd makeOrderAsPaidCommand) error {

}

func (s OrderService) orderById(id orders.ID) (orders.orders, error) {

}