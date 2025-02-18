package repository

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/maxwelbm/alkemy-g7.git/pkg/logger"

	"github.com/go-sql-driver/mysql"
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/maxwelbm/alkemy-g7.git/pkg/customerror"
)

type Carriers struct {
	db  *sql.DB
	log logger.Logger
}

func NewCarriersRepository(db *sql.DB, log logger.Logger) *Carriers {
	return &Carriers{db: db, log: log}
}

func (r *Carriers) GetByID(id int) (carrier model.Carries, err error) {
	row := r.db.QueryRow("SELECT `id`, `cid`, `company_name`, `address`, `telephone`, `locality_id` FROM `carriers` WHERE `id` = ?", id)

	err = row.Scan(&carrier.ID, &carrier.CID, &carrier.CompanyName, &carrier.Address, &carrier.Telephone, &carrier.LocalityID)

	if err != nil {
		if err == sql.ErrNoRows {
			err = customerror.NewCarrierError(customerror.ErrNotFound.Error(), "carrier", http.StatusNotFound)
		}

		return
	}

	return
}

func (r *Carriers) PostCarrier(carrier model.Carries) (id int64, err error) {
	result, err := r.db.Exec(
		"INSERT INTO `carriers` (`cid`, `company_name`, `address`, `telephone`, `locality_id`) VALUES (?, ?, ?, ?, ?)",
		carrier.CID, carrier.CompanyName, carrier.Address, carrier.Telephone, carrier.LocalityID,
	)

	if err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) {
			switch mysqlErr.Number {
			case 1062:
				err = customerror.NewCarrierError(customerror.ErrConflict.Error(), "cid", http.StatusConflict)
			default:
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
