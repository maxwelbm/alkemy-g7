package interfaces

import (
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
)

type ILocalityService interface {
	GetSellers(id int) (locality model.LocalitiesJSONSellers, err error)
	GetCarries(id int) (locality model.LocalitiesJSONCarries, err error)
	GetById(id int) (locality model.Locality, err error)
	CreateLocality(locality *model.Locality) (l model.Locality, err error)
}
