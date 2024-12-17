package service

import (
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/maxwelbm/alkemy-g7.git/internal/repository/interfaces"
)

func CreateServiceSellers(rp interfaces.ISellerRepo) *SellersService {
	return &SellersService{rp: rp}
}

type SellersService struct {
	rp interfaces.ISellerRepo
}

func (s *SellersService) FindAll() (v map[int]model.Seller, err error) {
	v, err = s.rp.Get()
	return
}