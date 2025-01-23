package service

import (
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/maxwelbm/alkemy-g7.git/internal/repository/interfaces"
	serviceInterface "github.com/maxwelbm/alkemy-g7.git/internal/service/interfaces"
)

func CreateServiceSellers(rp interfaces.ISellerRepo, rpl serviceInterface.ILocalityService) *SellersService {
	return &SellersService{rp: rp, rpl: rpl}
}

type SellersService struct {
	rp  interfaces.ISellerRepo
	rpl serviceInterface.ILocalityService
}

func (s *SellersService) GetAll() (sellers []model.Seller, err error) {
	sellers, err = s.rp.Get()
	return
}

func (s *SellersService) GetByID(id int) (seller model.Seller, err error) {
	seller, err = s.rp.GetByID(id)
	return
}

func (s *SellersService) CreateSeller(seller *model.Seller) (sl model.Seller, err error) {
	if err := seller.ValidateEmptyFields(seller); err != nil {
		return sl, err
	}

	_, err = s.rpl.GetByID(seller.Locality)
	if err != nil {
		return
	}

	sl, err = s.rp.Post(seller)
	
	return
}

func (s *SellersService) UpdateSeller(id int, seller *model.Seller) (sl model.Seller, err error) {
	if seller.Locality != 0 {
		_, err := s.rpl.GetByID(seller.Locality)
		if err != nil {
			return sl, err
		}
	}

	existSl, _ := s.GetByID(id)
	err = seller.ValidateUpdateFields(seller, &existSl)
	
	if err != nil {
		return sl, err
	}
	
	sl, err = s.rp.Patch(id, seller)
	
	return sl, err
}

func (s *SellersService) DeleteSeller(id int) error {
	err := s.rp.Delete(id)
	return err
}
