package handler

import (
	"encoding/json"
	"net/http"

	"github.com/bootcamp-go/web/response"
	"github.com/maxwelbm/alkemy-g7.git/internal/handler/responses"
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/maxwelbm/alkemy-g7.git/internal/service/interfaces"
	"github.com/maxwelbm/alkemy-g7.git/pkg/custom_error"
)

type PurchaseOrderHandler struct {
	svc interfaces.IPurchaseOrdersService
}

func NewPurchaseOrderHandler(svc interfaces.IPurchaseOrdersService) *PurchaseOrderHandler {
	return &PurchaseOrderHandler{svc}
}

func (h *PurchaseOrderHandler) HandlerCreatePurchaseOrder(w http.ResponseWriter, r *http.Request) {
	var reqBody model.PurchaseOrder

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(&reqBody)

	if err != nil {
		response.JSON(w, http.StatusUnprocessableEntity, responses.CreateResponseBody("JSON syntax error. Please verify your input.", nil))
		return
	}

	err = reqBody.ValidateEmptyFields()
	if err != nil {
		response.JSON(w, http.StatusUnprocessableEntity, responses.CreateResponseBody(err.Error(), nil))
		return
	}

	purchaseOrder, err := h.svc.CreatePurchaseOrder(reqBody)

	if err != nil {

		if err, ok := err.(*custom_error.BuyerError); ok {
			response.JSON(w, err.Code, responses.CreateResponseBody(err.Error(), nil))
			return
		}

		if err, ok := err.(*custom_error.GenericError); ok {
			response.JSON(w, err.Code, responses.CreateResponseBody(err.Error(), nil))
			return
		}

		if err, ok := err.(*custom_error.PurcahseOrderError); ok {
			response.JSON(w, err.Code, responses.CreateResponseBody(err.Error(), nil))
			return
		}

		response.JSON(w, http.StatusInternalServerError, responses.CreateResponseBody("Unable to create purchase order", nil))
		return
	}

	response.JSON(w, http.StatusCreated, responses.CreateResponseBody("", purchaseOrder))

}
