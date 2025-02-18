package repository

import (
	"database/sql"
	"fmt"
	"github.com/maxwelbm/alkemy-g7.git/pkg/logger"
	"net/http"

	"github.com/go-sql-driver/mysql"
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/maxwelbm/alkemy-g7.git/pkg/customerror"
)

type BuyerRepository struct {
	db  *sql.DB
	log logger.Logger
}

func (r *BuyerRepository) Delete(id int) (err error) {
	r.log.Log("BuyerRepository", "INFO", fmt.Sprintf("initializing Delete function with parameter %d", id))
	_, err = r.db.Exec("DELETE FROM buyers WHERE id = ?", id)
	if err != nil {
		if err.(*mysql.MySQLError).Number == 1451 {
			err = customerror.NewBuyerError(http.StatusConflict, customerror.ErrDependencies.Error(), "Buyer")
		}
		r.log.Log("BuyerRepository", "ERROR", fmt.Sprintf("Error: %v", err))
		return
	}
	r.log.Log("BuyerRepository", "INFO", fmt.Sprintf("Buyer with ID %d successfully deleted", id))
	return
}

func (r *BuyerRepository) Get() (buyers []model.Buyer, err error) {
	r.log.Log("BuyerRepository", "INFO", "initializing Get function")
	rows, err := r.db.Query("SELECT id, card_number_id,first_name,last_name FROM buyers")
	if err != nil {
		r.log.Log("BuyerRepository", "ERROR", fmt.Sprintf("Error: %v", err))
		return
	}

	defer rows.Close()

	for rows.Next() {
		var buyer model.Buyer
		err = rows.Scan(&buyer.ID, &buyer.CardNumberID, &buyer.FirstName, &buyer.LastName)

		if err != nil {
			r.log.Log("BuyerRepository", "ERROR", fmt.Sprintf("Error: %v", err))
			return
		}

		buyers = append(buyers, buyer)
	}
	r.log.Log("BuyerRepository", "INFO", fmt.Sprintf("retornando %v", buyers))
	return buyers, nil
}

func (r *BuyerRepository) GetByID(id int) (buyer model.Buyer, err error) {
	r.log.Log("BuyerRepository", "INFO", fmt.Sprintf("initializing GetByID function with parameter %d", id))
	row := r.db.QueryRow("SELECT id,card_number_id,first_name,last_name FROM buyers WHERE id=?", id)
	err = row.Scan(&buyer.ID, &buyer.CardNumberID, &buyer.FirstName, &buyer.LastName)

	if err != nil {
		if err == sql.ErrNoRows {
			err = customerror.NewBuyerError(http.StatusNotFound, customerror.ErrNotFound.Error(), "Buyer")
		}
		r.log.Log("BuyerRepository", "ERROR", fmt.Sprintf("Error: %v", err))
		return
	}
	r.log.Log("BuyerRepository", "INFO", fmt.Sprintf("returning %v", buyer))
	return
}

func (r *BuyerRepository) Post(newBuyer model.Buyer) (id int64, err error) {
	r.log.Log("BuyerRepository", "INFO", fmt.Sprintf("initializing Post function with parameter %v", newBuyer))
	prepare, err := r.db.Prepare("INSERT INTO buyers (card_number_id, first_name, last_name) VALUES (?,?,?)")
	if err != nil {
		r.log.Log("BuyerRepository", "ERROR", fmt.Sprintf("Error: %v", err))
		return
	}

	result, err := prepare.Exec(newBuyer.CardNumberID, newBuyer.FirstName, newBuyer.LastName)

	if err != nil {
		if err.(*mysql.MySQLError).Number == 1062 {
			err = customerror.NewBuyerError(http.StatusConflict, customerror.ErrConflict.Error(), "card_number_id")
		}
		r.log.Log("BuyerRepository", "ERROR", fmt.Sprintf("Error: %v", err))
		return
	}

	id, err = result.LastInsertId()
	r.log.Log("BuyerRepository", "INFO", fmt.Sprintf("returning inserted ID %d", id))
	return
}

