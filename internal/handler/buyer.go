package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"
	"github.com/maxwelbm/alkemy-g7.git/internal/handler/responses"
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/maxwelbm/alkemy-g7.git/internal/service"
	"github.com/maxwelbm/alkemy-g7.git/internal/service/interfaces"
	"github.com/maxwelbm/alkemy-g7.git/pkg/custom_error"
)

type BuyerHandler struct {
	svc interfaces.IBuyerservice
}

func NewBuyerHandler(svc *service.BuyerService) *BuyerHandler {
	return &BuyerHandler{svc}
}

func (bh *BuyerHandler) HandlerGetAllBuyers(w http.ResponseWriter, r *http.Request) {

	buyers, err := bh.svc.GetAllBuyer()
	if err != nil {
		response.JSON(w, http.StatusBadRequest, responses.CreateResponseBody("Unable to list Buyers", nil))
		return
	}

	response.JSON(w, http.StatusOK, responses.CreateResponseBody("", buyers))

}

func (bh *BuyerHandler) HandlerGetBuyerById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		response.JSON(w, http.StatusBadRequest, responses.CreateResponseBody("Invalid ID", nil))
		return
	}

	buyer, err := bh.svc.GetBuyerByID(id)

	if err != nil {

		if err, ok := err.(*custom_error.BuyerError); ok {
			response.JSON(w, err.Code, responses.CreateResponseBody(err.Error(), nil))
			return
		}

		response.JSON(w, http.StatusBadRequest, responses.CreateResponseBody("Unable to search for buyer", nil))
		return
	}

	response.JSON(w, http.StatusOK, responses.CreateResponseBody("", buyer))

}

func (bh *BuyerHandler) HandlerDeleteBuyerById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		response.JSON(w, http.StatusBadRequest, responses.CreateResponseBody("Invalid ID", nil))
		return
	}

	err = bh.svc.DeleteBuyerByID(id)

	if err != nil {

		if err, ok := err.(*custom_error.BuyerError); ok {
			response.JSON(w, err.Code, responses.CreateResponseBody(err.Error(), nil))
			return
		}

		response.JSON(w, http.StatusInternalServerError, responses.CreateResponseBody("Unable to delete buyer", nil))
		return
	}

	response.JSON(w, http.StatusNoContent, nil)

}

func (bh *BuyerHandler) HandlerCreateBuyer(w http.ResponseWriter, r *http.Request) {
	var reqBody model.Buyer

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(&reqBody)

	if err != nil {
		response.JSON(w, http.StatusUnprocessableEntity, responses.CreateResponseBody("JSON syntax error. Please verify your input.", nil))
		return
	}

	err = reqBody.ValidateEmptyFields(false)

	if err != nil {
		response.JSON(w, http.StatusUnprocessableEntity, responses.CreateResponseBody(err.Error(), nil))
		return
	}

	buyer, err := bh.svc.CreateBuyer(reqBody)

	if err != nil {
		if err, ok := err.(*custom_error.BuyerError); ok {
			response.JSON(w, err.Code, responses.CreateResponseBody(err.Error(), nil))
			return
		}

		response.JSON(w, http.StatusInternalServerError, responses.CreateResponseBody("Unable to create buyer", nil))
		return

	}

	response.JSON(w, http.StatusCreated, responses.CreateResponseBody("", buyer))

}

func (bh *BuyerHandler) HandlerUpdateBuyer(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		response.JSON(w, http.StatusBadRequest, responses.CreateResponseBody("Invalid ID", nil))
		return
	}

	var reqBody model.Buyer

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	err = decoder.Decode(&reqBody)

	if err != nil {
		response.JSON(w, http.StatusUnprocessableEntity, responses.CreateResponseBody("JSON syntax error. Please verify your input.", nil))
		return
	}

	err = reqBody.ValidateEmptyFields(true)

	if err != nil {
		response.JSON(w, http.StatusUnprocessableEntity, responses.CreateResponseBody(err.Error(), nil))
		return
	}

	buyer, err := bh.svc.UpdateBuyer(id, reqBody)

	if err != nil {

		if err, ok := err.(*custom_error.BuyerError); ok {
			response.JSON(w, err.Code, responses.CreateResponseBody(err.Error(), nil))
			return
		}

		response.JSON(w, http.StatusInternalServerError, responses.CreateResponseBody("Unable to update buyer", nil))
		return

	}

	response.JSON(w, http.StatusOK, responses.CreateResponseBody("", buyer))

}

func (bh *BuyerHandler) HandlerCountPurchaseOrderBuyer(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")

	if idStr == "" {
		count, err := bh.svc.CountPurchaseOrderBuyer()
		if err != nil {
			if err, ok := err.(*custom_error.BuyerError); ok {
				response.JSON(w, err.Code, responses.CreateResponseBody(err.Error(), nil))
				return
			}
			response.JSON(w, http.StatusInternalServerError, responses.CreateResponseBody("Unable to update buyer", nil))
			return
		}

		response.JSON(w, http.StatusOK, responses.CreateResponseBody("", count))
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.JSON(w, http.StatusBadRequest, responses.CreateResponseBody("Invalid ID", nil))
		return
	}

	count, err := bh.svc.CountPurchaseOrderByBuyerID(id)
	if err != nil {
		if err, ok := err.(*custom_error.BuyerError); ok {
			response.JSON(w, err.Code, responses.CreateResponseBody(err.Error(), nil))
			return
		}
		response.JSON(w, http.StatusInternalServerError, responses.CreateResponseBody("Unable to update buyer", nil))
		return
	}

	response.JSON(w, http.StatusOK, responses.CreateResponseBody("", count))

}
