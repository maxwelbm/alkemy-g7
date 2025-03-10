package interfaces

import "github.com/maxwelbm/alkemy-g7.git/internal/model"

type IProductService interface {
	GetAllProducts() ([]model.Product, error)
	GetProductByID(id int) (model.Product, error)
	CreateProduct(product model.Product) (model.Product, error)
	UpdateProduct(id int, product model.Product) (model.Product, error)
	DeleteProduct(id int) error 
}