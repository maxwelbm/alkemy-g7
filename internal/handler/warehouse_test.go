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
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/maxwelbm/alkemy-g7.git/internal/service"
	"github.com/maxwelbm/alkemy-g7.git/pkg/customError"
	"github.com/stretchr/testify/assert"
)

func setupWarehouse() *handler.WarehouseHandler {
	mockServiceWarehouse := new(service.WarehouseServiceMock)
	hd := handler.NewWareHouseHandler(mockServiceWarehouse)
	return hd
}
func TestHandlerGetAllWarehouse(t *testing.T) {

	t.Run("GetAllWarehouse return sucess", func(t *testing.T) {
		hd := setupWarehouse()

		expectedWarehouse := []model.WareHouse{{
			Id:                 1,
			WareHouseCode:      "test",
			Telephone:          "test",
			MinimunCapacity:    1,
			MinimunTemperature: 1,
			Address:            "test",
		}, {
			Id:                 2,
			WareHouseCode:      "test",
			Telephone:          "test",
			MinimunCapacity:    1,
			MinimunTemperature: 1,
			Address:            "test",
		}}

		request := httptest.NewRequest(http.MethodGet, "/api/v1/warehouses", nil)
		response := httptest.NewRecorder()
		mockServiceWarehouse := hd.Srv.(*service.WarehouseServiceMock)
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
		hd := setupWarehouse()

		request := httptest.NewRequest(http.MethodGet, "/api/v1/warehouses", nil)
		response := httptest.NewRecorder()
		mockServiceWarehouse := hd.Srv.(*service.WarehouseServiceMock)
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
		hd := setupWarehouse()
		mockServiceWarehouse := hd.Srv.(*service.WarehouseServiceMock)

		r := chi.NewRouter()
		r.Get("/api/v1/warehouses/{id}", hd.GetWareHouseById())

		expectedWarehouse := model.WareHouse{
			Id:                 1,
			WareHouseCode:      "test",
			Telephone:          "test",
			MinimunCapacity:    1,
			MinimunTemperature: 1,
			Address:            "test",
		}

		mockServiceWarehouse.On("GetByIdWareHouse", 1).Return(expectedWarehouse, nil)

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
		hd := setupWarehouse()
		mockServiceWarehouse := hd.Srv.(*service.WarehouseServiceMock)

		r := chi.NewRouter()
		r.Get("/api/v1/warehouses/{id}", hd.GetWareHouseById())

		mockServiceWarehouse.On("GetByIdWareHouse", 30).Return(model.WareHouse{}, customError.NewWareHouseError(customError.ErrNotFound.Error(), "warehouse", http.StatusNotFound))

		request := httptest.NewRequest(http.MethodGet, "/api/v1/warehouses/"+strconv.Itoa(30), nil)

		response := httptest.NewRecorder()

		r.ServeHTTP(response, request)

		expectedJson := `{"message":"warehouse not found"}`

		assert.Equal(t, http.StatusNotFound, response.Code)
		assert.JSONEq(t, expectedJson, response.Body.String())
		mockServiceWarehouse.AssertExpectations(t)
	})

	t.Run("GetByIdWarehouse id invalid", func(t *testing.T) {
		hd := setupWarehouse()

		request := httptest.NewRequest(http.MethodGet, "/api/v1/warehouses/th", nil)

		response := httptest.NewRecorder()
		handler := hd.GetWareHouseById()
		handler.ServeHTTP(response, request)

		expectedJson := `{"message":"invalid id"}`

		assert.Equal(t, http.StatusBadRequest, response.Code)
		assert.JSONEq(t, expectedJson, response.Body.String())

	})

	t.Run("GetByIdWarehouse when service fails returns internal server error", func(t *testing.T) {
		hd := setupWarehouse()
		mockServiceWarehouse := hd.Srv.(*service.WarehouseServiceMock)

		r := chi.NewRouter()
		r.Get("/api/v1/warehouses/{id}", hd.GetWareHouseById())

		mockServiceWarehouse.On("GetByIdWareHouse", 2).Return(model.WareHouse{}, errors.New("error"))

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
		hd := setupWarehouse()
		mockServiceWarehouse := hd.Srv.(*service.WarehouseServiceMock)

		r := chi.NewRouter()
		r.Delete("/api/v1/warehouses/{id}", hd.DeleteByIdWareHouse())

		mockServiceWarehouse.On("DeleteByIdWareHouse", 1).Return(nil)

		request := httptest.NewRequest(http.MethodDelete, "/api/v1/warehouses/"+strconv.Itoa(1), nil)

		response := httptest.NewRecorder()
		r.ServeHTTP(response, request)

		assert.Equal(t, http.StatusNoContent, response.Code)
		mockServiceWarehouse.AssertExpectations(t)
	})

	t.Run("DeleteByIdWarehouse not found", func(t *testing.T) {
		hd := setupWarehouse()
		mockServiceWarehouse := hd.Srv.(*service.WarehouseServiceMock)

		r := chi.NewRouter()
		r.Delete("/api/v1/warehouses/{id}", hd.DeleteByIdWareHouse())

		mockServiceWarehouse.On("DeleteByIdWareHouse", 30).Return(customError.NewWareHouseError(customError.ErrNotFound.Error(), "warehouse", http.StatusNotFound))

		request := httptest.NewRequest(http.MethodDelete, "/api/v1/warehouses/"+strconv.Itoa(30), nil)

		response := httptest.NewRecorder()

		r.ServeHTTP(response, request)

		expectedJson := `{"message":"warehouse not found"}`

		assert.Equal(t, http.StatusNotFound, response.Code)
		assert.JSONEq(t, expectedJson, response.Body.String())
		mockServiceWarehouse.AssertExpectations(t)
	})

	t.Run("DeleteByIdWarehouse id invalid", func(t *testing.T) {
		hd := setupWarehouse()

		request := httptest.NewRequest(http.MethodDelete, "/api/v1/warehouses/th", nil)

		response := httptest.NewRecorder()
		handler := hd.DeleteByIdWareHouse()
		handler.ServeHTTP(response, request)

		expectedJson := `{"message":"invalid id"}`

		assert.Equal(t, http.StatusBadRequest, response.Code)
		assert.JSONEq(t, expectedJson, response.Body.String())
	})
}

func TestHandlerPostWarehouse(t *testing.T) {
	t.Run("PostWarehouse create sucess", func(t *testing.T) {
		hd := setupWarehouse()
		mockServiceWarehouse := hd.Srv.(*service.WarehouseServiceMock)

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
			Id:                 1,
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
		hd := setupWarehouse()
		mockServiceWarehouse := hd.Srv.(*service.WarehouseServiceMock)

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
		hd := setupWarehouse()

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
		hd := setupWarehouse()
		mockServiceWarehouse := hd.Srv.(*service.WarehouseServiceMock)

		r := chi.NewRouter()
		r.Post("/api/v1/warehouses", hd.PostWareHouse())

		warehouse := model.WareHouse{
			WareHouseCode:      "warehouse_code",
			Address:            "address",
			Telephone:          "telephone",
			MinimunCapacity:    1,
			MinimunTemperature: 1,
		}

		mockServiceWarehouse.On("PostWareHouse", warehouse).Return(model.WareHouse{}, customError.NewWareHouseError(customError.ErrConflict.Error(), "warehouse_code", http.StatusConflict))

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
		hd := setupWarehouse()
		mockServiceWarehouse := hd.Srv.(*service.WarehouseServiceMock)

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
