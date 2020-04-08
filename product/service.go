package products

//Service interface
type Service interface {
	GetProductByID(param *getProductByIDRequest) (*Product, error)
	GetProducts(params *getProductsRequest) (*ProductList, error)
}

type service struct {
	repository Repository
}

//NewService instance
func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

//GetProducts method
func (s *service) GetProducts(params *getProductsRequest) (*ProductList, error) {
	products, err := s.repository.GetProducts(params)
	if err != nil {
		panic(err)
	}
	totalProducts, err := s.repository.GetTotalProducts()

	return &ProductList{Data: products, TotalRecords: totalProducts}, nil
}

//GetProductById method
func (s *service) GetProductByID(param *getProductByIDRequest) (*Product, error) {
	return s.repository.GetProductByID(param.ProductID)
}
