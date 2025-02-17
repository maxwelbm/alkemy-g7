package repository

import (
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/stretchr/testify/assert"
)

func TestPostInboundOrder(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewInboundService(db)

	inboundOrder := model.InboundOrder{
		OrderDate:      time.Date(2023, 10, 1, 0, 0, 0, 0, time.UTC),
		OrderNumber:    "ORD123",
		EmployeeID:     1,
		ProductBatchID: 1,
		WareHouseID:    1,
	}

	t.Run("should insert the inbound order and return it with the generated ID", func(t *testing.T) {
		mock.ExpectExec("INSERT INTO inbound_orders").
			WithArgs(inboundOrder.OrderDate, inboundOrder.OrderNumber, inboundOrder.EmployeeID, inboundOrder.ProductBatchID, inboundOrder.WareHouseID).
			WillReturnResult(sqlmock.NewResult(1, 1))

		result, err := repo.Post(inboundOrder)

		assert.NoError(t, err)
		assert.Equal(t, 1, result.ID)
		assert.Equal(t, inboundOrder.OrderDate, result.OrderDate)
		assert.Equal(t, inboundOrder.OrderNumber, result.OrderNumber)
		assert.Equal(t, inboundOrder.EmployeeID, result.EmployeeID)
		assert.Equal(t, inboundOrder.ProductBatchID, result.ProductBatchID)
		assert.Equal(t, inboundOrder.WareHouseID, result.WareHouseID)

		assert.NoError(t, mock.ExpectationsWereMet())
	})

}
