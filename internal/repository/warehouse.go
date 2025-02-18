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

func NewWareHouseRepository(db *sql.DB, log logger.Logger) *WarehouseMysql {
	return &WarehouseMysql{db, log}
}

type WarehouseMysql struct {
	db  *sql.DB
	log logger.Logger
}

func (r *WarehouseMysql) GetAllWareHouse() (w []model.WareHouse, err error) {
	r.log.Log("WareHouseRepository", "INFO", "initializing GetAllWareHouse function")
	rows, err := r.db.Query("SELECT w.id, w.warehouse_code, w.address, w.telephone, w.minimum_capacity, w.minimum_temperature FROM warehouses w")
	if err != nil {
		r.log.Log("WareHouseRepository", "ERROR", fmt.Sprintf("Error: %v", err))

		return
	}

	for rows.Next() {
		var warehouse model.WareHouse
		err = rows.Scan(&warehouse.ID, &warehouse.WareHouseCode, &warehouse.Address, &warehouse.Telephone, &warehouse.MinimunCapacity, &warehouse.MinimunTemperature)

		if err != nil {
			r.log.Log("WareHouseRepository", "ERROR", fmt.Sprintf("Error: %v", err))

			return
		}

		w = append(w, warehouse)
	}

	err = rows.Err()
	if err != nil {
		r.log.Log("WareHouseRepository", "ERROR", fmt.Sprintf("Error: %v", err))

		return
	}
	r.log.Log("WareHouseRepository", "INFO", "GetAllWareHouse completed successfully")

	return
}

func (r *WarehouseMysql) GetByIDWareHouse(id int) (w model.WareHouse, err error) {

	r.log.Log("WareHouseRepository", "INFO", "initializing GetByIDWareHouse function")

	row := r.db.QueryRow("SELECT w.id, w.warehouse_code, w.address, w.telephone, w.minimum_capacity, w.minimum_temperature FROM warehouses w WHERE w.id = ?", id)

	err = row.Scan(&w.ID, &w.WareHouseCode, &w.Address, &w.Telephone, &w.MinimunCapacity, &w.MinimunTemperature)
	if err != nil {
		if err == sql.ErrNoRows {
			r.log.Log("WareHouseRepository", "ERROR", fmt.Sprintf("Error: %v", err))
			err = customerror.NewWareHouseError(customerror.ErrNotFound.Error(), "warehouse", http.StatusNotFound)
		}
	}
	r.log.Log("WareHouseRepository", "INFO", "GetByIDWareHouse completed successfully")

	return
}

func (r *WarehouseMysql) PostWareHouse(warehouse model.WareHouse) (id int64, err error) {
	r.log.Log("WareHouseRepository", "INFO", "initializing PostWareHouse function")

	result, err := r.db.Exec(
		"INSERT INTO `warehouses` (`warehouse_code`, `address`, `telephone`, `minimum_capacity`, `minimum_temperature`) VALUES (?, ?, ?, ?, ?)",
		warehouse.WareHouseCode, warehouse.Address, warehouse.Telephone, warehouse.MinimunCapacity, warehouse.MinimunTemperature,
	)

	if err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) {
			r.log.Log("WareHouseRepository", "ERROR", fmt.Sprintf("Error: %v", err))

			switch mysqlErr.Number {
			case 1062:
				err = customerror.NewWareHouseError(customerror.ErrConflict.Error(), "warehouse_code", http.StatusConflict)
			}
			r.log.Log("WareHouseRepository", "ERROR", fmt.Sprintf("Error: %v", err))

			return
		}

		return
	}

	id, err = result.LastInsertId()
	if err != nil {
		r.log.Log("WareHouseRepository", "ERROR", fmt.Sprintf("Error: %v", err))

		return
	}
	r.log.Log("WareHouseRepository", "INFO", "PostWareHouse completed successfully")

	return
}

func (r *WarehouseMysql) UpdateWareHouse(id int, warehouse model.WareHouse) (err error) {
	r.log.Log("WareHouseRepository", "INFO", "initializing UpdateWareHouse function")

	_, err = r.db.Exec(
		"UPDATE warehouses w SET w.warehouse_code = ?, w.address = ?, w.telephone = ?, w.minimum_capacity = ?, w.minimum_temperature = ? WHERE w.id = ?",
		warehouse.WareHouseCode, warehouse.Address, warehouse.Telephone, warehouse.MinimunCapacity, warehouse.MinimunTemperature, warehouse.ID,
	)
	if err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) {
			r.log.Log("WareHouseRepository", "ERROR", fmt.Sprintf("Error: %v", err))

			switch mysqlErr.Number {
			case 1062:
				err = customerror.NewWareHouseError(customerror.ErrConflict.Error(), "warehouse_code", http.StatusConflict)
			}
		}
		r.log.Log("WareHouseRepository", "ERROR", fmt.Sprintf("Error: %v", err))

		return
	}
	r.log.Log("WareHouseRepository", "INFO", "UpdateWareHouse completed successfully")

	return
}

func (r *WarehouseMysql) DeleteByIDWareHouse(id int) (err error) {
	r.log.Log("WareHouseRepository", "INFO", "initializing DeleteByIDWareHouse function")

	_, err = r.db.Exec("DELETE FROM `warehouses` WHERE `id` = ?", id)

	if err != nil {
		r.log.Log("WareHouseRepository", "ERROR", fmt.Sprintf("Error: %v", err))
		return
	}

	r.log.Log("WareHouseRepository", "INFO", "DeleteWareHouse completed successfully")

	return
}
