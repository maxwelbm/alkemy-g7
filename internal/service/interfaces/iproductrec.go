package interfaces

import "github.com/maxwelbm/alkemy-g7.git/internal/model"

type IProductRecRepo interface {
	CreateProductRecords(pr model.ProductRecords) (model.ProductRecords, error)
	GetProductRecordById(id int) (model.ProductRecords, error)
	GetProductRecordReport(idProduct int) (model.ProductRecordsReport, error)

}