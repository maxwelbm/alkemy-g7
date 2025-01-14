package interfaces

import "github.com/maxwelbm/alkemy-g7.git/internal/model"

type IProductRecRepository interface {
	Create(pr model.ProductRecords) (model.ProductRecords, error)
	GetAll() ([]model.ProductRecords, error)
	GetById(id int) (model.ProductRecords, error)
	GetByIdProduct(idProduct int) ([]model.ProductRecords, error)
	GetAllReport() ([]model.ProductRecordsReport, error)
}