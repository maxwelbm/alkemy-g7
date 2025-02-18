package service

import (
	"fmt"

	"github.com/go-sql-driver/mysql"
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/maxwelbm/alkemy-g7.git/internal/repository/interfaces"
	servicesInterfaces "github.com/maxwelbm/alkemy-g7.git/internal/service/interfaces"
	"github.com/maxwelbm/alkemy-g7.git/pkg/customerror"
	"github.com/maxwelbm/alkemy-g7.git/pkg/logger"
)

type InboundOrderService struct {
	rp          interfaces.IInboundOrderRepository
	employeeSv  servicesInterfaces.IEmployeeService
	warehouseSv servicesInterfaces.IWarehouseService
	log         logger.Logger
}

func NewInboundOrderService(
	rp interfaces.IInboundOrderRepository,
	employeeSv servicesInterfaces.IEmployeeService,
	warehouseSv servicesInterfaces.IWarehouseService,
	log logger.Logger) *InboundOrderService {
	return &InboundOrderService{
		rp:          rp,
		employeeSv:  employeeSv,
		warehouseSv: warehouseSv,
		log:         log,
	}
}

func (i *InboundOrderService) Post(inboundOrder model.InboundOrder) (model.InboundOrder, error) {
	i.log.Log("InboundOrderService", "INFO", "initializing Post function for inbound order")

	isValid := inboundOrder.IsValid()

	if !isValid {
		i.log.Log("InboundOrderService", "ERROR", "invalid inbound order entry")
		return model.InboundOrder{}, customerror.InboundErrInvalidEntry
	}

	_, err := i.employeeSv.GetEmployeeByID(inboundOrder.EmployeeID)

	if err != nil {
		i.log.Log("InboundOrderService", "ERROR", fmt.Sprintf("invalid employee ID: %d, error: %v", inboundOrder.EmployeeID, err))
		return model.InboundOrder{}, customerror.InboundErrInvalidEmployee
	}

	_, err = i.warehouseSv.GetByIDWareHouse(inboundOrder.WareHouseID)

	if err != nil {
		i.log.Log("InboundOrderService", "ERROR", fmt.Sprintf("invalid warehouse ID: %d, error: %v", inboundOrder.WareHouseID, err))
		return model.InboundOrder{}, customerror.InboundErrInvalidWarehouse
	}

	entry, err := i.rp.Post(inboundOrder)

	if err != nil {
		i.log.Log("InboundOrderService", "ERROR", fmt.Sprintf("failed to create inbound order: %v", err))

		mysqlErr, ok := err.(*mysql.MySQLError)
		if !ok {
			i.log.Log("InboundOrderService", "ERROR", fmt.Sprintf("unexpected error: %v", err))
			return model.InboundOrder{}, err
		}

		switch mysqlErr.Number {
		case 1452:
			i.log.Log("InboundOrderService", "ERROR", "invalid product batch ID")
			return model.InboundOrder{}, customerror.InboundErrInvalidProductBatch
		case 1062:
			i.log.Log("InboundOrderService", "ERROR", "duplicated order number")
			return model.InboundOrder{}, customerror.InboundErrDuplicatedOrderNumber
		default:
			i.log.Log("InboundOrderService", "ERROR", fmt.Sprintf("MySQL error: %v", mysqlErr))
			return model.InboundOrder{}, err
		}
	}

	i.log.Log("InboundOrderService", "INFO", fmt.Sprintf("Post function finished successfully, created inbound order with ID: %d", entry.ID))

	return entry, nil
}
