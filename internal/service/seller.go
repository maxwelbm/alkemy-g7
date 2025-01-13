package service

import (
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/maxwelbm/alkemy-g7.git/internal/repository/interfaces"
)

func CreateServiceSellers(rp interfaces.ISellerRepo) *SellersService {
	return &SellersService{rp: rp}
}

func (s *SellersService) validateUpdateFields(sl *model.Seller, existSeller *model.Seller) {
	if sl.CID == 0 {
		sl.CID = existSeller.CID
	}
	if sl.Address == "" {
		sl.Address = existSeller.Address
	}
	if sl.CompanyName == "" {
		sl.CompanyName = existSeller.CompanyName
	}
	if sl.Telephone == "" {
		sl.Telephone = existSeller.Telephone
	}
}

func (s *SellersService) validateEmptyFields(sl model.Seller) error {
	if sl.CID == 0 || sl.Address == "" || sl.CompanyName == "" || sl.Telephone == "" {
		return model.ErrorNullAttribute
	}
	return nil
}

type SellersService struct {
	rp interfaces.ISellerRepo
}

func (s *SellersService) GetAll() (sellers []model.Seller, err error) {
	sellers, err = s.rp.Get()
	return
}

func (s *SellersService) GetById(id int) (seller model.Seller, err error) {
	seller, err = s.rp.GetById(id)
	return
}

func (s *SellersService) CreateSeller(seller *model.Seller) (sl model.Seller, err error) {
	if err := s.validateEmptyFields(*seller); err != nil {
		return sl, err
	}
	sl, err = s.rp.Post(seller)
	return
}

func (s *SellersService) UpdateSeller(id int, seller *model.Seller) (sl model.Seller, err error) {
	existSl, err := s.GetById(id)
	if err != nil {
		return
	}

	s.validateUpdateFields(seller, &existSl)
	sl, err = s.rp.Patch(id, seller)
	return sl, err
}

func (s *SellersService) DeleteSeller(id int) error {
	_, err := s.GetById(id)
	if err != nil {
		return err
	}

	err = s.rp.Delete(id)
	return err
}
