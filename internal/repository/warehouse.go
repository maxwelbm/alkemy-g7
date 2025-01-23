package repository

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/go-sql-driver/mysql"
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/maxwelbm/alkemy-g7.git/pkg/custom_error"
)

func NewWareHouseRepository(db *sql.DB) *WarehouseMysql {
	return &WarehouseMysql{db}
}

type WarehouseMysql struct {
	db *sql.DB
}

func (r *WarehouseMysql) GetAllWareHouse() (w []model.WareHouse, err error) {
	rows, err := r.db.Query("SELECT w.id, w.warehouse_code, w.address, w.telephone, w.minimum_capacity, w.minimum_temperature FROM warehouses w")
	if err != nil {
		return
	}

	for rows.Next() {
		var warehouse model.WareHouse
		err = rows.Scan(&warehouse.Id, &warehouse.WareHouseCode, &warehouse.Address, &warehouse.Telephone, &warehouse.MinimunCapacity, &warehouse.MinimunTemperature)
		if err != nil {
			return
		}

		w = append(w, warehouse)
	}

	err = rows.Err()
	if err != nil {
		return
	}

	return
}

func (r *WarehouseMysql) GetByIdWareHouse(id int) (w model.WareHouse, err error) {
	row := r.db.QueryRow("SELECT w.id, w.warehouse_code, w.address, w.telephone, w.minimum_capacity, w.minimum_temperature FROM warehouses w WHERE w.id = ?", id)

	err = row.Scan(&w.Id, &w.WareHouseCode, &w.Address, &w.Telephone, &w.MinimunCapacity, &w.MinimunTemperature)
	if err != nil {
		if err == sql.ErrNoRows {
			err = custom_error.NewWareHouseError(custom_error.ErrNotFound.Error(), "warehouse", http.StatusNotFound)
		}
		return
	}

	return
}

func (r *WarehouseMysql) PostWareHouse(warehouse model.WareHouse) (id int64, err error) {
	result, err := r.db.Exec(
		"INSERT INTO `warehouses` (`warehouse_code`, `address`, `telephone`, `minimum_capacity`, `minimum_temperature`) VALUES (?, ?, ?, ?, ?)",
		warehouse.WareHouseCode, warehouse.Address, warehouse.Telephone, warehouse.MinimunCapacity, warehouse.MinimunTemperature,
	)

	if err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) {
			switch mysqlErr.Number {
			case 1062:
				err = custom_error.NewWareHouseError(custom_error.ErrConflict.Error(), "warehouse_code", http.StatusConflict)
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

func (r *WarehouseMysql) UpdateWareHouse(id int, warehouse model.WareHouse) (err error) {

	_, err = r.db.Exec(
		"UPDATE warehouses w SET w.warehouse_code = ?, w.address = ?, w.telephone = ?, w.minimum_capacity = ?, w.minimum_temperature = ? WHERE w.id = ?",
		warehouse.WareHouseCode, warehouse.Address, warehouse.Telephone, warehouse.MinimunCapacity, warehouse.MinimunTemperature, warehouse.Id,
	)
	if err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) {
			switch mysqlErr.Number {
			case 1062:
				err = custom_error.NewWareHouseError(custom_error.ErrConflict.Error(), "warehouse_code", http.StatusConflict)
			default:
				// ...
			}
			return
		}

		return
	}

	return
}

func (r *WarehouseMysql) DeleteByIdWareHouse(id int) (err error) {

	_, err = r.db.Exec("DELETE FROM `warehouses` WHERE `id` = ?", id)

	if err != nil {
		return
	}

	return
}
