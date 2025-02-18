package service

import (
	"fmt"
	"time"

	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	repo "github.com/maxwelbm/alkemy-g7.git/internal/repository/interfaces"
	serv "github.com/maxwelbm/alkemy-g7.git/internal/service/interfaces"
	appErr "github.com/maxwelbm/alkemy-g7.git/pkg/customerror"
	"github.com/maxwelbm/alkemy-g7.git/pkg/logger"
)

type ProductRecService struct {
	ProductRecRepository repo.IProductRecRepository
	ProductSv            serv.IProductService
	log                  logger.Logger
}

func NewProductRecService(productRecRepo repo.IProductRecRepository, productServ serv.IProductService, logger logger.Logger) *ProductRecService {
	return &ProductRecService{
		ProductRecRepository: productRecRepo,
		ProductSv:            productServ,
		log:                  logger,
	}
}

func (prs *ProductRecService) CreateProductRecords(pr model.ProductRecords) (model.ProductRecords, error) {
	prs.log.Log("ProductRecService", "INFO", "CreateProductRecords function initializing")

	if err := pr.Validate(); err != nil {
		prs.log.Log("ProductRecService", "ERROR", "Validation error: "+err.Error())
		return model.ProductRecords{}, appErr.HandleError("product record", appErr.ErrorInvalid, err.Error())
	}

	if _, err := prs.ProductSv.GetProductByID(pr.ProductID); err != nil {
		prs.log.Log("ProductRecService", "ERROR", "Product not found with ID: "+string(pr.ProductID))
		return model.ProductRecords{}, err
	}

	pr.LastUpdateDate = time.Now()
	prs.log.Log("ProductRecService", "INFO", "Creating product record for Product ID: "+string(pr.ProductID))

	productRecord, err := prs.ProductRecRepository.Create(pr)

	if err != nil {
		prs.log.Log("ProductRecService", "ERROR", "Error creating product record: "+err.Error())
		return model.ProductRecords{}, err
	}

	prs.log.Log("ProductRecService", "INFO", fmt.Sprintf("Product record created successfully: %+v", productRecord))
	return productRecord, nil
}

func (prs *ProductRecService) GetProductRecordByID(id int) (model.ProductRecords, error) {
	prs.log.Log("ProductRecService", "INFO", "GetProductRecordByID function initializing for ID: "+string(id))

	productRecord, err := prs.ProductRecRepository.GetByID(id)
	if err != nil {
		prs.log.Log("ProductRecService", "ERROR", "Error retrieving product record with ID: "+string(id)+", error: "+err.Error())
		return model.ProductRecords{}, err
	}

	prs.log.Log("ProductRecService", "INFO", fmt.Sprintf("Retrieved product record: %v", productRecord))
	return productRecord, nil
}

func (prs *ProductRecService) GetProductRecordReport(idProduct int) ([]model.ProductRecordsReport, error) {
	prs.log.Log("ProductRecService", "INFO", "GetProductRecordReport function initializing for ProductID: "+string(idProduct))

	allReports, err := prs.ProductRecRepository.GetAllReport()
	if err != nil {
		prs.log.Log("ProductRecService", "ERROR", "Error retrieving product record reports: "+err.Error())
		return nil, err
	}

	var filteredReports []model.ProductRecordsReport

	if idProduct == 0 {
		prs.log.Log("ProductRecService", "INFO", "No ProductID provided, returning all reports.")
		return allReports, nil
	}

	if _, err := prs.ProductSv.GetProductByID(idProduct); err != nil {
		prs.log.Log("ProductRecService", "ERROR", "Product not found with ID: "+string(idProduct))
		return filteredReports, err
	}

	for _, report := range allReports {
		if report.ProductID == idProduct {
			filteredReports = append(filteredReports, report)
		}
	}

	prs.log.Log("ProductRecService", "INFO", fmt.Sprintf("Filtered reports for ProductID: %d, Count: %d", idProduct, len(filteredReports)))
	return filteredReports, nil
}
