package repository_test

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-sql-driver/mysql"
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/maxwelbm/alkemy-g7.git/internal/repository"
	"github.com/maxwelbm/alkemy-g7.git/pkg/customerror"
	"github.com/stretchr/testify/assert"
)

func TestWarehouseMysql_GetAllWareHouse(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer db.Close()

	rp := repository.NewWareHouseRepository(db, logMock)

	t.Run("Success GetAllWareHouse", func(t *testing.T) {
		expectedWarehouses := []model.WareHouse{
			{
				ID:                 1,
				WareHouseCode:      "WH001",
				Address:            "123 Main St",
				Telephone:          "1234567890",
				MinimunCapacity:    100,
				MinimunTemperature: 20,
			},
			{
				ID:                 2,
				WareHouseCode:      "WH002",
				Address:            "456 Elm St",
				Telephone:          "0987654321",
				MinimunCapacity:    200,
				MinimunTemperature: 10,
			},
		}

		rows := sqlmock.NewRows([]string{"id", "warehouse_code", "address", "telephone", "minimum_capacity", "minimum_temperature"}).
			AddRow(expectedWarehouses[0].ID, expectedWarehouses[0].WareHouseCode, expectedWarehouses[0].Address, expectedWarehouses[0].Telephone, expectedWarehouses[0].MinimunCapacity, expectedWarehouses[0].MinimunTemperature).
			AddRow(expectedWarehouses[1].ID, expectedWarehouses[1].WareHouseCode, expectedWarehouses[1].Address, expectedWarehouses[1].Telephone, expectedWarehouses[1].MinimunCapacity, expectedWarehouses[1].MinimunTemperature)

		expectedQuery := "SELECT w.id, w.warehouse_code, w.address, w.telephone, w.minimum_capacity, w.minimum_temperature FROM warehouses w"
		mock.ExpectQuery(expectedQuery).WillReturnRows(rows)

		warehouses, err := rp.GetAllWareHouse()

		assert.NoError(t, err)
		assert.Equal(t, expectedWarehouses, warehouses)
	})
	t.Run("Error GetAllWareHouse", func(t *testing.T) {
		expectedQuery := "SELECT w.id, w.warehouse_code, w.address, w.telephone, w.minimum_capacity, w.minimum_temperature FROM warehouses w"
		mock.ExpectQuery(expectedQuery).WillReturnError(errors.New("database error"))

		warehouses, err := rp.GetAllWareHouse()

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "database error")
		assert.Nil(t, warehouses)
	})

	t.Run("Empty Result GetAllWareHouse", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "warehouse_code", "address", "telephone", "minimum_capacity", "minimum_temperature"})

		expectedQuery := "SELECT w.id, w.warehouse_code, w.address, w.telephone, w.minimum_capacity, w.minimum_temperature FROM warehouses w"
		mock.ExpectQuery(expectedQuery).WillReturnRows(rows)

		warehouses, err := rp.GetAllWareHouse()

		assert.NoError(t, err)
		assert.Empty(t, warehouses)
	})
}

