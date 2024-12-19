package service

import (
	"errors"
	"reflect"
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/maxwelbm/alkemy-g7.git/internal/repository/interfaces"
)

func CreateServiceSellers(rp interfaces.ISellerRepo) *SellersService {
	return &SellersService{rp: rp}
}

type SellersService struct {
	rp interfaces.ISellerRepo
}

func (s *SellersService) ValidateFieldsUpdate(sl model.SellerUpdate) error {
    if sl.CID != nil && *sl.CID == 0 {
        return validateFormatInt(*sl.CID)
    }

    if sl.Address != nil && *sl.Address == "" {
        return validateFormatString(*sl.Address)
    }

    if sl.CompanyName != nil && *sl.CompanyName == "" {
        return validateFormatString(*sl.CompanyName)
    }

    if sl.Telephone != nil && *sl.Telephone == "" {
        return validateFormatString(*sl.Telephone)
    }

    return nil
}

func (s *SellersService) ValidateFields(seller model.Seller) error {
	if err := validateFormatString(seller.CompanyName); err != nil {
		return err
	} else if err := validateFormatString(seller.Address); err != nil {
		return err
	} else if err := validateFormatString(seller.Telephone); err != nil {
		return err
	} else if err := validateFormatInt(seller.CID); err != nil {
		return err
	}
	return nil
}

func validateFormatString(attribute string) error {
	err := errors.New("Invalid format or empty value, expected string attribute.")

	if reflect.TypeOf(attribute).Kind() != reflect.String {
		return err
	}
	if attribute == "" {
		return err
	}
	return nil
}

func validateFormatInt(attribute int) error {
	err := errors.New("Invalid format or empty value, expected int attribute.")

	if reflect.TypeOf(attribute).Kind() != reflect.Int {
		return err
	}
	if attribute == 0 {
		return err
	}
	return nil
}

func (s *SellersService) GetAll() (sellers []model.Seller, err error) {
	sellers, err = s.rp.Get()
	return
}

func (s *SellersService) GetById(id int) (seller model.Seller, err error) {
	seller, err = s.rp.GetById(id)
	return
}

func (s *SellersService) CreateSeller(seller model.Seller) (sl model.Seller, err error) {
	if err := s.ValidateFields(seller); err != nil {
		return sl, err
	}

	sl, err = s.rp.Post(seller)
	return
}

func (s *SellersService) UpdateSeller(id int, seller model.SellerUpdate) (sl model.Seller, err error) {
    if err := s.ValidateFieldsUpdate(seller); err != nil {
        return sl, err
    }

    sl, err = s.rp.Patch(id, seller)

    return sl, err
}

func (s *SellersService) Delete(id int) error {
	panic("unimplemented")
}