package interfaces

import "github.com/maxwelbm/alkemy-g7.git/internal/model"

type IProductRecService interface {
	CreateProductRecords(pr model.ProductRecords) (model.ProductRecords, error)
	GetProductRecordByID(id int) (model.ProductRecords, error)
	GetProductRecordReport(idProduct int) ([]model.ProductRecordsReport, error)
}
