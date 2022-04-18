package application

type productReadModel interface {
	AllProducts() ([]products.product, err)
}

type ProductService struct {
	repo      products.Repository //sorage
	readModel productReadModel    //of type
}

func NewProductService() ProductService {

}

func (s ProductService) AllProducts() {

}

func (s productService) AddProduct() {

}
