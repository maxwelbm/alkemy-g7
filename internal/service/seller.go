package service

import (
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
	var send bool
    if sl.CID != nil {
		if *sl.CID == 0 {
			return model.ErrorAttribute
		} else {
			send = true
		}
	}

    if sl.Address != nil {
		if *sl.Address == "" {
        	return model.ErrorAttribute
    	} else {
			send = true
		}
	}

    if sl.CompanyName != nil {
		if *sl.CompanyName == "" {
			return model.ErrorAttribute    
		} else {
			send = true
		}
	}

    if sl.Telephone != nil {
		if *sl.Telephone == "" {
			return model.ErrorAttribute    
		}
	} 

	if !send {
		return model.ErrorAttribute
	}

    return nil
}

func (s *SellersService) ValidateFieldsCreate(seller model.Seller) error {
	fieldsToValidate := []struct {
		value interface{}
		validateFunc func(attribute interface{}, t string) error
	}{
		{seller.CompanyName, validateFormatReflect},
		{seller.Address, validateFormatReflect},
		{seller.Telephone, validateFormatReflect},
		{seller.CID, validateFormatReflect},
	}

	for _, field := range fieldsToValidate {
		if err := field.validateFunc(field.value, reflect.TypeOf(field.value).String()); err != nil {
			return err
		}
	}
	return nil
}

func validateFormatReflect(attribute interface{}, t string) error {
	switch t {
	case reflect.String.String():
		if reflect.TypeOf(attribute).Kind() != reflect.String {
			return model.ErrorAttribute
		}
		if attribute == "" {
			return model.ErrorAttribute
		}
	case reflect.Int.String():
		if reflect.TypeOf(attribute).Kind() != reflect.Int {
			return model.ErrorAttribute
		}
		if attribute == 0 {
			return model.ErrorAttribute
		}
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
	if err := s.ValidateFieldsCreate(seller); err != nil {
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

func (s *SellersService) DeleteSeller(id int) error {
	err := s.rp.Delete(id)
	return err
}