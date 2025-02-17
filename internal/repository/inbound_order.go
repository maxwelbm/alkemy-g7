package repository

import (
	"database/sql"
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/maxwelbm/alkemy-g7.git/pkg/logger"
)

type InboundOrderService struct {
	db  *sql.DB
	log *logger.Logger
}

func NewInboundService(db *sql.DB, log *logger.Logger) *InboundOrderService {
	return &InboundOrderService{
		db:  db,
		log: log,
	}
}

func (i *InboundOrderService) Post(inboundOrder model.InboundOrder) (model.InboundOrder, error) {
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
		return inboundOrder, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return inboundOrder, err
	}

	inboundOrder.ID = int(id)

	return inboundOrder, nil
}
