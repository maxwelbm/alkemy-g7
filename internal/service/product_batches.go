package service

import (
	"fmt"

	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	irepo "github.com/maxwelbm/alkemy-g7.git/internal/repository/interfaces"
	"github.com/maxwelbm/alkemy-g7.git/internal/service/interfaces"
	"github.com/maxwelbm/alkemy-g7.git/pkg/customerror"
	"github.com/maxwelbm/alkemy-g7.git/pkg/logger"
)

type ProductBatchesService struct {
	Rp      irepo.IProductBatchesRepo
	SvcProd interfaces.IProductService
	SvcSec  interfaces.ISectionService
	log     logger.Logger
}

func CreateProductBatchesService(rp irepo.IProductBatchesRepo, SvcProd interfaces.IProductService, SvcSec interfaces.ISectionService, log logger.Logger) *ProductBatchesService {
	return &ProductBatchesService{Rp: rp, SvcProd: SvcProd, SvcSec: SvcSec, log: log}
}

func (s *ProductBatchesService) GetByID(id int) (prodBatches model.ProductBatches, err error) {
	s.log.Log("ProductBatchesService", "INFO", "initializing GetByID function with id parameter")
	prodBatches, err = s.Rp.GetByID(id)

	return
}

func (s *ProductBatchesService) Post(prodBatches *model.ProductBatches) (newProdBatches model.ProductBatches, err error) {
	s.log.Log("ProductBatchesService", "INFO", "initializing Post function with prodBatches parameter")

	if err = prodBatches.Validate(); err != nil {
		s.log.Log("ProductBatchesService", "ERROR", fmt.Sprintf("Error: %v", err))

		return model.ProductBatches{}, customerror.HandleError("product batches", customerror.ErrorInvalid, err.Error())
	}

	_, err = s.SvcProd.GetProductByID(prodBatches.ProductID)
	if err != nil {
		s.log.Log("ProductBatchesService", "ERROR", fmt.Sprintf("Error: %v", err))

		return
	}

	_, err = s.SvcSec.GetByID(prodBatches.SectionID)
	if err != nil {
		s.log.Log("ProductBatchesService", "ERROR", fmt.Sprintf("Error: %v", err))

		return
	}

	newProdBatches, err = s.Rp.Post(prodBatches)

	s.log.Log("ProductBatchesService", "INFO", "successfully executed post function")

	return
}
