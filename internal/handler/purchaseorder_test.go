package handler_test

import (
	"bytes"
	"errors"
	"github.com/maxwelbm/alkemy-g7.git/internal/mocks"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/maxwelbm/alkemy-g7.git/internal/handler"
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/maxwelbm/alkemy-g7.git/pkg/customerror"
	"github.com/stretchr/testify/assert"
)

func setupPurchaseOrder(t *testing.T) *handler.PurchaseOrderHandler {
	mockService := mocks.NewMockIPurchaseOrdersService(t)

	return handler.NewPurchaseOrderHandler(mockService, logMock)
}

func TestHandlerCreatePurchaseOrder(t *testing.T) {
	t.Run("Created Purchase Order successfully", func(t *testing.T) {
		hd := setupPurchaseOrder(t)

		dateString := "2025-01-01T00:00:00Z"
		layout := time.RFC3339

		parsedTime, err := time.Parse(layout, dateString)
		assert.NoError(t, err)
		createdOrder := model.PurchaseOrder{
			ID:              1,
			OrderNumber:     "ON001",
			OrderDate:       parsedTime,
			TrackingCode:    "TC001",
			BuyerID:         1,
			ProductRecordID: 1,
		}
		mockService := hd.Svc.(*mocks.MockIPurchaseOrdersService)
		mockService.On("CreatePurchaseOrder", model.PurchaseOrder{
			ID:              0,
			OrderNumber:     "ON001",
			OrderDate:       parsedTime,
			TrackingCode:    "TC001",
			BuyerID:         1,
			ProductRecordID: 1,
		}).Return(createdOrder, nil)

		body := []byte(`{
    "order_number": "ON001",
    "order_date": "2025-01-01T00:00:00Z",
    "tracking_code": "TC001",
    "buyer_id": 1,
    "product_record_id": 1
}`)

		request := httptest.NewRequest(http.MethodPost, "/purchaseorders", bytes.NewReader(body))
		response := httptest.NewRecorder()
		hd.HandlerCreatePurchaseOrder(response, request)

		expectedJson := `{
    "data": {
        "id": 1,
        "order_number": "ON001",
        "order_date": "2025-01-01T00:00:00Z",
        "tracking_code": "TC001",
        "buyer_id": 1,
        "product_record_id": 1
    }
}`
		assert.Equal(t, http.StatusCreated, response.Code)
		assert.JSONEq(t, expectedJson, response.Body.String())
		mockService.AssertExpectations(t)
	})

	t.Run("Error Buyer Not found", func(t *testing.T) {
		hd := setupPurchaseOrder(t)

		dateString := "2025-01-01T00:00:00Z"
		layout := time.RFC3339

		parsedTime, err := time.Parse(layout, dateString)
		assert.NoError(t, err)

		mockService := hd.Svc.(*mocks.MockIPurchaseOrdersService)
		mockService.On("CreatePurchaseOrder", model.PurchaseOrder{
			ID:              0,
			OrderNumber:     "ON001",
			OrderDate:       parsedTime,
			TrackingCode:    "TC001",
			BuyerID:         99,
			ProductRecordID: 1,
		}).Return(model.PurchaseOrder{}, customerror.NewBuyerError(http.StatusNotFound, customerror.ErrNotFound.Error(), "Buyer"))

		body := []byte(`{
    "order_number": "ON001",
    "order_date": "2025-01-01T00:00:00Z",
    "tracking_code": "TC001",
    "buyer_id": 99,
    "product_record_id": 1
}`)

		request := httptest.NewRequest(http.MethodPost, "/purchaseorders", bytes.NewReader(body))
		response := httptest.NewRecorder()
		hd.HandlerCreatePurchaseOrder(response, request)

		expectedJson := `{"message":"Buyer not found"}`

		assert.Equal(t, http.StatusNotFound, response.Code)
		assert.JSONEq(t, expectedJson, response.Body.String())
		mockService.AssertExpectations(t)
	})
	t.Run("Error Product Record Not found", func(t *testing.T) {
		hd := setupPurchaseOrder(t)

		dateString := "2025-01-01T00:00:00Z"
		layout := time.RFC3339

		parsedTime, err := time.Parse(layout, dateString)
		assert.NoError(t, err)

		mockService := hd.Svc.(*mocks.MockIPurchaseOrdersService)
		mockService.On("CreatePurchaseOrder", model.PurchaseOrder{
			ID:              0,
			OrderNumber:     "ON001",
			OrderDate:       parsedTime,
			TrackingCode:    "TC001",
			BuyerID:         99,
			ProductRecordID: 1,
		}).Return(model.PurchaseOrder{}, customerror.HandleError("product record", customerror.ErrorNotFound, ""))

		body := []byte(`{
    "order_number": "ON001",
    "order_date": "2025-01-01T00:00:00Z",
    "tracking_code": "TC001",
    "buyer_id": 99,
    "product_record_id": 1
}`)

		request := httptest.NewRequest(http.MethodPost, "/purchaseorders", bytes.NewReader(body))
		response := httptest.NewRecorder()
		hd.HandlerCreatePurchaseOrder(response, request)

		expectedJson := `{"message":"product record not found"}`

		assert.Equal(t, http.StatusNotFound, response.Code)
		assert.JSONEq(t, expectedJson, response.Body.String())
		mockService.AssertExpectations(t)
	})
	t.Run("Order Number Already exists", func(t *testing.T) {
		hd := setupPurchaseOrder(t)

		dateString := "2025-01-01T00:00:00Z"
		layout := time.RFC3339

		parsedTime, err := time.Parse(layout, dateString)
		assert.NoError(t, err)

		mockService := hd.Svc.(*mocks.MockIPurchaseOrdersService)
		mockService.On("CreatePurchaseOrder", model.PurchaseOrder{
			ID:              0,
			OrderNumber:     "ON001",
			OrderDate:       parsedTime,
			TrackingCode:    "TC001",
			BuyerID:         99,
			ProductRecordID: 1,
		}).Return(model.PurchaseOrder{}, customerror.NewPurcahseOrderError(http.StatusConflict, customerror.ErrConflict.Error(), "order_number"))

		body := []byte(`{
    "order_number": "ON001",
    "order_date": "2025-01-01T00:00:00Z",
    "tracking_code": "TC001",
    "buyer_id": 99,
    "product_record_id": 1
}`)

		request := httptest.NewRequest(http.MethodPost, "/purchaseorders", bytes.NewReader(body))
		response := httptest.NewRecorder()
		hd.HandlerCreatePurchaseOrder(response, request)

		expectedJson := `{"message":"order_number it already exists"}`

		assert.Equal(t, http.StatusConflict, response.Code)
		assert.JSONEq(t, expectedJson, response.Body.String())
		mockService.AssertExpectations(t)
	})

	t.Run("Error JSON syntax", func(t *testing.T) {
		hd := setupPurchaseOrder(t)

		body := []byte(`{
    "order_number": ,
    "order_date": "2025-01-01T00:00:00Z",
    "tracking_code": "TC001",
    "buyer_id": 99,
    "product_record_id": 1
}`)

		request := httptest.NewRequest(http.MethodPost, "/purchaseorders", bytes.NewReader(body))
		response := httptest.NewRecorder()
		hd.HandlerCreatePurchaseOrder(response, request)

		expectedJson := `{"message":"JSON syntax error. Please verify your input."}`

		assert.Equal(t, http.StatusUnprocessableEntity, response.Code)
		assert.JSONEq(t, expectedJson, response.Body.String())

	})

	t.Run("Error fields required", func(t *testing.T) {
		hd := setupPurchaseOrder(t)

		body := []byte(`{
    "order_number":"" ,
    "order_date": "2025-01-01T00:00:00Z",
    "tracking_code": "",
    "buyer_id": 99,
    "product_record_id": 1
}`)

		request := httptest.NewRequest(http.MethodPost, "/purchaseorders", bytes.NewReader(body))
		response := httptest.NewRecorder()
		hd.HandlerCreatePurchaseOrder(response, request)

		expectedJson := `{"message":"Field(s) order_number,tracking_code cannot be empty"}`

		assert.Equal(t, http.StatusUnprocessableEntity, response.Code)
		assert.JSONEq(t, expectedJson, response.Body.String())

	})

	t.Run("Unmapped Error", func(t *testing.T) {
		hd := setupPurchaseOrder(t)

		dateString := "2025-01-01T00:00:00Z"
		layout := time.RFC3339

		parsedTime, err := time.Parse(layout, dateString)
		assert.NoError(t, err)

		mockService := hd.Svc.(*mocks.MockIPurchaseOrdersService)
		mockService.On("CreatePurchaseOrder", model.PurchaseOrder{
			ID:              0,
			OrderNumber:     "ON001",
			OrderDate:       parsedTime,
			TrackingCode:    "TC001",
			BuyerID:         99,
			ProductRecordID: 1,
		}).Return(model.PurchaseOrder{}, errors.New("unmapped error"))

		body := []byte(`{
    "order_number": "ON001",
    "order_date": "2025-01-01T00:00:00Z",
    "tracking_code": "TC001",
    "buyer_id": 99,
    "product_record_id": 1
}`)

		request := httptest.NewRequest(http.MethodPost, "/purchaseorders", bytes.NewReader(body))
		response := httptest.NewRecorder()
		hd.HandlerCreatePurchaseOrder(response, request)

		expectedJson := `{"message":"Unable to create purchase order"}`

		assert.Equal(t, http.StatusInternalServerError, response.Code)
		assert.JSONEq(t, expectedJson, response.Body.String())
		mockService.AssertExpectations(t)

	})
}
