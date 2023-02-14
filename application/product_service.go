package application

type ProductService struct {
	Persistense ProductPersistenceInterface
}

func(s *ProductService) Get(id string) (ProductInterface, error) {
	product, err := s.Persistense.Get(id)
	if err != nil {
		return nil, err
	} 
	return product, nil
}