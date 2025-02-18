package repository

import (
	"database/sql"
	"fmt"

	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/maxwelbm/alkemy-g7.git/pkg/logger"
)

type InboundOrderService struct {
	db  *sql.DB
	log logger.Logger
}

func NewInboundService(db *sql.DB, log logger.Logger) *InboundOrderService {
	return &InboundOrderService{
		db:  db,
		log: log,
	}
}

func (i *InboundOrderService) Post(inboundOrder model.InboundOrder) (model.InboundOrder, error) {
	i.log.Log("InboundOrderService", "INFO", "initializing Post function for inbound order")

	query := `
		INSERT INTO inbound_orders (order_date, order_number, employee_id, product_batch_id, warehouse_id)
		VALUES (?, ?, ?, ?, ?)
	`

	result, err := i.db.Exec(
		query,
		inboundOrder.OrderDate,
		inboundOrder.OrderNumber,
		inboundOrder.EmployeeID,
		inboundOrder.ProductBatchID,
		inboundOrder.WareHouseID,
	)

	if err != nil {
		i.log.Log("InboundOrderService", "ERROR", fmt.Sprintf("failed to insert inbound order: %v", err))
		return inboundOrder, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		i.log.Log("InboundOrderService", "ERROR", fmt.Sprintf("failed to retrieve last insert ID: %v", err))
		return inboundOrder, err
	}

	inboundOrder.ID = int(id)

	i.log.Log("InboundOrderService", "INFO", fmt.Sprintf("Post function finished successfully, created inbound order with ID: %d", inboundOrder.ID))

	return inboundOrder, nil
}
