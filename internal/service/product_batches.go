package service

import (
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/maxwelbm/alkemy-g7.git/internal/repository"
	"github.com/maxwelbm/alkemy-g7.git/internal/service/interfaces"
	"github.com/maxwelbm/alkemy-g7.git/pkg/custom_error"
)

type ProductBatchesService struct {
	rp      repository.ProductBatchesRepository
	svcProd interfaces.IProductService
	svcSec  interfaces.ISectionService
}

func CreateProductBatchesService(rp repository.ProductBatchesRepository, svcProd interfaces.IProductService, svcSec interfaces.ISectionService) *ProductBatchesService {
	return &ProductBatchesService{rp: rp, svcProd: svcProd, svcSec: svcSec}
}

func (s *ProductBatchesService) GetById(id int) (prodBatches model.ProductBatches, err error) {
	prodBatches, err = s.rp.GetById(id)
	return
}

func (s *ProductBatchesService) Post(prodBatches *model.ProductBatches) (newProdBatches model.ProductBatches, err error) {
	if err = prodBatches.Validate(); err != nil {
		return model.ProductBatches{}, custom_error.HandleError("product batches", custom_error.ErrorInvalid, err.Error())
	}
	_, err = s.svcProd.GetProductByID(prodBatches.ProductID)
	if err != nil {
		return
	}

	_, err = s.svcSec.GetById(prodBatches.SectionID)
	if err != nil {
		return
	}

	newProdBatches, err = s.rp.Post(prodBatches)
	return
}
