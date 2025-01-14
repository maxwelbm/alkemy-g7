package service

import (
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/maxwelbm/alkemy-g7.git/internal/repository/interfaces"
	svc "github.com/maxwelbm/alkemy-g7.git/internal/service/interfaces"
)

type CarrierDefault struct {
	rp          interfaces.ICarriersRepo
	svcLocality svc.ILocalityService
}

func NewCarrierService(rp interfaces.ICarriersRepo, svcLocality svc.ILocalityService) *CarrierDefault {
	return &CarrierDefault{
		rp:          rp,
		svcLocality: svcLocality,
	}
}

func (cp *CarrierDefault) GetById(id int) (carrier model.Carries, err error) {
	carrier, err = cp.rp.GetById(id)
	return
}

func (cp *CarrierDefault) PostCarrier(newCarrier model.Carries) (carrier model.Carries, err error) {

	_, err = cp.svcLocality.GetById(newCarrier.LocalityId)

	if err != nil {
		return
	}

	id, err := cp.rp.PostCarrier(newCarrier)

	if err != nil {
		return
	}

	carrier, err = cp.GetById(int(id))

	return
}
