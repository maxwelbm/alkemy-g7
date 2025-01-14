package service

import (
	"github.com/go-sql-driver/mysql"
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/maxwelbm/alkemy-g7.git/internal/repository/interfaces"
	servicesInterfaces "github.com/maxwelbm/alkemy-g7.git/internal/service/interfaces"
	"github.com/maxwelbm/alkemy-g7.git/pkg/custom_error"
)

type InboundOrderService struct {
	rp          interfaces.IInboundOrderRepository
	employeeSv  servicesInterfaces.IEmployeeService
	warehouseSv servicesInterfaces.IWarehouseService
}

func NewInboundOrderService(
	rp interfaces.IInboundOrderRepository,
	employeeSv servicesInterfaces.IEmployeeService,
	warehouseSv servicesInterfaces.IWarehouseService) *InboundOrderService {
	return &InboundOrderService{
		rp:          rp,
		employeeSv:  employeeSv,
		warehouseSv: warehouseSv,
	}
}

func (i *InboundOrderService) Post(inboundOrder model.InboundOrder) (model.InboundOrder, error) {
	isValid := inboundOrder.IsValid()

	if !isValid {
		return model.InboundOrder{}, custom_error.InboundErrInvalidEntry
	}

	_, err := i.employeeSv.GetEmployeeById(inboundOrder.EmployeeId)

	if err != nil {
		return model.InboundOrder{}, custom_error.InboundErrInvalidEmployee
	}

	//@todo
	// productBatchExists := @todo

	_, err = i.warehouseSv.GetByIdWareHouse(inboundOrder.WareHouseId)

	if err != nil {
		return model.InboundOrder{}, custom_error.InboundErrInvalidWarehouse
	}

	entry, err := i.rp.Post(inboundOrder)

	if err != nil {
		mysqlErr, ok := err.(*mysql.MySQLError)
		if !ok {
			return model.InboundOrder{}, err
		}

		switch mysqlErr.Number {
		case 1452:
			return model.InboundOrder{}, custom_error.InboundErrInvalidProductBatch
		case 1062:
			return model.InboundOrder{}, custom_error.InboundErrDuplicatedOrderNumber
		default:
			return model.InboundOrder{}, err
		}
	}

	return entry, nil
}
