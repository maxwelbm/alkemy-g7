package service

import (
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/maxwelbm/alkemy-g7.git/internal/repository/interfaces"
)

func CreateServiceLocalities(rp interfaces.ILocalityRepo) *LocalitiesService {
	return &LocalitiesService{rp: rp}
}

func (s *LocalitiesService) validateEmptyFields(l model.Locality) error {
	if l.Locality == "" || l.Province == "" || l.Country == "" {
		return model.ErrorNullLocalityAttribute
	}
	return nil
}

type LocalitiesService struct {
	rp interfaces.ILocalityRepo
}

func (s *LocalitiesService) GetSellers(id int) (report []model.LocalitiesJSONSellers, err error) {
	if id != 0 {
		report, err = s.rp.GetReportSellersWithId(id)
		return
	}

	report, err = s.rp.GetSellers(id)
	return
}

func (s *LocalitiesService) GetCarriers(id int) (report []model.LocalitiesJSONCarriers, err error) {
	if id != 0 {
		report, err = s.rp.GetReportCarriersWithId(id)
		return
	}

	report, err = s.rp.GetCarriers(id)
	return
}

func (s *LocalitiesService) GetById(id int) (locality model.Locality, err error) {
	locality, err = s.rp.GetById(id)
	return
}

func (s *LocalitiesService) CreateLocality(locality *model.Locality) (l model.Locality, err error) {
	if err := s.validateEmptyFields(*locality); err != nil {
		return l, err
	}
	l, err = s.rp.CreateLocality(locality)
	return
}
