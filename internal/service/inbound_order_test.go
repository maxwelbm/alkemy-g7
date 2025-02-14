package service

import (
	"errors"
	"testing"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/maxwelbm/alkemy-g7.git/internal/mocks"
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/maxwelbm/alkemy-g7.git/pkg/customerror"
	"github.com/stretchr/testify/assert"
)

func TestPostInboundOrder(t *testing.T) {
	repo := mocks.NewMockIInboundOrderRepository(t)
	employeeSvc := mocks.NewMockIEmployeeService(t)
	warehouseSvc := mocks.NewMockIWarehouseService(t)

	service := NewInboundOrderService(repo, employeeSvc, warehouseSvc)

	inboundOrder := model.InboundOrder{
		ID:             1,
		OrderDate:      time.Date(2023, 10, 1, 0, 0, 0, 0, time.UTC),
		OrderNumber:    "ORD123",
		EmployeeID:     1,
		ProductBatchID: 1,
		WareHouseID:    1,
	}

	t.Run("should return the created inbound order and no error", func(t *testing.T) {
		employeeSvc.On("GetEmployeeByID", inboundOrder.EmployeeID).Return(model.Employee{ID: 1}, nil).Once()
		warehouseSvc.On("GetByIDWareHouse", inboundOrder.WareHouseID).Return(model.WareHouse{ID: 1}, nil).Once()
		repo.On("Post", inboundOrder).Return(inboundOrder, nil).Once()

		result, err := service.Post(inboundOrder)

		assert.NoError(t, err)
		assert.Equal(t, inboundOrder, result)

	})

	t.Run("should return error when inbound order is invalid", func(t *testing.T) {
		invalidInboundOrder := model.InboundOrder{
			OrderNumber: "",
		}

		result, err := service.Post(invalidInboundOrder)

		assert.Error(t, err)
		assert.Equal(t, customerror.InboundErrInvalidEntry, err)
		assert.Empty(t, result)
	})

	t.Run("should return error when employee does not exist", func(t *testing.T) {
		employeeSvc.On("GetEmployeeByID", inboundOrder.EmployeeID).Return(model.Employee{}, customerror.EmployeeErrNotFound).Once()

		result, err := service.Post(inboundOrder)

		assert.Error(t, err)
		assert.Equal(t, customerror.InboundErrInvalidEmployee, err)
		assert.Empty(t, result)
	})

	t.Run("should return error when warehouse does not exist", func(t *testing.T) {
		employeeSvc.On("GetEmployeeByID", inboundOrder.EmployeeID).Return(model.Employee{ID: 1}, nil).Once()
		warehouseSvc.On("GetByIDWareHouse", inboundOrder.WareHouseID).Return(model.WareHouse{}, customerror.InboundErrInvalidWarehouse).Once()

		result, err := service.Post(inboundOrder)

		assert.Error(t, err)
		assert.Equal(t, customerror.InboundErrInvalidWarehouse, err)
		assert.Empty(t, result)

	})

	t.Run("should return error when product batch does not exist", func(t *testing.T) {
		employeeSvc.On("GetEmployeeByID", inboundOrder.EmployeeID).Return(model.Employee{ID: 1}, nil).Once()
		warehouseSvc.On("GetByIDWareHouse", inboundOrder.WareHouseID).Return(model.WareHouse{ID: 1}, nil).Once()
		repo.On("Post", inboundOrder).Return(model.InboundOrder{}, &mysql.MySQLError{Number: 1452}).Once()

		result, err := service.Post(inboundOrder)

		assert.Error(t, err)
		assert.Equal(t, customerror.InboundErrInvalidProductBatch, err)
		assert.Empty(t, result)

	})

	t.Run("should return error when order number is duplicated", func(t *testing.T) {
		employeeSvc.On("GetEmployeeByID", inboundOrder.EmployeeID).Return(model.Employee{ID: 1}, nil).Once()
		warehouseSvc.On("GetByIDWareHouse", inboundOrder.WareHouseID).Return(model.WareHouse{ID: 1}, nil).Once()
		repo.On("Post", inboundOrder).Return(model.InboundOrder{}, &mysql.MySQLError{Number: 1062}).Once()

		result, err := service.Post(inboundOrder)

		assert.Error(t, err)
		assert.Equal(t, customerror.InboundErrDuplicatedOrderNumber, err)
		assert.Empty(t, result)

	})

	t.Run("should return error when unmapped sql error occurs", func(t *testing.T) {
		employeeSvc.On("GetEmployeeByID", inboundOrder.EmployeeID).Return(model.Employee{ID: 1}, nil).Once()
		warehouseSvc.On("GetByIDWareHouse", inboundOrder.WareHouseID).Return(model.WareHouse{ID: 1}, nil).Once()
		repo.On("Post", inboundOrder).Return(model.InboundOrder{}, &mysql.MySQLError{Number: 66}).Once()

		result, err := service.Post(inboundOrder)

		assert.Error(t, err)
		assert.Empty(t, result)

	})

	t.Run("should return error when an unexpected error occurs", func(t *testing.T) {
		employeeSvc.On("GetEmployeeByID", inboundOrder.EmployeeID).Return(model.Employee{ID: 1}, nil).Once()
		warehouseSvc.On("GetByIDWareHouse", inboundOrder.WareHouseID).Return(model.WareHouse{ID: 1}, nil).Once()
		repo.On("Post", inboundOrder).Return(model.InboundOrder{}, errors.New("unexpected error")).Once()

		result, err := service.Post(inboundOrder)

		assert.Error(t, err)
		assert.Equal(t, errors.New("unexpected error"), err)
		assert.Empty(t, result)

	})
}
