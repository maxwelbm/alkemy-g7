package repository

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/go-sql-driver/mysql"
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/maxwelbm/alkemy-g7.git/pkg/customError"
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

func (r *WarehouseMysql) GetByIDWareHouse(id int) (w model.WareHouse, err error) {
	row := r.db.QueryRow("SELECT w.id, w.warehouse_code, w.address, w.telephone, w.minimum_capacity, w.minimum_temperature FROM warehouses w WHERE w.id = ?", id)

	err = row.Scan(&w.Id, &w.WareHouseCode, &w.Address, &w.Telephone, &w.MinimunCapacity, &w.MinimunTemperature)
	if err != nil {
		if err == sql.ErrNoRows {
			err = customError.NewWareHouseError(customError.ErrNotFound.Error(), "warehouse", http.StatusNotFound)
		}
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
				err = customError.NewWareHouseError(customError.ErrConflict.Error(), "warehouse_code", http.StatusConflict)
			}
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
				err = customError.NewWareHouseError(customError.ErrConflict.Error(), "warehouse_code", http.StatusConflict)
			}
		}

		return
	}

	return
}

func (r *WarehouseMysql) DeleteByIDWareHouse(id int) (err error) {
	_, err = r.db.Exec("DELETE FROM `warehouses` WHERE `id` = ?", id)

	if err != nil {
		return
	}

	return
}
