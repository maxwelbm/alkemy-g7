// Code generated by mockery v2.52.1. DO NOT EDIT.

package mocks

import (
	model "github.com/maxwelbm/alkemy-g7.git/internal/model"
	mock "github.com/stretchr/testify/mock"
)

// MockIEmployeeService is an autogenerated mock type for the IEmployeeService type
type MockIEmployeeService struct {
	mock.Mock
}

// DeleteEmployee provides a mock function with given fields: id
func (_m *MockIEmployeeService) DeleteEmployee(id int) error {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for DeleteEmployee")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetEmployeeByID provides a mock function with given fields: id
func (_m *MockIEmployeeService) GetEmployeeByID(id int) (model.Employee, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for GetEmployeeByID")
	}

	var r0 model.Employee
	var r1 error
	if rf, ok := ret.Get(0).(func(int) (model.Employee, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(int) model.Employee); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(model.Employee)
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetEmployees provides a mock function with no fields
func (_m *MockIEmployeeService) GetEmployees() ([]model.Employee, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetEmployees")
	}

	var r0 []model.Employee
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]model.Employee, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []model.Employee); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Employee)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetInboundOrdersReportByEmployee provides a mock function with given fields: employeeID
func (_m *MockIEmployeeService) GetInboundOrdersReportByEmployee(employeeID int) (model.InboundOrdersReportByEmployee, error) {
	ret := _m.Called(employeeID)

	if len(ret) == 0 {
		panic("no return value specified for GetInboundOrdersReportByEmployee")
	}

	var r0 model.InboundOrdersReportByEmployee
	var r1 error
	if rf, ok := ret.Get(0).(func(int) (model.InboundOrdersReportByEmployee, error)); ok {
		return rf(employeeID)
	}
	if rf, ok := ret.Get(0).(func(int) model.InboundOrdersReportByEmployee); ok {
		r0 = rf(employeeID)
	} else {
		r0 = ret.Get(0).(model.InboundOrdersReportByEmployee)
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(employeeID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetInboundOrdersReports provides a mock function with no fields
func (_m *MockIEmployeeService) GetInboundOrdersReports() ([]model.InboundOrdersReportByEmployee, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetInboundOrdersReports")
	}

	var r0 []model.InboundOrdersReportByEmployee
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]model.InboundOrdersReportByEmployee, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []model.InboundOrdersReportByEmployee); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.InboundOrdersReportByEmployee)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InsertEmployee provides a mock function with given fields: employee
func (_m *MockIEmployeeService) InsertEmployee(employee model.Employee) (model.Employee, error) {
	ret := _m.Called(employee)

	if len(ret) == 0 {
		panic("no return value specified for InsertEmployee")
	}

	var r0 model.Employee
	var r1 error
	if rf, ok := ret.Get(0).(func(model.Employee) (model.Employee, error)); ok {
		return rf(employee)
	}
	if rf, ok := ret.Get(0).(func(model.Employee) model.Employee); ok {
		r0 = rf(employee)
	} else {
		r0 = ret.Get(0).(model.Employee)
	}

	if rf, ok := ret.Get(1).(func(model.Employee) error); ok {
		r1 = rf(employee)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateEmployee provides a mock function with given fields: id, employee
func (_m *MockIEmployeeService) UpdateEmployee(id int, employee model.Employee) (model.Employee, error) {
	ret := _m.Called(id, employee)

	if len(ret) == 0 {
		panic("no return value specified for UpdateEmployee")
	}

	var r0 model.Employee
	var r1 error
	if rf, ok := ret.Get(0).(func(int, model.Employee) (model.Employee, error)); ok {
		return rf(id, employee)
	}
	if rf, ok := ret.Get(0).(func(int, model.Employee) model.Employee); ok {
		r0 = rf(id, employee)
	} else {
		r0 = ret.Get(0).(model.Employee)
	}

	if rf, ok := ret.Get(1).(func(int, model.Employee) error); ok {
		r1 = rf(id, employee)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewMockIEmployeeService creates a new instance of MockIEmployeeService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockIEmployeeService(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockIEmployeeService {
	mock := &MockIEmployeeService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
