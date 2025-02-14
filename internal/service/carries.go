package service

import (
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/maxwelbm/alkemy-g7.git/internal/repository/interfaces"
	svc "github.com/maxwelbm/alkemy-g7.git/internal/service/interfaces"
)

type CarrierDefault struct {
	Rp          interfaces.ICarriersRepo
	SvcLocality svc.ILocalityService
}

func NewCarrierService(rp interfaces.ICarriersRepo, svcLocality svc.ILocalityService) *CarrierDefault {
	return &CarrierDefault{
		Rp:          rp,
		SvcLocality: svcLocality,
	}
}

func (cp *CarrierDefault) GetByID(id int) (carrier model.Carries, err error) {
	carrier, err = cp.Rp.GetByID(id)
	return
}

func (cp *CarrierDefault) PostCarrier(newCarrier model.Carries) (carrier model.Carries, err error) {
	_, err = cp.SvcLocality.GetByID(newCarrier.LocalityID)

	if err != nil {
		return
	}

	id, err := cp.Rp.PostCarrier(newCarrier)

	if err != nil {
		return
	}

	carrier, err = cp.GetByID(int(id))

	return
}
