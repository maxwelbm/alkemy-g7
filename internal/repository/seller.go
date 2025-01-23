package repository

import (
	"database/sql"
	"errors"
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	er "github.com/maxwelbm/alkemy-g7.git/pkg/custom_error"
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
		return sellers, er.ErrDefaultSQLSeller
	}

	return
}

func (rp *SellersRepository) GetByID(id int) (sl model.Seller, err error) {
	query := "SELECT `id`, `cid`, `company_name`, `address`, `telephone`, `locality_id` FROM `sellers` WHERE `id` = ?"
	row := rp.db.QueryRow(query, id)

	err = row.Scan(&sl.ID, &sl.CID, &sl.CompanyName, &sl.Address, &sl.Telephone, &sl.Locality)

	if errors.Is(err, sql.ErrNoRows) {
		err = er.ErrSellerNotFound
		return
	}
	
	return
}

func (rp *SellersRepository) Post(seller *model.Seller) (sl model.Seller, err error) {
	query := "INSERT INTO `sellers` (`cid`, `company_name`, `address`, `telephone`, `locality_id`) VALUES (?, ?, ?, ?, ?)"
	result, err := rp.db.Exec(query, (*seller).CID, (*seller).CompanyName, (*seller).Address, (*seller).Telephone, (*seller).Locality)
	err = rp.validateSQLError(err)
	
	if err != nil {
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		return
	}
	
	sl, _ = rp.GetByID(int(id))

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

	length := len(updates)
	
	if length > 0 {
		query = query + " " + strings.Join(updates, ", ") + " WHERE `id` = ?"
		args = append(args, id)
	}

	_, err = rp.db.Exec(query, args...)
	err = rp.validateSQLError(err)
	
	if err != nil {
		return sl, err
	}
	
	sl, _ = rp.GetByID(id)

	return sl, err
}

func (rp *SellersRepository) Delete(id int) error {
	query := "DELETE FROM `sellers` WHERE `id` = ?"
	_, err := rp.db.Exec(query, id)
	err = rp.validateSQLError(err)
	
	if err != nil {
		return err
	}
	
	return err
}

func (rp *SellersRepository) validateSQLError(err error) (e error) {
	if err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) {
			switch mysqlErr.Number {
			case 1062:
				e = er.ErrCIDSellerAlreadyExist
			case 1064:
				e = er.ErrInvalidSellerJSONFormat
			case 1048:
				e = er.ErrNullSellerAttribute
			case 1451:
				e = er.ErrNotSellerDelete
			default:
				e = er.ErrDefaultSQLSeller
			}
		}
	}
	
	return e
}
