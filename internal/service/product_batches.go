package service

import (
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/maxwelbm/alkemy-g7.git/internal/repository"
)

type ProductBatchesService struct {
	rp repository.ProductBatchesRepository
}

func CreateProductBatchesService(rp repository.ProductBatchesRepository) *ProductBatchesService {
	return &ProductBatchesService{rp: rp}
}

func (s *ProductBatchesService) GetById(id int) (prodBatches model.ProductBatches, err error) {
	prodBatches, err = s.rp.GetById(id)
	return
}

func (s *ProductBatchesService) Post(prodBatches *model.ProductBatches) (newProdBatches model.ProductBatches, err error) {
	newProdBatches, err = s.rp.Post(prodBatches)
	return
}
