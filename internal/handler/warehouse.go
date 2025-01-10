package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"
	responses "github.com/maxwelbm/alkemy-g7.git/internal/handler/responses"
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/maxwelbm/alkemy-g7.git/internal/service/interfaces"
	"github.com/maxwelbm/alkemy-g7.git/pkg/custom_error"
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

		fmt.Println(wareHouse)
		if err != nil {
			response.JSON(w, http.StatusInternalServerError, responses.CreateResponseBody(err.Error(), nil))
			return
		}

		response.JSON(w, http.StatusOK, responses.CreateResponseBody("", wareHouse))

	}
}

func (h *WarehouseHandler) GetWareHouseById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))

		if err != nil {
			response.JSON(w, http.StatusBadRequest, responses.CreateResponseBody("invalid id", nil))
			return
		}

		warehouse, err := h.srv.GetByIdWareHouse(id)

		if err != nil {
			if err, ok := err.(*custom_error.WareHouseError); ok {
				response.JSON(w, err.Code, responses.CreateResponseBody(err.Error(), nil))
				return
			}

			response.JSON(w, http.StatusInternalServerError, responses.CreateResponseBody("unable to search warehouse", nil))
			return
		}

		response.JSON(w, http.StatusOK, responses.CreateResponseBody("", warehouse))

	}
}

func (h *WarehouseHandler) DeleteByIdWareHouse() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))

		if err != nil {
			response.JSON(w, http.StatusBadRequest, responses.CreateResponseBody("invalid id", nil))
			return
		}

		err = h.srv.DeleteByIdWareHouse(id)

		if err != nil {
			if err, ok := err.(*custom_error.WareHouseError); ok {
				response.JSON(w, err.Code, responses.CreateResponseBody(err.Error(), nil))
				return
			}
		}

		response.JSON(w, http.StatusNoContent, nil)
	}
}

func (h *WarehouseHandler) PostWareHouse() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var reqBody model.WareHouse

		if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
			response.JSON(w, http.StatusBadRequest, responses.CreateResponseBody("invalid request body", nil))
			return
		}

		err := reqBody.ValidateEmptyFields(false)

		if err != nil {
			response.JSON(w, http.StatusUnprocessableEntity, responses.CreateResponseBody(err.Error(), nil))
			return
		}

		warehouse, err := h.srv.PostWareHouse(reqBody)

		if err != nil {
			if err, ok := err.(*custom_error.WareHouseError); ok {
				response.JSON(w, err.Code, responses.CreateResponseBody(err.Error(), nil))
				return
			}

			response.JSON(w, http.StatusInternalServerError, responses.CreateResponseBody("unable to post warehouse", nil))
			return
		}

		response.JSON(w, http.StatusCreated, responses.CreateResponseBody("", warehouse))
	}
}

func (h *WarehouseHandler) UpdateWareHouse() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var reqBody model.WareHouse
		id, err := strconv.Atoi(chi.URLParam(r, "id"))

		if err != nil {
			response.JSON(w, http.StatusBadRequest, responses.CreateResponseBody("invalid id", nil))
			return
		}

		if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
			response.JSON(w, http.StatusBadRequest, responses.CreateResponseBody("invalid request body", nil))
			return
		}

		err = reqBody.ValidateEmptyFields(true)

		if err != nil {
			response.JSON(w, http.StatusUnprocessableEntity, responses.CreateResponseBody(err.Error(), nil))
			return
		}

		warehouse, err := h.srv.UpdateWareHouse(id, reqBody)

		if err != nil {
			if err, ok := err.(*custom_error.WareHouseError); ok {
				response.JSON(w, err.Code, responses.CreateResponseBody(err.Error(), nil))
				return
			}

			response.JSON(w, http.StatusInternalServerError, responses.CreateResponseBody("unable to update warehouse", nil))
			return
		}

		response.JSON(w, http.StatusOK, responses.CreateResponseBody("", warehouse))
	}
}
