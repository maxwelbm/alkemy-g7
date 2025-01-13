package repository

import (
	"database/sql"

	"github.com/go-sql-driver/mysql"
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/maxwelbm/alkemy-g7.git/pkg/custom_error"
)

type BuyerRepository struct {
	db *sql.DB
}

func (r BuyerRepository) Delete(id int) (err error) {

	_, err = r.db.Exec("DELETE FROM buyers WHERE id = ?", id)
	if err != nil {

		if err.(*mysql.MySQLError).Number == 1451 {
			err = custom_error.CustomError{Object: id, Err: custom_error.Conflict}
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

	for rows.Next() {
		var buyer model.Buyer

		err = rows.Scan(&buyer.Id, &buyer.CardNumberId, &buyer.FirstName, &buyer.LastName)
		if err != nil {
			return
		}

		buyers = append(buyers, buyer)
	}

	return buyers, nil

}

func (r *BuyerRepository) GetById(id int) (buyer model.Buyer, err error) {

	row := r.db.QueryRow("SELECT id, card_number_id,first_name,last_name FROM buyers WHERE id= ?", id)

	err = row.Scan(&buyer.Id, &buyer.CardNumberId, &buyer.FirstName, &buyer.LastName)

	if err != nil {
		if err == sql.ErrNoRows {
			err = custom_error.CustomError{Object: id, Err: custom_error.NotFound}
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

	result, err := prepare.Exec(newBuyer.CardNumberId, newBuyer.FirstName, newBuyer.LastName)

	if err != nil {
		if err.(*mysql.MySQLError).Number == 1062 {
			err = custom_error.CustomError{Object: id, Err: custom_error.Conflict}
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

	_, err = prepare.Exec(newBuyer.CardNumberId, newBuyer.FirstName, newBuyer.LastName, id)

	if err != nil {
		if err.(*mysql.MySQLError).Number == 1062 {
			err = custom_error.CustomError{Object: id, Err: custom_error.Conflict}
		}
		return
	}

	return
}

func NewBuyerRepository(db *sql.DB) *BuyerRepository {
	return &BuyerRepository{db: db}
}
