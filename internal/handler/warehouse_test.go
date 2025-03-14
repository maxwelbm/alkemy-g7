package handler_test

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/maxwelbm/alkemy-g7.git/internal/handler"
	"github.com/maxwelbm/alkemy-g7.git/internal/mocks"
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/maxwelbm/alkemy-g7.git/pkg/customerror"

	"github.com/stretchr/testify/assert"
)

func setupWarehouse(t *testing.T) *handler.WarehouseHandler {
	mockServiceWarehouse := mocks.NewMockIWarehouseService(t)
	hd := handler.NewWareHouseHandler(mockServiceWarehouse, logMock)
	return hd
}
func TestHandlerGetAllWarehouse(t *testing.T) {

	t.Run("GetAllWarehouse return sucess", func(t *testing.T) {
		hd := setupWarehouse(t)

		expectedWarehouse := []model.WareHouse{{
			ID:                 1,
			WareHouseCode:      "test",
			Telephone:          "test",
			MinimunCapacity:    1,
			MinimunTemperature: 1,
			Address:            "test",
		}, {
			ID:                 2,
			WareHouseCode:      "test",
			Telephone:          "test",
			MinimunCapacity:    1,
			MinimunTemperature: 1,
			Address:            "test",
		}}

		request := httptest.NewRequest(http.MethodGet, "/api/v1/warehouses", nil)
		response := httptest.NewRecorder()
		mockServiceWarehouse := hd.Srv.(*mocks.MockIWarehouseService)
		mockServiceWarehouse.On("GetAllWareHouse").Return(expectedWarehouse, nil)

		handler := hd.GetAllWareHouse()
		handler.ServeHTTP(response, request)

		assert.Equal(t, http.StatusOK, response.Code)

		expectedJson := `{
		"data": [
			{
				"id": 1,
				"warehouse_code": "test",
				"telephone": "test",
				"minimun_capacity": 1,
				"minimun_temperature": 1,
				"address": "test"
			},
			{
				"id": 2,
				"warehouse_code": "test",
				"telephone": "test",
				"minimun_capacity": 1,
				"minimun_temperature": 1,
				"address": "test"
			}
		]}`

		assert.Equal(t, http.StatusOK, response.Code)
		assert.JSONEq(t, expectedJson, response.Body.String())
		mockServiceWarehouse.AssertExpectations(t)

	})

	t.Run("GetAllWarehouse return error", func(t *testing.T) {
		hd := setupWarehouse(t)

		request := httptest.NewRequest(http.MethodGet, "/api/v1/warehouses", nil)
		response := httptest.NewRecorder()
		mockServiceWarehouse := hd.Srv.(*mocks.MockIWarehouseService)
		mockServiceWarehouse.On("GetAllWareHouse").Return([]model.WareHouse{}, errors.New("not found warehouses"))

		handler := hd.GetAllWareHouse()
		handler.ServeHTTP(response, request)

		assert.Equal(t, http.StatusInternalServerError, response.Code)

		expectedJson := `{"message":"not found warehouses"}`

		assert.Equal(t, http.StatusInternalServerError, response.Code)
		assert.JSONEq(t, expectedJson, response.Body.String())
	})
}

