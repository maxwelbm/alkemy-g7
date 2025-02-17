package repository

import (
	"database/sql"
	"errors"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	appErr "github.com/maxwelbm/alkemy-g7.git/pkg/customerror"
	"github.com/stretchr/testify/assert"
)

func TestProductRecRepository_Create(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer db.Close()

	repo := NewProductRecRepository(db, logMock)

	productRec := model.ProductRecords{
		ID:             1,
		LastUpdateDate: time.Time{},
		PurchasePrice:  100.50,
		SalePrice:      150.75,
		ProductID:      101,
	}

	t.Run("successful creation of a product rec", func(t *testing.T) {
		mock.ExpectExec("INSERT INTO product_records (last_update_date, purchase_price, sale_price, product_id) VALUES (?, ?, ?, ?)").WithArgs(&productRec.LastUpdateDate, &productRec.PurchasePrice, &productRec.SalePrice, &productRec.ProductID).
			WillReturnResult(sqlmock.NewResult(1, 1))

		res, err := repo.Create(productRec)

		assert.Equal(t, productRec, res)

		assert.NoError(t, err)
	})
	t.Run("error to executing query", func(t *testing.T) {
		expectedError := errors.New("error to executing query")

		mock.ExpectExec("INSERT INTO product_records (last_update_date, purchase_price, sale_price, product_id) VALUES (?, ?, ?, ?)").WithArgs(productRec.LastUpdateDate, productRec.PurchasePrice, productRec.SalePrice, productRec.ProductID).WillReturnError(expectedError)

		_, err := repo.Create(productRec)

		assert.EqualError(t, expectedError, err.Error())

	})
	t.Run("error test get last inserted id", func(t *testing.T) {
		expectedError := errors.New("error test get last inserted id")

		mock.ExpectExec("INSERT INTO product_records (last_update_date, purchase_price, sale_price, product_id) VALUES (?, ?, ?, ?)").WithArgs(productRec.LastUpdateDate, productRec.PurchasePrice, productRec.SalePrice, productRec.ProductID).WillReturnResult(sqlmock.NewErrorResult(expectedError))

		_, err := repo.Create(productRec)

		assert.EqualError(t, expectedError, err.Error())

	})
}

func TestProductRecRepository_GetByID(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer db.Close()

	productRec := model.ProductRecords{
		ID:             1,
		LastUpdateDate: time.Time{},
		PurchasePrice:  100.50,
		SalePrice:      150.75,
		ProductID:      101,
	}

	repo := NewProductRecRepository(db, logMock)

	query := `
	SELECT
	id,
	last_update_date, 
	product_id, 
	purchase_price, 
	sale_price
	FROM product_records 
	WHERE id = ?
	`

	t.Run("Getting a product successfully", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "last_update_date", "product_id", "purchase_price", "sale_price"}).
			AddRow(productRec.ID, productRec.LastUpdateDate, productRec.ProductID, productRec.PurchasePrice, productRec.SalePrice)

		mock.ExpectQuery(query).WithArgs(productRec.ID).WillReturnRows(rows)

		res, err := repo.GetByID(productRec.ID)

		assert.NoError(t, err)

		assert.Equal(t, productRec, res)

	})

	t.Run("Error product not found", func(t *testing.T) {
		expectedErr := appErr.HandleError("product record", appErr.ErrorNotFound, "")

		mock.ExpectQuery(query).WithArgs(productRec.ID).WillReturnError(sql.ErrNoRows)

		_, err := repo.GetByID(productRec.ID)

		assert.EqualError(t, expectedErr, err.Error())

	})

	t.Run("Error not maped in row", func(t *testing.T) {
		expectedErr := errors.New("not maped error")

		mock.ExpectQuery(query).WithArgs(productRec.ID).WillReturnError(expectedErr)

		_, err := repo.GetByID(productRec.ID)

		assert.EqualError(t, expectedErr, err.Error())

	})

}
func TestProductRecRepository_GetByIDProduct(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer db.Close()

	productRec := model.ProductRecords{
		ID:             1,
		LastUpdateDate: time.Time{},
		PurchasePrice:  100.50,
		SalePrice:      150.75,
		ProductID:      101,
	}

	repo := NewProductRecRepository(db, logMock)

	query := `
	SELECT
	id,
	last_update_date, 
	product_id, 
	purchase_price, 
	sale_price
	FROM product_records
	where product_id = ?
	`

	t.Run("Getting a product successfully", func(t *testing.T) {
		var expected []model.ProductRecords

		expected = append(expected, productRec)

		rows := sqlmock.NewRows([]string{"id", "last_update_date", "product_id", "purchase_price", "sale_price"}).
			AddRow(productRec.ID, productRec.LastUpdateDate, productRec.ProductID, productRec.PurchasePrice, productRec.SalePrice)

		mock.ExpectQuery(query).WithArgs(productRec.ProductID).WillReturnRows(rows)

		res, err := repo.GetByIDProduct(productRec.ProductID)

		assert.NoError(t, err)

		assert.Equal(t, expected, res)

	})

	t.Run("Error executing query", func(t *testing.T) {
		expectedErr := errors.New("not maped error in loop")

		mock.ExpectQuery(query).WithArgs(productRec.ID).WillReturnError(expectedErr)

		_, err := repo.GetByIDProduct(productRec.ID)

		assert.EqualError(t, expectedErr, err.Error())

	})

	t.Run("Error trying to convert row to struct", func(t *testing.T) {
		expectedErr := errors.New("Error trying to convert row to struct")
		rows := sqlmock.NewRows([]string{"id", "last_update_date", "product_id", "purchase_price", "sale_price"}).
			AddRow(productRec.ID, productRec.LastUpdateDate, nil, productRec.PurchasePrice, productRec.SalePrice)

		mock.ExpectQuery(query).WithArgs(productRec.ID).WillReturnRows(rows)

		_, err := repo.GetByIDProduct(productRec.ID)

		assert.EqualError(t, expectedErr, err.Error())

	})

	t.Run("Error not maped in row", func(t *testing.T) {
		expectedErr := errors.New("not maped error")
		rows := sqlmock.NewRows([]string{"id", "last_update_date", "product_id", "purchase_price", "sale_price"}).
			AddRow(productRec.ID, productRec.LastUpdateDate, productRec.ProductID, productRec.PurchasePrice, productRec.SalePrice).RowError(0, expectedErr)

		mock.ExpectQuery(query).WithArgs(productRec.ID).WillReturnRows(rows)

		_, err := repo.GetByIDProduct(productRec.ID)

		assert.EqualError(t, expectedErr, err.Error())

	})
}

