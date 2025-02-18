package handler

import (
	"encoding/json"
	"fmt"
	"github.com/maxwelbm/alkemy-g7.git/pkg/logger"
	"net/http"

	"github.com/bootcamp-go/web/response"
	"github.com/maxwelbm/alkemy-g7.git/internal/handler/responses"
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/maxwelbm/alkemy-g7.git/internal/service/interfaces"
	"github.com/maxwelbm/alkemy-g7.git/pkg/customerror"
)

type PurchaseOrderHandler struct {
	Svc interfaces.IPurchaseOrdersService
	log logger.Logger
}

func NewPurchaseOrderHandler(svc interfaces.IPurchaseOrdersService, log logger.Logger) *PurchaseOrderHandler {
	return &PurchaseOrderHandler{svc, log}
}

// HandlerCreatePurchaseOrder handles the creation of a new purchase order.
// @Summary Create a new purchase order
// @Description This endpoint allows you to create a new purchase order by providing the necessary details in the request body.
// @Tags PurchaseOrder
// @Accept json
// @Produce json
// @Param purchaseOrder body model.PurchaseOrder true "Purchase Order"
// @Success 201 {object} model.PurchaseOrderResponseSwagger{data=model.PurchaseOrder} "Purchase order created successfully"
// @Failure 404 {object} model.ErrorResponseSwagger "Buyer or ProductRec not found"
// @Failure 409 {object} model.ErrorResponseSwagger "Order number already exists"
// @Failure 422 {object} model.ErrorResponseSwagger "JSON syntax error Or Mandatory fields not filled in"
// @Failure 500 {object} model.ErrorResponseSwagger "Internal Server Error"
// @Router /purchaseOrders [post]
func (h *PurchaseOrderHandler) HandlerCreatePurchaseOrder(w http.ResponseWriter, r *http.Request) {
	h.log.Log("PurchaseOrderHandler", "INFO", "initializing Request CreatePurchaseOrder")
	var reqBody model.PurchaseOrder

	decoder := json.NewDecoder(r.Body)

	decoder.DisallowUnknownFields()
	err := decoder.Decode(&reqBody)

	if err != nil {
		h.log.Log("PurchaseOrderHandler", "ERROR", fmt.Sprintf("Error: %v", err))
		response.JSON(w, http.StatusUnprocessableEntity, responses.CreateResponseBody("JSON syntax error. Please verify your input.", nil))

		return
	}
	h.log.Log("PurchaseOrderHandler", "INFO", "Validating fields received ")
	err = reqBody.ValidateEmptyFields()

	if err != nil {
		h.log.Log("PurchaseOrderHandler", "ERROR", fmt.Sprintf("Error: %v", err))
		response.JSON(w, http.StatusUnprocessableEntity, responses.CreateResponseBody(err.Error(), nil))

		return
	}

	purchaseOrder, err := h.Svc.CreatePurchaseOrder(reqBody)

	if err != nil {
		if err, ok := err.(*customerror.BuyerError); ok {
			h.log.Log("PurchaseOrderHandler", "ERROR", fmt.Sprintf("Error: %v", err))
			response.JSON(w, err.Code, responses.CreateResponseBody(err.Error(), nil))

			return
		}

		if err, ok := err.(*customerror.GenericError); ok {
			h.log.Log("PurchaseOrderHandler", "ERROR", fmt.Sprintf("Error: %v", err))
			response.JSON(w, err.Code, responses.CreateResponseBody(err.Error(), nil))

			return
		}

		if err, ok := err.(*customerror.PurcahseOrderError); ok {
			h.log.Log("PurchaseOrderHandler", "ERROR", fmt.Sprintf("Error: %v", err))
			response.JSON(w, err.Code, responses.CreateResponseBody(err.Error(), nil))

			return
		}
		h.log.Log("PurchaseOrderHandler", "ERROR", fmt.Sprintf("Error: %v", err))
		response.JSON(w, http.StatusInternalServerError, responses.CreateResponseBody("Unable to create purchase order", nil))

		return
	}
	h.log.Log("PurchaseOrderHandler", "INFO", "Purchase created successfull")
	response.JSON(w, http.StatusCreated, responses.CreateResponseBody("", purchaseOrder))
}
