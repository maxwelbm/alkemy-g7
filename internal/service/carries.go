package service

import (
	"fmt"

	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/maxwelbm/alkemy-g7.git/internal/repository/interfaces"
	svc "github.com/maxwelbm/alkemy-g7.git/internal/service/interfaces"
	"github.com/maxwelbm/alkemy-g7.git/pkg/logger"
)

type CarrierDefault struct {
	Rp          interfaces.ICarriersRepo
	SvcLocality svc.ILocalityService
	log         logger.Logger
}

func NewCarrierService(rp interfaces.ICarriersRepo, svcLocality svc.ILocalityService, log logger.Logger) *CarrierDefault {
	return &CarrierDefault{
		Rp:          rp,
		SvcLocality: svcLocality,
		log:         log,
	}
}

func (cp *CarrierDefault) GetByID(id int) (carrier model.Carries, err error) {
	cp.log.Log("CarrierService", "INFO", "initializing GetByID function")
	carrier, err = cp.Rp.GetByID(id)
	cp.log.Log("CarrierService", "INFO", "GetByIDCarrier completed successfully")
	return
}

func (cp *CarrierDefault) PostCarrier(newCarrier model.Carries) (carrier model.Carries, err error) {
	cp.log.Log("CarrierService", "INFO", "initializing PostCarrier function")
	_, err = cp.SvcLocality.GetByID(newCarrier.LocalityID)

	if err != nil {
		cp.log.Log("CarrierService", "ERROR", fmt.Sprintf("Error: %v", err))
		return
	}

	id, err := cp.Rp.PostCarrier(newCarrier)

	if err != nil {
		cp.log.Log("CarrierService", "ERROR", fmt.Sprintf("Error: %v", err))
		return
	}

	carrier, err = cp.GetByID(int(id))

	cp.log.Log("CarrierService", "INFO", "PostCarrier completed successfully")
	return
}
