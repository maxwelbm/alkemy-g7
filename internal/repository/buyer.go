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

func (r BuyerRepository) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM `buyers` WHERE `id` = ?", id)
	if err != nil {
		return err
	}

	return err
}

func (br *BuyerRepository) Get() (map[int]model.Buyer, error) {

	// if len(br.dbBuyer.TbBuyer) == 0 {
	// 	return nil, fmt.Errorf("no buyers found")
	// }
	// return br.dbBuyer.TbBuyer, nil

	return nil, nil

}

func (br *BuyerRepository) GetById(id int) (model.Buyer, error) {
	// buyer, ok := br.dbBuyer.TbBuyer[id]

	// if !ok {
	// 	return model.Buyer{}, &custom_error.CustomError{Object: id, Err: custom_error.NotFound}
	// }

	// return buyer, nil

	return model.Buyer{}, nil
}

func (br *BuyerRepository) Post(newBuyer model.Buyer) (model.Buyer, error) {
	// BuyerExists := isCardNumberIdExists(newBuyer.CardNumberId, br)
	// lastId := getLastIdBuyer(br.dbBuyer.TbBuyer)

	// if BuyerExists {
	// 	return model.Buyer{}, &custom_error.CustomError{Object: newBuyer, Err: custom_error.Conflict}
	// }
	// buyer := model.Buyer{
	// 	Id:           lastId,
	// 	CardNumberId: newBuyer.CardNumberId,
	// 	FirstName:    newBuyer.FirstName,
	// 	LastName:     newBuyer.LastName,
	// }

	// br.dbBuyer.TbBuyer[buyer.Id] = buyer

	// return br.GetById(buyer.Id)

	return model.Buyer{}, nil

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
