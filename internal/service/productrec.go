package service

import (
	"time"

	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	repo "github.com/maxwelbm/alkemy-g7.git/internal/repository/interfaces"
	serv "github.com/maxwelbm/alkemy-g7.git/internal/service/interfaces"
)

type ProductRecService struct {
	ProductRecRepository repo.IProductRecRepository
	ProductSv            serv.IProductService
}

func NewProductRecService(productRecRepo repo.IProductRecRepository, productServ serv.IProductService) *ProductRecService {
	return &ProductRecService{
		ProductRecRepository: productRecRepo,
		ProductSv:            productServ,
	}
}

func (prs *ProductRecService) CreateProductRecords(pr model.ProductRecords) (model.ProductRecords, error) {
	if err := pr.Validate(); err != nil {
		return model.ProductRecords{}, err
	}

	if _, err := prs.ProductSv.GetProductById(pr.ProductId); err != nil {
		return model.ProductRecords{}, err
	}

	pr.LastUpdateDate = time.Now()

	productRecord, err := prs.ProductRecRepository.Create(pr)
	if err != nil {
		return model.ProductRecords{}, err
	}

	return productRecord, nil
}

func (prs *ProductRecService) GelAllProductRecordsReport() ([]model.ProductRecordsReport, error) {
	return nil, nil
}
