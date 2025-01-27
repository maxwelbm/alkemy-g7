package service

import (
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/stretchr/testify/mock"
)

type MockBuyerService struct {
	mock.Mock
}

// CountPurchaseOrderBuyer implements interfaces.IBuyerservice.
func (b *MockBuyerService) CountPurchaseOrderBuyer() (countBuyerPurchaseOrder []model.BuyerPurchaseOrder, err error) {
	args := b.Called()
	return args.Get(0).([]model.BuyerPurchaseOrder), args.Error(1)
}

// CountPurchaseOrderByBuyerID implements interfaces.IBuyerservice.
func (b *MockBuyerService) CountPurchaseOrderByBuyerID(id int) (countBuyerPurchaseOrder model.BuyerPurchaseOrder, err error) {
	args := b.Called(id)
	return args.Get(0).(model.BuyerPurchaseOrder), args.Error(1)
}

// CreateBuyer implements interfaces.IBuyerservice.
func (b *MockBuyerService) CreateBuyer(newBuyer model.Buyer) (buyer model.Buyer, err error) {
	args := b.Called(newBuyer)

	buyer = args.Get(0).(model.Buyer)
	err = args.Error(1)

	return
}

// DeleteBuyerByID implements interfaces.IBuyerservice.
func (b *MockBuyerService) DeleteBuyerByID(id int) (err error) {
	args := b.Called(id)

	err = args.Error(0)

	return
}

// GetBuyerByID implements interfaces.IBuyerservice.
func (b *MockBuyerService) GetBuyerByID(id int) (buyer model.Buyer, err error) {
	args := b.Called(id)
	buyer = args.Get(0).(model.Buyer)
	err = args.Error(1)

	return
}

// UpdateBuyer implements interfaces.IBuyerservice.
func (b *MockBuyerService) UpdateBuyer(id int, newBuyer model.Buyer) (buyer model.Buyer, err error) {
	args := b.Called(id, newBuyer)

	buyer = args.Get(0).(model.Buyer)
	err = args.Error(1)

	return
}

func (b *MockBuyerService) GetAllBuyer() (buyers []model.Buyer, err error) {
	args := b.Called()
	buyers = args.Get(0).([]model.Buyer)
	err = args.Error(1)

	return
}

func NewMockBuyerService() *MockBuyerService {
	return &MockBuyerService{}
}
