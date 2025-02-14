package repository_test

import (
	"database/sql"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/maxwelbm/alkemy-g7.git/internal/repository"
	"github.com/maxwelbm/alkemy-g7.git/pkg/customerror"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestBuyerRepository_Post(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer db.Close()

	rp := repository.NewBuyerRepository(db)

	t.Run("Verifies successful addition of a buyer", func(t *testing.T) {
		buyer := model.Buyer{FirstName: "Ac", LastName: "Milan", CardNumberID: "4321"}

		mock.ExpectPrepare("INSERT INTO buyers (card_number_id, first_name, last_name) VALUES (?,?,?)").ExpectExec().
			WithArgs(buyer.CardNumberID, buyer.FirstName, buyer.LastName).
			WillReturnResult(sqlmock.NewResult(1, 1))

		id, err := rp.Post(buyer)

		expectedId := 1

		errMock := mock.ExpectationsWereMet()

		assert.NoError(t, errMock)
		assert.NoError(t, err)
		assert.Equal(t, int64(expectedId), id)

	})

	t.Run("Verifies error addition of a buyer", func(t *testing.T) {
		buyer := model.Buyer{FirstName: "Ac", LastName: "Milan", CardNumberID: "4321"}

		mock.ExpectPrepare("INSERT INTO buyers (card_number_id, first_name, last_name) VALUES (?,?,?)").WillReturnError(errors.New("prepare statement error"))

		_, err := rp.Post(buyer)

		errMock := mock.ExpectationsWereMet()

		assert.NoError(t, errMock)
		assert.Error(t, err)

	})

	t.Run("return MySQLError type 1062", func(t *testing.T) {
		buyer := model.Buyer{FirstName: "Ac", LastName: "Milan", CardNumberID: "4321"}

		mock.ExpectPrepare("INSERT INTO buyers (card_number_id, first_name, last_name) VALUES (?,?,?)")

		mock.ExpectExec("INSERT INTO buyers (card_number_id, first_name, last_name) VALUES (?,?,?)").
			WithArgs(buyer.CardNumberID, buyer.FirstName, buyer.LastName).
			WillReturnError(&mysql.MySQLError{
				Number:   1062,
				SQLState: [5]byte{'2', '3', '0', '0', '0'},
				Message:  "Duplicate entry",
			})

		_, err := rp.Post(buyer)

		errMock := mock.ExpectationsWereMet()

		assert.NoError(t, errMock)
		assert.Error(t, err)

	})
}

func TestBuyerRepository_GetByID(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer db.Close()

	rp := repository.NewBuyerRepository(db)

	t.Run("retrieving a buyer by ID when the buyer exists", func(t *testing.T) {
		buyerID := 1
		buyer := model.Buyer{
			ID:           buyerID,
			CardNumberID: "4321",
			FirstName:    "Ac",
			LastName:     "Milan",
		}

		rows := sqlmock.NewRows([]string{"id", "card_number_id", "first_name", "last_name"}).
			AddRow(buyer.ID, buyer.CardNumberID, buyer.FirstName, buyer.LastName)

		mock.ExpectQuery("SELECT id,card_number_id,first_name,last_name FROM buyers WHERE id=?").
			WithArgs(buyerID).
			WillReturnRows(rows)

		result, err := rp.GetByID(buyerID)
		assert.NoError(t, err)

		errMock := mock.ExpectationsWereMet()
		assert.NoError(t, errMock)

		assert.Equal(t, buyer, result)
	})

	t.Run("Tests retrieving a buyer by ID when the buyer does not exist", func(t *testing.T) {
		buyerID := 99

		mock.ExpectQuery("SELECT id,card_number_id,first_name,last_name FROM buyers WHERE id=?").
			WithArgs(buyerID).
			WillReturnError(sql.ErrNoRows)

		buyer, err := rp.GetByID(buyerID)
		errorExpected := customerror.NewBuyerError(http.StatusNotFound, customerror.ErrNotFound.Error(), "Buyer")
		errorMock := mock.ExpectationsWereMet()

		assert.NoError(t, errorMock)
		assert.Error(t, errorExpected, err)
		assert.Equal(t, model.Buyer{}, buyer)

	})
}

func TestBuyerRepository_Get(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer db.Close()

	rp := repository.NewBuyerRepository(db)

	t.Run("Tests listing buyers when there are existing records", func(t *testing.T) {
		expectedBuyers := []model.Buyer{
			model.Buyer{
				ID:           1,
				CardNumberID: "4321",
				FirstName:    "Ac",
				LastName:     "Milan",
			},
			model.Buyer{
				ID:           2,
				CardNumberID: "43211",
				FirstName:    "Aca",
				LastName:     "Milana",
			},
		}

		rows := mock.NewRows([]string{"id", "card_number_id", "first_name", "last_name"})
		for _, buyer := range expectedBuyers {
			rows.AddRow(buyer.ID, buyer.CardNumberID, buyer.FirstName, buyer.LastName)
		}

		mock.ExpectQuery("SELECT id, card_number_id,first_name,last_name FROM buyers").WillReturnRows(rows)

		buyers, err := rp.Get()
		errMock := mock.ExpectationsWereMet()
		assert.NoError(t, err)
		assert.Equal(t, expectedBuyers, buyers)
		assert.NoError(t, errMock)

	})

	t.Run("Return errors an listing buyers", func(t *testing.T) {
		mock.ExpectQuery("SELECT id, card_number_id,first_name,last_name FROM buyers").WillReturnError(errors.New("unmapped error"))

		buyers, err := rp.Get()
		mockErr := mock.ExpectationsWereMet()
		assert.Error(t, err)
		assert.Equal(t, []model.Buyer(nil), buyers)
		assert.NoError(t, mockErr)

	})
}

func TestBuyerRepository_Update(t *testing.T) {

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer db.Close()

	rp := repository.NewBuyerRepository(db)

	t.Run("Confirms that a buyer's details can be successfully updated.", func(t *testing.T) {

		buyerID := 1
		buyer := model.Buyer{FirstName: "Ac", LastName: "Milan", CardNumberID: "4321"}

		mock.ExpectPrepare("UPDATE buyers SET card_number_id = ?, first_name = ?, last_name = ? WHERE id = ?").
			ExpectExec().
			WithArgs(buyer.CardNumberID, buyer.FirstName, buyer.LastName, buyerID).
			WillReturnResult(sqlmock.NewResult(1, 1))

		err := rp.Update(buyerID, buyer)
		mockErr := mock.ExpectationsWereMet()

		assert.NoError(t, err)
		assert.NoError(t, mockErr)
	})

	t.Run("Error mysql 1062 duplicate entry", func(t *testing.T) {

		buyerID := 1
		buyer := model.Buyer{FirstName: "Ac", LastName: "Milan", CardNumberID: "4321"}

		mock.ExpectPrepare("UPDATE buyers SET card_number_id = ?, first_name = ?, last_name = ? WHERE id = ?").
			ExpectExec().
			WithArgs(buyer.CardNumberID, buyer.FirstName, buyer.LastName, buyerID).
			WillReturnError(&mysql.MySQLError{
				Number:   1062,
				SQLState: [5]byte{'2', '3', '0', '0', '0'},
				Message:  "Duplicate entry",
			})

		err := rp.Update(buyerID, buyer)
		mockErr := mock.ExpectationsWereMet()

		expectedError := customerror.NewBuyerError(http.StatusNotFound, customerror.ErrConflict.Error(), "card_number_id")

		assert.Error(t, expectedError, err)
		assert.NoError(t, mockErr)
	})

	t.Run("Error unmapped in db", func(t *testing.T) {
		buyerID := 1
		buyer := model.Buyer{FirstName: "Ac", LastName: "Milan", CardNumberID: "4321"}

		mock.ExpectPrepare("UPDATE buyers SET card_number_id = ?, first_name = ?, last_name = ? WHERE id = ?").
			WillReturnError(errors.New("unmapped Error"))

		err := rp.Update(buyerID, buyer)
		mockErr := mock.ExpectationsWereMet()
		assert.Error(t, err)
		assert.NoError(t, mockErr)

	})
}

func TestBuyerRepository_CountPurchaseOrderBuyers(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer db.Close()

	rp := repository.NewBuyerRepository(db)

	t.Run("count Purchase Order Buyers success", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "card_number_id", "first_name", "last_name", "purchase_orders_count"}).
			AddRow(1, "1234-5678", "John", "Doe", 3).
			AddRow(2, "8765-4321", "Jane", "Doe", 5)

		mock.ExpectQuery("SELECT b.id, b.card_number_id, b.first_name, b.last_name, COUNT(po.id) as purchase_orders_count FROM buyers b LEFT JOIN purchase_orders po ON po.buyer_id = b.id GROUP BY b.id").
			WillReturnRows(rows)

		count, err := rp.CountPurchaseOrderBuyers()

		expectedCount := []model.BuyerPurchaseOrder{
			{ID: 1, CardNumberID: "1234-5678", FirstName: "John", LastName: "Doe", PurchaseOrdersCount: 3},
			{ID: 2, CardNumberID: "8765-4321", FirstName: "Jane", LastName: "Doe", PurchaseOrdersCount: 5},
		}
		errMock := mock.ExpectationsWereMet()
		assert.NoError(t, errMock)
		assert.NoError(t, err)
		assert.Equal(t, expectedCount, count)

	})

	t.Run("error count Purchase Order Buyers", func(t *testing.T) {

		mock.ExpectQuery("SELECT b.id, b.card_number_id, b.first_name, b.last_name, COUNT(po.id) as purchase_orders_count FROM buyers b LEFT JOIN purchase_orders po ON po.buyer_id = b.id GROUP BY b.id").
			WillReturnError(errors.New("unmapped error"))

		count, err := rp.CountPurchaseOrderBuyers()

		expectedCount := []model.BuyerPurchaseOrder(nil)
		errMock := mock.ExpectationsWereMet()
		assert.NoError(t, errMock)
		assert.Error(t, err)
		assert.Equal(t, expectedCount, count)

	})
}

