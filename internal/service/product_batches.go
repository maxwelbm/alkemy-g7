package service

import (
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	irepo "github.com/maxwelbm/alkemy-g7.git/internal/repository/interfaces"
	"github.com/maxwelbm/alkemy-g7.git/internal/service/interfaces"
	"github.com/maxwelbm/alkemy-g7.git/pkg/customerror"
)

type ProductBatchesService struct {
	Rp      irepo.IProductBatchesRepo
	SvcProd interfaces.IProductService
	SvcSec  interfaces.ISectionService
}

func CreateProductBatchesService(rp irepo.IProductBatchesRepo, SvcProd interfaces.IProductService, SvcSec interfaces.ISectionService) *ProductBatchesService {
	return &ProductBatchesService{Rp: rp, SvcProd: SvcProd, SvcSec: SvcSec}
}

func (s *ProductBatchesService) GetByID(id int) (prodBatches model.ProductBatches, err error) {
	prodBatches, err = s.Rp.GetByID(id)
	return
}

func (s *ProductBatchesService) Post(prodBatches *model.ProductBatches) (newProdBatches model.ProductBatches, err error) {
	if err = prodBatches.Validate(); err != nil {
		return model.ProductBatches{}, customerror.HandleError("product batches", customerror.ErrorInvalid, err.Error())
	}

	_, err = s.SvcProd.GetProductByID(prodBatches.ProductID)
	if err != nil {
		return
	}

	_, err = s.SvcSec.GetByID(prodBatches.SectionID)
	if err != nil {
		return
	}

	newProdBatches, err = s.Rp.Post(prodBatches)

	return
}
