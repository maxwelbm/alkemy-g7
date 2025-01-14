package interfaces

import (
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
)

type ILocalityRepo interface {
	GetReportSellersWithId(id int) (locality []model.LocalitiesJSONSellers, err error)
	GetSellers(id int) (report []model.LocalitiesJSONSellers, err error)
	GetReportCarriersWithId(id int) (locality []model.LocalitiesJSONCarriers, err error)
	GetCarriers(id int) (report []model.LocalitiesJSONCarriers, err error)
	GetById(id int) (model.Locality, error)
	CreateLocality(l *model.Locality) (model.Locality, error)
}
