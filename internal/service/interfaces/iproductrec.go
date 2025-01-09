package interfaces

import "github.com/maxwelbm/alkemy-g7.git/internal/model"

type IProductRecRepo interface {
	CreateProductRecords(pr model.ProductRecords) (model.ProductRecords, error)
	GelAllProductRecordsReport() ([]model.ProductRecordsReport, error)

}