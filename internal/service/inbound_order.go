package service

import (
	"github.com/go-sql-driver/mysql"
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/maxwelbm/alkemy-g7.git/internal/repository/interfaces"
	servicesInterfaces "github.com/maxwelbm/alkemy-g7.git/internal/service/interfaces"
	"github.com/maxwelbm/alkemy-g7.git/pkg/customerror"
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
		return model.InboundOrder{}, customerror.InboundErrInvalidEntry
	}

	_, err := i.employeeSv.GetEmployeeByID(inboundOrder.EmployeeID)

	if err != nil {
		return model.InboundOrder{}, customerror.InboundErrInvalidEmployee
	}

	//@todo
	// productBatchExists := @todo

	_, err = i.warehouseSv.GetByIDWareHouse(inboundOrder.WareHouseID)

	if err != nil {
		return model.InboundOrder{}, customerror.InboundErrInvalidWarehouse
	}

	entry, err := i.rp.Post(inboundOrder)

	if err != nil {
		mysqlErr, ok := err.(*mysql.MySQLError)
		if !ok {
			return model.InboundOrder{}, err
		}

		switch mysqlErr.Number {
		case 1452:
			return model.InboundOrder{}, customerror.InboundErrInvalidProductBatch
		case 1062:
			return model.InboundOrder{}, customerror.InboundErrDuplicatedOrderNumber
		default:
			return model.InboundOrder{}, err
		}
	}

	return entry, nil
}
