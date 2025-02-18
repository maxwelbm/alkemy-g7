package repository

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/maxwelbm/alkemy-g7.git/pkg/logger"

	"github.com/go-sql-driver/mysql"
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/maxwelbm/alkemy-g7.git/pkg/customerror"
)

type PurchaseOrderRepository struct {
	db  *sql.DB
	log logger.Logger
}

// Post implements interfaces.IPurchaseOrdersRepo.
func (p *PurchaseOrderRepository) Post(newPurchaseOrder model.PurchaseOrder) (id int64, err error) {
	p.log.Log("PurchaseOrderRepository", "INFO", fmt.Sprintf("initializing Post function with parameter %v", newPurchaseOrder))
	prepare, err := p.db.Prepare("INSERT INTO purchase_orders (order_number, order_date, tracking_code, buyer_id, product_record_id) VALUES(?,?,?,?,?)")

	if err != nil {
		p.log.Log("PurchaseOrderRepository", "ERROR", fmt.Sprintf("Error:  %v", err))
		return
	}

	result, err := prepare.Exec(&newPurchaseOrder.OrderNumber, &newPurchaseOrder.OrderDate, &newPurchaseOrder.TrackingCode, &newPurchaseOrder.BuyerID, &newPurchaseOrder.ProductRecordID)

	if err != nil {
		if err.(*mysql.MySQLError).Number == 1062 {
			err = customerror.NewPurcahseOrderError(http.StatusConflict, customerror.ErrConflict.Error(), "order_number")
		}

		p.log.Log("PurchaseOrderRepository", "ERROR", fmt.Sprintf("Error:  %v", err))

		return
	}

	id, err = result.LastInsertId()
	p.log.Log("PurchaseOrderRepository", "INFO", fmt.Sprintf("returning inserted ID %d", id))

	return
}

func (p *PurchaseOrderRepository) GetByID(id int) (purchaseOrder model.PurchaseOrder, err error) {
	p.log.Log("PurchaseOrderRepository", "INFO", fmt.Sprintf("initializing GetByID function with parameter %v", id))
	row := p.db.QueryRow("SELECT id, order_number, order_date, tracking_code, buyer_id, product_record_id FROM purchase_orders WHERE id = ?", id)

	err = row.Scan(&purchaseOrder.ID, &purchaseOrder.OrderNumber, &purchaseOrder.OrderDate, &purchaseOrder.TrackingCode, &purchaseOrder.BuyerID, &purchaseOrder.ProductRecordID)

	if err != nil {
		if err == sql.ErrNoRows {
			err = customerror.NewPurcahseOrderError(http.StatusNotFound, customerror.ErrNotFound.Error(), "Purchase Order")
		}

		p.log.Log("PurchaseOrderRepository", "ERROR", fmt.Sprintf("Error:  %v", err))

		return
	}

	p.log.Log("PurchaseOrderRepository", "INFO", fmt.Sprintf("returning PurchaseOrder:  %v", purchaseOrder))

	return
}

func NewPurchaseOrderRepository(db *sql.DB, log logger.Logger) *PurchaseOrderRepository {
	return &PurchaseOrderRepository{db: db, log: log}
}
