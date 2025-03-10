// Code generated by mockery v2.52.1. DO NOT EDIT.

package mocks

import (
	model "github.com/maxwelbm/alkemy-g7.git/internal/model"
	mock "github.com/stretchr/testify/mock"
)

// MockIBuyerRepo is an autogenerated mock type for the IBuyerRepo type
type MockIBuyerRepo struct {
	mock.Mock
}

// CountPurchaseOrderBuyers provides a mock function with no fields
func (_m *MockIBuyerRepo) CountPurchaseOrderBuyers() ([]model.BuyerPurchaseOrder, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for CountPurchaseOrderBuyers")
	}

	var r0 []model.BuyerPurchaseOrder
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]model.BuyerPurchaseOrder, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []model.BuyerPurchaseOrder); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.BuyerPurchaseOrder)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CountPurchaseOrderByBuyerID provides a mock function with given fields: id
func (_m *MockIBuyerRepo) CountPurchaseOrderByBuyerID(id int) (model.BuyerPurchaseOrder, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for CountPurchaseOrderByBuyerID")
	}

	var r0 model.BuyerPurchaseOrder
	var r1 error
	if rf, ok := ret.Get(0).(func(int) (model.BuyerPurchaseOrder, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(int) model.BuyerPurchaseOrder); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(model.BuyerPurchaseOrder)
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: id
func (_m *MockIBuyerRepo) Delete(id int) error {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with no fields
func (_m *MockIBuyerRepo) Get() ([]model.Buyer, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 []model.Buyer
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]model.Buyer, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []model.Buyer); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Buyer)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: id
func (_m *MockIBuyerRepo) GetByID(id int) (model.Buyer, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for GetByID")
	}

	var r0 model.Buyer
	var r1 error
	if rf, ok := ret.Get(0).(func(int) (model.Buyer, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(int) model.Buyer); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(model.Buyer)
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Post provides a mock function with given fields: newBuyer
func (_m *MockIBuyerRepo) Post(newBuyer model.Buyer) (int64, error) {
	ret := _m.Called(newBuyer)

	if len(ret) == 0 {
		panic("no return value specified for Post")
	}

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(model.Buyer) (int64, error)); ok {
		return rf(newBuyer)
	}
	if rf, ok := ret.Get(0).(func(model.Buyer) int64); ok {
		r0 = rf(newBuyer)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(model.Buyer) error); ok {
		r1 = rf(newBuyer)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: id, newBuyer
func (_m *MockIBuyerRepo) Update(id int, newBuyer model.Buyer) error {
	ret := _m.Called(id, newBuyer)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(int, model.Buyer) error); ok {
		r0 = rf(id, newBuyer)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewMockIBuyerRepo creates a new instance of MockIBuyerRepo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockIBuyerRepo(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockIBuyerRepo {
	mock := &MockIBuyerRepo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
