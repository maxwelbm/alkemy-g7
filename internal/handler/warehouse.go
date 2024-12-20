package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"
	"github.com/maxwelbm/alkemy-g7.git/internal/handler/request"
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/maxwelbm/alkemy-g7.git/internal/service/interfaces"
)

type WarehouseHandler struct {
	srv interfaces.IWarehouseService
}

func NewWareHouseHandler(srv interfaces.IWarehouseService) *WarehouseHandler {
	return &WarehouseHandler{srv: srv}
}

func (h *WarehouseHandler) GetAllWareHouse() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		wareHouse, err := h.srv.GetAllWareHouse()

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var data []model.WareHouse

		for _, value := range wareHouse {
			data = append(data, model.WareHouse{
				Id:                 value.Id,
				Address:            value.Address,
				Telephone:          value.Telephone,
				WareHouseCode:      value.WareHouseCode,
				MinimunCapacity:    value.MinimunCapacity,
				MinimunTemperature: value.MinimunTemperature,
			})
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"data": data,
		})

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

func (h *WarehouseHandler) DeleteByIdWareHouse() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))

		if err != nil {
			response.JSON(w, http.StatusBadRequest, "invalid id")
			return
		}

		err = h.srv.DeleteByIdWareHouse(id)

		if err != nil {
			response.JSON(w, http.StatusNotFound, "warehouse not found")
			return
		}

		response.JSON(w, http.StatusNoContent, "")
	}
}

func (h *WarehouseHandler) PostWareHouse() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var reqBody request.RequestBody
		if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
			response.JSON(w, http.StatusBadRequest, "invalid request body")
			return
		}

		if !request.IsValidateFields(reqBody) {
			response.JSON(w, http.StatusUnprocessableEntity, "invalid request body")
			return
		}
		warehouse, err := h.srv.PostWareHouse(model.WareHouse{
			Address:            reqBody.Address,
			Telephone:          reqBody.Telephone,
			WareHouseCode:      reqBody.WareHouseCode,
			MinimunCapacity:    reqBody.MinimunCapacity,
			MinimunTemperature: reqBody.MinimunTemperature,
		})

		if err != nil {
			response.JSON(w, http.StatusBadRequest, err.Error())
			return
		}

		response.JSON(w, http.StatusCreated, map[string]any{
			"data": warehouse,
		})
	}
}