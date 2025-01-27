package repository

import (
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/stretchr/testify/mock"
)

type SellerMockRepository struct {
	mock.Mock
}

func (rp *SellerMockRepository) Get() (sellers []model.Seller, err error) {
	args := rp.Called()
	sellers = args.Get(0).([]model.Seller)
	err = args.Error(1)

	return
}

func (rp *SellerMockRepository) GetById(id int) (sl model.Seller, err error) {
	args := rp.Called(id)
	sl = args.Get(0).(model.Seller)
	err = args.Error(1)

	return
}

func (rp *SellerMockRepository) Post(seller *model.Seller) (sl model.Seller, err error) {
	args := rp.Called(seller)
	sl = args.Get(0).(model.Seller)
	err = args.Error(1)

	return
}

func (rp *SellerMockRepository) Patch(id int, seller *model.Seller) (sl model.Seller, err error) {
	args := rp.Called(id, seller)
	sl = args.Get(0).(model.Seller)
	err = args.Error(1)

	return
}

func (rp *SellerMockRepository) Delete(id int) (err error) {
	args := rp.Called(id)
	err = args.Error(0)

	return
}
