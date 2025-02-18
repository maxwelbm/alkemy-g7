package service

import (
	"fmt"
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/maxwelbm/alkemy-g7.git/internal/repository/interfaces"
	serviceInterface "github.com/maxwelbm/alkemy-g7.git/internal/service/interfaces"
	"github.com/maxwelbm/alkemy-g7.git/pkg/logger"
)

func CreateServiceSellers(rp interfaces.ISellerRepo, rpl serviceInterface.ILocalityService, log logger.Logger) *SellersService {
	return &SellersService{Rp: rp, Rpl: rpl, log: log}
}

type SellersService struct {
	Rp  interfaces.ISellerRepo
	Rpl serviceInterface.ILocalityService
	log logger.Logger
}

func (s *SellersService) GetAll() (sellers []model.Seller, err error) {
	sellers, err = s.Rp.Get()

	s.log.Log("SellersService", "INFO", fmt.Sprintf("Retrieved sellers: %+v", sellers))

	return
}

func (s *SellersService) GetByID(id int) (seller model.Seller, err error) {
	seller, err = s.Rp.GetByID(id)

	s.log.Log("SellersService", "INFO", fmt.Sprintf("Retrieved seller: %+v", seller))

	return
}

func (s *SellersService) CreateSeller(seller *model.Seller) (sl model.Seller, err error) {
	if err := seller.ValidateEmptyFields(seller); err != nil {
		s.log.Log("SellersService", "ERROR", fmt.Sprintf("Error: %v", err))

		return sl, err
	}

	_, err = s.Rpl.GetByID(seller.Locality)
	if err != nil {
		s.log.Log("SellersService", "ERROR", fmt.Sprintf("Error: %v", err))

		return
	}

	sl, err = s.Rp.Post(seller)

	s.log.Log("SellersService", "INFO", fmt.Sprintf("Created seller: %+v", sl))

	return
}

func (s *SellersService) UpdateSeller(id int, seller *model.Seller) (sl model.Seller, err error) {
	if seller.Locality != 0 {
		_, err := s.Rpl.GetByID(seller.Locality)
		if err != nil {
			s.log.Log("SellersService", "ERROR", fmt.Sprintf("Error: %v", err))

			return sl, err
		}
	}

	existSl, _ := s.GetByID(id)
	err = seller.ValidateUpdateFields(seller, &existSl)

	if err != nil {
		s.log.Log("SellersService", "ERROR", fmt.Sprintf("Error: %v", err))

		return
	}

	sl, err = s.Rp.Patch(id, seller)

	s.log.Log("SellersService", "INFO", fmt.Sprintf("Updated seller: %+v", sl))

	return sl, err
}

func (s *SellersService) DeleteSeller(id int) error {
	err := s.Rp.Delete(id)

	s.log.Log("SellersService", "INFO", fmt.Sprintf("Removed seller with ID: %d", id))

	return err
}
