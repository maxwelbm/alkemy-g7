package service

import (
	"fmt"
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/maxwelbm/alkemy-g7.git/internal/repository/interfaces"
	"github.com/maxwelbm/alkemy-g7.git/pkg/logger"
)

func CreateServiceLocalities(rp interfaces.ILocalityRepo, log logger.Logger) *LocalitiesService {
	return &LocalitiesService{Rp: rp, log: log}
}

type LocalitiesService struct {
	Rp  interfaces.ILocalityRepo
	log logger.Logger
}

func (s *LocalitiesService) GetSellers(id int) (report []model.LocalitiesJSONSellers, err error) {
	if id != 0 {
		report, err = s.Rp.GetReportSellersWithID(id)
		return
	}

	report, err = s.Rp.GetSellers(id)

	s.log.Log("LocalitiesService", "INFO", fmt.Sprintf("Retrieved report sellers: %+v", report))

	return
}

func (s *LocalitiesService) GetCarriers(id int) (report []model.LocalitiesJSONCarriers, err error) {
	if id != 0 {
		report, err = s.Rp.GetReportCarriersWithID(id)
		return
	}

	report, err = s.Rp.GetCarriers(id)

	s.log.Log("LocalitiesService", "INFO", fmt.Sprintf("Retrieved report carriers: %+v", report))

	return
}

func (s *LocalitiesService) GetByID(id int) (locality model.Locality, err error) {
	locality, err = s.Rp.GetByID(id)

	s.log.Log("LocalitiesService", "INFO", fmt.Sprintf("Retrieved locality by ID: %+v", locality))

	return
}

func (s *LocalitiesService) CreateLocality(locality *model.Locality) (l model.Locality, err error) {
	if err := locality.ValidateEmptyFields(locality); err != nil {
		s.log.Log("LocalitiesService", "ERROR", fmt.Sprintf("Error: %+v", err))

		return l, err
	}

	l, err = s.Rp.CreateLocality(locality)

	s.log.Log("LocalitiesService", "INFO", fmt.Sprintf("Created locality: %+v", l))

	return
}
