package service

import (
	"errors"

	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/maxwelbm/alkemy-g7.git/internal/repository/interfaces"
)

type ProductService struct {
	ProductRepository interfaces.IProductsRepo
}

func NewProductService(repo interfaces.IProductsRepo) *ProductService {
	return &ProductService{ProductRepository: repo}
}

func (ps *ProductService) GetAllProducts() ([]model.Product, error) {
	products, err := ps.ProductRepository.GetAll()
	var productSlice []model.Product
	if err != nil {
		return productSlice, err
	}

	for _, product := range products {
		productSlice = append(productSlice, product)
	}
	

	return productSlice, nil
}

func (ps *ProductService) GetProductById(id int) (model.Product, error) {
	product, err := ps.ProductRepository.GetById(id)
	if err != nil {
		return model.Product{}, err
	}
	return product, nil
}

func (ps *ProductService) CreateProduct(product model.Product) (model.Product, error) {
	err := product.Validate()

	if err != nil {
		return model.Product{}, err
	}

	productsList, _ := ps.ProductRepository.GetAll()
	existsByCode := existsByProductCode(product.ProductCode, productsList)

	if existsByCode {
		return model.Product{}, errors.New("já existe um produto com esse código")
	}

	productDb, err := ps.ProductRepository.Create(product)

	if err != nil {
		return model.Product{}, err
	}
	return productDb, nil
}

func (ps *ProductService) UpdateProduct(id int, product model.Product) (model.Product, error) {
	return model.Product{}, nil
}

func (ps *ProductService) DeleteProduct(id int) error {
	err := ps.ProductRepository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func existsByProductCode(productCode string, products map[int]model.Product) bool {
	for _, product := range products {
		if product.ProductCode == productCode {
			return true
		}
	}
	return false
}
