package repository

import (
	"database/sql"

	"github.com/maxwelbm/alkemy-g7.git/internal/model"
)

type InboundOrderService struct {
	db *sql.DB
}

func NewInboundService(db *sql.DB) *InboundOrderService {
	return &InboundOrderService{
		db: db,
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
