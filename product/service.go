package products

//Service interface
type Service interface {
	GetProductByID(param *getProductByIDRequest) (*Product, error)
}

type service struct {
	repository Repository
}

//NewService instance
func NewService(repository Repository) Service{
	return &service{
		repository: repository,
	}
}

//GetProductById method
func (s *service) GetProductByID(param *getProductByIDRequest) (*Product, error) {
	return s.repository.GetProductByID(param.ProductID)
}