func TestHandlerGetWarehouseById(t *testing.T) {
	t.Run("GetByIdWareHouse return sucess", func(t *testing.T) {
		hd := setupWarehouse(t)
		mockServiceWarehouse := hd.Srv.(*mocks.MockIWarehouseService)

		r := chi.NewRouter()
		r.Get("/api/v1/warehouses/{id}", hd.GetWareHouseByID())

		expectedWarehouse := model.WareHouse{
			ID:                 1,
			WareHouseCode:      "test",
			Telephone:          "test",
			MinimunCapacity:    1,
			MinimunTemperature: 1,
			Address:            "test",
		}

		mockServiceWarehouse.On("GetByIDWareHouse", 1).Return(expectedWarehouse, nil)

		request := httptest.NewRequest(http.MethodGet, "/api/v1/warehouses/"+strconv.Itoa(1), nil)

		response := httptest.NewRecorder()
		r.ServeHTTP(response, request)

		expectedJson := `{
		"data": {
			"id": 1,
			"warehouse_code": "test",
			"telephone": "test",
			"minimun_capacity": 1,
			"minimun_temperature": 1,
			"address": "test"
		}
	}`

		assert.Equal(t, http.StatusOK, response.Code)
		assert.JSONEq(t, expectedJson, response.Body.String())
		mockServiceWarehouse.AssertExpectations(t)

	})

	t.Run("GetByIdWareHouse not found", func(t *testing.T) {
		hd := setupWarehouse(t)
		mockServiceWarehouse := hd.Srv.(*mocks.MockIWarehouseService)

		r := chi.NewRouter()
		r.Get("/api/v1/warehouses/{id}", hd.GetWareHouseByID())

		mockServiceWarehouse.On("GetByIDWareHouse", 30).Return(model.WareHouse{}, customerror.NewWareHouseError(customerror.ErrNotFound.Error(), "warehouse", http.StatusNotFound))

		request := httptest.NewRequest(http.MethodGet, "/api/v1/warehouses/"+strconv.Itoa(30), nil)

		response := httptest.NewRecorder()

		r.ServeHTTP(response, request)

		expectedJson := `{"message":"warehouse not found"}`

		assert.Equal(t, http.StatusNotFound, response.Code)
		assert.JSONEq(t, expectedJson, response.Body.String())
		mockServiceWarehouse.AssertExpectations(t)
	})

	t.Run("GetByIdWarehouse id invalid", func(t *testing.T) {
		hd := setupWarehouse(t)

		request := httptest.NewRequest(http.MethodGet, "/api/v1/warehouses/th", nil)

		response := httptest.NewRecorder()

		handler := hd.GetWareHouseByID()

		handler.ServeHTTP(response, request)

		expectedJson := `{"message":"invalid id"}`

		assert.Equal(t, http.StatusBadRequest, response.Code)
		assert.JSONEq(t, expectedJson, response.Body.String())

	})

	t.Run("GetByIdWarehouse when service fails returns internal server error", func(t *testing.T) {
		hd := setupWarehouse(t)
		mockServiceWarehouse := hd.Srv.(*mocks.MockIWarehouseService)

		r := chi.NewRouter()
		r.Get("/api/v1/warehouses/{id}", hd.GetWareHouseByID())

		mockServiceWarehouse.On("GetByIDWareHouse", 2).Return(model.WareHouse{}, errors.New("internal server error"))

		request := httptest.NewRequest(http.MethodGet, "/api/v1/warehouses/"+strconv.Itoa(2), nil)

		response := httptest.NewRecorder()

		r.ServeHTTP(response, request)

		expectedJson := `{"message":"unable to search warehouse"}`

		assert.Equal(t, http.StatusInternalServerError, response.Code)
		assert.JSONEq(t, expectedJson, response.Body.String())
		mockServiceWarehouse.AssertExpectations(t)
	})
}

