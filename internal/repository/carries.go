package repository

import (
	"database/sql"
	"errors"
	"fmt"
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
	r.log.Log("CarriesRepository", "INFO", "initializing GetByID function")

	err = row.Scan(&carrier.ID, &carrier.CID, &carrier.CompanyName, &carrier.Address, &carrier.Telephone, &carrier.LocalityID)

	if err != nil {
		if err == sql.ErrNoRows {
			r.log.Log("CarriesRepository", "ERROR", fmt.Sprintf("Error: %v", err))

			err = customerror.NewCarrierError(customerror.ErrNotFound.Error(), "carrier", http.StatusNotFound)
		}
		r.log.Log("CarriesRepository", "ERROR", fmt.Sprintf("Error: %v", err))

		return
	}
	r.log.Log("CarriesRepository", "INFO", "GetByIDCarries completed successfully")

	return
}

func (r *Carriers) PostCarrier(carrier model.Carries) (id int64, err error) {
	r.log.Log("CarriesRepository", "INFO", "initializing PostCarrier function")

	result, err := r.db.Exec(
		"INSERT INTO `carriers` (`cid`, `company_name`, `address`, `telephone`, `locality_id`) VALUES (?, ?, ?, ?, ?)",
		carrier.CID, carrier.CompanyName, carrier.Address, carrier.Telephone, carrier.LocalityID,
	)

	if err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) {
			r.log.Log("CarriesRepository", "ERROR", fmt.Sprintf("Error: %v", err))

			switch mysqlErr.Number {
			case 1062:
				err = customerror.NewCarrierError(customerror.ErrConflict.Error(), "cid", http.StatusConflict)
			default:
			}
			r.log.Log("CarriesRepository", "ERROR", fmt.Sprintf("Error: %v", err))

			return
		}
		r.log.Log("CarriesRepository", "ERROR", fmt.Sprintf("Error: %v", err))

		return
	}

	id, err = result.LastInsertId()
	if err != nil {
		r.log.Log("CarriesRepository", "ERROR", fmt.Sprintf("Error: %v", err))

		return
	}
	r.log.Log("CarriesRepository", "INFO", "PostCarrier completed successfully")

	return
}
