package repository_test

import (
	"database/sql"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-sql-driver/mysql"
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/maxwelbm/alkemy-g7.git/internal/repository"
	"github.com/maxwelbm/alkemy-g7.git/pkg/customerror"
	"github.com/stretchr/testify/assert"
)

func TestProductBatchesRepository_Post(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	rp := repository.CreateProductBatchesRepository(db, logMock)

	t.Run("given a valid product batches then create it with no error", func(t *testing.T) {
		createdPB := model.ProductBatches{
			ID:                 1,
			BatchNumber:        "B01",
			CurrentQuantity:    10,
			CurrentTemperature: 10.00,
			MinimumTemperature: 5.00,
			DueDate:            time.Time{},
			InitialQuantity:    5,
			ManufacturingDate:  time.Time{},
			ManufacturingHour:  10,
			ProductID:          1,
			SectionID:          1,
		}

		mock.ExpectExec("INSERT INTO product_batches (batch_number, current_quantity, current_temperature, minimum_temperature, due_date, initial_quantity, manufacturing_date, manufacturing_hour, product_id, section_id) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)").WithArgs(&createdPB.BatchNumber, &createdPB.CurrentQuantity, &createdPB.CurrentTemperature, &createdPB.MinimumTemperature, &createdPB.DueDate, &createdPB.InitialQuantity, &createdPB.ManufacturingDate, &createdPB.ManufacturingHour, &createdPB.ProductID, &createdPB.SectionID).WillReturnResult(sqlmock.NewResult(1, 1))

		mock.ExpectQuery("SELECT `id`, `batch_number`, `current_quantity`, `current_temperature`, `minimum_temperature`, `due_date`, `initial_quantity`, `manufacturing_date`, `manufacturing_hour`, `product_id`, `section_id` FROM `product_batches` WHERE `id` = ?").WithArgs(&createdPB.ID).WillReturnRows(sqlmock.NewRows([]string{"id", "batch_number", "current_quantity", "current_temperature", "minimum_temperature", "due_date", "initial_quantity", "manufacturing_date", "manufacturing_hour", "product_id", "section_id"}).AddRow(&createdPB.ID, &createdPB.BatchNumber, &createdPB.CurrentQuantity, &createdPB.CurrentTemperature, &createdPB.MinimumTemperature, &createdPB.DueDate, &createdPB.InitialQuantity, &createdPB.ManufacturingDate, &createdPB.ManufacturingHour, &createdPB.ProductID, &createdPB.SectionID))

		pb, err := rp.Post(&createdPB)
		mockErr := mock.ExpectationsWereMet()

		assert.Equal(t, createdPB, pb)
		assert.NoError(t, err)
		assert.NoError(t, mockErr)

	})

	t.Run("given a duplicate product batches then return MySQLError type 1062", func(t *testing.T) {
		createdPB := model.ProductBatches{
			ID:                 1,
			BatchNumber:        "B01",
			CurrentQuantity:    10,
			CurrentTemperature: 10.00,
			MinimumTemperature: 5.00,
			DueDate:            time.Time{},
			InitialQuantity:    5,
			ManufacturingDate:  time.Time{},
			ManufacturingHour:  10,
			ProductID:          1,
			SectionID:          1,
		}

		mock.ExpectExec("INSERT INTO product_batches (batch_number, current_quantity, current_temperature, minimum_temperature, due_date, initial_quantity, manufacturing_date, manufacturing_hour, product_id, section_id) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)").WithArgs(&createdPB.BatchNumber, &createdPB.CurrentQuantity, &createdPB.CurrentTemperature, &createdPB.MinimumTemperature, &createdPB.DueDate, &createdPB.InitialQuantity, &createdPB.ManufacturingDate, &createdPB.ManufacturingHour, &createdPB.ProductID, &createdPB.SectionID).WillReturnError(&mysql.MySQLError{Number: 1062})

		_, err := rp.Post(&createdPB)
		mockErr := mock.ExpectationsWereMet()

		assert.Error(t, err)
		assert.NoError(t, mockErr)
	})
}

func TestProductBatchesRepository_GetByID(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	rp := repository.CreateProductBatchesRepository(db, logMock)

	t.Run("given a valid id then return the product batch with no error", func(t *testing.T) {
		pbID := 1

		createdPB := model.ProductBatches{
			ID:                 1,
			BatchNumber:        "B01",
			CurrentQuantity:    10,
			CurrentTemperature: 10.00,
			MinimumTemperature: 5.00,
			DueDate:            time.Time{},
			InitialQuantity:    5,
			ManufacturingDate:  time.Time{},
			ManufacturingHour:  10,
			ProductID:          1,
			SectionID:          1,
		}

		rows := sqlmock.NewRows([]string{"id", "batch_number", "current_quantity", "current_temperature", "minimum_temperature", "due_date", "initial_quantity", "manufacturing_date", "manufacturing_hour", "product_id", "section_id"}).AddRow(&createdPB.ID, &createdPB.BatchNumber, &createdPB.CurrentQuantity, &createdPB.CurrentTemperature, &createdPB.MinimumTemperature, &createdPB.DueDate, &createdPB.InitialQuantity, &createdPB.ManufacturingDate, &createdPB.ManufacturingHour, &createdPB.ProductID, &createdPB.SectionID)

		mock.ExpectQuery("SELECT `id`, `batch_number`, `current_quantity`, `current_temperature`, `minimum_temperature`, `due_date`, `initial_quantity`, `manufacturing_date`, `manufacturing_hour`, `product_id`, `section_id` FROM `product_batches` WHERE `id` = ?").WithArgs(&createdPB.ID).WillReturnRows(rows)

		pbItem, err := rp.GetByID(pbID)
		mockErr := mock.ExpectationsWereMet()

		assert.NoError(t, mockErr)
		assert.NoError(t, err)
		assert.Equal(t, createdPB, pbItem)

	})

	t.Run("given an invalid id then return the error", func(t *testing.T) {
		pbID := 99
		createdPB := model.ProductBatches{}
		expectedErr := customerror.HandleError("product batches", customerror.ErrorNotFound, "")

		mock.ExpectQuery("SELECT `id`, `batch_number`, `current_quantity`, `current_temperature`, `minimum_temperature`, `due_date`, `initial_quantity`, `manufacturing_date`, `manufacturing_hour`, `product_id`, `section_id` FROM `product_batches` WHERE `id` = ?").WithArgs(pbID).WillReturnError(sql.ErrNoRows)

		pbItem, err := rp.GetByID(pbID)
		mockErr := mock.ExpectationsWereMet()

		assert.NoError(t, mockErr)
		assert.EqualError(t, expectedErr, err.Error())
		assert.Equal(t, createdPB, pbItem)

	})
}
