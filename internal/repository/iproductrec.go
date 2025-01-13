package repository

import "github.com/maxwelbm/alkemy-g7.git/internal/model"

type IProductRecService interface {
	CreateProductRecords(pr model.ProductRecords) (model.ProductRecords, error)
	GelAllProductRecordsReport() ([]model.ProductRecordsReport, error)
}