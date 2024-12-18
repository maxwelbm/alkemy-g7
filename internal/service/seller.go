package service

import (
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/maxwelbm/alkemy-g7.git/internal/repository"
)

func CreateServiceSellers(rp repository.SellersRepository) *SellersService {
	return &SellersService{rp: rp}
}

type SellersService struct {
	rp repository.SellersRepository
}

func (s *SellersService) GetAll() (v map[int]model.Seller, err error) {
	v, err = s.rp.Get()
	return
}

func (s *SellersService) GetByID(id int) (seller model.Seller, err error) {
	seller, err = s.rp.GetByID(id)
	return
}