func TestHandlerDeleteByIdWarehouse(t *testing.T) {
	t.Run("DeleteByIdWarehouse return sucess", func(t *testing.T) {
		hd := setupWarehouse(t)
		mockServiceWarehouse := hd.Srv.(*mocks.MockIWarehouseService)

		r := chi.NewRouter()
		r.Delete("/api/v1/warehouses/{id}", hd.DeleteByIDWareHouse())

		mockServiceWarehouse.On("DeleteByIDWareHouse", 1).Return(nil)

		request := httptest.NewRequest(http.MethodDelete, "/api/v1/warehouses/"+strconv.Itoa(1), nil)

		response := httptest.NewRecorder()
		r.ServeHTTP(response, request)

		assert.Equal(t, http.StatusNoContent, response.Code)
		mockServiceWarehouse.AssertExpectations(t)
	})

	t.Run("DeleteByIdWarehouse not found", func(t *testing.T) {
		hd := setupWarehouse(t)
		mockServiceWarehouse := hd.Srv.(*mocks.MockIWarehouseService)

		r := chi.NewRouter()
		r.Delete("/api/v1/warehouses/{id}", hd.DeleteByIDWareHouse())

		mockServiceWarehouse.On("DeleteByIDWareHouse", 30).Return(customerror.NewWareHouseError(customerror.ErrNotFound.Error(), "warehouse", http.StatusNotFound))

		request := httptest.NewRequest(http.MethodDelete, "/api/v1/warehouses/"+strconv.Itoa(30), nil)

		response := httptest.NewRecorder()

		r.ServeHTTP(response, request)

		expectedJson := `{"message":"warehouse not found"}`

		assert.Equal(t, http.StatusNotFound, response.Code)
		assert.JSONEq(t, expectedJson, response.Body.String())
		mockServiceWarehouse.AssertExpectations(t)
	})

	t.Run("DeleteByIdWarehouse id invalid", func(t *testing.T) {
		hd := setupWarehouse(t)

		request := httptest.NewRequest(http.MethodDelete, "/api/v1/warehouses/th", nil)

		response := httptest.NewRecorder()

		handler := hd.DeleteByIDWareHouse()

		handler.ServeHTTP(response, request)

		expectedJson := `{"message":"invalid id"}`

		assert.Equal(t, http.StatusBadRequest, response.Code)
		assert.JSONEq(t, expectedJson, response.Body.String())
	})
}

