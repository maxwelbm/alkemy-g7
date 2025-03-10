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
	"github.com/maxwelbm/alkemy-g7.git/pkg/customerror"
	"github.com/maxwelbm/alkemy-g7.git/pkg/logger"
)

type WarehouseHandler struct {
	Srv interfaces.IWarehouseService
	log logger.Logger
}

func NewWareHouseHandler(srv interfaces.IWarehouseService, log logger.Logger) *WarehouseHandler {
	return &WarehouseHandler{Srv: srv, log: log}
}

// GetAllWareHouse retrieves all warehouses.
// @Summary Retrieve all warehouses
// @Description Fetch all registered warehouses from the database
// @Tags Warehouses
// @Produce json
// @Success 200 {object} model.WareHousesResponseSwagger
// @Failure 500 {object} model.ErrorResponseSwagger "Unable to search warehouse"
// @Router /warehouses [get]
func (h *WarehouseHandler) GetAllWareHouse() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		h.log.Log("WarehouseHandler", "INFO", "initializing GetAllWareHouse function")
		wareHouse, err := h.Srv.GetAllWareHouse()
		if err != nil {
			h.log.Log("WarehouseHandler", "ERROR", fmt.Sprintf("Error: %v", err))
			response.JSON(w, http.StatusInternalServerError, responses.CreateResponseBody(err.Error(), nil))
			return
		}

		h.log.Log("WarehouseHandler", "INFO", "GetAllWareHouse completed successfully")
		response.JSON(w, http.StatusOK, responses.CreateResponseBody("", wareHouse))
	}
}

func (h *WarehouseHandler) GetWareHouseByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		h.log.Log("WarehouseHandler", "INFO", "initializing GetWareHouseByID function")
		id, err := strconv.Atoi(chi.URLParam(r, "id"))

		if err != nil {
			h.log.Log("WarehouseHandler", "ERROR", fmt.Sprintf("Error: %v", err))
			response.JSON(w, http.StatusBadRequest, responses.CreateResponseBody("invalid id", nil))
			return
		}

		warehouse, err := h.Srv.GetByIDWareHouse(id)

		if err != nil {
			if err, ok := err.(*customerror.WareHouseError); ok {
				h.log.Log("WarehouseHandler", "ERROR", fmt.Sprintf("Error: %v", err))
				response.JSON(w, err.Code, responses.CreateResponseBody(err.Error(), nil))
				return
			}
			h.log.Log("WarehouseHandler", "ERROR", fmt.Sprintf("Error: %v", err))
			response.JSON(w, http.StatusInternalServerError, responses.CreateResponseBody("unable to search warehouse", nil))

			return
		}

		h.log.Log("WarehouseHandler", "INFO", "GetWareHouseByID completed successfully")
		response.JSON(w, http.StatusOK, responses.CreateResponseBody("", warehouse))
	}
}

// DeleteByIDWareHouse deletes a warehouse by its ID.
// @Summary Delete a warehouse
// @Description Delete a warehouse by its ID
// @Tags Warehouses
// @Produce json
// @Param id path int true "Warehouse ID"
// @Success 204 "No Content"
// @Failure 400 {object} model.ErrorResponseSwagger "Invalid ID"
// @Failure 500 {object} model.ErrorResponseSwagger "Unable to delete warehouse"
// @Router /warehouses/{id} [delete]
func (h *WarehouseHandler) DeleteByIDWareHouse() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		h.log.Log("WarehouseHandler", "INFO", "initializing DeleteByIDWareHouse function")
		id, err := strconv.Atoi(chi.URLParam(r, "id"))

		if err != nil {
			h.log.Log("WarehouseHandler", "ERROR", fmt.Sprintf("Error: %v", err))
			response.JSON(w, http.StatusBadRequest, responses.CreateResponseBody("invalid id", nil))
			return
		}

		err = h.Srv.DeleteByIDWareHouse(id)

		if err != nil {
			if err, ok := err.(*customerror.WareHouseError); ok {
				h.log.Log("WarehouseHandler", "ERROR", fmt.Sprintf("Error: %v", err))
				response.JSON(w, err.Code, responses.CreateResponseBody(err.Error(), nil))
				return
			}
		}

		h.log.Log("WarehouseHandler", "INFO", "DeleteByIDWareHouse completed successfully")
		response.JSON(w, http.StatusNoContent, nil)
	}
}

