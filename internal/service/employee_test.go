package service

import (
	"errors"
	"testing"

	"github.com/go-sql-driver/mysql"
	"github.com/maxwelbm/alkemy-g7.git/internal/mocks"
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/maxwelbm/alkemy-g7.git/pkg/customerror"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestInsertEmployee(t *testing.T) {
	employeeRepo := mocks.NewMockIEmployeeRepo(t)
	warehouseRepo := mocks.NewMockIWarehouseRepo(t)
	employeeSv := CreateEmployeeService(employeeRepo, warehouseRepo, mocks.MockLog{})
	t.Run("should create the employee case everything is ok", func(t *testing.T) {
		warehouseRepo.On("GetByIDWareHouse", 1).Return(model.WareHouse{}, nil).Once()
		employeeRepo.On("Post", mock.Anything).Return(model.Employee{ID: 10, CardNumberID: "#123", FirstName: "Bruce", LastName: "Wayne", WarehouseID: 1}, nil).Once()

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
		warehouseRepo.On("GetByIDWareHouse", 2).Return(model.WareHouse{}, customerror.ErrNotFound).Once()

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

	t.Run("should return an error in case of database error", func(t *testing.T) {
		warehouseRepo.On("GetByIDWareHouse", 1).Return(model.WareHouse{}, nil)
		employeeRepo.On("Post", mock.Anything).Return(model.Employee{}, &mysql.MySQLError{Number: 1062})

		validEntry := model.Employee{
			CardNumberID: "#123",
			FirstName:    "Bruce",
			LastName:     "Wayne",
			WarehouseID:  1,
		}

		employee, err := employeeSv.InsertEmployee(validEntry)

		assert.Error(t, err)
		assert.EqualValues(t, customerror.EmployeeErrDuplicatedCardNumber, err)
		assert.Empty(t, employee)
	})
}

func TestGetEmployees(t *testing.T) {
	employeeRepo := mocks.NewMockIEmployeeRepo(t)
	employeeSv := CreateEmployeeService(employeeRepo, nil, mocks.MockLog{})

	t.Run("should return all the employees", func(t *testing.T) {
		employeeRepo.On("Get", mock.Anything).Return([]model.Employee{{CardNumberID: "#123", ID: 1, FirstName: "Bruce", LastName: "Wayne", WarehouseID: 1}, {ID: 2, CardNumberID: "#234", FirstName: "Yami", LastName: "Sukehiro", WarehouseID: 2}}, nil).Once()

		employee, err := employeeSv.GetEmployees()

		assert.NotNil(t, employee)
		assert.Len(t, employee, 2)
		assert.Nil(t, err)
	})

	t.Run("should return the error in case of repository err", func(t *testing.T) {
		employeeRepo.On("Get", mock.Anything).Return([]model.Employee{}, customerror.EmployeeErrInvalid).Once()

		employee, err := employeeSv.GetEmployees()

		assert.Nil(t, employee)
		assert.Error(t, err)
	})
}

func TestGetEmployeeByID(t *testing.T) {
	employeeRepo := mocks.NewMockIEmployeeRepo(t)
	employeeSv := CreateEmployeeService(employeeRepo, nil, mocks.MockLog{})

	t.Run("should return the employee by id", func(t *testing.T) {
		mockEmployee := model.Employee{ID: 1, CardNumberID: "#123", FirstName: "Jack", LastName: "Chan", WarehouseID: 2}
		employeeRepo.On("GetByID", mock.Anything).Return(mockEmployee, nil).Once()

		employee, err := employeeSv.GetEmployeeByID(1)

		assert.ObjectsAreEqualValues(employee, mockEmployee)
		assert.Nil(t, err)
	})
}

func TestUpdateEmployee(t *testing.T) {
	employeeRepo := mocks.NewMockIEmployeeRepo(t)
	warehouseRepo := mocks.NewMockIWarehouseRepo(t)
	employeeSv := CreateEmployeeService(employeeRepo, warehouseRepo, mocks.MockLog{})

	validEntry := model.Employee{
		WarehouseID:  2,
		FirstName:    "Renato",
		LastName:     "Moicano",
		CardNumberID: "#456",
	}

	t.Run("should merge the entry data with the existing employee data and return the updated employee", func(t *testing.T) {
		existingEmployeeMock := model.Employee{ID: 1, CardNumberID: "#123", FirstName: "Islam", LastName: "Makhachev", WarehouseID: 1}
		employeeRepo.On("GetByID", mock.Anything).Return(existingEmployeeMock, nil).Once()
		employeeRepo.On("Update", mock.Anything, mock.Anything).Return(model.Employee{ID: existingEmployeeMock.ID, CardNumberID: validEntry.CardNumberID, FirstName: validEntry.FirstName, LastName: validEntry.LastName, WarehouseID: validEntry.WarehouseID}, nil)

		warehouseRepo.On("GetByIDWareHouse", mock.Anything).Return(model.WareHouse{}, nil).Once()

		employee, err := employeeSv.UpdateEmployee(1, validEntry)

		assert.Nil(t, err)
		assert.NotEmpty(t, employee)
		assert.Equal(t, validEntry.WarehouseID, employee.WarehouseID)
		assert.NotEqual(t, existingEmployeeMock.FirstName, employee.FirstName)
		assert.NotEqual(t, existingEmployeeMock.LastName, employee.LastName)
		assert.Equal(t, validEntry.CardNumberID, employee.CardNumberID)
	})

	t.Run("should return an error in case an empty employeee", func(t *testing.T) {
		invalidEntry := model.Employee{}

		employee, err := employeeSv.UpdateEmployee(1, invalidEntry)

		assert.Error(t, err)
		assert.Empty(t, employee)
	})

	t.Run("should return an error in case of new warehouseid does not exist", func(t *testing.T) {
		warehouseRepo.On("GetByIDWareHouse", mock.Anything).Return(model.WareHouse{}, errors.New("")).Once()

		employee, err := employeeSv.UpdateEmployee(1, validEntry)

		assert.Error(t, err)
		assert.Empty(t, employee)
	})

	t.Run("should return an error in case of employeeid does not exist", func(t *testing.T) {
		warehouseRepo.On("GetByIDWareHouse", mock.Anything).Return(model.WareHouse{}, nil).Once()
		employeeRepo.On("GetByID", mock.Anything).Return(model.Employee{}, customerror.EmployeeErrNotFound).Once()

		employee, err := employeeSv.UpdateEmployee(1, validEntry)

		assert.Error(t, err)
		assert.Empty(t, employee)
	})
}

func TestDeleteEmployee(t *testing.T) {
	employeeRepo := mocks.NewMockIEmployeeRepo(t)
	employeeSv := CreateEmployeeService(employeeRepo, nil, mocks.MockLog{})

	t.Run("should return nil case success", func(t *testing.T) {
		employeeRepo.On("GetByID", 1).Return(model.Employee{}, nil).Once()
		employeeRepo.On("Delete", mock.Anything).Return(nil).Once()

		err := employeeSv.DeleteEmployee(1)

		assert.Nil(t, err)
	})

	t.Run("should return an error case employee id in case of employee id does not exist", func(t *testing.T) {
		employeeRepo.On("GetByID", 1).Return(model.Employee{}, customerror.EmployeeErrNotFound).Once()

		err := employeeSv.DeleteEmployee(1)

		assert.Error(t, err)
	})
}

func TestGetInboundOrdersReports(t *testing.T) {
	employeeRepo := mocks.NewMockIEmployeeRepo(t)
	employeeSv := CreateEmployeeService(employeeRepo, nil, mocks.MockLog{})

	t.Run("should return all inbound orders reports", func(t *testing.T) {
		employeeRepo.On("GetInboundOrdersReports", mock.Anything).Return([]model.InboundOrdersReportByEmployee{
			{ID: 1, CardNumberID: "123", FirstName: "Islam", LastName: "Makhachev", WarehouseID: 1, InboundOrdersCount: 3},
			{ID: 2, CardNumberID: "456", FirstName: "Jon", LastName: "Jones", WarehouseID: 2, InboundOrdersCount: 0}}, nil)

		inboundOrders, err := employeeSv.GetInboundOrdersReports()

		assert.Nil(t, err)
		assert.NotEmpty(t, inboundOrders)
		assert.Len(t, inboundOrders, 2)
	})
}

func TestGetInboundOrdersReportByEmployee(t *testing.T) {
	employeeRepo := mocks.NewMockIEmployeeRepo(t)
	employeeSv := CreateEmployeeService(employeeRepo, nil, mocks.MockLog{})

	t.Run("should return the inbound orders by employee", func(t *testing.T) {
		mockData := model.InboundOrdersReportByEmployee{ID: 1, CardNumberID: "#123", FirstName: "Jon", LastName: "Jones", WarehouseID: 2, InboundOrdersCount: 30}
		employeeRepo.On("GetByID", mock.Anything).Return(model.Employee{}, nil).Once()
		employeeRepo.On("GetInboundOrdersReportByEmployee", mock.Anything).Return(mockData, nil).Once()
		inboundOrders, err := employeeSv.GetInboundOrdersReportByEmployee(1)

		assert.Nil(t, err)
		assert.NotEmpty(t, inboundOrders)
		assert.EqualValues(t, inboundOrders, mockData)
	})
	t.Run("should return an error in case of negative employeeid", func(t *testing.T) {
		inboundOrders, err := employeeSv.GetInboundOrdersReportByEmployee(-1)

		assert.Error(t, err)
		assert.Empty(t, inboundOrders)
		assert.EqualValues(t, err, customerror.EmployeeErrInvalid)
	})
	t.Run("should return an error in case of employee does not exist", func(t *testing.T) {
		employeeRepo.On("GetByID", mock.Anything).Return(model.Employee{}, customerror.EmployeeErrNotFound).Once()
		inboundOrders, err := employeeSv.GetInboundOrdersReportByEmployee(1)

		assert.Error(t, err)
		assert.Empty(t, inboundOrders)
		assert.EqualValues(t, err, customerror.EmployeeErrNotFound)
	})
	t.Run("should return an error in case of repository method fails", func(t *testing.T) {
		repoMockErr := errors.New("something went wrong")
		employeeRepo.On("GetByID", mock.Anything).Return(model.Employee{}, nil).Once()
		employeeRepo.On("GetInboundOrdersReportByEmployee", mock.Anything).Return(model.InboundOrdersReportByEmployee{}, repoMockErr).Once()
		inboundOrders, err := employeeSv.GetInboundOrdersReportByEmployee(1)

		assert.Error(t, err)
		assert.Empty(t, inboundOrders)
		assert.EqualValues(t, err, repoMockErr)
	})

}