func TestHandlerPostWarehouse(t *testing.T) {
	t.Run("PostWarehouse create sucess", func(t *testing.T) {
		hd := setupWarehouse(t)
		mockServiceWarehouse := hd.Srv.(*mocks.MockIWarehouseService)

		r := chi.NewRouter()
		r.Post("/api/v1/warehouses", hd.PostWareHouse())

		warehouse := model.WareHouse{
			WareHouseCode:      "warehouse_code",
			Address:            "address",
			Telephone:          "telephone",
			MinimunCapacity:    1,
			MinimunTemperature: 1,
		}

		mockServiceWarehouse.On("PostWareHouse", warehouse).Return(model.WareHouse{
			ID:                 1,
			WareHouseCode:      "warehouse_code",
			Address:            "address",
			Telephone:          "telephone",
			MinimunCapacity:    1,
			MinimunTemperature: 1,
		}, nil)

		reqBody := []byte(`{
			"warehouse_code": "warehouse_code",
			"address": "address",
			"telephone": "telephone",
			"minimun_capacity": 1,
			"minimun_temperature": 1}`,
		)

		request := httptest.NewRequest(http.MethodPost, "/api/v1/warehouses", bytes.NewReader(reqBody))

		response := httptest.NewRecorder()
		r.ServeHTTP(response, request)

		assert.Equal(t, http.StatusCreated, response.Code)

		expectedJson := `{
		"data": {
			"id": 1,
			"warehouse_code": "warehouse_code",
			"telephone": "telephone",
			"minimun_capacity": 1,
			"minimun_temperature": 1,
			"address": "address"
		}
	}`

		assert.JSONEq(t, expectedJson, response.Body.String())
		mockServiceWarehouse.AssertExpectations(t)
	})

	t.Run("PostWarehouse required fields not found", func(t *testing.T) {
		hd := setupWarehouse(t)
		mockServiceWarehouse := hd.Srv.(*mocks.MockIWarehouseService)

		r := chi.NewRouter()
		r.Post("/api/v1/warehouses", hd.PostWareHouse())

		reqBody := []byte(`{
			}`,
		)

		request := httptest.NewRequest(http.MethodPost, "/api/v1/warehouses", bytes.NewReader(reqBody))

		response := httptest.NewRecorder()

		expectedJson := `{"message":"Field(s) address, telephone, warehouse_code, minimun_capacity, minimun_temperature cannot be empty or invalid"}`

		r.ServeHTTP(response, request)

		assert.Equal(t, http.StatusUnprocessableEntity, response.Code)
		assert.JSONEq(t, expectedJson, response.Body.String())
		mockServiceWarehouse.AssertExpectations(t)
	})

	t.Run("PostWarehouse invalid request body", func(t *testing.T) {
		hd := setupWarehouse(t)

		r := chi.NewRouter()
		r.Post("/api/v1/warehouses", hd.PostWareHouse())

		reqBody := []byte(`{
			"warehouse_code": "warehouse_code",
			"address": "",
			"telephone": "",
			"minimun_capacity": 1,
			"minimun_temperature": 1`,
		)

		request := httptest.NewRequest(http.MethodPost, "/api/v1/warehouses", bytes.NewReader(reqBody))

		response := httptest.NewRecorder()

		expectedJson := `{"message":"invalid request body"}`

		r.ServeHTTP(response, request)

		assert.Equal(t, http.StatusBadRequest, response.Code)
		assert.JSONEq(t, expectedJson, response.Body.String())
	})

	t.Run("PostWarehouse warehouse_code conflit", func(t *testing.T) {
		hd := setupWarehouse(t)
		mockServiceWarehouse := hd.Srv.(*mocks.MockIWarehouseService)

		r := chi.NewRouter()
		r.Post("/api/v1/warehouses", hd.PostWareHouse())

		warehouse := model.WareHouse{
			WareHouseCode:      "warehouse_code",
			Address:            "address",
			Telephone:          "telephone",
			MinimunCapacity:    1,
			MinimunTemperature: 1,
		}

		mockServiceWarehouse.On("PostWareHouse", warehouse).Return(model.WareHouse{}, customerror.NewWareHouseError(customerror.ErrConflict.Error(), "warehouse_code", http.StatusConflict))

		reqBody := []byte(`{
			"warehouse_code": "warehouse_code",
			"address": "address",
			"telephone": "telephone",
			"minimun_capacity": 1,
			"minimun_temperature": 1}`,
		)

		request := httptest.NewRequest(http.MethodPost, "/api/v1/warehouses", bytes.NewReader(reqBody))

		response := httptest.NewRecorder()

		expectedJson := `{"message":"warehouse_code it already exists"}`

		r.ServeHTTP(response, request)

		assert.Equal(t, http.StatusConflict, response.Code)
		assert.JSONEq(t, expectedJson, response.Body.String())
		mockServiceWarehouse.AssertExpectations(t)
	})

	t.Run("PostWarehouse when service fails returns internal server error", func(t *testing.T) {
		hd := setupWarehouse(t)
		mockServiceWarehouse := hd.Srv.(*mocks.MockIWarehouseService)

		r := chi.NewRouter()
		r.Post("/api/v1/warehouses", hd.PostWareHouse())

		warehouse := model.WareHouse{
			WareHouseCode:      "warehouse_code",
			Address:            "address",
			Telephone:          "telephone",
			MinimunCapacity:    1,
			MinimunTemperature: 1,
		}

		mockServiceWarehouse.On("PostWareHouse", warehouse).Return(model.WareHouse{}, errors.New("internal server error"))

		reqBody := []byte(`{
			"warehouse_code": "warehouse_code",
			"address": "address",
			"telephone": "telephone",
			"minimun_capacity": 1,
			"minimun_temperature": 1}`,
		)

		request := httptest.NewRequest(http.MethodPost, "/api/v1/warehouses", bytes.NewReader(reqBody))

		response := httptest.NewRecorder()

		expectedJson := `{"message":"unable to post warehouse"}`

		r.ServeHTTP(response, request)

		assert.Equal(t, http.StatusInternalServerError, response.Code)
		assert.JSONEq(t, expectedJson, response.Body.String())
		mockServiceWarehouse.AssertExpectations(t)
	})
}

