package interfaces

import "github.com/maxwelbm/alkemy-g7.git/internal/model"

type IProductBatchesRepo interface {
	GetById(id int) (model.ProductBatches, error)
	Post(prodBatches *model.ProductBatches) (model.ProductBatches, error)
}