func (r *BuyerRepository) Update(id int, newBuyer model.Buyer) (err error) {
	r.log.Log("BuyerRepository", "INFO", fmt.Sprintf("initializing Update function with parameter %v and %d", newBuyer, id))
	prepare, err := r.db.Prepare("UPDATE buyers SET card_number_id = ?, first_name = ?, last_name = ? WHERE id = ?")
	if err != nil {
		r.log.Log("BuyerRepository", "ERROR", fmt.Sprintf("Error: %v", err))
		return
	}

	_, err = prepare.Exec(newBuyer.CardNumberID, newBuyer.FirstName, newBuyer.LastName, id)

	if err != nil {
		if err.(*mysql.MySQLError).Number == 1062 {
			err = customerror.NewBuyerError(http.StatusNotFound, customerror.ErrConflict.Error(), "card_number_id")
		}
		r.log.Log("BuyerRepository", "ERROR", fmt.Sprintf("Error: %v", err))
		return
	}
	r.log.Log("BuyerRepository", "INFO", "Update completed successfully")
	return
}

func (r *BuyerRepository) CountPurchaseOrderByBuyerID(id int) (countBuyerPurchaseOrder model.BuyerPurchaseOrder, err error) {
	r.log.Log("BuyerRepository", "INFO", fmt.Sprintf("initializing CountPurchaseOrderByBuyerID function with parameter  %d", id))
	row := r.db.QueryRow("SELECT b.id, b.card_number_id, b.first_name, b.last_name, COUNT(po.id) as purchase_orders_count FROM buyers b LEFT JOIN purchase_orders po ON po.buyer_id = b.id WHERE b.id = ? GROUP BY b.id", id)
	err = row.Scan(&countBuyerPurchaseOrder.ID, &countBuyerPurchaseOrder.CardNumberID, &countBuyerPurchaseOrder.FirstName, &countBuyerPurchaseOrder.LastName, &countBuyerPurchaseOrder.PurchaseOrdersCount)

	if err != nil {
		if err == sql.ErrNoRows {
			err = customerror.NewBuyerError(http.StatusNotFound, customerror.ErrNotFound.Error(), "Buyer")
		}
		r.log.Log("BuyerRepository", "ERROR", fmt.Sprintf("Error: %v", err))
		return
	}
	r.log.Log("BuyerRepository", "INFO", fmt.Sprintf("Count done successfully, return: %v", countBuyerPurchaseOrder))
	return
}

func (r *BuyerRepository) CountPurchaseOrderBuyers() (countBuyerPurchaseOrder []model.BuyerPurchaseOrder, err error) {
	r.log.Log("BuyerRepository", "INFO", fmt.Sprintf("initializing CountPurchaseOrderBuyers function"))
	rows, err := r.db.Query("SELECT b.id, b.card_number_id, b.first_name, b.last_name, COUNT(po.id) as purchase_orders_count FROM buyers b LEFT JOIN purchase_orders po ON po.buyer_id = b.id GROUP BY b.id")

	if err != nil {
		r.log.Log("BuyerRepository", "ERROR", fmt.Sprintf("Error: %v", err))
		return
	}

	defer rows.Close()

	for rows.Next() {
		var buyerPurchaseOrder model.BuyerPurchaseOrder
		err = rows.Scan(&buyerPurchaseOrder.ID, &buyerPurchaseOrder.CardNumberID, &buyerPurchaseOrder.FirstName, &buyerPurchaseOrder.LastName, &buyerPurchaseOrder.PurchaseOrdersCount)

		if err != nil {
			r.log.Log("BuyerRepository", "ERROR", fmt.Sprintf("Error: %v", err))
			return
		}

		countBuyerPurchaseOrder = append(countBuyerPurchaseOrder, buyerPurchaseOrder)
	}
	r.log.Log("BuyerRepository", "INFO", fmt.Sprintf("Count done successfully, return: %v", countBuyerPurchaseOrder))
	return
}

func NewBuyerRepository(db *sql.DB, log logger.Logger) *BuyerRepository {
	return &BuyerRepository{db: db, log: log}
}
