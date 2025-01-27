package interfaces

import "github.com/maxwelbm/alkemy-g7.git/internal/model"

type IProductsRepo interface {
	GetAll() (map[int]model.Product, error)
	GetByID(id int) (model.Product, error)
	Create(product model.Product) (model.Product, error)
	Update(id int, product model.Product) (model.Product, error)
	Delete(id int) error
}
