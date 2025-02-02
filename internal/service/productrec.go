package service

import (
	"time"

	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	repo "github.com/maxwelbm/alkemy-g7.git/internal/repository/interfaces"
	serv "github.com/maxwelbm/alkemy-g7.git/internal/service/interfaces"
	appErr "github.com/maxwelbm/alkemy-g7.git/pkg/customerror"
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
		return model.ProductRecords{}, appErr.HandleError("product record", appErr.ErrorInvalid, err.Error())
	}

	if _, err := prs.ProductSv.GetProductByID(pr.ProductID); err != nil {
		return model.ProductRecords{}, err
	}

	pr.LastUpdateDate = time.Now()

	productRecord, err := prs.ProductRecRepository.Create(pr)

	if err != nil {
		return model.ProductRecords{}, err
	}

	return productRecord, nil
}

func (prs *ProductRecService) GetProductRecordByID(id int) (model.ProductRecords, error) {
	productRecord, err := prs.ProductRecRepository.GetByID(id)
	if err != nil {
		return model.ProductRecords{}, err
	}

	return productRecord, nil
}

func (prs *ProductRecService) GetProductRecordReport(idProduct int) ([]model.ProductRecordsReport, error) {
	allReports, err := prs.ProductRecRepository.GetAllReport()

	var filteredReports []model.ProductRecordsReport

	if err != nil {
		return nil, err
	}

	if idProduct == 0 {
		return allReports, nil
	}

	if _, err := prs.ProductSv.GetProductByID(idProduct); err != nil {
		return filteredReports, err
	}

	for _, report := range allReports {
		if report.ProductID == idProduct {
			filteredReports = append(filteredReports, report)
		}
	}

	return filteredReports, nil
}
