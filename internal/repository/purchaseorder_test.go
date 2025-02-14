package repository_test

import (
	"database/sql"
	"errors"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-sql-driver/mysql"
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/maxwelbm/alkemy-g7.git/internal/repository"
	"github.com/stretchr/testify/assert"
)

func TestPurchaseOrderRepository_Post(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	rp := repository.NewPurchaseOrderRepository(db)

	t.Run("Verifies successful addition of a purchase Order", func(t *testing.T) {
		purchaseOrderID := 1
		createdPurchaseOrder := model.PurchaseOrder{
			OrderNumber:     "ON001",
			OrderDate:       time.Time{},
			TrackingCode:    "TC001",
			BuyerID:         1,
			ProductRecordID: 1,
		}

		mock.ExpectPrepare("INSERT INTO purchase_orders (order_number, order_date, tracking_code, buyer_id, product_record_id) VALUES(?,?,?,?,?)").
			ExpectExec().
			WithArgs(createdPurchaseOrder.OrderNumber, createdPurchaseOrder.OrderDate, createdPurchaseOrder.TrackingCode,
				createdPurchaseOrder.BuyerID, createdPurchaseOrder.ProductRecordID,
			).WillReturnResult(sqlmock.NewResult(1, 1))

		ID, err := rp.Post(createdPurchaseOrder)
		MockErr := mock.ExpectationsWereMet()

		assert.Equal(t, int64(purchaseOrderID), ID)
		assert.NoError(t, err)
		assert.NoError(t, MockErr)

	})

	t.Run("return MySQLError type 1062 an create purchase order", func(t *testing.T) {
		purchaseOrderID := 0
		createdPurchaseOrder := model.PurchaseOrder{
			OrderNumber:     "ON001",
			OrderDate:       time.Time{},
			TrackingCode:    "TC001",
			BuyerID:         1,
			ProductRecordID: 1,
		}

		mock.ExpectPrepare("INSERT INTO purchase_orders (order_number, order_date, tracking_code, buyer_id, product_record_id) VALUES(?,?,?,?,?)").
			ExpectExec().
			WithArgs(createdPurchaseOrder.OrderNumber, createdPurchaseOrder.OrderDate, createdPurchaseOrder.TrackingCode,
				createdPurchaseOrder.BuyerID, createdPurchaseOrder.ProductRecordID,
			).WillReturnError(&mysql.MySQLError{
			Number: 1062,
		})

		ID, err := rp.Post(createdPurchaseOrder)
		MockErr := mock.ExpectationsWereMet()

		assert.Equal(t, int64(purchaseOrderID), ID)
		assert.Error(t, err)
		assert.NoError(t, MockErr)

	})

	t.Run("return error unmapped prepare an create purchase order", func(t *testing.T) {
		purchaseOrderID := 0
		createdPurchaseOrder := model.PurchaseOrder{
			OrderNumber:     "ON001",
			OrderDate:       time.Time{},
			TrackingCode:    "TC001",
			BuyerID:         1,
			ProductRecordID: 1,
		}

		mock.ExpectPrepare("INSERT INTO purchase_orders (order_number, order_date, tracking_code, buyer_id, product_record_id) VALUES(?,?,?,?,?)").
			WillReturnError(errors.New("error prepare"))

		ID, err := rp.Post(createdPurchaseOrder)
		MockErr := mock.ExpectationsWereMet()

		assert.Equal(t, int64(purchaseOrderID), ID)
		assert.Error(t, err)
		assert.NoError(t, MockErr)

	})
}

func TestPurchaseOrderRepository_GetByID(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	rp := repository.NewPurchaseOrderRepository(db)

	t.Run("retrieve existing Purchase Order", func(t *testing.T) {
		purchaseOrderID := 1
		ExpectedPurchaseOrder := model.PurchaseOrder{
			ID:              purchaseOrderID,
			OrderNumber:     "ON001",
			OrderDate:       time.Time{},
			TrackingCode:    "TC001",
			BuyerID:         1,
			ProductRecordID: 1,
		}

		rows := sqlmock.NewRows([]string{"id", "order_number", "order_date", "tracking_code", "buyer_id", "product_record_id"}).
			AddRow(ExpectedPurchaseOrder.ID, ExpectedPurchaseOrder.OrderNumber, ExpectedPurchaseOrder.OrderDate, ExpectedPurchaseOrder.TrackingCode,
				ExpectedPurchaseOrder.BuyerID, ExpectedPurchaseOrder.ProductRecordID)

		mock.ExpectQuery("SELECT id, order_number, order_date, tracking_code, buyer_id, product_record_id FROM purchase_orders WHERE id = ?").
			WithArgs(purchaseOrderID).WillReturnRows(rows)

		purchase, err := rp.GetByID(purchaseOrderID)
		mockErr := mock.ExpectationsWereMet()

		assert.NoError(t, mockErr)
		assert.NoError(t, err)
		assert.Equal(t, ExpectedPurchaseOrder, purchase)

	})

	t.Run("error sql NoRows an search Purchase Order", func(t *testing.T) {
		purchaseOrderID := 1
		ExpectedPurchaseOrder := model.PurchaseOrder{}

		mock.ExpectQuery("SELECT id, order_number, order_date, tracking_code, buyer_id, product_record_id FROM purchase_orders WHERE id = ?").
			WithArgs(purchaseOrderID).WillReturnError(sql.ErrNoRows)

		purchase, err := rp.GetByID(purchaseOrderID)
		mockErr := mock.ExpectationsWereMet()

		assert.NoError(t, mockErr)
		assert.Error(t, err)
		assert.Equal(t, ExpectedPurchaseOrder, purchase)

	})
}
