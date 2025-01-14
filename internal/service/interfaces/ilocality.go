package interfaces

import (
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
)

type ILocalityService interface {
	GetSellers(id int) (report []model.LocalitiesJSONSellers, err error)
	GetCarriers(id int) (report []model.LocalitiesJSONCarriers, err error)
	GetById(id int) (locality model.Locality, err error)
	CreateLocality(locality *model.Locality) (l model.Locality, err error)
}