func TestHandlerUpdateWarehouse(t *testing.T) {
	t.Run("UpdateWarehouse return sucess", func(t *testing.T) {
		hd := setupWarehouse(t)
		mockServiceWarehouse := hd.Srv.(*mocks.MockIWarehouseService)

		warehouse := model.WareHouse{
			ID:                 1,
			WareHouseCode:      "warehouse_code",
			Address:            "Update Address",
			Telephone:          "telephone",
			MinimunCapacity:    1,
			MinimunTemperature: 1,
		}

		mockServiceWarehouse.On("UpdateWareHouse", 1, model.WareHouse{
			Address: "Update Address",
		}).Return(warehouse, nil)

		body := []byte(`{
			"address": "Update Address"
		}`)

		request := httptest.NewRequest(http.MethodPatch, "/api/v1/warehouses/1", bytes.NewReader(body))
		response := httptest.NewRecorder()

		expectedJson := `{
			"data": {
				"id": 1,
				"warehouse_code": "warehouse_code",
				"address": "Update Address",
				"telephone": "telephone",
				"minimun_capacity": 1,
				"minimun_temperature": 1
			}
		}`

		r := chi.NewRouter()
		r.Patch("/api/v1/warehouses/{id}", hd.UpdateWareHouse())

		r.ServeHTTP(response, request)

		assert.Equal(t, http.StatusOK, response.Code)
		assert.JSONEq(t, expectedJson, response.Body.String())
		mockServiceWarehouse.AssertExpectations(t)
	})

	t.Run("UpdateWarehouse invalid id", func(t *testing.T) {
		hd := setupWarehouse(t)

		request := httptest.NewRequest(http.MethodPatch, "/api/v1/warehouses/th", nil)
		response := httptest.NewRecorder()

		expectedJson := `{"message":"invalid id"}`

		hd.UpdateWareHouse().ServeHTTP(response, request)

		assert.Equal(t, http.StatusBadRequest, response.Code)
		assert.JSONEq(t, expectedJson, response.Body.String())
	})

	t.Run("UpdateWarehouse invalid request body", func(t *testing.T) {
		hd := setupWarehouse(t)

		r := chi.NewRouter()
		r.Patch("/api/v1/warehouses/{id}", hd.UpdateWareHouse())

		reqBody := []byte(`{
			"warehouse_code": "warehouse_code",
			"address": "",
			"telephone": "",
			"minimun_capacity": 1,
			"minimun_temperature": 1`,
		)

		request := httptest.NewRequest(http.MethodPatch, "/api/v1/warehouses/1", bytes.NewReader(reqBody))
		response := httptest.NewRecorder()

		expectedJson := `{"message":"invalid request body"}`

		r.ServeHTTP(response, request)

		assert.Equal(t, http.StatusBadRequest, response.Code)
		assert.JSONEq(t, expectedJson, response.Body.String())
	})

	t.Run("UpdateWarehouse internal server error", func(t *testing.T) {
		hd := setupWarehouse(t)
		mockServiceWarehouse := hd.Srv.(*mocks.MockIWarehouseService)

		mockServiceWarehouse.On("UpdateWareHouse", 1, model.WareHouse{
			Address: "Update Address",
		}).Return(model.WareHouse{}, errors.New("internal server error"))

		body := []byte(`{
			"address": "Update Address"
		}`)

		request := httptest.NewRequest(http.MethodPatch, "/api/v1/warehouses/1", bytes.NewReader(body))
		response := httptest.NewRecorder()

		expectedJson := `{"message":"unable to update warehouse"}`

		r := chi.NewRouter()
		r.Patch("/api/v1/warehouses/{id}", hd.UpdateWareHouse())

		r.ServeHTTP(response, request)

		assert.Equal(t, http.StatusInternalServerError, response.Code)
		assert.JSONEq(t, expectedJson, response.Body.String())
		mockServiceWarehouse.AssertExpectations(t)
	})
}
