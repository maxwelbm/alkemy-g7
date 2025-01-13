package interfaces

import (
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
)

type ILocalityRepo interface {
	GetSellers(id int) (locality model.LocalitiesJSONSellers, err error)
	GetCarries(id int) (locality model.LocalitiesJSONCarries, err error)
	GetById(id int) (model.Locality, error)
	Post(l *model.Locality) (model.Locality, error)
}
