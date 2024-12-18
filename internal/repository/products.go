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
	id := getLastIdProduct(pr.DB.TbProducts)
	product.ID = id
	pr.DB.TbProducts[product.ID] = product
	return pr.DB.TbProducts[product.ID], nil

}

func (pr *ProductRepository) Update(id int, product model.Product) (model.Product, error) {
	return model.Product{}, nil
}

func (pr *ProductRepository) Delete(id int) error {
	_, exists := pr.DB.TbProducts[id]

	if !exists {
		return errors.New("produto não encontrado")
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
