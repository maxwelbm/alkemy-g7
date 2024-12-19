package handler

import (
	"net/http"

	"github.com/bootcamp-go/web/response"
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

		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    wareHouse,
		})
	}
}
