package repository

import (
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/maxwelbm/alkemy-g7.git/pkg/custom_error"
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
		return product, custom_error.CustomError{Object: id, Err: custom_error.NotFound}
	}
	return product, nil
}

func (pr *ProductRepository) Create(product model.Product) (model.Product, error) {
	id := getLastIdProduct(pr.DB.TbProducts)
	product.ID = id + 1
	pr.DB.TbProducts[product.ID] = product
	return pr.DB.TbProducts[product.ID], nil

}

func (pr *ProductRepository) Update(id int, product model.Product) (model.Product, error) {
	pr.DB.TbProducts[id] = product
	productUpdated := pr.DB.TbProducts[id]
	return productUpdated, nil
}

func (pr *ProductRepository) Delete(id int) error {
	_, exists := pr.DB.TbProducts[id]

	if !exists {
		return custom_error.CustomError{Object: id, Err: custom_error.NotFound}
	}

	delete(pr.DB.TbProducts, id)
	return nil
}

func getLastIdProduct(products map[int]model.Product) (lastId int) {
	lastId = 0

	for _, product := range products {
		if product.ID > lastId {
			lastId = product.ID
		}

	}
	return

}
