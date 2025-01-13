package service

import (
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/maxwelbm/alkemy-g7.git/internal/repository/interfaces"
)

func CreateServiceLocalities(rp interfaces.ILocalityRepo) *LocalitiesService {
	return &LocalitiesService{rp: rp}
}

func (s *LocalitiesService) validateEmptyFields(l model.Locality) error {
	if l.ID == "" || l.Locality == "" || l.Province == "" {
		return model.ErrorNullLocalityAttribute
	}
	return nil
}

type LocalitiesService struct {
	rp interfaces.ILocalityRepo
}

func (s *LocalitiesService) GetCarries(id int) (locality model.LocalitiesJSONCarries, err error) {
	panic("unimplemented")
}

func (s *LocalitiesService) GetSellers(id int) (locality model.LocalitiesJSONSellers, err error) {
	panic("unimplemented")
}

func (s *LocalitiesService) GetById(id int) (locality model.Locality, err error) {
	locality, err = s.rp.GetById(id)
	return
}

func (s *LocalitiesService) CreateLocality(locality *model.Locality) (l model.Locality, err error) {
	if err := s.validateEmptyFields(*locality); err != nil {
		return l, err
	}
	l, err = s.rp.Post(locality)
	return
}
