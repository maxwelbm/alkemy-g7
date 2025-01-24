package handler_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/maxwelbm/alkemy-g7.git/internal/handler"
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/maxwelbm/alkemy-g7.git/internal/service"
	"github.com/stretchr/testify/assert"
)

func setupWarehouse() *handler.WarehouseHandler {
	mockServiceWarehouse := new(service.WarehouseServiceMock)
	hd := handler.NewWareHouseHandler(mockServiceWarehouse)
	return hd
}
func TestHandlerGetAllWarehouse(t *testing.T) {

	t.Run("GetAllWarehouse", func(t *testing.T) {
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
}
