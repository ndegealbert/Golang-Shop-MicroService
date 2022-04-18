package application

type productReadModel interface {
	AllProducts() ([]products.product, err)
}

type ProductService struct {
	repo      products.Repository //storage
	readModel productReadModel    //of type
}

func NewProductService() ProductService {

}

func (s ProductService) AllProducts() {

}

type AddProductCommand struct {
	ID            string
	Name          string
	Description   string
	PriceCent     units
	priceCurrency string
}

func (s ProductService) AddProduct(cmd AddProductCommand) error {
	//NewPrice(cmd.PriceCent, cmd.priceCurrency)
	//newProduct(products.ID(cmd.ID), cmd.Description, cmd.Name, cmd.price)
	price.NewPrice(cmd.PriceCent, cmd.priceCurrency)
	products.newProduct(products.ID(cmd.ID), cmd.Description, cmd.Name, price)

	s.repo.Save

}