// PostWareHouse creates a new warehouse.
// @Summary Create a new warehouse
// @Description Create a new warehouse
// @Tags Warehouses
// @Accept json
// @Produce json
// @Param warehouse body model.WareHouse true "Warehouse details"
// @Success 201 {object} model.WareHousesResponseSwagger{data=model.WareHouse}
// @Failure 400 {object} model.ErrorResponseSwagger "Invalid request body"
// @Failure 422 {object} model.ErrorResponseSwagger "JSON syntax error Or Mandatory fields not filled in"
// @Failure 500 {object} model.ErrorResponseSwagger "Unable to post warehouse"
// @Router /warehouses [post]
func (h *WarehouseHandler) PostWareHouse() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		h.log.Log("WarehouseHandler", "INFO", "initializing PostWareHouse function")
		var reqBody model.WareHouse

		if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
			h.log.Log("WarehouseHandler", "ERROR", fmt.Sprintf("Error: %v", err))
			response.JSON(w, http.StatusBadRequest, responses.CreateResponseBody("invalid request body", nil))
			return
		}

		err := reqBody.ValidateEmptyFields(false)

		if err != nil {
			h.log.Log("WarehouseHandler", "ERROR", fmt.Sprintf("Error: %v", err))
			response.JSON(w, http.StatusUnprocessableEntity, responses.CreateResponseBody(err.Error(), nil))
			return
		}

		warehouse, err := h.Srv.PostWareHouse(reqBody)

		if err != nil {
			if err, ok := err.(*customerror.WareHouseError); ok {
				h.log.Log("WarehouseHandler", "ERROR", fmt.Sprintf("Error: %v", err))
				response.JSON(w, err.Code, responses.CreateResponseBody(err.Error(), nil))
				return
			}

			h.log.Log("WarehouseHandler", "ERROR", fmt.Sprintf("Error: %v", err))
			response.JSON(w, http.StatusInternalServerError, responses.CreateResponseBody("unable to post warehouse", nil))

			return
		}

		h.log.Log("WarehouseHandler", "INFO", "PostWareHouse completed successfully")

		response.JSON(w, http.StatusCreated, responses.CreateResponseBody("", warehouse))
	}
}

// UpdateWareHouse updates a warehouse by its ID.
// @Summary Update a warehouse
// @Description Update a warehouse by its ID
// @Tags Warehouses
// @Accept json
// @Produce json
// @Param id path int true "Warehouse ID"
// @Param warehouse body model.WareHouse true "Updated warehouse details"
// @Success 200 {object} model.WareHousesResponseSwagger{data=model.WareHouse}
// @Failure 400 {object} model.ErrorResponseSwagger "Invalid ID or Invalid request body"
// @Failure 422 {object} model.ErrorResponseSwagger "JSON syntax error Or Mandatory fields not filled in"
// @Failure 500 {object} model.ErrorResponseSwagger "Unable to update warehouse"
// @Router /warehouses/{id} [put]
func (h *WarehouseHandler) UpdateWareHouse() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		h.log.Log("WarehouseHandler", "INFO", "initializing UpdateWareHouse function")
		var reqBody model.WareHouse

		id, err := strconv.Atoi(chi.URLParam(r, "id"))

		if err != nil {
			h.log.Log("WarehouseHandler", "ERROR", fmt.Sprintf("Error: %v", err))
			response.JSON(w, http.StatusBadRequest, responses.CreateResponseBody("invalid id", nil))
			return
		}

		if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
			h.log.Log("WarehouseHandler", "ERROR", fmt.Sprintf("Error: %v", err))
			response.JSON(w, http.StatusBadRequest, responses.CreateResponseBody("invalid request body", nil))
			return
		}

		err = reqBody.ValidateEmptyFields(true)

		if err != nil {
			h.log.Log("WarehouseHandler", "ERROR", fmt.Sprintf("Error: %v", err))
			response.JSON(w, http.StatusUnprocessableEntity, responses.CreateResponseBody(err.Error(), nil))
			return
		}

		warehouse, err := h.Srv.UpdateWareHouse(id, reqBody)

		if err != nil {
			if err, ok := err.(*customerror.WareHouseError); ok {
				h.log.Log("WarehouseHandler", "ERROR", fmt.Sprintf("Error: %v", err))
				response.JSON(w, err.Code, responses.CreateResponseBody(err.Error(), nil))
				return
			}

			response.JSON(w, http.StatusInternalServerError, responses.CreateResponseBody("unable to update warehouse", nil))
			return
		}

		h.log.Log("WarehouseHandler", "INFO", "UpdateWareHouse completed successfully")
		response.JSON(w, http.StatusOK, responses.CreateResponseBody("", warehouse))
	}
}
