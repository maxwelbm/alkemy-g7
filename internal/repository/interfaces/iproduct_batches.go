package interfaces

import "github.com/maxwelbm/alkemy-g7.git/internal/model"

type IProductBatchesRepo interface {
	GetByID(id int) (model.ProductBatches, error)
	Post(prodBatches *model.ProductBatches) (model.ProductBatches, error)
}
