package repository

import (
	"database/sql"
	"net/http"

	"github.com/go-sql-driver/mysql"
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/maxwelbm/alkemy-g7.git/pkg/customError"
)

type PurchaseOrderRepository struct {
	db *sql.DB
}

// Post implements interfaces.IPurchaseOrdersRepo.
func (p *PurchaseOrderRepository) Post(newPurchaseOrder model.PurchaseOrder) (id int64, err error) {
	prepare, err := p.db.Prepare("INSERT INTO purchase_orders (order_number, order_date, tracking_code, buyer_id, product_record_id) VALUES(?,?,?,?,?)")
	if err != nil {
		return
	}

	result, err := prepare.Exec(&newPurchaseOrder.OrderNumber, &newPurchaseOrder.OrderDate, &newPurchaseOrder.TrackingCode, &newPurchaseOrder.BuyerID, &newPurchaseOrder.ProductRecordID)

	if err != nil {
		if err.(*mysql.MySQLError).Number == 1062 {
			err = customError.NewPurcahseOrderError(http.StatusConflict, customError.ErrConflict.Error(), "order_number")
		}

		return
	}

	id, err = result.LastInsertId()

	return
}

func (p *PurchaseOrderRepository) GetByID(id int) (purchaseOrder model.PurchaseOrder, err error) {
	row := p.db.QueryRow("SELECT id, order_number, order_date, tracking_code, buyer_id, product_record_id FROM purchase_orders WHERE id = ?", id)

	err = row.Scan(&purchaseOrder.ID, &purchaseOrder.OrderNumber, &purchaseOrder.OrderDate, &purchaseOrder.TrackingCode, &purchaseOrder.BuyerID, &purchaseOrder.ProductRecordID)

	if err != nil {
		if err == sql.ErrNoRows {
			err = customError.NewPurcahseOrderError(http.StatusNotFound, customError.ErrNotFound.Error(), "Purchase Order")
		}

		return
	}

	return
}

func NewPurchaseOrderRepository(db *sql.DB) *PurchaseOrderRepository {
	return &PurchaseOrderRepository{db: db}
}