func TestBuyerRepository_CountPurchaseOrderByBuyerID(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer db.Close()

	rp := repository.NewBuyerRepository(db)

	t.Run("count Purchase Order Buyer existing success", func(t *testing.T) {
		buyerID := 1

		rows := sqlmock.NewRows([]string{"id", "card_number_id", "first_name", "last_name", "purchase_orders_count"}).
			AddRow(1, "1234-5678", "John", "Doe", 3)

		mock.ExpectQuery("SELECT b.id, b.card_number_id, b.first_name, b.last_name, COUNT(po.id) as purchase_orders_count FROM buyers b LEFT JOIN purchase_orders po ON po.buyer_id = b.id WHERE b.id = ? GROUP BY b.id").
			WithArgs(buyerID).
			WillReturnRows(rows)

		count, err := rp.CountPurchaseOrderByBuyerID(buyerID)

		expectedCount := model.BuyerPurchaseOrder{ID: 1, CardNumberID: "1234-5678", FirstName: "John", LastName: "Doe", PurchaseOrdersCount: 3}

		errMock := mock.ExpectationsWereMet()
		assert.NoError(t, errMock)
		assert.NoError(t, err)
		assert.Equal(t, expectedCount, count)

	})

	t.Run("error sql noRows count Purchase Order Buyer by id", func(t *testing.T) {

		buyerID := 99

		mock.ExpectQuery("SELECT b.id, b.card_number_id, b.first_name, b.last_name, COUNT(po.id) as purchase_orders_count FROM buyers b LEFT JOIN purchase_orders po ON po.buyer_id = b.id WHERE b.id = ? GROUP BY b.id").
			WillReturnError(sql.ErrNoRows)

		count, err := rp.CountPurchaseOrderByBuyerID(buyerID)

		expectedCount := model.BuyerPurchaseOrder{}
		errMock := mock.ExpectationsWereMet()
		assert.NoError(t, errMock)
		assert.Error(t, err)
		assert.Equal(t, expectedCount, count)

	})
}

func TestBuyerRepository_Delete(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer db.Close()

	rp := repository.NewBuyerRepository(db)

	t.Run("Delete Buyer exisiting success", func(t *testing.T) {
		buyerID := 1

		mock.ExpectExec("DELETE FROM buyers WHERE id = ?").
			WithArgs(buyerID).
			WillReturnResult(sqlmock.NewResult(0, 1))

		err := rp.Delete(buyerID)

		mockErr := mock.ExpectationsWereMet()

		assert.NoError(t, mockErr)
		assert.NoError(t, err)
	})

	t.Run("Delete Buyer DependencyError", func(t *testing.T) {
		buyerID := 1

		mock.ExpectExec("DELETE FROM buyers WHERE id = ?").
			WithArgs(1).
			WillReturnError(&mysql.MySQLError{Number: 1451})

		err := rp.Delete(buyerID)

		mockErr := mock.ExpectationsWereMet()

		assert.NoError(t, mockErr)
		assert.Error(t, err)
	})
}