func TestWarehouseMysql_GetByIDWareHouse(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rp := repository.NewWareHouseRepository(db, logMock)

	t.Run("Success GetByIDWareHouse", func(t *testing.T) {
		expectedWarehouse := model.WareHouse{
			ID:                 1,
			WareHouseCode:      "WH001",
			Address:            "123 Main St",
			Telephone:          "1234567890",
			MinimunCapacity:    100,
			MinimunTemperature: 20,
		}

		expectedQuery := "SELECT w.id, w.warehouse_code, w.address, w.telephone, w.minimum_capacity, w.minimum_temperature FROM warehouses w WHERE w.id = ?"
		mock.ExpectQuery(expectedQuery).
			WithArgs(expectedWarehouse.ID).
			WillReturnRows(sqlmock.NewRows([]string{"id", "warehouse_code", "address", "telephone", "minimum_capacity", "minimum_temperature"}).
				AddRow(expectedWarehouse.ID, expectedWarehouse.WareHouseCode, expectedWarehouse.Address, expectedWarehouse.Telephone, expectedWarehouse.MinimunCapacity, expectedWarehouse.MinimunTemperature))

		warehouse, err := rp.GetByIDWareHouse(expectedWarehouse.ID)

		assert.NoError(t, err)
		assert.Equal(t, expectedWarehouse, warehouse)
	})

	t.Run("Error GetByIDWareHouse", func(t *testing.T) {
		expectedQuery := "SELECT w.id, w.warehouse_code, w.address, w.telephone, w.minimum_capacity, w.minimum_temperature FROM warehouses w WHERE w.id = ?"
		mock.ExpectQuery(expectedQuery).
			WithArgs(1).
			WillReturnError(errors.New("database error"))

		warehouse, err := rp.GetByIDWareHouse(1)

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "database error")
		assert.Equal(t, model.WareHouse{}, warehouse)
	})

	t.Run("NotFound GetByIDWareHouse", func(t *testing.T) {
		expectedQuery := "SELECT w.id, w.warehouse_code, w.address, w.telephone, w.minimum_capacity, w.minimum_temperature FROM warehouses w WHERE w.id = ?"
		mock.ExpectQuery(expectedQuery).
			WithArgs(1).
			WillReturnError(sql.ErrNoRows)

		warehouse, err := rp.GetByIDWareHouse(1)

		assert.Error(t, err)
		assert.IsType(t, &customerror.WareHouseError{}, err)
		assert.Equal(t, model.WareHouse{}, warehouse)
	})
}

func TestWarehouseMysql_PostWareHouse(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rp := repository.NewWareHouseRepository(db, logMock)

	t.Run("Success PostWareHouse", func(t *testing.T) {
		warehouse := model.WareHouse{
			WareHouseCode:      "WH001",
			Address:            "123 Main St",
			Telephone:          "1234567890",
			MinimunCapacity:    100,
			MinimunTemperature: 20,
		}

		expectedQuery := "INSERT INTO `warehouses` (`warehouse_code`, `address`, `telephone`, `minimum_capacity`, `minimum_temperature`) VALUES (?, ?, ?, ?, ?)"
		mock.ExpectExec(expectedQuery).
			WithArgs(warehouse.WareHouseCode, warehouse.Address, warehouse.Telephone, warehouse.MinimunCapacity, warehouse.MinimunTemperature).
			WillReturnResult(sqlmock.NewResult(1, 1))

		id, err := rp.PostWareHouse(warehouse)

		assert.NoError(t, err)
		assert.Equal(t, int64(1), id)
	})

	t.Run("Error PostWareHouse", func(t *testing.T) {
		warehouse := model.WareHouse{
			WareHouseCode:      "WH001",
			Address:            "123 Main St",
			Telephone:          "1234567890",
			MinimunCapacity:    100,
			MinimunTemperature: 20,
		}

		expectedQuery := "INSERT INTO `warehouses` (`warehouse_code`, `address`, `telephone`, `minimum_capacity`, `minimum_temperature`) VALUES (?, ?, ?, ?, ?)"
		mock.ExpectExec(expectedQuery).
			WithArgs(warehouse.WareHouseCode, warehouse.Address, warehouse.Telephone, warehouse.MinimunCapacity, warehouse.MinimunTemperature).
			WillReturnError(errors.New("database error"))

		id, err := rp.PostWareHouse(warehouse)

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "database error")
		assert.Equal(t, int64(0), id)
	})

	t.Run("Conflict PostWareHouse", func(t *testing.T) {
		warehouse := model.WareHouse{
			WareHouseCode:      "WH001",
			Address:            "123 Main St",
			Telephone:          "1234567890",
			MinimunCapacity:    100,
			MinimunTemperature: 20,
		}

		expectedQuery := "INSERT INTO `warehouses` (`warehouse_code`, `address`, `telephone`, `minimum_capacity`, `minimum_temperature`) VALUES (?, ?, ?, ?, ?)"
		mockErr := &mysql.MySQLError{Number: 1062}
		mock.ExpectExec(expectedQuery).
			WithArgs(warehouse.WareHouseCode, warehouse.Address, warehouse.Telephone, warehouse.MinimunCapacity, warehouse.MinimunTemperature).
			WillReturnError(mockErr)

		id, err := rp.PostWareHouse(warehouse)

		assert.Error(t, err)
		assert.IsType(t, &customerror.WareHouseError{}, err)
		assert.Equal(t, int64(0), id)
	})
}

