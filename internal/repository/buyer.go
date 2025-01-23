package repository

import (
	"database/sql"
	"net/http"

	"github.com/go-sql-driver/mysql"
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/maxwelbm/alkemy-g7.git/pkg/custom_error"
)

type BuyerRepository struct {
	db *sql.DB
}

func (r *BuyerRepository) Delete(id int) (err error) {
	_, err = r.db.Exec("DELETE FROM buyers WHERE id = ?", id)
	if err != nil {
		if err.(*mysql.MySQLError).Number == 1451 {
			err = custom_error.NewBuyerError(http.StatusConflict, custom_error.ErrDependencies.Error(), "Buyer")
		}

		return
	}

	return
}

func (r *BuyerRepository) Get() (buyers []model.Buyer, err error) {
	rows, err := r.db.Query("SELECT id, card_number_id,first_name,last_name FROM buyers")
	if err != nil {
		return
	}

	defer rows.Close()

	for rows.Next() {
		var buyer model.Buyer
		err = rows.Scan(&buyer.ID, &buyer.CardNumberID, &buyer.FirstName, &buyer.LastName)

		if err != nil {
			return
		}

		buyers = append(buyers, buyer)
	}

	return buyers, nil
}

func (r *BuyerRepository) GetByID(id int) (buyer model.Buyer, err error) {
	row := r.db.QueryRow("SELECT id, card_number_id,first_name,last_name FROM buyers WHERE id= ?", id)
	err = row.Scan(&buyer.ID, &buyer.CardNumberID, &buyer.FirstName, &buyer.LastName)

	if err != nil {
		if err == sql.ErrNoRows {
			err = custom_error.NewBuyerError(http.StatusNotFound, custom_error.ErrNotFound.Error(), "Buyer")
		}

		return
	}

	return
}

func (r *BuyerRepository) Post(newBuyer model.Buyer) (id int64, err error) {
	prepare, err := r.db.Prepare("INSERT INTO buyers (card_number_id, first_name, last_name) VALUES (?,?,?)")
	if err != nil {
		return
	}

	result, err := prepare.Exec(newBuyer.CardNumberID, newBuyer.FirstName, newBuyer.LastName)

	if err != nil {
		if err.(*mysql.MySQLError).Number == 1062 {
			err = custom_error.NewBuyerError(http.StatusConflict, custom_error.ErrConflict.Error(), "card_number_id")
		}

		return
	}

	id, err = result.LastInsertId()

	return
}

func (r *BuyerRepository) Update(id int, newBuyer model.Buyer) (err error) {
	prepare, err := r.db.Prepare("UPDATE buyers SET card_number_id = ?, first_name = ?, last_name = ? WHERE id = ?")
	if err != nil {
		return
	}

	_, err = prepare.Exec(newBuyer.CardNumberID, newBuyer.FirstName, newBuyer.LastName, id)

	if err != nil {
		if err.(*mysql.MySQLError).Number == 1062 {
			err = custom_error.NewBuyerError(http.StatusNotFound, custom_error.ErrConflict.Error(), "card_number_id")
		}

		return
	}

	return
}

func (r *BuyerRepository) CountPurchaseOrderByBuyerID(id int) (countBuyerPurchaseOrder model.BuyerPurchaseOrder, err error) {
	row := r.db.QueryRow("SELECT b.id, b.card_number_id, b.first_name, b.last_name, COUNT(po.id) as purchase_orders_count FROM buyers b LEFT JOIN purchase_orders po ON po.buyer_id = b.id WHERE b.id = ? GROUP BY b.id", id)
	err = row.Scan(&countBuyerPurchaseOrder.ID, &countBuyerPurchaseOrder.CardNumberID, &countBuyerPurchaseOrder.FirstName, &countBuyerPurchaseOrder.LastName, &countBuyerPurchaseOrder.PurchaseOrdersCount)

	if err != nil {
		if err == sql.ErrNoRows {
			err = custom_error.NewBuyerError(http.StatusNotFound, custom_error.ErrNotFound.Error(), "Buyer")
		}

		return
	}

	return
}

func (r *BuyerRepository) CountPurchaseOrderBuyers() (countBuyerPurchaseOrder []model.BuyerPurchaseOrder, err error) {
	rows, err := r.db.Query("SELECT b.id, b.card_number_id, b.first_name, b.last_name, COUNT(po.id) as purchase_orders_count FROM buyers b LEFT JOIN purchase_orders po ON po.buyer_id = b.id GROUP BY b.id")

	if err != nil {
		return
	}

	defer rows.Close()

	for rows.Next() {
		var buyerPurchaseOrder model.BuyerPurchaseOrder
		err = rows.Scan(&buyerPurchaseOrder.ID, &buyerPurchaseOrder.CardNumberID, &buyerPurchaseOrder.FirstName, &buyerPurchaseOrder.LastName, &buyerPurchaseOrder.PurchaseOrdersCount)

		if err != nil {
			return
		}

		countBuyerPurchaseOrder = append(countBuyerPurchaseOrder, buyerPurchaseOrder)
	}

	return
}

func NewBuyerRepository(db *sql.DB) *BuyerRepository {
	return &BuyerRepository{db: db}
}
