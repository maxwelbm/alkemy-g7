package service

import (
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/maxwelbm/alkemy-g7.git/internal/repository/interfaces"
)

type ProductService struct {
	ProductRepository interfaces.IProductsRepo
}

func NewProductService(repo interfaces.IProductsRepo) *ProductService {
	return &ProductService{ProductRepository: repo}
}

func (ps *ProductService) GetAllProducts() (map[int]model.Product, error) {
	productsList, err := ps.ProductRepository.GetAll()
	if err != nil {
		return productsList, err
	}

	return productsList, nil
}

func (ps *ProductService) GetProductById(id int) (model.Product, error) {
	product, err := ps.ProductRepository.GetById(id)
	if err != nil {
		return model.Product{}, err
	}
	return product, nil
}

func (ps *ProductService) CreateProduct(product model.Product) (model.Product, error) {
	return model.Product{}, nil
}

func (ps *ProductService) UpdateProduct(id int, product model.Product) (model.Product, error) {
	return model.Product{}, nil
}

func (ps *ProductService) DeleteProduct(id int) error {
	return nil
}
