package service

import (
	"fmt"
	"time"

	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	repo "github.com/maxwelbm/alkemy-g7.git/internal/repository/interfaces"
	serv "github.com/maxwelbm/alkemy-g7.git/internal/service/interfaces"
	appErr "github.com/maxwelbm/alkemy-g7.git/pkg/custom_error"
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
		return model.ProductRecords{},  appErr.HandleError("product record", appErr.ErrInvalid, err.Error())
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

func (prs *ProductRecService) GetProductRecordById(id int) (model.ProductRecords, error) {
	productRecord, err := prs.ProductRecRepository.GetById(id)
	fmt.Print(id)
	if err != nil {
		return model.ProductRecords{}, err
	}

	return productRecord, nil
}

func (prs *ProductRecService) GetProductRecordReport(idProduct int) (model.ProductRecordsReport, error) {
	var productReport model.ProductRecordsReport
	product, err := prs.ProductSv.GetProductById(idProduct)
	if err != nil {
		return productReport, err
	}

	productRecordList, err := prs.ProductRecRepository.GetByIdProduct(idProduct)
	if err != nil {
		return productReport, err
	}

	productReport.Description = product.Description
	productReport.ProductId = product.ID
	productReport.RecordsCount = len(productRecordList)

	return productReport, nil
}