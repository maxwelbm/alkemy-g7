package repository

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/go-sql-driver/mysql"
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/maxwelbm/alkemy-g7.git/pkg/custom_error"
)

type Carriers struct {
	db *sql.DB
}

func NewCarriersRepository(db *sql.DB) *Carriers {
	return &Carriers{db: db}
}

func (r *Carriers) GetById(id int) (carrier model.Carries, err error) {
	row := r.db.QueryRow("SELECT `id`,`cid`, `company_name`, `address`, `telephone`, `locality_id` FROM `carriers` WHERE `id` = ?", id)

	err = row.Scan(&carrier.Id, &carrier.CID, &carrier.CompanyName, &carrier.Address, &carrier.Telephone, &carrier.LocalityId)
	if err != nil {
		if err == sql.ErrNoRows {
			err = custom_error.NewCarrierError(custom_error.ErrNotFound.Error(), "carrier", http.StatusNotFound)
		}
		return
	}

	return

}

func (r *Carriers) PostCarrier(carrier model.Carries) (id int64, err error) {
	result, err := r.db.Exec(
		"INSERT INTO `carriers` (`cid`, `company_name`, `address`, `telephone`, `locality_id`) VALUES (?, ?, ?, ?, ?)",
		carrier.CID, carrier.CompanyName, carrier.Address, carrier.Telephone, carrier.LocalityId,
	)

	if err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) {
			switch mysqlErr.Number {
			case 1062:
				err = custom_error.NewCarrierError(custom_error.ErrConflict.Error(), "cid", http.StatusConflict)
			default:
				// ...
			}
			return
		}
		return
	}

	id, err = result.LastInsertId()
	if err != nil {
		return
	}
	return
}
