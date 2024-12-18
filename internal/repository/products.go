package repository

import (
	"errors"

	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/maxwelbm/alkemy-g7.git/pkg/database"
)

type ProductRepository struct {
	DB database.Database
}

func NewProductRepository(db database.Database) *ProductRepository {
	return &ProductRepository{DB: db}
}

func (pr *ProductRepository) GetAll() (map[int]model.Product, error) {
	productList := pr.DB.TbProducts
	return productList, nil
}

func (pr *ProductRepository) GetById(id int) (model.Product, error) {

	product, exists := pr.DB.TbProducts[id]

	if !exists {
		return product, errors.New("produto não encontrado")
	}
	return product, nil
}

func (pr *ProductRepository) Create(product model.Product) (model.Product, error) {
	return model.Product{}, nil
}

func (pr *ProductRepository) Update(id int, product model.Product) (model.Product, error) {
	return model.Product{}, nil
}

func (pr *ProductRepository) Delete(id int) error {
	return nil
}
