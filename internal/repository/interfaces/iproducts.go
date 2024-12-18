package interfaces

import "github.com/maxwelbm/alkemy-g7.git/internal/model"

type IProductsRepo interface {
	Get() (map[int]model.Products, error)
	GetById(id int) (model.Products, error)
	Post(product model.Products) (model.Products, error)
	Update(id int, product model.Products) (model.Products, error)
	Delete(id int) error
}
