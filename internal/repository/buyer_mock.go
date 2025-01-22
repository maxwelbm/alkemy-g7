package repository

import (
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/stretchr/testify/mock"
)

type MockBuyerRepo struct {
	mock.Mock
}

// CountPurchaseOrderBuyers implements interfaces.IBuyerRepo.
func (m *MockBuyerRepo) CountPurchaseOrderBuyers() (countBuyerPurchaseOrder []model.BuyerPurchaseOrder, err error) {
	panic("unimplemented")
}

// CountPurchaseOrderByBuyerId implements interfaces.IBuyerRepo.
func (m *MockBuyerRepo) CountPurchaseOrderByBuyerId(id int) (countBuyerPurchaseOrder model.BuyerPurchaseOrder, err error) {
	panic("unimplemented")
}

// Delete implements interfaces.IBuyerRepo.
func (m *MockBuyerRepo) Delete(id int) (err error) {
	panic("unimplemented")
}

// Get implements interfaces.IBuyerRepo.
func (m *MockBuyerRepo) Get() (buyers []model.Buyer, err error) {
	args := m.Called()

	buyers = args.Get(0).([]model.Buyer)

	err = args.Error(1)

	return
}

// GetById implements interfaces.IBuyerRepo.
func (m *MockBuyerRepo) GetById(id int) (buyer model.Buyer, err error) {
	panic("unimplemented")
}

// Post implements interfaces.IBuyerRepo.
func (m *MockBuyerRepo) Post(newBuyer model.Buyer) (id int64, err error) {
	panic("unimplemented")
}

// Update implements interfaces.IBuyerRepo.
func (m *MockBuyerRepo) Update(id int, newBuyer model.Buyer) (err error) {
	panic("unimplemented")
}
