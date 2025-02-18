package handler_test

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/maxwelbm/alkemy-g7.git/internal/handler"
	"github.com/maxwelbm/alkemy-g7.git/internal/mocks"
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/maxwelbm/alkemy-g7.git/pkg/customerror"
	"github.com/stretchr/testify/assert"
)

func setupProductBatches(t *testing.T) *handler.ProductBatchesController {
	mockService := mocks.NewMockIProductBatchesService(t)

	return handler.CreateProductBatchesHandler(mockService, logMock)
}

func TestHandler_Post(t *testing.T) {
	t.Run("given a valid product batches then create it with no error", func(t *testing.T) {
		hd := setupProductBatches(t)

		dateString := "2025-01-01T00:00:00Z"
		layout := time.RFC3339

		parsedTime, err := time.Parse(layout, dateString)
		assert.NoError(t, err)

		createdPB := model.ProductBatches{
			ID:                 1,
			BatchNumber:        "B01",
			CurrentQuantity:    10,
			CurrentTemperature: 10.00,
			MinimumTemperature: 5.00,
			DueDate:            parsedTime,
			InitialQuantity:    5,
			ManufacturingDate:  parsedTime,
			ManufacturingHour:  10,
			ProductID:          1,
			SectionID:          1,
		}

		mockService := hd.Sv.(*mocks.MockIProductBatchesService)
		mockService.On("Post", &model.ProductBatches{
			ID:                 0,
			BatchNumber:        "B01",
			CurrentQuantity:    10,
			CurrentTemperature: 10.00,
			MinimumTemperature: 5.00,
			DueDate:            parsedTime,
			InitialQuantity:    5,
			ManufacturingDate:  parsedTime,
			ManufacturingHour:  10,
			ProductID:          1,
			SectionID:          1,
		}).Return(createdPB, nil)

		reqBody := []byte(`{
		"batch_number": "B01",
		"current_quantity": 10,
		"current_temperature": 10.00,
		"minimum_temperature": 5.00,
		"due_date": "2025-01-01T00:00:00Z",
		"initial_quantity": 5,
		"manufacturing_date": "2025-01-01T00:00:00Z",
		"manufacturing_hour": 10,
		"product_id": 1,
		"section_id": 1}`)

		request := httptest.NewRequest(http.MethodPost, "/reportProducts", bytes.NewReader(reqBody))
		response := httptest.NewRecorder()
		hd.Post(response, request)

		expectedJson := `{
		"data": {
		"ID": 1,
		"BatchNumber": "B01",
		"CurrentQuantity": 10,
		"CurrentTemperature": 10.00,
		"MinimumTemperature": 5.00,
		"DueDate": "2025-01-01T00:00:00Z",
		"InitialQuantity": 5,
		"ManufacturingDate": "2025-01-01T00:00:00Z",
		"ManufacturingHour": 10,
		"ProductID": 1,
		"SectionID": 1}
		}`
		assert.Equal(t, http.StatusCreated, response.Code)
		assert.JSONEq(t, expectedJson, response.Body.String())
		mockService.AssertExpectations(t)
	})

	t.Run("given an invalid product batches then return an error", func(t *testing.T) {
		hd := setupProductBatches(t)

		dateString := "2025-01-01T00:00:00Z"
		layout := time.RFC3339

		parsedTime, err := time.Parse(layout, dateString)
		assert.NoError(t, err)

		mockService := hd.Sv.(*mocks.MockIProductBatchesService)
		mockService.On("Post", &model.ProductBatches{
			ID:                 0,
			BatchNumber:        "B01",
			CurrentQuantity:    10,
			CurrentTemperature: 10.00,
			MinimumTemperature: 5.00,
			DueDate:            parsedTime,
			InitialQuantity:    5,
			ManufacturingDate:  parsedTime,
			ManufacturingHour:  10,
			ProductID:          1,
			SectionID:          1,
		}).Return(model.ProductBatches{}, customerror.HandleError("product batches", customerror.ErrorNotFound, ""))

		reqBody := []byte(`{
		"batch_number": "B01",
		"current_quantity": 10,
		"current_temperature": 10.00,
		"minimum_temperature": 5.00,
		"due_date": "2025-01-01T00:00:00Z",
		"initial_quantity": 5,
		"manufacturing_date": "2025-01-01T00:00:00Z",
		"manufacturing_hour": 10,
		"product_id": 1,
		"section_id": 1}`)

		request := httptest.NewRequest(http.MethodPost, "/reportProducts", bytes.NewReader(reqBody))
		response := httptest.NewRecorder()
		hd.Post(response, request)

		expectedJson := `{"message": "product batches not found"}`
		assert.Equal(t, http.StatusNotFound, response.Code)
		assert.JSONEq(t, expectedJson, response.Body.String())
		mockService.AssertExpectations(t)
	})

	t.Run("return an internal error", func(t *testing.T) {
		hd := setupProductBatches(t)

		dateString := "2025-01-01T00:00:00Z"
		layout := time.RFC3339

		parsedTime, err := time.Parse(layout, dateString)
		assert.NoError(t, err)

		mockService := hd.Sv.(*mocks.MockIProductBatchesService)
		mockService.On("Post", &model.ProductBatches{
			ID:                 0,
			BatchNumber:        "B01",
			CurrentQuantity:    10,
			CurrentTemperature: 10.00,
			MinimumTemperature: 5.00,
			DueDate:            parsedTime,
			InitialQuantity:    5,
			ManufacturingDate:  parsedTime,
			ManufacturingHour:  10,
			ProductID:          1,
			SectionID:          1,
		}).Return(model.ProductBatches{}, errors.New("internal error"))

		reqBody := []byte(`{
		"batch_number": "B01",
		"current_quantity": 10,
		"current_temperature": 10.00,
		"minimum_temperature": 5.00,
		"due_date": "2025-01-01T00:00:00Z",
		"initial_quantity": 5,
		"manufacturing_date": "2025-01-01T00:00:00Z",
		"manufacturing_hour": 10,
		"product_id": 1,
		"section_id": 1}`)

		request := httptest.NewRequest(http.MethodPost, "/reportProducts", bytes.NewReader(reqBody))
		response := httptest.NewRecorder()
		hd.Post(response, request)

		expectedJson := `{"message": "Unable to create product batches"}`
		assert.Equal(t, http.StatusInternalServerError, response.Code)
		assert.JSONEq(t, expectedJson, response.Body.String())
		mockService.AssertExpectations(t)
	})

	t.Run("given an invalid request body then return error", func(t *testing.T) {
		hd := setupProductBatches(t)

		reqBody := []byte(`{}`)

		request := httptest.NewRequest(http.MethodPost, "/reportProducts", bytes.NewReader(reqBody))
		response := httptest.NewRecorder()
		hd.Post(response, request)

		expectedJson := `{"message": "request body cannot be empty"}`
		assert.Equal(t, http.StatusUnprocessableEntity, response.Code)
		assert.JSONEq(t, expectedJson, response.Body.String())
	})

	t.Run("given an invalid request body then return error", func(t *testing.T) {
		hd := setupProductBatches(t)

		reqBody := []byte(`{"batch_number": "B01",
		"current_quantity": 10,
		"current_temperature": 10.00,
		"minimum_temperature": "5.00"}`)

		request := httptest.NewRequest(http.MethodPost, "/reportProducts", bytes.NewReader(reqBody))
		response := httptest.NewRecorder()
		hd.Post(response, request)

		expectedJson := `{"message": "invalid request body"}`
		assert.Equal(t, http.StatusUnprocessableEntity, response.Code)
		assert.JSONEq(t, expectedJson, response.Body.String())
	})
}