func TestProductRecRepository_GetAll(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer db.Close()

	productRec := model.ProductRecords{
		ID:             1,
		LastUpdateDate: time.Time{},
		PurchasePrice:  100.50,
		SalePrice:      150.75,
		ProductID:      101,
	}

	query := `
	SELECT
	id,
	last_update_date, 
	product_id, 
	purchase_price, 
	sale_price
	FROM product_records`

	repo := NewProductRecRepository(db, logMock)

	t.Run("Getting a list product rec successfully", func(t *testing.T) {
		var expected []model.ProductRecords

		expected = append(expected, productRec)

		rows := sqlmock.NewRows([]string{"id", "last_update_date", "product_id", "purchase_price", "sale_price"}).
			AddRow(productRec.ID, productRec.LastUpdateDate, productRec.ProductID, productRec.PurchasePrice, productRec.SalePrice)

		mock.ExpectQuery(query).WillReturnRows(rows)

		res, err := repo.GetAll()

		assert.NoError(t, err)

		assert.Equal(t, expected, res)

	})

	t.Run("Error executing query", func(t *testing.T) {
		expectedErr := errors.New("not maped error in loop")

		mock.ExpectQuery(query).WillReturnError(expectedErr)

		_, err := repo.GetAll()

		assert.EqualError(t, expectedErr, err.Error())

	})
	t.Run("Error trying to convert row to struct", func(t *testing.T) {
		expectedErr := errors.New("Error trying to convert row to struct")
		rows := sqlmock.NewRows([]string{"id", "last_update_date", "product_id", "purchase_price", "sale_price"}).
			AddRow(productRec.ID, productRec.LastUpdateDate, nil, productRec.PurchasePrice, productRec.SalePrice)

		mock.ExpectQuery(query).WillReturnRows(rows)

		_, err := repo.GetAll()

		assert.EqualError(t, expectedErr, err.Error())

	})

	t.Run("Error not maped in row", func(t *testing.T) {
		expectedErr := errors.New("not maped error")
		rows := sqlmock.NewRows([]string{"id", "last_update_date", "product_id", "purchase_price", "sale_price"}).
			AddRow(productRec.ID, productRec.LastUpdateDate, productRec.ProductID, productRec.PurchasePrice, productRec.SalePrice).RowError(0, expectedErr)

		mock.ExpectQuery(query).WillReturnRows(rows)

		_, err := repo.GetAll()

		assert.EqualError(t, expectedErr, err.Error())

	})
}

func TestProductRecRepository_GetAllReport(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer db.Close()

	productReport := model.ProductRecordsReport{
		ProductID:    101,
		Description:  "T-Shirt",
		RecordsCount: 3,
	}

	query := `
	SELECT  
	p.id, 
	p.description, 
	count(p.id) as record_count 
	FROM products p
	inner join product_records pr on pr.product_id = p.id
	GROUP by p.id, p.description`

	repo := NewProductRecRepository(db, logMock)

	t.Run("Getting a list product rec successfully", func(t *testing.T) {
		var expected []model.ProductRecordsReport

		expected = append(expected, productReport)

		rows := sqlmock.NewRows([]string{"id", "description", "record_count"}).
			AddRow(productReport.ProductID, productReport.Description, productReport.RecordsCount)

		mock.ExpectQuery(query).WillReturnRows(rows)

		res, err := repo.GetAllReport()

		assert.NoError(t, err)

		assert.Equal(t, expected, res)

	})

	t.Run("Error executing query", func(t *testing.T) {
		expectedErr := errors.New("not maped error in loop")

		mock.ExpectQuery(query).WillReturnError(expectedErr)

		_, err := repo.GetAllReport()

		assert.EqualError(t, expectedErr, err.Error())

	})
	t.Run("Error trying to convert row to struct", func(t *testing.T) {
		expectedErr := errors.New("Error trying to convert row to struct")
		rows := sqlmock.NewRows([]string{"id", "description", "record_count"}).
			AddRow(nil, productReport.Description, productReport.RecordsCount)

		mock.ExpectQuery(query).WillReturnRows(rows)

		_, err := repo.GetAllReport()

		assert.EqualError(t, expectedErr, err.Error())

	})

	t.Run("Error not maped in row", func(t *testing.T) {
		expectedErr := errors.New("not maped error")
		rows := sqlmock.NewRows([]string{"id", "description", "record_count"}).
			AddRow(productReport.ProductID, productReport.Description, productReport.RecordsCount).
			RowError(0, expectedErr)

		mock.ExpectQuery(query).WillReturnRows(rows)

		_, err := repo.GetAllReport()

		assert.EqualError(t, expectedErr, err.Error())

	})
}
