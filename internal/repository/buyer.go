package repository

import (
	"database/sql"
	"errors"

	"github.com/go-sql-driver/mysql"
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/maxwelbm/alkemy-g7.git/pkg/custom_error"
)

type BuyerRepository struct {
	db *sql.DB
}

func (r BuyerRepository) Delete(id int) (err error) {

	_, err = r.db.Exec("DELETE FROM buyer WHERE id = ?", id)
	if err != nil {
		return
	}

	return
}

func (r *BuyerRepository) Get() (buyers []model.Buyer, err error) {

	rows, err := r.db.Query("SELECT id, card_number_id,first_name,last_name FROM buyer")
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

	row := r.db.QueryRow("SELECT id, card_number_id,first_name,last_name FROM buyer WHERE id= ?", id)

	err = row.Scan(&buyer.Id, &buyer.CardNumberId, &buyer.FirstName, &buyer.LastName)

	if err != nil {
		if err == sql.ErrNoRows {
			err = custom_error.NotFound
		}
		return
	}

	return
}

func (r *BuyerRepository) Post(newBuyer model.Buyer) (id int64, err error) {

	prepare, err := r.db.Prepare("INSERT INTO buyer (card_number_id, first_name, last_name) VALUES (?,?,?)")

	if err != nil {
		return
	}

	result, err := prepare.Exec(newBuyer.CardNumberId, newBuyer.FirstName, newBuyer.LastName)

	if err != nil {
		if err.(*mysql.MySQLError).Number == 1062 {
			err = custom_error.Conflict
		}
		return
	}

	id, err = result.LastInsertId()

	return

}

func (r *BuyerRepository) Update(id int, newBuyer model.Buyer) error {

	_, err := r.db.Exec(
		"UPDATE `buyers` SET `card_number_id` = ?, `first_name` = ?, `last_name` = ? WHERE `id` = ?",
		newBuyer.CardNumberId, newBuyer.FirstName, newBuyer.LastName, newBuyer.Id,
	)
	if err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) {
			switch mysqlErr.Number {
			case 1062:
				err = custom_error.Conflict
			default:

			}
			return err
		}
	}

	return err
}

func NewBuyerRepository(db *sql.DB) *BuyerRepository {
	return &BuyerRepository{db: db}
}

// func isCardNumberIdExists(CardNumberId string, br *BuyerRepository) bool {

// 	for _, b := range br.dbBuyer.TbBuyer {
// 		if b.CardNumberId == CardNumberId {
// 			return true
// 		}
// 	}

// 	return false
// }
