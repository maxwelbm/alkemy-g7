package service

import (
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/maxwelbm/alkemy-g7.git/internal/repository/interfaces"
)

func CreateServiceLocalities(rp interfaces.ILocalityRepo) *LocalitiesService {
	return &LocalitiesService{rp: rp}
}

type LocalitiesService struct {
	rp interfaces.ILocalityRepo
}

func (s *LocalitiesService) GetSellers(id int) (report []model.LocalitiesJSONSellers, err error) {
	if id != 0 {
		report, err = s.rp.GetReportSellersWithID(id)
		return
	}

	report, err = s.rp.GetSellers(id)
	
	return
}

func (s *LocalitiesService) GetCarriers(id int) (report []model.LocalitiesJSONCarriers, err error) {
	if id != 0 {
		report, err = s.rp.GetReportCarriersWithID(id)
		return
	}

	report, err = s.rp.GetCarriers(id)
	
	return
}

func (s *LocalitiesService) GetByID(id int) (locality model.Locality, err error) {
	locality, err = s.rp.GetByID(id)
	return
}

func (s *LocalitiesService) CreateLocality(locality *model.Locality) (l model.Locality, err error) {
	if err := locality.ValidateEmptyFields(locality); err != nil {
		return l, err
	}
	
	l, err = s.rp.CreateLocality(locality)
	
	return
}
