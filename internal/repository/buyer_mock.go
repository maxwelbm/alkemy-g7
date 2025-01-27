package repository

import (
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/stretchr/testify/mock"
)

type MockBuyerRepo struct {
	mock.Mock
}

func (m *MockBuyerRepo) CountPurchaseOrderBuyers() (countBuyerPurchaseOrder []model.BuyerPurchaseOrder, err error) {
	args := m.Called()
	return args.Get(0).([]model.BuyerPurchaseOrder), args.Error(1)
}

func (m *MockBuyerRepo) CountPurchaseOrderByBuyerID(id int) (countBuyerPurchaseOrder model.BuyerPurchaseOrder, err error) {
	args := m.Called(id)
	return args.Get(0).(model.BuyerPurchaseOrder), args.Error(1)
}

func (m *MockBuyerRepo) Delete(id int) (err error) {
	args := m.Called(id)

	err = args.Error(0)

	return
}

func (m *MockBuyerRepo) Get() (buyers []model.Buyer, err error) {
	args := m.Called()

	buyers = args.Get(0).([]model.Buyer)

	err = args.Error(1)

	return
}

func (m *MockBuyerRepo) GetByID(id int) (buyer model.Buyer, err error) {
	args := m.Called(id)

	buyer = args.Get(0).(model.Buyer)
	err = args.Error(1)

	return
}

func (m *MockBuyerRepo) Post(newBuyer model.Buyer) (id int64, err error) {
	args := m.Called(newBuyer)

	id = args.Get(0).(int64)
	err = args.Error(1)

	return
}

func (m *MockBuyerRepo) Update(id int, newBuyer model.Buyer) (err error) {
	args := m.Called(id, newBuyer)

	err = args.Error(0)

	return
}
