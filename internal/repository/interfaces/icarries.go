package interfaces

import "github.com/maxwelbm/alkemy-g7.git/internal/model"

type ICarriersRepo interface {
	PostCarrier(newCarrier model.Carries) (id int64, err error)
	GetByID(id int) (carrier model.Carries, err error)
}
