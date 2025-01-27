package service

import (
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/maxwelbm/alkemy-g7.git/internal/repository"
	"github.com/maxwelbm/alkemy-g7.git/internal/service/interfaces"
	"github.com/maxwelbm/alkemy-g7.git/pkg/customerror"
)

type ProductBatchesService struct {
	rp      repository.ProductBatchesRepository
	svcProd interfaces.IProductService
	svcSec  interfaces.ISectionService
}

func CreateProductBatchesService(rp repository.ProductBatchesRepository, svcProd interfaces.IProductService, svcSec interfaces.ISectionService) *ProductBatchesService {
	return &ProductBatchesService{rp: rp, svcProd: svcProd, svcSec: svcSec}
}

func (s *ProductBatchesService) GetByID(id int) (prodBatches model.ProductBatches, err error) {
	prodBatches, err = s.rp.GetByID(id)
	return
}

func (s *ProductBatchesService) Post(prodBatches *model.ProductBatches) (newProdBatches model.ProductBatches, err error) {
	if err = prodBatches.Validate(); err != nil {
		return model.ProductBatches{}, customerror.HandleError("product batches", customerror.ErrorInvalid, err.Error())
	}

	_, err = s.svcProd.GetProductByID(prodBatches.ProductID)
	if err != nil {
		return
	}

	_, err = s.svcSec.GetByID(prodBatches.SectionID)
	if err != nil {
		return
	}

	newProdBatches, err = s.rp.Post(prodBatches)

	return
}
