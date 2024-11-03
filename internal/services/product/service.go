package product

type Service struct {
}

func NewSerivce() *Service {
	return &Service{}
}

func (*Service) List() []Product {
	return allProducts
}
