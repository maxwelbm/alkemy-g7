package repository

import (
	"database/sql"
	"errors"
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
)

func CreateRepositorySellers(db *sql.DB) *SellersRepository {
	return &SellersRepository{db}
}

type SellersRepository struct {
	db *sql.DB
}

func (rp *SellersRepository) Get() (sellers []model.Seller, err error) {
	query := "SELECT `id`, `cid`, `company_name`, `address`, `telephone`, `locality_id` FROM `sellers`"
	rows, err := rp.db.Query(query)
	if err != nil {
		return
	}

	for rows.Next() {
		var seller model.Seller
		err = rows.Scan(&seller.ID, &seller.CID, &seller.CompanyName, &seller.Address, &seller.Telephone, &seller.Locality)
		if err != nil {
			return
		}
		sellers = append(sellers, seller)
	}

	err = rows.Err()
	if err != nil {
		return
	}

	return
}

func (rp *SellersRepository) GetById(id int) (sl model.Seller, err error) {
	query := "SELECT `id`, `cid`, `company_name`, `address`, `telephone`, `locality_id` FROM `sellers` WHERE `id` = ?"
	row := rp.db.QueryRow(query, id)

	err = row.Scan(&sl.ID, &sl.CID, &sl.CompanyName, &sl.Address, &sl.Telephone, &sl.Locality)

	if errors.Is(err, sql.ErrNoRows) {
		err = model.ErrorSellerNotFound
		return
	}
	return
}

func (rp *SellersRepository) Post(seller *model.Seller) (sl model.Seller, err error) {
	query := "INSERT INTO `sellers` (`cid`, `company_name`, `address`, `telephone`, `locality_id`) VALUES (?, ?, ?, ?, ?)"
	result, err := rp.db.Exec(query, (*seller).CID, (*seller).CompanyName, (*seller).Address, (*seller).Telephone, (*seller).Locality)
	if err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) {
			switch mysqlErr.Number {
			case 1062:
				err = model.ErrorCIDAlreadyExist
			case 1064:
				err = model.ErrorInvalidJSONFormat
			case 1048:
				err = model.ErrorNullAttribute
			}
			return
		}
	}

	id, err := result.LastInsertId()
	if err != nil {
		return
	}
	sl, _ = rp.GetById(int(id))

	return
}

func (rp *SellersRepository) Patch(id int, seller *model.Seller) (sl model.Seller, err error) {
	query := "UPDATE `sellers` SET"

	var updates []string
	var args []interface{}

	if seller.CID != 0 {
		updates = append(updates, "`cid` = ?")
		args = append(args, (*seller).CID)
	}
	if seller.CompanyName != "" {
		updates = append(updates, "`company_name` = ?")
		args = append(args, (*seller).CompanyName)
	}
	if seller.Address != "" {
		updates = append(updates, "`address` = ?")
		args = append(args, (*seller).Address)
	}
	if seller.Telephone != "" {
		updates = append(updates, "`telephone` = ?")
		args = append(args, (*seller).Telephone)
	}
	if seller.Locality != 0 {
		updates = append(updates, "`locality_id` = ?")
		args = append(args, (*seller).Locality)
	}

	if len(updates) > 0 {
		query = query + " " + strings.Join(updates, ", ") + " WHERE `id` = ?"
		args = append(args, id)
	} else {
		err = model.ErrorNullAttribute
		return
	}

	_, err = rp.db.Exec(query, args...)
	if err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) {
			switch mysqlErr.Number {
			case 1062:
				err = model.ErrorCIDAlreadyExist
			case 1064:
				err = model.ErrorInvalidJSONFormat
			case 1048:
				err = model.ErrorNullAttribute
			}

			return
		}
	}
	sl, _ = rp.GetById(int(id))

	return
}

func (rp *SellersRepository) Delete(id int) error {
	query := "DELETE FROM `sellers` WHERE `id` = ?"
	_, err := rp.db.Exec(query, id)
	return err
}
