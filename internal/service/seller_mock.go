package service

import (
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/stretchr/testify/mock"
)

type SellerMockService struct {
	mock.Mock
}

func (s *SellerMockService) GetAll() (sellers []model.Seller, err error) {
	args := s.Called()
	sellers = args.Get(0).([]model.Seller)
	err = args.Error(1)

	return
}

func (s *SellerMockService) GetById(id int) (seller model.Seller, err error) {
	args := s.Called(id)
	seller = args.Get(0).(model.Seller)
	err = args.Error(1)

	return
}

func (s *SellerMockService) CreateSeller(seller *model.Seller) (sl model.Seller, err error) {
	args := s.Called(seller)
	sl = args.Get(0).(model.Seller)
	err = args.Error(1)

	return
}

func (s *SellerMockService) UpdateSeller(id int, seller *model.Seller) (sl model.Seller, err error) {
	args := s.Called(id, seller)
	sl = args.Get(0).(model.Seller)
	err = args.Error(1)

	return
}

func (s *SellerMockService) DeleteSeller(id int) error {
	args := s.Called(id)
	err := args.Error(0)

	return err
}