func TestWarehouseMysql_UpdateWareHouse(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rp := repository.NewWareHouseRepository(db, logMock)

	t.Run("Success UpdateWareHouse", func(t *testing.T) {
		id := 1
		warehouse := model.WareHouse{
			ID:                 id,
			WareHouseCode:      "WH001",
			Address:            "123 Main St",
			Telephone:          "1234567890",
			MinimunCapacity:    100,
			MinimunTemperature: 20,
		}

		expectedQuery := "UPDATE warehouses w SET w.warehouse_code = ?, w.address = ?, w.telephone = ?, w.minimum_capacity = ?, w.minimum_temperature = ? WHERE w.id = ?"
		mock.ExpectExec(expectedQuery).
			WithArgs(warehouse.WareHouseCode, warehouse.Address, warehouse.Telephone, warehouse.MinimunCapacity, warehouse.MinimunTemperature, warehouse.ID).
			WillReturnResult(sqlmock.NewResult(1, 1))

		err := rp.UpdateWareHouse(id, warehouse)

		assert.NoError(t, err)
	})

	t.Run("Error UpdateWareHouse", func(t *testing.T) {
		id := 1
		warehouse := model.WareHouse{
			ID:                 id,
			WareHouseCode:      "WH001",
			Address:            "123 Main St",
			Telephone:          "1234567890",
			MinimunCapacity:    100,
			MinimunTemperature: 20,
		}

		expectedQuery := "UPDATE warehouses w SET w.warehouse_code = ?, w.address = ?, w.telephone = ?, w.minimum_capacity = ?, w.minimum_temperature = ? WHERE w.id = ?"
		mock.ExpectExec(expectedQuery).
			WithArgs(warehouse.WareHouseCode, warehouse.Address, warehouse.Telephone, warehouse.MinimunCapacity, warehouse.MinimunTemperature, warehouse.ID).
			WillReturnError(errors.New("database error"))

		err := rp.UpdateWareHouse(id, warehouse)

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "database error")
	})

	t.Run("Conflict UpdateWareHouse", func(t *testing.T) {
		id := 1
		warehouse := model.WareHouse{
			ID:                 id,
			WareHouseCode:      "WH001",
			Address:            "123 Main St",
			Telephone:          "1234567890",
			MinimunCapacity:    100,
			MinimunTemperature: 20,
		}

		expectedQuery := "UPDATE warehouses w SET w.warehouse_code = ?, w.address = ?, w.telephone = ?, w.minimum_capacity = ?, w.minimum_temperature = ? WHERE w.id = ?"
		mockErr := &mysql.MySQLError{Number: 1062}
		mock.ExpectExec(expectedQuery).
			WithArgs(warehouse.WareHouseCode, warehouse.Address, warehouse.Telephone, warehouse.MinimunCapacity, warehouse.MinimunTemperature, warehouse.ID).
			WillReturnError(mockErr)

		err := rp.UpdateWareHouse(id, warehouse)

		assert.Error(t, err)
		assert.IsType(t, &customerror.WareHouseError{}, err)
	})
}

func TestWarehouseMysql_DeleteByIDWareHouse(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rp := repository.NewWareHouseRepository(db, logMock)

	t.Run("Success DeleteByIDWareHouse", func(t *testing.T) {
		id := 1

		expectedQuery := "DELETE FROM `warehouses` WHERE `id` = ?"
		mock.ExpectExec(expectedQuery).
			WithArgs(id).
			WillReturnResult(sqlmock.NewResult(1, 1))

		err := rp.DeleteByIDWareHouse(id)

		assert.NoError(t, err)
	})

	t.Run("Error DeleteByIDWareHouse", func(t *testing.T) {
		id := 1

		expectedQuery := "DELETE FROM `warehouses` WHERE `id` = ?"
		mock.ExpectExec(expectedQuery).
			WithArgs(id).
			WillReturnError(errors.New("database error"))

		err := rp.DeleteByIDWareHouse(id)

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "database error")
	})
}
