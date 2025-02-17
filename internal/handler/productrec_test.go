package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/maxwelbm/alkemy-g7.git/internal/mocks"
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/maxwelbm/alkemy-g7.git/pkg/customerror"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateProductRecServ(t *testing.T) {
	t.Run("Success Create Product Record", func(t *testing.T) {
		productRecServiceMock := new(mocks.MockIProductRecService)
		prh := NewProductRecHandler(productRecServiceMock)

		productRecord := model.ProductRecords{
			ID:            1,
			ProductID:     1,
			PurchasePrice: 11.9,
			SalePrice:     32.4,
		}

		productRecServiceMock.On("CreateProductRecords", productRecord).Return(productRecord, nil)

		body, _ := json.Marshal(productRecord)
		req := httptest.NewRequest("POST", "/product-records", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		res := httptest.NewRecorder()

		prh.CreateProductRecServ(res, req)

		assert.Equal(t, http.StatusCreated, res.Code)

		var response map[string]interface{}
		json.Unmarshal(res.Body.Bytes(), &response)
		assert.NotNil(t, response["data"])
	})

	t.Run("Error Creating Product Record - Invalid JSON", func(t *testing.T) {
		productRecServiceMock := new(mocks.MockIProductRecService)
		prh := NewProductRecHandler(productRecServiceMock)
		req := httptest.NewRequest("POST", "/product-records", bytes.NewBuffer([]byte("invalid json")))
		req.Header.Set("Content-Type", "application/json")
		res := httptest.NewRecorder()

		prh.CreateProductRecServ(res, req)

		assert.Equal(t, http.StatusUnprocessableEntity, res.Code)
		var response map[string]interface{}
		json.Unmarshal(res.Body.Bytes(), &response)
		assert.Equal(t, "json mal formatado ou invalido", response["message"])
	})

	t.Run("Error Creating Product Record - Service Error", func(t *testing.T) {
		productRecServiceMock := new(mocks.MockIProductRecService)
		prh := NewProductRecHandler(productRecServiceMock)
		productRecord := model.ProductRecords{
			ID:            1,
			ProductID:     1,
			PurchasePrice: 11.9,
			SalePrice:     32.4,
		}

		productRecServiceMock.On("CreateProductRecords", productRecord).Return(model.ProductRecords{}, &customerror.GenericError{Code: http.StatusInternalServerError, Message: "Unable to create product record"})

		body, _ := json.Marshal(productRecord)
		req := httptest.NewRequest("POST", "/product-records", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		res := httptest.NewRecorder()

		prh.CreateProductRecServ(res, req)

		assert.Equal(t, http.StatusInternalServerError, res.Code)
		var response map[string]interface{}
		json.Unmarshal(res.Body.Bytes(), &response)
		assert.Equal(t, " Unable to create product record", response["message"])
	})

	t.Run("Error Creating Product Record - Internal server error", func(t *testing.T) {
		productRecServiceMock := new(mocks.MockIProductRecService)
		prh := NewProductRecHandler(productRecServiceMock)
		productRecord := model.ProductRecords{
			ID:            1,
			ProductID:     1,
			PurchasePrice: 11.9,
			SalePrice:     32.4,
		}

		productRecServiceMock.On("CreateProductRecords", productRecord).Return(model.ProductRecords{}, errors.New("An error"))

		body, _ := json.Marshal(productRecord)
		req := httptest.NewRequest("POST", "/product-records", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		res := httptest.NewRecorder()

		prh.CreateProductRecServ(res, req)

		assert.Equal(t, http.StatusInternalServerError, res.Code)
		var response map[string]interface{}
		json.Unmarshal(res.Body.Bytes(), &response)
		assert.Equal(t, "unknow server error", response["message"])
	})
}

func TestGetProductRecReport(t *testing.T) {
	t.Run("Success Get Product Record Report with valid ID", func(t *testing.T) {
		productRecServiceMock := new(mocks.MockIProductRecService)
		prh := NewProductRecHandler(productRecServiceMock)

		productId := 1
		mockReports := []model.ProductRecordsReport{
			{ProductID: 1, Description: "Product A", RecordsCount: 2},
			{ProductID: 2, Description: "Product B", RecordsCount: 3},
		}
		productRecServiceMock.On("GetProductRecordReport", productId).Return(mockReports, nil)

		req := httptest.NewRequest("GET", "/product-records/report?id=1", nil)
		res := httptest.NewRecorder()

		prh.GetProductRecReport(res, req)

		assert.Equal(t, http.StatusOK, res.Code)

		var response map[string]interface{}
		json.Unmarshal(res.Body.Bytes(), &response)
		assert.NotNil(t, response["data"])
		assert.Equal(t, 2, len(response["data"].([]interface{})))
	})

	t.Run("Error Get Product Record Report - Invalid Parameter", func(t *testing.T) {
		productRecServiceMock := new(mocks.MockIProductRecService)
		prh := NewProductRecHandler(productRecServiceMock)

		req := httptest.NewRequest("GET", "/product-records/report?id=abc", nil)
		res := httptest.NewRecorder()

		prh.GetProductRecReport(res, req)

		assert.Equal(t, http.StatusBadRequest, res.Code)
		var response map[string]interface{}
		json.Unmarshal(res.Body.Bytes(), &response)
		assert.Equal(t, "Invalid Parameter", response["message"])
	})

	t.Run("Error Get Product Record Report - Service Error", func(t *testing.T) {
		productRecServiceMock := new(mocks.MockIProductRecService)
		prh := NewProductRecHandler(productRecServiceMock)

		productId := 1
		productRecServiceMock.On("GetProductRecordReport", productId).Return(nil, &customerror.GenericError{Code: http.StatusInternalServerError, Message: "Internal Server Error"})

		req := httptest.NewRequest("GET", "/product-records/report?id=1", nil)
		res := httptest.NewRecorder()

		prh.GetProductRecReport(res, req)

		assert.Equal(t, http.StatusInternalServerError, res.Code)
		var response map[string]interface{}
		json.Unmarshal(res.Body.Bytes(), &response)
		assert.Equal(t, " Internal Server Error", response["message"])
	})

	t.Run("Error Creating Product Record - Internal server error", func(t *testing.T) {
		productRecServiceMock := new(mocks.MockIProductRecService)
		prh := NewProductRecHandler(productRecServiceMock)


		productRecServiceMock.On("GetProductRecordReport", mock.Anything).Return(nil, errors.New("An error"))

		req := httptest.NewRequest("GET", "/product-records/report?id=1", nil)
		res := httptest.NewRecorder()

		prh.GetProductRecReport(res, req)

		assert.Equal(t, http.StatusInternalServerError, res.Code)
		var response map[string]interface{}
		json.Unmarshal(res.Body.Bytes(), &response)
		assert.Equal(t, "Internal Server Error", response["message"])
	})

	t.Run("Success Get Product Record Report with no itens in slice", func(t *testing.T) {
		productRecServiceMock := new(mocks.MockIProductRecService)
		prh := NewProductRecHandler(productRecServiceMock)

		productId := 1
		expected := "{\"message\":\"empty list\"}"
		mockReports := []model.ProductRecordsReport{}

		productRecServiceMock.On("GetProductRecordReport", productId).Return(mockReports, nil)

		req := httptest.NewRequest("GET", "/product-records/report?id=1", nil)
		res := httptest.NewRecorder()

		prh.GetProductRecReport(res, req)

		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, res.Body.String(), expected)

	})
}
