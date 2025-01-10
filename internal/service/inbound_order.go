package service

import (
	"github.com/go-sql-driver/mysql"
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/maxwelbm/alkemy-g7.git/internal/repository/interfaces"
	"github.com/maxwelbm/alkemy-g7.git/pkg/custom_error"
)

type InboundOrderService struct {
	rp          interfaces.IInboundOrderRepository
	employeeRp  interfaces.IEmployeeRepo
	warehouseRp interfaces.IWarehouseRepo
}

func NewInboundOrderService(
	rp interfaces.IInboundOrderRepository,
	employeeRp interfaces.IEmployeeRepo,
	warehouseRp interfaces.IWarehouseRepo) *InboundOrderService {
	return &InboundOrderService{
		rp:          rp,
		employeeRp:  employeeRp,
		warehouseRp: warehouseRp,
	}
}

func (i *InboundOrderService) Post(inboundOrder model.InboundOrder) (model.InboundOrder, error) {
	isValid := inboundOrder.IsValid()

	if !isValid {
		return model.InboundOrder{}, custom_error.InboundErrInvalidEntry
	}

	_, err := i.employeeRp.GetById(inboundOrder.EmployeeId)

	if err != nil {
		return model.InboundOrder{}, custom_error.InboundErrInvalidEmployee
	}

	//@todo
	// productBatchExists := @todo

	_, err = i.warehouseRp.GetByIdWareHouse(inboundOrder.WareHouseId)

	if err != nil {
		return model.InboundOrder{}, custom_error.InboundErrInvalidWarehouse
	}

	entry, err := i.rp.Post(inboundOrder)

	if err != nil {
		if err.(*mysql.MySQLError).Number == 1062 {
			err = custom_error.InboundErrDuplicatedOrderNumber
		}

		return model.InboundOrder{}, err
	}

	return entry, nil
}
