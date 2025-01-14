package interfaces

import "github.com/maxwelbm/alkemy-g7.git/internal/model"

type ICarrierService interface {
	PostCarrier(newCarrier model.Carries) (carrier model.Carries, err error)
	GetById(id int) (carrier model.Carries, err error)
}
