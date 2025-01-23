package service

import (
	"testing"

	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/maxwelbm/alkemy-g7.git/pkg/custom_error"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var employeeRepo = new(employeeRepositoryMock)
var warehouseRepo = new(warehouseRepositoryMock)
var employeeSv = CreateEmployeeService(employeeRepo, warehouseRepo)

func TestInsertEmployee(t *testing.T) {
	t.Run("should create the employee case everything is ok", func(t *testing.T) {
		warehouseRepo.On("GetByIdWareHouse", 1).Return(model.WareHouse{}, nil)
		employeeRepo.On("Post", mock.Anything).Return(model.Employee{ID: 10, CardNumberID: "#123", FirstName: "Bruce", LastName: "Wayne", WarehouseID: 1}, nil)

		validEntry := model.Employee{
			CardNumberID: "#123",
			FirstName:    "Bruce",
			LastName:     "Wayne",
			WarehouseID:  1,
		}

		newEmployee, err := employeeSv.InsertEmployee(validEntry)

		assert.Nil(t, err)
		assert.NotEmpty(t, newEmployee)
		assert.NotZero(t, newEmployee.ID)
	})

	t.Run("should return an error in case of invalid entry", func(t *testing.T) {
		invalidEntry := model.Employee{
			CardNumberID: "#123",
			LastName:     "Wayne",
		}

		employee, err := employeeSv.InsertEmployee(invalidEntry)

		assert.NotNil(t, err)
		assert.Error(t, err)
		assert.Empty(t, employee)
	})

	t.Run("should return an error in case of warehouseNotFound", func(t *testing.T) {
		warehouseRepo.On("GetByIdWareHouse", 2).Return(model.WareHouse{}, custom_error.ErrNotFound).Once()
		employeeRepo.On("Post", mock.Anything).Return(model.Employee{ID: 10, CardNumberID: "#123", FirstName: "Bruce", LastName: "Wayne", WarehouseID: 1}, nil)

		validEntry := model.Employee{
			CardNumberID: "#123",
			FirstName:    "Bruce",
			LastName:     "Wayne",
			WarehouseID:  2,
		}

		employee, err := employeeSv.InsertEmployee(validEntry)

		assert.NotNil(t, err)
		assert.Error(t, err)
		assert.Empty(t, employee)
	})

}

// Mock
type employeeRepositoryMock struct {
	mock.Mock
}

func (m *employeeRepositoryMock) Get() ([]model.Employee, error) {
	args := m.Called()
	return args.Get(0).([]model.Employee), args.Error(1)
}

func (m *employeeRepositoryMock) GetByID(id int) (model.Employee, error) {
	args := m.Called(id)
	return args.Get(0).(model.Employee), args.Error(1)
}

func (m *employeeRepositoryMock) Update(id int, employee model.Employee) (model.Employee, error) {
	args := m.Called(id, employee)
	return args.Get(0).(model.Employee), args.Error(1)
}

func (m *employeeRepositoryMock) Post(employee model.Employee) (model.Employee, error) {
	args := m.Called(employee)
	return args.Get(0).(model.Employee), args.Error(1)
}

func (m *employeeRepositoryMock) Delete(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *employeeRepositoryMock) GetInboundOrdersReportByEmployee(employeeID int) (model.InboundOrdersReportByEmployee, error) {
	args := m.Called(employeeID)
	return args.Get(0).(model.InboundOrdersReportByEmployee), args.Error(1)
}

func (m *employeeRepositoryMock) GetInboundOrdersReports() ([]model.InboundOrdersReportByEmployee, error) {
	args := m.Called()
	return args.Get(0).([]model.InboundOrdersReportByEmployee), args.Error(1)
}

type warehouseRepositoryMock struct {
	mock.Mock
}

func (m *warehouseRepositoryMock) GetAllWareHouse() (w []model.WareHouse, err error) {
	return nil, nil
}

func (m *warehouseRepositoryMock) GetByIdWareHouse(id int) (w model.WareHouse, err error) {
	args := m.Called(id)
	return args.Get(0).(model.WareHouse), args.Error(1)
}

func (m *warehouseRepositoryMock) PostWareHouse(warehouse *model.WareHouse) (id int64, err error) {
	return 0, nil
}
func (m *warehouseRepositoryMock) UpdateWareHouse(id int, warehouse *model.WareHouse) (err error) {
	return nil
}
func (w *warehouseRepositoryMock) DeleteByIdWareHouse(id int) error {
	return nil
}
