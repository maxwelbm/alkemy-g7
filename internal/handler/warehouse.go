package handler

import (
	"net/http"
	"strconv"

	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/maxwelbm/alkemy-g7.git/internal/service"
)

type WarehouseHandler struct {
	srv service.WareHouseDefault
}

func NewWareHouseHandler(srv service.WareHouseDefault) *WarehouseHandler {
	return &WarehouseHandler{srv: srv}
}

func (h *WarehouseHandler) GetAllWareHouse() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		wareHouse, err := h.srv.GetAllWareHouse()

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var data []model.WareHouseJson

		for _, value := range wareHouse {
			data = append(data, model.WareHouseJson{
				Id:                 value.Id,
				Address:            value.Address,
				Telephone:          value.Telephone,
				WareHouseCode:      value.WareHouseCode,
				MinimunCapacity:    value.MinimunCapacity,
				MinimunTemperature: value.MinimunTemperature,
			})
		}
		responseJson := model.WareHouseRes{Data: data}

		response.JSON(w, http.StatusOK, responseJson)

	}
}

func (h *WarehouseHandler) GetWareHouseById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))

		if err != nil {
			response.JSON(w, http.StatusBadRequest, "invalid id")
			return
		}

		warehouse, err := h.srv.GetByIdWareHouse(id)

		if err != nil {
			response.JSON(w, http.StatusNotFound, "warehouse not found")
			return
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"data": warehouse,
		})

	}

}
