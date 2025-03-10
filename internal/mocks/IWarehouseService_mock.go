// Code generated by mockery v2.52.1. DO NOT EDIT.

package mocks

import (
	model "github.com/maxwelbm/alkemy-g7.git/internal/model"
	mock "github.com/stretchr/testify/mock"
)

// MockIWarehouseService is an autogenerated mock type for the IWarehouseService type
type MockIWarehouseService struct {
	mock.Mock
}

// DeleteByIDWareHouse provides a mock function with given fields: id
func (_m *MockIWarehouseService) DeleteByIDWareHouse(id int) error {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for DeleteByIDWareHouse")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllWareHouse provides a mock function with no fields
func (_m *MockIWarehouseService) GetAllWareHouse() ([]model.WareHouse, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetAllWareHouse")
	}

	var r0 []model.WareHouse
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]model.WareHouse, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []model.WareHouse); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.WareHouse)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByIDWareHouse provides a mock function with given fields: id
func (_m *MockIWarehouseService) GetByIDWareHouse(id int) (model.WareHouse, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for GetByIDWareHouse")
	}

	var r0 model.WareHouse
	var r1 error
	if rf, ok := ret.Get(0).(func(int) (model.WareHouse, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(int) model.WareHouse); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(model.WareHouse)
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PostWareHouse provides a mock function with given fields: warehouse
func (_m *MockIWarehouseService) PostWareHouse(warehouse model.WareHouse) (model.WareHouse, error) {
	ret := _m.Called(warehouse)

	if len(ret) == 0 {
		panic("no return value specified for PostWareHouse")
	}

	var r0 model.WareHouse
	var r1 error
	if rf, ok := ret.Get(0).(func(model.WareHouse) (model.WareHouse, error)); ok {
		return rf(warehouse)
	}
	if rf, ok := ret.Get(0).(func(model.WareHouse) model.WareHouse); ok {
		r0 = rf(warehouse)
	} else {
		r0 = ret.Get(0).(model.WareHouse)
	}

	if rf, ok := ret.Get(1).(func(model.WareHouse) error); ok {
		r1 = rf(warehouse)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateWareHouse provides a mock function with given fields: id, warehouse
func (_m *MockIWarehouseService) UpdateWareHouse(id int, warehouse model.WareHouse) (model.WareHouse, error) {
	ret := _m.Called(id, warehouse)

	if len(ret) == 0 {
		panic("no return value specified for UpdateWareHouse")
	}

	var r0 model.WareHouse
	var r1 error
	if rf, ok := ret.Get(0).(func(int, model.WareHouse) (model.WareHouse, error)); ok {
		return rf(id, warehouse)
	}
	if rf, ok := ret.Get(0).(func(int, model.WareHouse) model.WareHouse); ok {
		r0 = rf(id, warehouse)
	} else {
		r0 = ret.Get(0).(model.WareHouse)
	}

	if rf, ok := ret.Get(1).(func(int, model.WareHouse) error); ok {
		r1 = rf(id, warehouse)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewMockIWarehouseService creates a new instance of MockIWarehouseService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockIWarehouseService(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockIWarehouseService {
	mock := &MockIWarehouseService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
